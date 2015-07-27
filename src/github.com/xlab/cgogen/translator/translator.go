package translator

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"

	"go/token"

	"github.com/cznic/c/internal/cc"
	"github.com/cznic/c/internal/xc"
)

type Translator struct {
	out       io.Writer
	rules     Rules
	acceptRxs RxMap
	ignoreRxs RxMap
}

type RxMap map[RuleTarget][]*regexp.Regexp

func New(rules Rules, out io.Writer) (*Translator, error) {
	t := &Translator{
		rules: rules,
		out:   out,
	}
	if rxMap, err := getRuleActionRxs(rules, ActionAccept); err != nil {
		return nil, err
	} else {
		t.acceptRxs = rxMap
	}
	if rxMap, err := getRuleActionRxs(rules, ActionIgnore); err != nil {
		return nil, err
	} else {
		t.ignoreRxs = rxMap
	}
	return t, nil
}

func getRuleActionRxs(rules Rules, action RuleAction) (RxMap, error) {
	rxMap := make(RxMap, len(rules))
	for target, specs := range rules {
		for _, spec := range specs {
			if spec.Action != action {
				continue
			}
			rx, err := regexp.Compile(spec.From)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("translator: %s rules: invalid regexp %s", target, spec.From))
			}
			rxMap[target] = append(rxMap[target], rx)
		}
	}
	return rxMap, nil
}

type definePosList []definePos
type definePos struct {
	Pos  token.Pos
	Name []byte
}

func (s definePosList) Len() int      { return len(s) }
func (s definePosList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s definePosList) Less(i, j int) bool {
	if s[i].Pos != s[j].Pos {
		return s[i].Pos < s[j].Pos
	} else {
		return bytes.Compare(s[i].Name, s[j].Name) < 0
	}
}

func (t *Translator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(t.out, format, args...)
}

func (t *Translator) Learn(macros []int) {
	var defList definePosList
	srcLineByID := make(map[int]string)
	defineByID := make(map[int]int)

	for _, id := range macros {
		name := xc.Dict.S(id)
		if !t.isAcceptableName(TargetConst, name) {
			continue
		}
		pos, tokList, uTokList, ok := cc.ExpandDefine(id)
		if !ok || len(tokList) == 0 {
			continue
		}

		defineByID[id] = tokList[0].Val
		defList = append(defList, definePos{
			Pos:  pos,
			Name: name,
		})

		srcParts := make([]string, len(uTokList))
		for i, v := range uTokList {
			srcParts[i] = cc.TokSrc(v)
		}
		srcLineByID[id] = strings.Join(srcParts, " ")
	}

	sort.Sort(definePosList(defList))

	t.Printf("const (")
	for _, def := range defList {
		id := xc.Dict.ID(def.Name)
		pos := xc.FileSet.Position(def.Pos)
		t.Printf("\n// %s:%d:%d\n//   > define %s %v\n%s = %s",
			narrowPath(pos.Filename), pos.Line, pos.Offset, def.Name, srcLineByID[id],
			def.Name, xc.Dict.S(defineByID[id]))
	}
	t.Printf("\n)\n\n")
}

func (t *Translator) transformName(target RuleTarget, name []byte) []byte {
	for _, rule := range t.rules[target] {
		switch rule.Action {
		// case ActionCut:
		// TODO: rx precompile
		}
	}
}

func (t *Translator) isAcceptableName(target RuleTarget, name []byte) bool {
	if rxs, ok := t.ignoreRxs[target]; ok {
		for _, rx := range rxs {
			if rx.Match(name) {
				return false
			}
		}
	}
	if rxs, ok := t.acceptRxs[target]; ok {
		for _, rx := range rxs {
			if rx.Match(name) {
				return true
			}
		}
	}
	return false
}

func (t *Translator) Translate(unit *cc.TranslationUnit, macros []string) {

}
