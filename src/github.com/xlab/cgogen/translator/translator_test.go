package translator

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/xlab/cgogen/parser"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestWebInclude(t *testing.T) {
	return
	pCfg := parser.NewConfig("test/web_include_test.h")
	pCfg.SysIncludePaths = []string{"/usr/include"}
	pCfg.WebIncludesEnabled = true
	p, err := parser.New(pCfg)
	if err != nil {
		t.Fatal(err)
	}
	unit, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()
	cfg := &Config{
		ConstRules: ConstRules{
			ConstEnum:    ConstEvalFull,
			ConstDeclare: ConstExpand,
		},
	}
	tl, err := New(cfg, buf)
	if err != nil {
		t.Fatal(err)
	}
	if err := tl.Learn(unit); err != nil {
		t.Fatal(err)
	}

	// tl.Report()
}

func TestLearn(t *testing.T) {
	pCfg := parser.NewConfig("test/translator_test.h")
	pCfg.SysIncludePaths = []string{"/usr/include"}
	p, err := parser.New(pCfg)
	if err != nil {
		t.Fatal(err)
	}
	unit, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()

	cfg := &Config{
		Rules: Rules{
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
		},
		ConstRules: ConstRules{
			ConstEnum:    ConstEvalFull,
			ConstDeclare: ConstExpand,
		},
	}

	tl, err := New(cfg, buf)
	if err != nil {
		t.Fatal(err)
	}
	if err := tl.Learn(unit); err != nil {
		t.Fatal(err)
	}

	// tl.Report()
}
