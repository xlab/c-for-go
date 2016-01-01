// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate.go
//go:generate golex -o trigraphs.go trigraphs.l
//go:generate golex -o scanner.go scanner.l
//go:generate stringer -type Kind
//go:generate stringer -type Linkage
//go:generate stringer -type Namespace
//go:generate stringer -type Scope
//go:generate go run generate.go -2

// Package cc is a C99 compiler front end.
//
// Links
//
// Referenced from elsewhere:
//
//  [0]: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf
//  [1]: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1406.pdf
//  [2]: https://github.com/rsc/c2go/blob/fc8cbfad5a47373828c81c7a56cccab8b221d310/cc/cc.y
package cc

import (
	"bufio"
	"bytes"
	"fmt"
	"go/token"
	"os"

	"github.com/cznic/golex/lex"
	"github.com/cznic/mathutil"
	"github.com/xlab/c/xc"
)

type tweaks struct {
	enableDefineOmitCommaBeforeDDD bool // #define foo(a, b...)
	enableDlrInIdentifiers         bool // foo$bar
	enableEmptyDefine              bool // #define
	enableTrigraphs                bool // ??=define foo(bar)
	enableUndefExtraTokens         bool // #undef foo(bar)
	enableWarnings                 bool // #warning
}

func exampleAST(rule int, src string) interface{} {
	report := xc.NewReport()
	report.IgnoreErrors = true
	lx, err := newLexer(
		fmt.Sprintf("example%v.c", rule),
		len(src)+1, // Plus final injected NL
		bytes.NewBufferString(src),
		report,
		&tweaks{},
	)
	lx.model = &Model{ // 64 bit
		Items: map[Kind]ModelItem{
			Ptr:               {8, 8, nil},
			Void:              {0, 1, nil},
			Char:              {1, 1, nil},
			SChar:             {1, 1, nil},
			UChar:             {1, 1, nil},
			Short:             {2, 2, nil},
			UShort:            {2, 2, nil},
			Int:               {4, 4, nil},
			UInt:              {4, 4, nil},
			Long:              {8, 8, nil},
			ULong:             {8, 8, nil},
			LongLong:          {8, 8, nil},
			ULongLong:         {8, 8, nil},
			Float:             {4, 4, nil},
			Double:            {8, 8, nil},
			LongDouble:        {8, 8, nil},
			Bool:              {1, 1, nil},
			FloatComplex:      {8, 8, nil},
			DoubleComplex:     {16, 16, nil},
			LongDoubleComplex: {16, 16, nil},
		},
		lx:     lx,
		report: report,
	}

	lx.model.init()
	if err != nil {
		panic(err)
	}

	lx.exampleRule = rule
	yyParse(lx)
	return lx.example
}

func ppParseString(fn, src string, report *xc.Report, tweaks *tweaks) (*PreprocessingFile, error) {
	sz := len(src)
	lx, err := newLexer(fn, int(sz)+1, bytes.NewBufferString(src), report, tweaks)
	if err != nil {
		return nil, err
	}

	lx.Unget(lex.NewChar(token.Pos(lx.File.Base()), PREPROCESSING_FILE))
	yyParse(lx)
	return lx.preprocessingFile, nil
}

func ppParse(fn string, report *xc.Report, tweaks *tweaks) (*PreprocessingFile, error) {
	o := xc.Files.Once(fn, func() interface{} {
		f, err := os.Open(fn)
		if err != nil {
			return err
		}

		defer f.Close()

		fi, err := os.Stat(fn)
		if err != nil {
			return err
		}

		sz := fi.Size()
		if sz > mathutil.MaxInt-1 {
			return fmt.Errorf("%s: file size too big: %v", fn, sz)
		}

		lx, err := newLexer(fn, int(sz)+1, bufio.NewReader(f), report, tweaks)
		if err != nil {
			return err
		}

		lx.Unget(lex.NewChar(token.Pos(lx.File.Base()), PREPROCESSING_FILE))
		yyParse(lx)
		return lx.preprocessingFile
	})
	switch r := o.Value(); x := r.(type) {
	case error:
		return nil, x
	case *PreprocessingFile:
		return x, nil
	default:
		panic("internal error")
	}
}

// Opt is a configuration/setup function that can be passed to the Parser
// function.
type Opt func(*lexer)

// EnableDefineOmitCommaBeforeDDD makes the parser accept non standard
//
//	#define foo(a, b...)
func EnableDefineOmitCommaBeforeDDD() Opt {
	return func(l *lexer) { l.tweaks.enableDefineOmitCommaBeforeDDD = true }
}

// EnableDlrInIdentifiers makes the parser accept non standard
//
//	int foo$bar
func EnableDlrInIdentifiers() Opt {
	return func(l *lexer) { l.tweaks.enableDlrInIdentifiers = true }
}

// EnableEmptyDefine makes the parser accept non standard
//
//	#define
func EnableEmptyDefine() Opt {
	return func(l *lexer) { l.tweaks.enableEmptyDefine = true }
}

// EnableUndefExtraTokens makes the parser accept non standard
//
//	#undef foo(bar)
func EnableUndefExtraTokens() Opt {
	return func(l *lexer) { l.tweaks.enableUndefExtraTokens = true }
}

// SysIncludePaths option configures where to search for system include files.
// (<name.h>)
func SysIncludePaths(paths []string) Opt {
	return func(l *lexer) { l.sysIncludePaths = fromSlashes(paths) }
}

// IncludePaths option configures where to search for include files. ("name.h")
func IncludePaths(paths []string) Opt {
	return func(l *lexer) { l.includePaths = fromSlashes(paths) }
}

// YyDebug sets the parser debug level.
func YyDebug(n int) Opt {
	return func(*lexer) { yyDebug = n }
}

// Cpp registers a preprocessor hook function which is called for every line,
// or group of lines the preprocessor produces before it is consumed by the
// parser. The token slice must not be modified by the hook.
func Cpp(f func([]xc.Token)) Opt {
	return func(lx *lexer) { lx.cpp = f }
}

// ErrLimit limits the number of calls to the error reporting methods.  After
// the limit is reached, all errors are reported using log.Print and then
// log.Fatal() is called with a message about too many errors.  To disable
// error limit, set ErrLimit to value less or equal zero.  Default value is 10.
func ErrLimit(n int) Opt {
	return func(lx *lexer) { lx.report.ErrLimit = n }
}

// Trigraphs enables processing of trigraphs.
func Trigraphs() Opt { return func(lx *lexer) { lx.tweaks.enableTrigraphs = true } }

func crashOnError() Opt { return func(lx *lexer) { lx.report.PanicOnError = true } }
func nopOpt() Opt       { return func(*lexer) {} }

type DefinesMap map[string][]xc.Token

// Parse defines any macros in predefine. Then Parse preprocesses and parses
// the translation unit consisting of files in paths. The m communicates the
// scalar types model and opts allow to amend parser behavior.
func Parse(predefine string, paths []string, m *Model, opts ...Opt) (*TranslationUnit, DefinesMap, error) {
	if m == nil {
		return nil, nil, fmt.Errorf("invalid nil model passed")
	}

	fromSlashes(paths)
	report := xc.NewReport()
	lx0 := &lexer{tweaks: &tweaks{enableWarnings: true}, report: report}
	for _, opt := range opts {
		opt(lx0)
	}
	if err := report.Errors(true); err != nil {
		return nil, nil, err
	}

	m.lx = lx0
	m.init()
	m.report = report
	if err := m.sanityCheck(); err != nil {
		report.Err(0, "%s", err.Error())
		return nil, nil, report.Errors(true)
	}

	tweaks := lx0.tweaks
	predefined, err := ppParseString("", predefine, report, tweaks)
	if err != nil {
		return nil, nil, err
	}

	ch := make(chan []xc.Token, 1000)
	macros := newMacros()
	stop := make(chan int, 1)
	go func() {
		defer close(ch)

		newPP(ch, lx0.includePaths, lx0.sysIncludePaths, macros, false, m, tweaks).preprocessingFile(predefined)
		for _, path := range paths {
			select {
			case <-stop:
				return
			default:
			}
			pf, err := ppParse(path, report, tweaks)
			if err != nil {
				report.Err(0, err.Error())
				return
			}

			newPP(ch, lx0.includePaths, lx0.sysIncludePaths, macros, true, m, tweaks).preprocessingFile(pf)
		}
	}()

	if err := report.Errors(true); err != nil { // Do not parse if preprocessing already failed.
		go func() {
			for range ch { // Drain.
			}
		}()
		stop <- 1
		return nil, nil, err
	}

	lx := newSimpleLexer(lx0.cpp, report, tweaks)
	lx.ch = ch
	lx.state = lsTranslationUnit0
	lx.model = m
	m.lx = lx
	yyParse(lx)

	defines := make(DefinesMap, len(macros.m))
	for id, macro := range macros.m {
		if !macro.isFnLike {
			name := xc.Dict.S(id)
			toks := decodeTokens(macro.repl, nil)
			defines[string(name)] = toks
		}
	}
	return lx.translationUnit, defines, report.Errors(true)
}
