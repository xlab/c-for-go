package translator

import (
	"bufio"
	"os"
	"testing"

	"github.com/xlab/cgogen/parser"
)

func TestLearn(t *testing.T) {
	pCfg := parser.NewConfig("/usr/local/Cellar/libvpx/1.4.0/include/vpx/vpx_encoder.h")
	pCfg.SysIncludePaths = []string{"/usr/include"}
	p, err := parser.New(pCfg)
	if err != nil {
		t.Fatal(err)
	}
	_, macros, err := p.Parse()
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
			RuleSpec{From: "vpx_", To: "_", Action: ActionReplace},
			RuleSpec{From: "_abi", Transform: TransformUpper},
			RuleSpec{From: "_img", To: "_image", Action: ActionReplace},
			RuleSpec{From: "_fmt", To: "_format", Action: ActionReplace},
			RuleSpec{From: "_([^_]+)", To: "$1", Action: ActionReplace, Transform: TransformTitle},
		},
	}
	tl, err := New(rules, buf)
	if err != nil {
		t.Fatal(err)
	}
	tl.Learn(macros)
}
