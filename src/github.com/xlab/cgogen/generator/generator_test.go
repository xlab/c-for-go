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
		originFile    = "test/test.h"
		goHelpersFile = "test/test_helpers.go"
		cHelpersFile  = "test/test_helpers.c"
		resultFile    = "test/test.go"
		//
		goHelpersBuf = new(bytes.Buffer)
		cHelpersBuf  = new(bytes.Buffer)
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
		gen.MonitorAndWriteHelpers(goHelpersBuf, cHelpersBuf)

		if goHelpersBuf.Len() > 0 {
			buf := goHelpersBuf.Bytes()
			goHelpersFmt, err := imports.Process(goHelpersFile, buf, nil)
			if assert.NoError(err) {
				assert.NoError(ioutil.WriteFile(goHelpersFile, goHelpersFmt, 0644))
			} else {
				assert.NoError(ioutil.WriteFile(goHelpersFile, buf, 0644))
			}
		}
		if cHelpersBuf.Len() > 0 {
			assert.NoError(ioutil.WriteFile(cHelpersFile, cHelpersBuf.Bytes(), 0644))
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
	p, err := parser.New(&parser.Config{
		TargetPaths:  []string{originHeader},
		IncludePaths: []string{"/usr/include"},
	})
	if err != nil {
		return nil, err
	}

	unit, err := p.Parse()
	if err != nil {
		return nil, err
	}

	t, err := tl.New(&tl.Config{
		ConstRules: tl.ConstRules{
			tl.ConstEnum:    tl.ConstEvalFull,
			tl.ConstDeclare: tl.ConstExpand,
		},
		Rules: tl.Rules{
			tl.TargetGlobal: {
				tl.RuleSpec{From: "test_", Action: tl.ActionAccept},
				tl.RuleSpec{Transform: tl.TransformExport},
			},
			tl.TargetType: {
				tl.RuleSpec{From: "_t$", Action: tl.ActionReplace},
			},
			tl.TargetPrivate: {
				tl.RuleSpec{Transform: tl.TransformUnexport},
			},
			tl.TargetPostGlobal: {
				tl.RuleSpec{Load: "snakecase"},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if err := t.Learn(unit); err != nil {
		return nil, err
	}

	genCfg := &Config{
		PackageName: "test",
		Includes:    []string{filepath.Base(originHeader)},
	}
	return New(genCfg.PackageName, genCfg, t)
}
