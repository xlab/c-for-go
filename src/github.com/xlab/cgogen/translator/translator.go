package translator

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"

	"github.com/cznic/c/internal/cc"
	"github.com/cznic/c/internal/xc"
)

type Translator struct {
	out         io.Writer
	rules       Rules
	compiledRxs map[RuleAction]RxMap
	constRules  ConstRules
	typemap     CTypeMap

	valueMap     map[string]Value
	exprMap      map[string]Expression
	tagMap       map[string]CDecl
	defines      []CDecl
	typedefs     []CDecl
	declarations []CDecl
}

type RxMap map[RuleTarget][]Rx

type Rx struct {
	From *regexp.Regexp
	To   []byte
	//
	Transform RuleTransform
}

type Config struct {
	Rules      Rules
	ConstRules ConstRules
	Typemap    CTypeMap
}

func New(cfg *Config, out io.Writer) (*Translator, error) {
	if cfg == nil {
		cfg = &Config{}
	}
	t := &Translator{
		out:         out,
		rules:       cfg.Rules,
		constRules:  cfg.ConstRules,
		typemap:     cfg.Typemap,
		compiledRxs: make(map[RuleAction]RxMap),
		valueMap:    make(map[string]Value),
		exprMap:     make(map[string]Expression),
		tagMap:      make(map[string]CDecl),
	}
	for _, action := range ruleActions {
		if rxMap, err := getRuleActionRxs(t.rules, action); err != nil {
			return nil, err
		} else {
			t.compiledRxs[action] = rxMap
		}
	}
	return t, nil
}

func getRuleActionRxs(rules Rules, action RuleAction) (RxMap, error) {
	rxMap := make(RxMap, len(rules))
	for target, specs := range rules {
		for _, spec := range specs {
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
				return nil, errors.New(fmt.Sprintf("translator: %s rules: invalid regexp %s", target, spec.From))
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

func (t *Translator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(t.out, format, args...)
}

func (t *Translator) Learn(unit *cc.TranslationUnit) error {
	for id := range cc.Macros {
		name := xc.Dict.S(id)
		if !t.IsAcceptableName(TargetDefine, name) {
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
			Name:       string(name),
			Expression: tokList[0].S(),
			Src:        strings.Join(srcParts, " "),
		})
	}

	sort.Sort(declList(t.defines))
	for unit != nil {
		unit = t.walkTranslationUnit(unit)
	}
	sort.Sort(declList(t.declarations))
	sort.Sort(declList(t.typedefs))
	return xc.Compilation.Errors(true)
}

func (t *Translator) Report() {
	t.Printf("[!] TAGS:\n")
	for tag, decl := range t.tagMap {
		t.Printf("%s refers to %v\n", tag, decl)
	}

	t.Printf("\n\n\n[!] TYPEDEFs:\n")
	for _, decl := range t.typedefs {
		t.Printf("%v\n", decl)
	}

	t.Printf("\n\n\n[!] DECLARATIONS:\n")
	for _, decl := range t.declarations {
		t.Printf("%v\n", decl)
	}

	t.Printf("\n\n\n[!] const (")
	for _, line := range t.defines {
		t.Printf("\n// %s\n//   > define %s %v\n%s = %s",
			srcLocation(line.Pos), line.Name, line.Src,
			t.TransformName(TargetDefine, line.Name), line.Expression)
	}
	t.Printf("\n)\n\n")
}

func (t *Translator) TransformName(target RuleTarget, str string) []byte {
	var name []byte
	if target != TargetGlobal {
		// apply global rules first
		name = t.TransformName(TargetGlobal, str)
	} else {
		name = []byte(str)
	}

	for _, rx := range t.compiledRxs[ActionReplace][target] {
		indices := rx.From.FindAllSubmatchIndex(name, -1)
		reference := make([]byte, 0, len(name))
		reference = append(reference, name...)

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
			// if len(rx.Transform) > 0 {
			// 	log.Println("doing", rx.Transform, "at", string(name[idx[0]:idx[1]]), "in", string(name))
			// }
			buf := rx.From.Expand([]byte{}, rx.To, reference, idx)
			switch rx.Transform {
			case TransformLower:
				buf = bytes.ToLower(buf)
			case TransformTitle:
				buf = bytes.Title(buf)
			case TransformUpper:
				buf = bytes.ToUpper(buf)
			}
			name = replaceBytes(name, idx, buf)
		}
	}

	return name
}

func (t *Translator) TranslateSpec(spec CTypeSpec) (GoTypeSpec, error) {
	if gospec, ok := t.typemap[spec]; ok {
		return gospec, nil
	}
	if gospec, ok := builtinCTypeMap[spec]; ok {
		return gospec, nil
	}
	return GoTypeSpec{}, errors.New("no explicit rule found")
}

func (t *Translator) IsAcceptableName(target RuleTarget, name []byte) bool {
	if rxs, ok := t.compiledRxs[ActionIgnore][target]; ok {
		for _, rx := range rxs {
			if rx.From.Match(name) {
				return false
			}
		}
	}
	if rxs, ok := t.compiledRxs[ActionAccept][target]; ok {
		for _, rx := range rxs {
			if rx.From.Match(name) {
				return true
			}
		}
	}
	if target != TargetGlobal {
		// fallback to global rules
		return t.IsAcceptableName(TargetGlobal, name)
	}
	return false
}

func (t *Translator) Translate() {
	// TODO
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

func (t *Translator) Declarations() []CDecl {
	return t.declarations
}

func (t *Translator) Typedefs() []CDecl {
	return t.typedefs
}
