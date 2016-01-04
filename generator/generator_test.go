package generator

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xlab/cgogen/parser"
	tl "github.com/xlab/cgogen/translator"
	"golang.org/x/tools/imports"
)

func TestProxies(t *testing.T) {
	assert := assert.New(t)
	var (
		originFile    = "test/foo.h"
		goHelpersFile = "test/cgo_helpers.go"
		chHelpersFile = "test/cgo_helpers.h"
		ccHelpersFile = "test/cgo_helpers.c"
		resultFile    = "test/foo.go"
		//
		goHelpersBuf = new(bytes.Buffer)
		chHelpersBuf = new(bytes.Buffer)
		ccHelpersBuf = new(bytes.Buffer)
		buf          = new(bytes.Buffer)
	)
	// prepare to generate
	gen, err := getGenerator(originFile)
	if !assert.NoError(err) {
		return
	}
	var wg sync.WaitGroup
	defer func() {
		gen.Close()
		wg.Wait()
	}()

	go func() {
		wg.Add(1)
		gen.MonitorAndWriteHelpers(goHelpersBuf, chHelpersBuf, ccHelpersBuf)

		if goHelpersBuf.Len() > 0 {
			buf := goHelpersBuf.Bytes()
			goHelpersFmt, err := imports.Process(goHelpersFile, buf, nil)
			if assert.NoError(err) {
				assert.NoError(ioutil.WriteFile(goHelpersFile, goHelpersFmt, 0644))
			} else {
				assert.NoError(ioutil.WriteFile(goHelpersFile, buf, 0644))
			}
		}
		if chHelpersBuf.Len() > 0 {
			assert.NoError(ioutil.WriteFile(chHelpersFile, chHelpersBuf.Bytes(), 0644))
		}
		if ccHelpersBuf.Len() > 0 {
			assert.NoError(ioutil.WriteFile(ccHelpersFile, ccHelpersBuf.Bytes(), 0644))
		}

		wg.Done()
	}()

	gen.WritePackageHeader(buf)
	gen.WriteIncludes(buf)
	gen.WriteConst(buf)
	gen.WriteTypedefs(buf)
	gen.WriteDeclares(buf)
	fmtBuf, err := imports.Process(resultFile, buf.Bytes(), nil)
	if !assert.NoError(err) {
		fmtBuf = buf.Bytes()
	}
	assert.NoError(ioutil.WriteFile(resultFile, fmtBuf, 0644))
}

func getGenerator(originHeader string) (*Generator, error) {
	unit, defines, err := parser.ParseWith(&parser.Config{
		SourcesPaths: []string{originHeader},
		IncludePaths: []string{"/usr/include"},
	})
	if err != nil {
		return nil, err
	}

	t, err := tl.New(&tl.Config{
		ConstRules: tl.ConstRules{
			tl.ConstEnum: tl.ConstEvalFull,
			tl.ConstDecl: tl.ConstExpand,
		},
		PtrTips: tl.PtrTips{
			tl.TipScopeFunction: {
				{Target: "_message$", Tips: tl.Tips{
					tl.TipPtrRef, // object reference in message's functions
				}},
				{Target: "_bytes$", Self: tl.TipPtrArr},
			},
			tl.TipScopeStruct: {
				{Target: "_attachment_t$", Tips: tl.Tips{
					tl.TipPtrRef, // attachment's data raw reference
				}},
			},
		},
		MemTips: tl.MemTips{
			{Target: "_attachment_t$", Self: tl.TipMemRaw},
		},
		Rules: tl.Rules{
			tl.TargetGlobal: {
				tl.RuleSpec{From: "(?i)foo_", Action: tl.ActionAccept},
				tl.RuleSpec{Transform: tl.TransformLower},
				tl.RuleSpec{From: "foo_", To: "_", Action: tl.ActionReplace},
				tl.RuleSpec{Transform: tl.TransformExport},
			},
			tl.TargetType: {
				tl.RuleSpec{From: "_t$", Action: tl.ActionReplace},
			},
			tl.TargetPrivate: {
				tl.RuleSpec{Transform: tl.TransformUnexport},
			},
			tl.TargetPostGlobal: {
				tl.RuleSpec{From: "_id?|$", Transform: tl.TransformUpper},
				tl.RuleSpec{Load: "snakecase"},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	t.Learn(unit, defines)
	// t.Report()

	genCfg := &Config{
		PackageName: "foo",
		Includes:    []string{filepath.Base(originHeader)},
	}
	return New(genCfg.PackageName, genCfg, t)
}
