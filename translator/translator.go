package translator

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/dlclark/regexp2"
	"modernc.org/cc/v4"
)

type Translator struct {
	validations        Validations
	rules              Rules
	compiledRxs        map[RuleAction]RxMap
	compiledPtrTipRxs  PtrTipRxMap
	compiledTypeTipRxs TypeTipRxMap
	compiledMemTipRxs  MemTipRxList
	constRules         ConstRules
	typemap            CTypeMap
	builtinTypemap     CTypeMap
	// builtinTypemap2 is an optional
	// typemap with overriden const (u)char rules
	builtinTypemap2 CTypeMap
	ignoredFiles    map[string]struct{}

	valueMap  map[string]Value
	exprMap   map[string]string
	tagMap    map[string]*CDecl
	lenFields map[string]string

	defines  []*CDecl
	typedefs []*CDecl
	declares []*CDecl

	typedefsSet    map[string]struct{}
	typedefKinds   map[string]CTypeKind
	transformCache *NameTransformCache

	ptrTipCache  *TipCache
	typeTipCache *TipCache
	memTipCache  *TipCache
}

type RxMap map[RuleTarget][]Rx

type Rx struct {
	From *regexp.Regexp
	To   []byte
	//
	Transform RuleTransform
}

type PtrTipRxMap map[TipScope][]TipSpecRx
type TypeTipRxMap map[TipScope][]TipSpecRx
type MemTipRxList []TipSpecRx

type TipSpecRx struct {
	Target  *regexp.Regexp
	Default Tip
	tips    Tips
	self    Tip
}

func (t TipSpecRx) TipAt(i int) Tip {
	if i < len(t.tips) {
		if tip := t.tips[i]; tip.IsValid() {
			return tip
		}
	}
	return t.Default
}

func (t TipSpecRx) HasTip(q Tip) bool {
	for _, tip := range t.tips {
		if tip.IsValid() && tip == q {
			return true
		}
	}
	return false
}

func (t TipSpecRx) Self() Tip {
	if t.self.IsValid() {
		return t.self
	}
	return t.Default
}

type Config struct {
	Validations        Validations       `yaml:"Validations"`
	Rules              Rules             `yaml:"Rules"`
	ConstRules         ConstRules        `yaml:"ConstRules"`
	PtrTips            PtrTips           `yaml:"PtrTips"`
	TypeTips           TypeTips          `yaml:"TypeTips"`
	MemTips            MemTips           `yaml:"MemTips"`
	Typemap            CTypeMap          `yaml:"Typemap"`
	ConstCharIsString  *bool             `yaml:"ConstCharIsString"`
	ConstUCharIsString *bool             `yaml:"ConstUCharIsString"`
	LenFields          map[string]string `yaml:"LenFields"`

	IgnoredFiles []string `yaml:"-"`
	LongIs64Bit  bool     `yaml:"-"`
}

func New(cfg *Config) (*Translator, error) {
	if cfg == nil {
		cfg = &Config{}
	}

	var constCharAsString bool
	if cfg.ConstCharIsString == nil {
		constCharAsString = true
	} else {
		constCharAsString = *cfg.ConstCharIsString
	}

	var constUCharAsString bool
	if cfg.ConstUCharIsString == nil {
		constUCharAsString = false
	} else {
		constUCharAsString = *cfg.ConstUCharIsString
	}

	t := &Translator{
		validations:        cfg.Validations,
		rules:              cfg.Rules,
		constRules:         cfg.ConstRules,
		typemap:            cfg.Typemap,
		builtinTypemap:     getCTypeMap(constCharAsString, constUCharAsString, cfg.LongIs64Bit),
		builtinTypemap2:    getCTypeMap(true, false, cfg.LongIs64Bit),
		compiledRxs:        make(map[RuleAction]RxMap),
		compiledPtrTipRxs:  make(PtrTipRxMap),
		compiledTypeTipRxs: make(TypeTipRxMap),
		valueMap:           make(map[string]Value),
		exprMap:            make(map[string]string),
		tagMap:             make(map[string]*CDecl),
		lenFields:          cfg.LenFields,
		typedefsSet:        make(map[string]struct{}),
		typedefKinds:       make(map[string]CTypeKind),
		ignoredFiles:       make(map[string]struct{}),
		transformCache:     &NameTransformCache{},
		ptrTipCache:        &TipCache{},
		typeTipCache:       &TipCache{},
		memTipCache:        &TipCache{},
	}
	for _, p := range cfg.IgnoredFiles {
		t.ignoredFiles[p] = struct{}{}
	}
	for _, action := range ruleActions {
		if rxMap, err := getRuleActionRxs(t.rules, action); err != nil {
			return nil, err
		} else {
			t.compiledRxs[action] = rxMap
		}
	}
	if rxMap, err := getPtrTipRxs(cfg.PtrTips); err != nil {
		return nil, err
	} else {
		t.compiledPtrTipRxs = rxMap
	}
	if rxMap, err := getTypeTipRxs(cfg.TypeTips); err != nil {
		return nil, err
	} else {
		t.compiledTypeTipRxs = rxMap
	}
	if rxList, err := getMemTipRxs(cfg.MemTips); err != nil {
		return nil, err
	} else {
		t.compiledMemTipRxs = rxList
	}

	for _, v := range t.validations {
		if _, err := regexp2.Compile(v.MatchedFunc, 0); err != nil {
			return nil, fmt.Errorf("translator: %s, invalid regexp %s", err.Error(), v.MatchedFunc)
		}
	}

	return t, nil
}

func getRuleActionRxs(rules Rules, action RuleAction) (RxMap, error) {
	rxMap := make(RxMap, len(rules))
	for target, specs := range rules {
		for _, spec := range specs {
			if len(spec.Load) > 0 {
				if s, ok := builtinRules[spec.Load]; ok {
					spec.LoadSpec(s)
				} else {
					return nil, fmt.Errorf("no builtin rule found: %s", spec.Load)
				}
			}
			if spec.Action == ActionNone {
				spec.Action = ActionReplace
				spec.To = "${_src}"
				if len(spec.From) == 0 {
					spec.From = "(?P<_src>.*)"
				} else {
					spec.From = fmt.Sprintf("(?P<_src>%s)", spec.From)
				}
			} else if len(spec.From) == 0 {
				spec.From = "(.*)"
			}
			if spec.Action != action {
				continue
			}
			rxFrom, err := regexp.Compile(spec.From)
			if err != nil {
				return nil, fmt.Errorf("translator: %s rules: invalid regexp %s", target, spec.From)
			}
			rx := Rx{From: rxFrom, To: []byte(spec.To)}
			if spec.Action == ActionReplace {
				rx.Transform = spec.Transform
			}
			rxMap[target] = append(rxMap[target], rx)
		}
	}
	return rxMap, nil
}

func getPtrTipRxs(tips PtrTips) (PtrTipRxMap, error) {
	rxMap := make(PtrTipRxMap, len(tips))
	for scope, specs := range tips {
		for _, spec := range specs {
			if len(spec.Target) == 0 {
				continue
			}
			rx, err := regexp.Compile(spec.Target)
			if err != nil {
				return nil, fmt.Errorf("translator: ptr tip in %s scope: invalid regexp %s",
					scope, spec.Target)
			}
			specRx := TipSpecRx{
				Target:  rx,
				Default: spec.Default,
				tips:    spec.Tips,
				self:    spec.Self,
			}
			rxMap[scope] = append(rxMap[scope], specRx)
			if scope == TipScopeStruct {
				rxMap[TipScopeType] = append(rxMap[TipScopeType], specRx)
			}
		}
	}
	return rxMap, nil
}

func getTypeTipRxs(tips TypeTips) (TypeTipRxMap, error) {
	rxMap := make(TypeTipRxMap, len(tips))
	for scope, specs := range tips {
		for _, spec := range specs {
			if len(spec.Target) == 0 {
				continue
			}
			rx, err := regexp.Compile(spec.Target)
			if err != nil {
				return nil, fmt.Errorf("translator: type tip in %s scope: invalid regexp %s",
					scope, spec.Target)
			}
			specRx := TipSpecRx{
				Target:  rx,
				Default: spec.Default,
				tips:    spec.Tips,
				self:    spec.Self,
			}
			rxMap[scope] = append(rxMap[scope], specRx)
			if scope == TipScopeStruct {
				rxMap[TipScopeType] = append(rxMap[TipScopeType], specRx)
			}
		}
	}
	return rxMap, nil
}

func getMemTipRxs(tips MemTips) (MemTipRxList, error) {
	var list MemTipRxList
	for _, spec := range tips {
		if len(spec.Target) == 0 {
			continue
		}
		rx, err := regexp.Compile(spec.Target)
		if err != nil {
			return nil, fmt.Errorf("translator: mem tip: invalid regexp %s", spec.Target)
		}
		specRx := TipSpecRx{
			Target:  rx,
			Default: spec.Default,
			tips:    spec.Tips,
			self:    spec.Self,
		}
		list = append(list, specRx)
	}
	return list, nil
}

type declList []*CDecl

func (s declList) Len() int      { return len(s) }
func (s declList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s declList) Less(i, j int) bool {
	if s[i].Position.Filename != s[j].Position.Filename {
		return s[i].Position.Filename < s[j].Position.Filename
	} else if s[i].Position.Offset != s[j].Position.Offset {
		return s[i].Position.Offset < s[j].Position.Offset
	} else {
		return s[i].Name < s[j].Name
	}
}

func (t *Translator) Learn(ast *cc.AST) {
	t.walkTranslationUnit(ast.TranslationUnit)
	t.resolveTypedefs(t.typedefs)
	sort.Sort(declList(t.declares))
	sort.Sort(declList(t.typedefs))
	t.collectDefines(t.declares, ast.Macros)
	sort.Sort(declList(t.defines))
}

// This has been left intentionally.
//
// func (t *Translator) Report() {
// 	fmt.Printf("[!] TAGS:\n")
// 	for tag, decl := range t.tagMap {
// 		fmt.Printf("%s refers to %v\n", tag, decl)
// 	}
// 	fmt.Printf("\n\n\n[!] TYPEDEFs:\n")
// 	for _, decl := range t.typedefs {
// 		fmt.Printf("%v\n", decl)
// 	}
// 	fmt.Printf("\n\n\n[!] DECLARATIONS:\n")
// 	for _, decl := range t.declares {
// 		fmt.Printf("%v\n", decl)
// 	}
// 	fmt.Printf("\n\n\n[!] const (")
// 	for _, line := range t.defines {
// 		fmt.Printf("\n// %s\n//   > define %s %v\n%s = %s",
// 			SrcLocation(line.Pos), line.Name, line.Src,
// 			t.TransformName(TargetConst, string(line.Name)), line.Expression)
// 	}
// 	fmt.Printf("\n)\n\n")
// }

func (t *Translator) collectDefines(declares []*CDecl, defines map[string]*cc.Macro) {
	seen := make(map[string]struct{}, len(defines)+len(declares))

	// traverse declared constants because macro can reference them,
	// so we need to collect a map of valid references beforehand
	for _, decl := range declares {
		if t.IsTokenIgnored(decl.Position) {
			continue
		}
		if decl.Spec.Kind() == EnumKind {
			enumMembers := decl.Spec.(*CEnumSpec).Members
			for _, member := range enumMembers {
				if !t.IsAcceptableName(TargetConst, member.Name) {
					continue
				}
				seen[member.Name] = struct{}{}
			}
		}
	}

	// double traverse because macros can depend on each other and the map
	// brings a randomized order of them.
	for _, macro := range defines {
		if t.IsTokenIgnored(macro.Position()) {
			continue
		} else if macro.IsFnLike {
			continue
		}
		name := string(macro.Name.SrcStr())
		if !t.IsAcceptableName(TargetConst, name) {
			continue
		}
		seen[name] = struct{}{}
	}

	for _, macro := range defines {
		if t.IsTokenIgnored(macro.Position()) {
			continue
		}
		name := string(macro.Name.SrcStr())
		if _, ok := seen[name]; !ok {
			continue
		}
		expand := false
		if t.constRules[ConstDefines] == ConstExpand {
			expand = true
		}

		if !expand {
			switch macro.Type().Kind() {
			case cc.InvalidKind: // unresolved value -> try to expand
				expand = true
			case cc.Bool: // ban bools
				continue
			}
			if !expand {
				t.defines = append(t.defines, &CDecl{
					IsDefine: true,
					Name:     name,
					Value:    Value(macro.Value()),
					Position: macro.Position(),
				})
				continue
			}
		} else if macro.Type().Kind() == cc.Bool {
			// ban bools
			continue
		}
		tokens := macro.ReplacementList()
		srcParts := make([]string, 0, len(tokens))
		exprParts := make([]string, 0, len(tokens))
		valid := true

		// TODO: some state machine
		needsTypecast := false
		typecastValue := false
		typecastValueParens := 0

		for _, token := range tokens {
			src := token.SrcStr()
			srcParts = append(srcParts, src)
			switch token.Ch {
			case rune(cc.IDENTIFIER):
				if _, ok := seen[src]; ok {
					// const reference
					exprParts = append(exprParts, string(t.TransformName(TargetConst, src, true)))
				} else if _, ok := t.typedefsSet[src]; ok {
					// type reference
					needsTypecast = true
					exprParts = append(exprParts, string(t.TransformName(TargetType, src, true)))
				} else {
					// an unresolved reference
					valid = false
					break
				}
			default:
				// TODO: state machine
				const (
					lparen = rune(40)
					rparen = rune(41)
				)
				switch {
				case needsTypecast && token.Ch == rparen:
					typecastValue = true
					needsTypecast = false
					exprParts = append(exprParts, src+"(")
				case typecastValue && token.Ch == lparen:
					typecastValueParens++
				case typecastValue && token.Ch == rparen:
					if typecastValueParens == 0 {
						typecastValue = false
						exprParts = append(exprParts, ")"+src)
					} else {
						typecastValueParens--
					}
				default:
					// somewhere in the world a kitten died because of this
					if token.Ch == '~' {
						src = "^"
					}
					if runes := []rune(src); len(runes) > 0 && isNumeric(runes) {
						// TODO(xlab): better const handling
						src = readNumeric(runes)
					}
					exprParts = append(exprParts, src)
				}
			}
			if !valid {
				break
			}
		}
		if typecastValue {
			// still in typecast value, need to close paren
			exprParts = append(exprParts, ")")
			typecastValue = false
		}
		if !valid {
			if macro.Type().Kind() == cc.InvalidKind {
				// fallback to the evaluated value
				t.defines = append(t.defines, &CDecl{
					IsDefine: true,
					Name:     name,
					Value:    Value(macro.Value),
					Position: macro.Position(),
				})
			}
			continue
		}
		t.defines = append(t.defines, &CDecl{
			IsDefine:   true,
			Name:       name,
			Expression: strings.Join(exprParts, " "),
			Src:        strings.Join(srcParts, " "),
			Position:   macro.Position(),
		})
	}
}

func (t *Translator) resolveTypedefs(typedefs []*CDecl) {
	for _, decl := range typedefs {
		if t.IsTokenIgnored(decl.Position) {
			continue
		}
		if decl.Spec.Kind() != TypeKind {
			t.typedefKinds[decl.Name] = decl.Spec.Kind()
			continue
		}
		if goSpec := t.TranslateSpec(decl.Spec); goSpec.IsPlain() {
			t.typedefKinds[decl.Name] = PlainTypeKind
		} else if goSpec.Kind != TypeKind {
			t.typedefKinds[decl.Name] = goSpec.Kind
		}
	}
}

func (t *Translator) TransformName(target RuleTarget, str string, publicOpt ...bool) []byte {
	if len(str) == 0 {
		return emptyStr
	}
	targetVisibility := NoTarget
	if len(publicOpt) > 0 {
		if publicOpt[0] {
			targetVisibility = TargetPublic
		} else {
			targetVisibility = TargetPrivate
		}
	}
	if name, ok := t.transformCache.Get(target, targetVisibility, str); ok {
		return name
	}

	var name []byte
	switch target {
	case TargetGlobal, TargetPostGlobal, TargetPrivate, TargetPublic:
		name = []byte(str)
	default:
		// apply the global rules first
		name = t.TransformName(TargetGlobal, str)
	}

	for _, rx := range t.compiledRxs[ActionReplace][target] {
		indices := rx.From.FindAllSubmatchIndex(name, -1)
		reference := make([]byte, len(name))
		copy(reference, name)

		// Itrate submatches backwards since we need to insert expanded
		// versions into the original name and doing so from beginning will shift indices
		// for latter inserts.
		//
		// Example flow:
		// doing title at _partitions in vpx_error_resilient_partitions
		// doing title at _resilient in vpx_error_resilientPartitions
		// doing title at _error in vpx_errorResilientPartitions
		// -> vpxErrorResilientPartitions
		for i := len(indices) - 1; i >= 0; i-- {
			idx := indices[i]
			buf := rx.From.Expand([]byte{}, rx.To, reference, idx)
			switch rx.Transform {
			case TransformLower:
				buf = bytes.ToLower(buf)
			case TransformTitle, TransformExport:
				if len(buf) > 0 {
					buf[0] = bytes.ToUpper(buf[:1])[0]
				}
			case TransformUnexport:
				if len(buf) > 0 {
					buf[0] = bytes.ToLower(buf[:1])[0]
				}
			case TransformUpper:
				buf = bytes.ToUpper(buf)
			}
			name = replaceBytes(name, idx, buf)
		}
	}
	switch target {
	case TargetGlobal, TargetPostGlobal, TargetPrivate, TargetPublic:
	default:
		// apply post-global & visibility rules in the end
		name = t.TransformName(TargetPostGlobal, string(name))
		switch targetVisibility {
		case TargetPrivate, TargetPublic:
			name = t.TransformName(targetVisibility, string(name))
		}
		if isBuiltinName(name) {
			name = rewriteName(name)
		}
		t.transformCache.Set(target, targetVisibility, str, name)
		return name
	}
	return name
}

func (t *Translator) lookupSpec(spec CTypeSpec, useConstCharTypemap bool) (GoTypeSpec, bool) {
	if gospec, ok := t.typemap[spec]; ok {
		return gospec, true
	}
	if useConstCharTypemap {
		if gospec, ok := t.builtinTypemap2[spec]; ok {
			return gospec, true
		}
	} else {
		if gospec, ok := t.builtinTypemap[spec]; ok {
			return gospec, true
		}
	}

	if spec.Const {
		spec.Const = false
		if gospec, ok := t.typemap[spec]; ok {
			return gospec, true
		}
		if gospec, ok := t.builtinTypemap[spec]; ok {
			return gospec, true
		}
	}
	return GoTypeSpec{}, false
}

func (t *Translator) PtrTipRx(scope TipScope, name string) (TipSpecRx, bool) {
	if rx, ok := t.ptrTipCache.Get(scope, name); ok {
		return rx, true
	}
	for _, rx := range t.compiledPtrTipRxs[scope] {
		if rx.Target.MatchString(name) {
			t.ptrTipCache.Set(scope, name, rx)
			return rx, true
		}
	}
	if scope != TipScopeAny {
		for _, rx := range t.compiledPtrTipRxs[TipScopeAny] {
			if rx.Target.MatchString(name) {
				t.ptrTipCache.Set(scope, name, rx)
				return rx, true
			}
		}
	}
	return TipSpecRx{}, false
}

func (t *Translator) TypeTipRx(scope TipScope, name string) (TipSpecRx, bool) {
	if rx, ok := t.typeTipCache.Get(scope, name); ok {
		return rx, true
	}
	for _, rx := range t.compiledTypeTipRxs[scope] {
		if rx.Target.MatchString(name) {
			t.typeTipCache.Set(scope, name, rx)
			return rx, true
		}
	}
	if scope != TipScopeAny {
		for _, rx := range t.compiledTypeTipRxs[TipScopeAny] {
			if rx.Target.MatchString(name) {
				t.typeTipCache.Set(scope, name, rx)
				return rx, true
			}
		}
	}
	return TipSpecRx{}, false
}

func (t *Translator) MemTipRx(name string) (TipSpecRx, bool) {
	if rx, ok := t.memTipCache.Get(TipScopeType, name); ok {
		return rx, true
	}
	if rx, ok := t.memTipCache.Get(TipScopeStruct, name); ok {
		return rx, true
	}
	for _, rx := range t.compiledMemTipRxs {
		if rx.Target.MatchString(name) {
			t.memTipCache.Set(TipScopeType, name, rx)
			t.memTipCache.Set(TipScopeStruct, name, rx)
			return rx, true
		}
	}
	return TipSpecRx{}, false
}

func (t *Translator) TipRxsForSpec(scope TipScope,
	name string, spec CType) (ptr, typ, mem TipSpecRx) {
	var ptrOk, typOk, memOk bool
	if tag := spec.GetBase(); len(tag) > 0 {
		ptr, ptrOk = t.PtrTipRx(scope, tag)
		typ, typOk = t.TypeTipRx(scope, tag)
		mem, memOk = t.MemTipRx(tag)
	}
	if !ptrOk {
		ptr, _ = t.PtrTipRx(scope, name)
	}
	if !typOk {
		typ, _ = t.TypeTipRx(scope, name)
	}
	if !memOk {
		mem, _ = t.MemTipRx(name)
	}
	return
}

func (t *Translator) TranslateSpec(spec CType, tips ...Tip) GoTypeSpec {
	var ptrTip Tip
	var typeTip Tip
	for _, tip := range tips {
		switch tip.Kind() {
		case TipKindPtr:
			ptrTip = tip
		case TipKindType:
			typeTip = tip
		}
	}
	if len(spec.OuterArrays()) > 0 {
		ptrTip = TipPtrSRef
	}

	switch spec.Kind() {
	case TypeKind:
		typeSpec := spec.(*CTypeSpec)
		lookupSpec := CTypeSpec{
			Base:     typeSpec.Base,
			Const:    typeSpec.Const,
			Unsigned: typeSpec.Unsigned,
			Short:    typeSpec.Short,
			Long:     typeSpec.Long,
			Complex:  typeSpec.Complex,
			// Arrays skip
			// VarArrays skip
			Pointers: typeSpec.Pointers,
		}
		if typeSpec.Base == "char" {
			lookupSpec.Unsigned = typeSpec.Unsigned
			lookupSpec.Signed = typeSpec.Signed
		}

		if gospec, ok := t.lookupSpec(lookupSpec, typeTip == TipTypeString); ok {
			tag := typeSpec.CGoName()
			if gospec == UnsafePointerSpec && len(tag) > 0 {
				if decl, ok := t.tagMap[tag]; ok {
					if decl.Spec.GetPointers() <= gospec.Pointers {
						gospec.Pointers = gospec.Pointers - decl.Spec.GetPointers()
					}
				}
			}
			gospec.OuterArr.Prepend(typeSpec.OuterArr)
			gospec.InnerArr.Prepend(typeSpec.InnerArr)
			if gospec.Pointers == 0 && gospec.Slices > 0 {
				switch ptrTip {
				case TipPtrRef:
					gospec.Slices--
					gospec.Pointers++
				case TipPtrSRef:
					gospec.Pointers = gospec.Slices
					gospec.Slices = 0
				}
			}
			if gospec.Kind == TypeKind {
				if gospec.IsPlain() {
					gospec.Kind = PlainTypeKind
				} else if gospec.Kind = t.typedefKinds[gospec.Base]; gospec.Kind == TypeKind {
					if kind := t.typedefKinds[lookupSpec.Base]; kind != PlainTypeKind {
						gospec.Kind = kind
					}
				}
			}
			if typeTip != TipTypePlain {
				if t.IsAcceptableName(TargetType, typeSpec.Raw) {
					gospec.Raw = string(t.TransformName(TargetType, typeSpec.Raw))
					if gospec.Base != "unsafe.Pointer" {
						gospec.splitPointers(ptrTip, typeSpec.Pointers)
					}
				}
			}

			return gospec
		}
		wrapper := GoTypeSpec{
			Kind:     t.typedefKinds[lookupSpec.Base],
			OuterArr: typeSpec.OuterArr,
			InnerArr: typeSpec.InnerArr,
		}
		if lookupSpec.Pointers > 0 {
			for lookupSpec.Pointers > 0 {
				switch ptrTip {
				case TipPtrSRef:
					if lookupSpec.Pointers > 1 {
						wrapper.Slices++
					} else {
						wrapper.Pointers++
					}
				case TipPtrRef:
					wrapper.Pointers++
				default:
					wrapper.Slices++
				}
				lookupSpec.Pointers--
				if gospec, ok := t.lookupSpec(lookupSpec, typeTip == TipTypeString); ok {
					tag := typeSpec.CGoName()
					if gospec == UnsafePointerSpec && len(tag) > 0 {
						if decl, ok := t.tagMap[tag]; ok {
							if decl.Spec.GetPointers() <= gospec.Pointers {
								gospec.Pointers = gospec.Pointers - decl.Spec.GetPointers()
							}
						}
					}
					gospec.Slices += wrapper.Slices
					gospec.Pointers += wrapper.Pointers
					gospec.OuterArr.Prepend(wrapper.OuterArr)
					gospec.InnerArr.Prepend(wrapper.InnerArr)
					if gospec.Kind == TypeKind {
						if gospec.IsPlain() {
							gospec.Kind = PlainTypeKind
						} else if wrapper.Kind == TypeKind || wrapper.Kind == PlainTypeKind {
							if kind := t.typedefKinds[lookupSpec.Base]; kind != PlainTypeKind {
								gospec.Kind = kind
							}
						} else {
							gospec.Kind = wrapper.Kind
						}
					}
					if typeTip != TipTypePlain {
						if t.IsAcceptableName(TargetType, typeSpec.Raw) {
							gospec.Raw = string(t.TransformName(TargetType, typeSpec.Raw))
							if gospec.Base != "unsafe.Pointer" {
								gospec.Pointers = 0
								gospec.Slices = 0
								gospec.splitPointers(ptrTip, typeSpec.Pointers)
							}
						}
					}
					return gospec
				}
			}
		}
		if t.IsAcceptableName(TargetType, typeSpec.Raw) {
			wrapper.Raw = string(t.TransformName(TargetType, typeSpec.Raw))
		}
		wrapper.Base = string(t.TransformName(TargetType, lookupSpec.Base))
		switch wrapper.Kind {
		case TypeKind:
			if wrapper.IsPlain() {
				wrapper.Kind = PlainTypeKind
			}
		case FunctionKind:
			wrapper.Slices = 0
			wrapper.OuterArr = ArraySpec("")
			wrapper.InnerArr = ArraySpec("")
			wrapper.Pointers = spec.GetPointers()
		}
		return wrapper
	case FunctionKind:
		wrapper := GoTypeSpec{
			Kind:     spec.Kind(),
			Pointers: spec.GetPointers(),
		}
		fspec := spec.(*CFunctionSpec)
		decl, known := t.tagMap[fspec.Raw]
		if !known {
			// another try using the type
			decl, known = t.tagMap[fspec.Typedef]
		}
		if known {
			if decl.Spec.GetPointers() <= wrapper.Pointers {
				wrapper.Pointers = wrapper.Pointers - decl.Spec.GetPointers()
			}
		}
		if t.IsAcceptableName(TargetType, fspec.Raw) {
			wrapper.Raw = string(t.TransformName(TargetType, fspec.Raw))
		}
		wrapper.Base = "func"
		return wrapper
	default:
		wrapper := GoTypeSpec{
			Kind:     spec.Kind(),
			OuterArr: spec.OuterArrays(),
			InnerArr: spec.InnerArrays(),
		}
		switch wrapper.Kind {
		case OpaqueStructKind, StructKind, FunctionKind, UnionKind:
			decl, tagKnown := t.tagMap[spec.GetTag()]
			if !tagKnown && wrapper.Kind == FunctionKind {
				// another try using the type
				fspec := spec.(*CFunctionSpec)
				decl, tagKnown = t.tagMap[fspec.Typedef]
			}
			if tagKnown {
				if !decl.Spec.IsOpaque() {
					wrapper.Kind = StructKind
				}
				// count in the pointers of the base type under typedef
				// e.g. typedef struct a_T* a; int f(a *arg);
				// -> go_F(C.a *arg) but not go_F(C.a **arg) because of residual pointers.
				wrapper.Pointers = 0
				wrapper.splitPointers(ptrTip, spec.GetPointers()-decl.Spec.GetPointers())
			} else {
				wrapper.splitPointers(ptrTip, spec.GetPointers())
			}
		default:
			wrapper.splitPointers(ptrTip, spec.GetPointers())
		}
		if base := spec.GetBase(); len(base) > 0 {
			wrapper.Raw = string(t.TransformName(TargetType, base))
		} else if cgoName := spec.CGoName(); len(cgoName) > 0 {
			wrapper.Raw = "C." + cgoName
		}
		return wrapper
	}
}

func (t *Translator) NormalizeSpecPointers(spec CType) CType {
	if spec.Kind() == OpaqueStructKind {
		decl, tagKnown := t.tagMap[spec.GetTag()]
		if tagKnown {
			spec := spec.Copy()
			spec.SetPointers(spec.GetPointers() - decl.Spec.GetPointers())
			return spec
		}
	}
	return spec
}

func (t *Translator) CGoSpec(spec CType, asArg bool) CGoSpec {
	cgo := CGoSpec{
		Pointers: spec.GetPointers(),
		OuterArr: spec.OuterArrays(),
		InnerArr: spec.InnerArrays(),
	}
	if decl, ok := t.tagMap[spec.GetTag()]; ok {
		// count in the pointers of the base type under typedef
		cgo.Pointers = spec.GetPointers() - decl.Spec.GetPointers()
	}
	if typ, ok := spec.(*CTypeSpec); ok {
		if typ.Base == "void*" {
			if len(typ.Raw) == 0 {
				cgo.Base = "unsafe.Pointer"
				return cgo
			} else if !asArg {
				cgo.Pointers++
			}
		}
	}
	cgo.Base = "C." + spec.CGoName()
	return cgo
}

func (t *Translator) registerTagsOf(decl *CDecl) {
	switch decl.Spec.Kind() {
	case TypeKind, EnumKind, StructKind, OpaqueStructKind, UnionKind:
		tag := decl.Spec.GetTag()
		if decl.Spec.Kind() == TypeKind {
			tag = decl.Spec.CGoName()
		}
		if len(tag) > 0 {
			prev, hasPrev := t.tagMap[tag]
			switch {
			case !hasPrev:
				// first time seen -> store the tag
				t.tagMap[tag] = decl
			case decl.IsTypedef && !prev.IsTypedef:
				// a typedef for tag prevails simple tag declarations
				t.tagMap[tag] = decl
			case decl.Spec.IsComplete() && !prev.Spec.IsComplete():
				// replace an opaque struct with the complete template
				t.tagMap[tag] = decl
			}
		}
	case FunctionKind:
		if tag := decl.Spec.GetTag(); len(tag) > 0 {
			if _, ok := t.tagMap[tag]; !ok {
				t.tagMap[tag] = decl
			}
		}
	}
	switch typ := decl.Spec.(type) {
	case *CStructSpec:
		for _, m := range typ.Members {
			switch m.Spec.Kind() {
			case TypeKind, StructKind, OpaqueStructKind, UnionKind:
				tag := decl.Spec.GetTag()
				if decl.Spec.Kind() == TypeKind {
					tag = decl.Spec.CGoName()
				}
				if len(tag) > 0 {
					if prev, ok := t.tagMap[tag]; !ok {
						// first time seen -> store the tag
						t.tagMap[tag] = m
					} else if m.Spec.IsComplete() && !prev.Spec.IsComplete() {
						// replace an opaque struct with the complete template
						t.tagMap[tag] = m
					}
				}
			case FunctionKind:
				if tag := decl.Spec.GetTag(); len(tag) > 0 {
					if _, ok := t.tagMap[tag]; !ok {
						t.tagMap[tag] = decl
					}
				}
			}
		}
	}
}

func (t *Translator) IsAcceptableName(target RuleTarget, name string) bool {
	if rxs, ok := t.compiledRxs[ActionAccept][target]; ok {
		for _, rx := range rxs {
			if rx.From.MatchString(name) {
				// try to find explicit ignore rules
				if rxs, ok := t.compiledRxs[ActionIgnore][target]; ok {
					for _, rx := range rxs {
						if rx.From.MatchString(name) {
							// found an ignore rule, ignore the name
							return false
						}
					}
				}
				// no ignore rules found, accept the name
				return true
			}
		}
	}
	// try to find explicit ignore rules
	if rxs, ok := t.compiledRxs[ActionIgnore][target]; ok {
		for _, rx := range rxs {
			if rx.From.MatchString(name) {
				// found an ignore rule, ignore the name
				return false
			}
		}
	}
	if target != TargetGlobal {
		// we don't have any specific rules for this target, check global rules
		return t.IsAcceptableName(TargetGlobal, name)
	}
	// default to ignore
	return false
}

func (t *Translator) TagMap() map[string]*CDecl {
	return t.tagMap
}

func (t *Translator) LenFields() map[string]string {
	return t.lenFields
}

func (t *Translator) ExpressionMap() map[string]string {
	return t.exprMap
}

func (t *Translator) ValueMap() map[string]Value {
	return t.valueMap
}

func (t *Translator) Defines() []*CDecl {
	return t.defines
}

func (t *Translator) Declares() []*CDecl {
	return t.declares
}

func (t *Translator) Typedefs() []*CDecl {
	return t.typedefs
}

func (t *Translator) GetLibrarySymbolValidation(name string) (string, string, bool) {
	for _, v := range t.validations {
		if v.MatchFunc(name) {
			return v.ValidateFunc, v.Ret, true
		}
	}
	return "", "", false
}
