// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate.go
//go:generate golex -o trigraphs.go trigraphs.l
//go:generate golex -o scanner.go scanner.l
//go:generate stringer -type Kind
//go:generate stringer -type Namespace
//go:generate stringer -type ScalarType
//go:generate stringer -type Scope
//go:generate go run generate.go -2

// Package cc is a C compiler front end.
//
// Links
//
// Referenced from elsewhere:
//
//  [0]: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf
//  [1]: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1406.pdf
package cc

import (
	"fmt"
	"strings"

	"github.com/cznic/c/internal/xc"
	"github.com/cznic/golex/lex"
)

// Opt is a configuration/setup function.
type Opt func(lx *lexer)

// SysIncludePaths option configures where to search for system include files.
// (<name.h>)
func SysIncludePaths(paths []string) Opt {
	return func(*lexer) { sysIncludePaths = fromSlashes(paths) }
}

// IncludePaths option configures where to search for include files. ("name.h")
func IncludePaths(paths []string) Opt {
	return func(*lexer) { includePaths = fromSlashes(paths) }
}

// YyDebug sets the parser debug level.
func YyDebug(n int) Opt {
	return func(*lexer) { yyDebug = n }
}

// Cpp registers a cpp hook function which is called for every line the
// preprocessor produces before it is consumed by the parser.
func Cpp(f func([]xc.Token)) Opt {
	return func(lx *lexer) { lx.cpp = f }
}

// Parse clears any existing macros and define any macros in predefine. Then
// Parse preprocesses and parses the translation unit consisting of files in
// paths. The m communicates the scalar types model and opts allow to amend the
// parser. Parse is not reentrant nor safe to call concurrently.
func Parse(predefine string, paths []string, m Model, opts ...Opt) (*TranslationUnit, error) {
	fromSlashes(paths)
	compilation.ClearErrors()
	Macros = map[int]*macro{}
	if err := m.sanityCheck(); err != nil {
		compilation.Err(0, "%s", err.Error())
		return nil, compilation.Errors(true)
	}

	model = m
	lx := newTULexer()
	for _, opt := range opts {
		opt(lx)
	}
	if err := compilation.Errors(true); err != nil {
		return nil, err
	}

	lx.ch = make(chan []xc.Token, 1000)
	lx0 := newTULexer()
	ctx := newEvalCtx(lx0, lx.ch)

	go func() {
		defer close(lx.ch)

		predefines(predefine, lx.ch)
		for _, path := range paths {
			ppFile(xc.Token{}, path).preprocess(ctx)
		}
	}()

	if err := compilation.Errors(true); err != nil { // Do not parse if preprocessing failed.
		return nil, err
	}

	lx.state = lsTranslationUnit0
	r := yyParse(lx)
	if err := compilation.Errors(true); err != nil {
		return nil, err
	}

	if r != 0 {
		panic("internal error")
	}

	return lx.tu, nil
}

func parsePreprocessingFile(lx *lexer) *PreprocessingFile {
	lx.unget(lex.NewChar(0, PREPROCESSING_FILE))
	lx.state = lsPreprocess
	if n := yyParse(lx); n != 0 {
		return nil
	}

	return lx.ast
}

func exampleAST(rule int, src string) interface{} {
	isExample = true
	defer func() { isExample = false }()

	scanner := newUtf8src(strings.NewReader(src), fileset.AddFile(fmt.Sprintf("example%v.c", rule), -1, len(src)))
	lx := newLexer(scanner, false)

	defer lx.close()

	lx.exampleRule = rule
	yyParse(lx)
	return lx.example
}
