package translator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xlab/cgogen/parser"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestLearn(t *testing.T) {
	cfg := parser.NewConfig("test/translator_test.h")
	cfg.IncludePaths = []string{"/usr/local/include", "/usr/include"}

	unit, macros, err := parser.ParseWith(cfg)
	if err != nil {
		t.Fatal(err)
	}
	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()
	rules := Rules{
		TargetGlobal: {
			RuleSpec{From: "(?i)VPX_", Action: ActionAccept},
			RuleSpec{Transform: TransformLower},
		},
		TargetConst: {
			RuleSpec{From: "_INLINE$", Action: ActionIgnore},
			RuleSpec{From: "vpx_", To: "_", Action: ActionReplace},
			RuleSpec{From: "_abi", Transform: TransformUpper},
			RuleSpec{From: "_img", To: "_image", Action: ActionReplace},
			RuleSpec{From: "_fmt", To: "_format", Action: ActionReplace},
			RuleSpec{From: "_([^_]+)", To: "$1", Action: ActionReplace, Transform: TransformTitle},
		},
	}
	constRules := ConstRules{
		ConstEnum: ConstEvalFull,
		ConstDecl: ConstExpand,
	}
	tl, err := New(&Config{
		Rules:      rules,
		ConstRules: constRules,
	})
	if err != nil {
		t.Fatal(err)
	}
	tl.Learn(unit, macros)
	//	tl.Report()
}

func ss(n []uint64) []ArraySizeSpec {
	ss := make([]ArraySizeSpec, 0, len(n))
	for i := range n {
		ss = append(ss, ArraySizeSpec{N: n[i]})
	}
	return ss
}

func TestCGoSpecAtLevel(t *testing.T) {
	assert := assert.New(t)
	tbl := []struct {
		Spec     CGoSpec
		Expected []string
	}{
		{CGoSpec{},
			[]string{"", "", ""}},
		{CGoSpec{Base: "char"},
			[]string{"char", "char", "char"}},
		{CGoSpec{Base: "char", Pointers: 1},
			[]string{"*char", "char", "char"}},
		{CGoSpec{Base: "char", Pointers: 2},
			[]string{"**char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: ss([]uint64{1})},
			[]string{"*char", "char", "char"}},
		{CGoSpec{Base: "char", Arrays: ss([]uint64{1, 2})},
			[]string{"*[2]char", "[2]char", "char"}},
		{CGoSpec{Base: "char", Arrays: ss([]uint64{1, 2, 3})},
			[]string{"*[2][3]char", "[2][3]char", "[3]char", "char"}},
		{CGoSpec{Base: "char", Arrays: ss([]uint64{1, 2, 3}), Pointers: 1},
			[]string{"*[2][3]*char", "[2][3]*char", "[3]*char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: ss([]uint64{1, 2, 3}), Pointers: 2},
			[]string{"*[2][3]**char", "[2][3]**char", "[3]**char", "**char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: ss([]uint64{1, 2}), Pointers: 3},
			[]string{"*[2]***char", "[2]***char", "***char", "**char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: ss([]uint64{1, 2}), Pointers: 1},
			[]string{"*[2]*char", "[2]*char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: ss([]uint64{1}), Pointers: 2},
			[]string{"***char", "**char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: ss([]uint64{1}), Pointers: 1},
			[]string{"**char", "*char", "char"}},
	}
	for _, test := range tbl {
		for level, exp := range test.Expected {
			assert.Equal(exp, test.Spec.AtLevel(uint8(level)), fmt.Sprintf("at level %d", level))
		}
	}
}
