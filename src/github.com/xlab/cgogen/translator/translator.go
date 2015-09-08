package translator

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/cznic/c/internal/cc"
	"github.com/cznic/c/internal/xc"
)

type Translator struct {
	rules             Rules
	prefixEnums       bool
	compiledRxs       map[RuleAction]RxMap
	compiledPtrTipRxs PtrTipRxMap
	compiledMemTipRxs MemTipRxList
	constRules        ConstRules
	typemap           CTypeMap

	valueMap map[string]Value
	exprMap  map[string]Expression
	tagMap   map[string]CDecl

	defines  []CDecl
	typedefs []CDecl
	declares []CDecl

	typedefsSet    map[string]struct{}
	transformCache *NameTransformCache
	ptrTipCache    *TipCache
	memTipCache    *TipCache
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
		exprMap:           make(map[string]Expression),
		tagMap:            make(map[string]CDecl),
		typedefsSet:       make(map[string]struct{}),
		transformCache:    &NameTransformCache{},
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
					spec = s
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

type declList []CDecl

func (s declList) Len() int      { return len(s) }
func (s declList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s declList) Less(i, j int) bool {
	if s[i].Pos != s[j].Pos {
		return s[i].Pos < s[j].Pos
	} else {
		return s[i].Name < s[j].Name
	}
}

func (t *Translator) Learn(unit *cc.TranslationUnit) error {
	for id := range cc.Macros {
		name := string(xc.Dict.S(id))
		if !t.IsAcceptableName(TargetConst, name) {
			continue
		}
		pos, tokList, uTokList, ok := cc.ExpandDefine(id)
		if !ok || len(tokList) == 0 {
			continue
		}
		srcParts := make([]string, len(uTokList))
		for i, v := range uTokList {
			srcParts[i] = cc.TokSrc(v)
		}
		t.defines = append(t.defines, CDecl{
			Pos:        pos,
			IsDefine:   true,
			Name:       name,
			Expression: tokList[0].S(),
			Src:        strings.Join(srcParts, " "),
		})
	}

	sort.Sort(declList(t.defines))
	for unit != nil {
		unit = t.walkTranslationUnit(unit)
	}
	sort.Sort(declList(t.declares))
	sort.Sort(declList(t.typedefs))
	return xc.Compilation.Errors(true)
}

func (t *Translator) TransformName(target RuleTarget, str string) []byte {
	if len(str) == 0 {
		return emptyStr
	}
	if name, ok := t.transformCache.Get(target, str); ok {
		return name
	}

	var name []byte
	if target != TargetGlobal && target != TargetPostGlobal {
		// apply global rules first
		name = t.TransformName(TargetGlobal, str)
	} else {
		name = []byte(str)
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
	if target != TargetGlobal && target != TargetPostGlobal {
		// apply post-global rules in the end
		name = t.TransformName(TargetPostGlobal, string(name))
	}
	t.transformCache.Set(target, str, name)
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
		spec := spec.(*CTypeSpec)
		lookupSpec := CTypeSpec{
			Base:     spec.Base,
			Const:    spec.Const,
			Unsigned: spec.Unsigned,
			Short:    spec.Short,
			Long:     spec.Long,
			// Arrays skip
			// VarArrays skip
			Pointers: spec.Pointers,
		}
		if gospec, ok := t.lookupSpec(lookupSpec); ok {
			gospec.Arrays = spec.GetArrays()
			if ptrTip == TipPtrRef {
				if gospec.Pointers == 0 && gospec.Slices > 0 {
					gospec.Slices--
					gospec.Pointers++
				}
			}
			return gospec
		}
		wrapper := GoTypeSpec{
			Arrays: spec.GetArrays(),
		}
		if lookupSpec.Pointers > 0 {
			for lookupSpec.Pointers > 0 {
				if ptrTip == TipPtrRef {
					if lookupSpec.Pointers > 1 {
						wrapper.Slices++
					} else {
						wrapper.Pointers++
					}
				} else {
					wrapper.Slices++
				}

				lookupSpec.Pointers--
				if gospec, ok := t.lookupSpec(lookupSpec); ok {
					gospec.Slices += wrapper.Slices
					gospec.Pointers += wrapper.Pointers
					gospec.Arrays = wrapper.Arrays + gospec.Arrays
					gospec.Pointers += spec.GetVarArrays()
					// if !ptrTip.IsValid() {
					// 	ptrTip = TipPtrArr
					// }
					// log.Println(spec, "->", lookupSpec, "(found) wrapped", wrapper,
					// 	"ptrTip:", ptrTip, "-> return:", gospec)
					return gospec
				}
			}
		}
		wrapper.Pointers += spec.GetVarArrays()
		wrapper.Base = string(t.TransformName(TargetType, lookupSpec.Base))
		return wrapper
	case FunctionKind:
		wrapper := GoTypeSpec{
			Arrays: spec.GetArrays(),
		}
		wrapper.splitPointers(ptrTip, spec.GetPointers())
		wrapper.Pointers += spec.GetVarArrays()
		wrapper.Base = "func"
		return wrapper
	default:
		wrapper := GoTypeSpec{
			Arrays: spec.GetArrays(),
		}
		wrapper.splitPointers(ptrTip, spec.GetPointers())
		wrapper.Pointers += spec.GetVarArrays()
		if fallback, ok := t.IsBaseDefined(spec); ok && t.IsAcceptableName(TargetType, spec.GetBase()) {
			wrapper.Base = string(t.TransformName(TargetType, spec.GetBase()))
		} else {
			wrapper.Base = fallback
		}
		return wrapper
	}
}

func (t *Translator) CGoSpec(spec CType) CGoSpec {
	cgo := CGoSpec{
		Pointers: spec.GetPointers(),
	}
	cgo.Pointers += spec.GetVarArrays()
	cgo.Arrays = GetArraySizes(spec.GetArrays())

	if spec, ok := spec.(*CTypeSpec); ok {
		switch base := spec.Base; base {
		case "int", "short", "long", "char":
			cgo.Base = "C."
			if spec.Unsigned {
				cgo.Base += "u"
			} else if spec.Signed {
				cgo.Base += "s"
			}
			switch {
			case spec.Long:
				cgo.Base += "long"
				if spec.Base == "long" {
					cgo.Base += "long"
				}
			case spec.Short:
				cgo.Base += "short"
			default:
				cgo.Base += spec.Base
			}
			return cgo
		case "void":
			return cgo
		default:
			cgo.Base = "C." + spec.Base
			return cgo
		}
	}
	cgo.Base = t.cgoName(spec)
	return cgo
}

func (t *Translator) cgoName(spec CType) string {
	name := spec.GetBase()
	if _, ok := t.tagMap[name]; ok {
		switch spec.Kind() {
		case StructKind:
			return "C.struct_" + name
		case UnionKind:
			return "C.union_" + name
		case EnumKind:
			return "C.enum_" + name
		}
	}
	if len(name) > 0 {
		return "C." + name
	}
	return ""
}

func (t *Translator) IsBaseDefined(spec CType) (fallback string, ok bool) {
	name := spec.GetBase()
	switch spec.Kind() {
	case StructKind:
		fallback = "C.struct_" + name
	case UnionKind:
		fallback = "C.union_" + name
	case EnumKind:
		fallback = "C.enum_" + name
	default:
		fallback = "C." + name
	}
	decl, ok := t.tagMap[name]
	if ok && decl.IsTemplate() {
		return
	}
	if _, ok = t.typedefsSet[name]; ok {
		return
	}
	return
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

func (t *Translator) TagMap() map[string]CDecl {
	return t.tagMap
}

func (t *Translator) ExpressionMap() map[string]Expression {
	return t.exprMap
}

func (t *Translator) ValueMap() map[string]Value {
	return t.valueMap
}

func (t *Translator) Defines() []CDecl {
	return t.defines
}

func (t *Translator) Declares() []CDecl {
	return t.declares
}

func (t *Translator) Typedefs() []CDecl {
	return t.typedefs
}
