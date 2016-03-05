package translator

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/cznic/cc"
	"github.com/cznic/xc"
)

type Translator struct {
	rules             Rules
	prefixEnums       bool
	compiledRxs       map[RuleAction]RxMap
	compiledPtrTipRxs PtrTipRxMap
	compiledMemTipRxs MemTipRxList
	constRules        ConstRules
	typemap           CTypeMap
	fileScope         *cc.Bindings

	valueMap map[string]Value
	exprMap  map[string]string
	tagMap   map[string]*CDecl

	defines  []*CDecl
	typedefs []*CDecl
	declares []*CDecl

	typedefsSet  map[string]struct{}
	typedefKinds map[string]CTypeKind
	//	translateCache *SpecTranslateCache
	transformCache *NameTransformCache
	typeCache      *TypeCache

	ptrTipCache *TipCache
	memTipCache *TipCache
}

type RxMap map[RuleTarget][]Rx

type Rx struct {
	From *regexp.Regexp
	To   []byte
	//
	Transform RuleTransform
}

type PtrTipRxMap map[TipScope][]TipSpecRx
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

func (t TipSpecRx) Self() Tip {
	if t.self.IsValid() {
		return t.self
	}
	return t.Default
}

type Config struct {
	Rules      Rules      `yaml:"Rules"`
	ConstRules ConstRules `yaml:"ConstRules"`
	PtrTips    PtrTips    `yaml:"PtrTips"`
	MemTips    MemTips    `yaml:"MemTips"`
	Typemap    CTypeMap   `yaml:"Typemap"`
}

func New(cfg *Config) (*Translator, error) {
	if cfg == nil {
		cfg = &Config{}
	}
	t := &Translator{
		rules:             cfg.Rules,
		constRules:        cfg.ConstRules,
		typemap:           cfg.Typemap,
		compiledRxs:       make(map[RuleAction]RxMap),
		compiledPtrTipRxs: make(PtrTipRxMap),
		valueMap:          make(map[string]Value),
		exprMap:           make(map[string]string),
		tagMap:            make(map[string]*CDecl),
		typedefsSet:       make(map[string]struct{}),
		typedefKinds:      make(map[string]CTypeKind),
		transformCache:    &NameTransformCache{},
		typeCache:         &TypeCache{},
		ptrTipCache:       &TipCache{},
		memTipCache:       &TipCache{},
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
	if rxList, err := getMemTipRxs(cfg.MemTips); err != nil {
		return nil, err
	} else {
		t.compiledMemTipRxs = rxList
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
	if s[i].Pos != s[j].Pos {
		return s[i].Pos < s[j].Pos
	} else {
		return s[i].Name < s[j].Name
	}
}

func (t *Translator) Learn(unit *cc.TranslationUnit) {
	t.walkTranslationUnit(unit)
	t.resolveTypedefs(t.typedefs)
	sort.Sort(declList(t.declares))
	sort.Sort(declList(t.typedefs))
	t.collectDefines(unit.Macros)
	sort.Sort(declList(t.defines))
}

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

func (t *Translator) collectDefines(defines map[int]*cc.Macro) {
	seen := make(map[string]struct{}, len(defines))
	for _, macro := range defines {
		if macro.IsFnLike {
			continue
		}
		name := string(macro.DefTok.S())
		if !t.IsAcceptableName(TargetConst, name) {
			continue
		}
		seen[name] = struct{}{}

		expand := false
		if t.constRules[ConstDefines] == ConstExpand {
			expand = true
		}

		if !expand {
			switch val := macro.Value.(type) {
			case nil: // unresolved value -> try to expand
				expand = true
			case bool: // ban bools
				continue
			case cc.StringLitID:
				macro.Value = fmt.Sprintf(`"%s"`, xc.Dict.S(int(val)))
			}
			if !expand {
				t.defines = append(t.defines, &CDecl{
					IsDefine: true,
					Name:     name,
					Value:    Value(macro.Value),
					Pos:      macro.DefTok.Pos(),
				})
				continue
			}
		} else if _, ok := macro.Value.(bool); ok {
			// ban bools
			continue
		}
		tokens := macro.ReplacementToks()
		if len(tokens) == 0 {
			t.defines = append(t.defines, &CDecl{
				IsDefine: true,
				Name:     name,
				Value:    Value(macro.Value),
				Pos:      macro.DefTok.Pos(),
			})
			continue
		}

		srcParts := make([]string, 0, len(tokens))
		exprParts := make([]string, 0, len(tokens))
		valid := true

		// TODO: some state machine
		needsTypecast := false
		typecastValue := false
		typecastValueParens := 0

		for _, token := range tokens {
			src := cc.TokSrc(token)
			srcParts = append(srcParts, src)
			switch token.Rune {
			case cc.IDENTIFIER:
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
				case needsTypecast && token.Rune == rparen:
					typecastValue = true
					needsTypecast = false
					exprParts = append(exprParts, src+"(")
				case typecastValue && token.Rune == lparen:
					typecastValueParens++
				case typecastValue && token.Rune == rparen:
					if typecastValueParens == 0 {
						typecastValue = false
						exprParts = append(exprParts, ")"+src)
					} else {
						typecastValueParens--
					}
				default:
					if runes := []rune(src); len(runes) > 0 && isNumeric(runes) {
						// TODO(xlab): better const handling
						// should be resolved by switching to the upstream CC
						src = readNumeric(runes)
					}
					exprParts = append(exprParts, src)
				}
			}
			if !valid {
				break
			}
		}
		if !valid {
			continue
		}
		t.defines = append(t.defines, &CDecl{
			IsDefine:   true,
			Name:       name,
			Expression: strings.Join(exprParts, " "),
			Src:        strings.Join(srcParts, " "),
			Pos:        macro.DefTok.Pos(),
		})
	}
}

func (t *Translator) resolveTypedefs(typedefs []*CDecl) {
	for _, decl := range typedefs {
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

func (t *Translator) lookupSpec(spec CTypeSpec) (GoTypeSpec, bool) {
	if gospec, ok := t.typemap[spec]; ok {
		return gospec, true
	}
	if gospec, ok := builtinCTypeMap[spec]; ok {
		return gospec, true
	}
	if spec.Const {
		spec.Const = false
		if gospec, ok := t.typemap[spec]; ok {
			return gospec, true
		}
		if gospec, ok := builtinCTypeMap[spec]; ok {
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

func (t *Translator) MemTipRx(name string) (TipSpecRx, bool) {
	if rx, ok := t.memTipCache.Get(TipScopeStruct, name); ok {
		return rx, true
	}
	for _, rx := range t.compiledMemTipRxs {
		if rx.Target.MatchString(name) {
			t.memTipCache.Set(TipScopeStruct, name, rx)
			return rx, true
		}
	}
	return TipSpecRx{}, false
}

func (t *Translator) PtrMemTipRxForSpec(scope TipScope, name string, spec CType) (ptr, mem TipSpecRx) {
	var ptrOk, memOk bool
	if tag := spec.GetBase(); len(tag) > 0 {
		ptr, ptrOk = t.PtrTipRx(scope, tag)
		mem, memOk = t.MemTipRx(tag)
	}
	if !ptrOk {
		ptr, _ = t.PtrTipRx(scope, name)
	}
	if !memOk {
		mem, _ = t.MemTipRx(name)
	}
	return
}

func (t *Translator) TranslateSpec(spec CType, ptrTips ...Tip) GoTypeSpec {
	var ptrTip Tip
	if len(ptrTips) > 0 {
		ptrTip = ptrTips[0]
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
		if gospec, ok := t.lookupSpec(lookupSpec); ok {
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
			if t.IsAcceptableName(TargetType, typeSpec.Raw) {
				gospec.Raw = string(t.TransformName(TargetType, typeSpec.Raw))
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
				if gospec, ok := t.lookupSpec(lookupSpec); ok {
					gospec.Slices += wrapper.Slices
					gospec.Pointers += wrapper.Pointers
					gospec.OuterArr.Prepend(wrapper.OuterArr)
					gospec.InnerArr.Prepend(wrapper.InnerArr)
					//					gospec.Pointers += spec.GetVarArrays()
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
					if t.IsAcceptableName(TargetType, typeSpec.Raw) {
						gospec.Raw = string(t.TransformName(TargetType, typeSpec.Raw))
					}
					return gospec
				}
			}
		}
		//		wrapper.Pointers += spec.GetVarArrays()
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
			// pointers won't work with Go functions
			wrapper.Pointers = 0
			wrapper.Slices = 0
			wrapper.OuterArr = ArraySpec("")
			wrapper.InnerArr = ArraySpec("")
		}
		return wrapper
	case FunctionKind:
		wrapper := GoTypeSpec{
			Kind: spec.Kind(),
		}
		// won't work with Go functions anyways.
		// wrapper.splitPointers(ptrTip, spec.GetPointers())
		// wrapper.Pointers += spec.GetVarArrays()
		wrapper.Base = "func"
		return wrapper
	default:
		kind := spec.Kind()
		if kind == OpaqueStructKind {
			if decl, ok := t.tagMap[spec.GetTag()]; ok {
				if !decl.Spec.IsOpaque() {
					kind = StructKind
				}
			}
		}
		wrapper := GoTypeSpec{
			Kind:     kind,
			OuterArr: spec.OuterArrays(),
			InnerArr: spec.InnerArrays(),
		}
		wrapper.splitPointers(ptrTip, spec.GetPointers())
		//		wrapper.Pointers += spec.GetVarArrays()
		if base := spec.GetBase(); len(base) > 0 {
			wrapper.Raw = string(t.TransformName(TargetType, base))
		} else if cgoName := spec.CGoName(); len(cgoName) > 0 {
			wrapper.Raw = "C." + cgoName
		}
		return wrapper
	}
}

func (t *Translator) CGoSpec(spec CType) CGoSpec {
	cgo := CGoSpec{
		Pointers: spec.GetPointers(),
		OuterArr: spec.OuterArrays(),
		InnerArr: spec.InnerArrays(),
	}
	if typ, ok := spec.(*CTypeSpec); ok {
		if typ.Base == "void" {
			cgo.Base = "unsafe.Pointer"
			cgo.Pointers--
			return cgo
		}
	}
	if spec.Kind() == TypeKind && spec.IsOpaque() {
		cgo.Base = "byte"
		return cgo
	}
	cgo.Base = "C." + spec.CGoName()
	return cgo
}

func (t *Translator) registerTagsOf(decl *CDecl) {
	switch decl.Spec.Kind() {
	case EnumKind, StructKind, UnionKind:
		if tag := decl.Spec.GetTag(); len(tag) > 0 {
			if prev, ok := t.tagMap[tag]; !ok {
				// first time seen -> store the tag
				t.tagMap[tag] = decl
			} else if !prev.Spec.IsComplete() {
				// overwrite with a template declaration
				t.tagMap[tag] = decl
			}
		}
	}
	switch typ := decl.Spec.(type) {
	case *CStructSpec:
		for _, m := range typ.Members {
			if m.Spec.Kind() != StructKind {
				continue
			}
			if tag := m.Spec.GetTag(); len(tag) > 0 {
				if prev, ok := t.tagMap[tag]; !ok {
					// first time seen -> store the tag
					t.tagMap[tag] = m
				} else if !prev.Spec.IsComplete() {
					// overwrite with a template declaration
					t.tagMap[tag] = m
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
