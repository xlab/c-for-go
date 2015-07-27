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
		TargetConst: {
			RuleSpec{From: "(?i)VPX_", Action: ActionAccept},
		},
	}
	tl, err := New(rules, buf)
	if err != nil {
		t.Fatal(err)
	}
	tl.Learn(macros)
}
