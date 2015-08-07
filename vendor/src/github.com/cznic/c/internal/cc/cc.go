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
	"errors"
	"fmt"
	"net/url"

	"github.com/cznic/c/internal/xc"
	"github.com/cznic/mathutil"
)

// Warning: this file has been altered by xlab.

type ParseConfig struct {
	Predefined         string
	Input              []byte
	Paths              []string
	SysIncludePaths    []string
	IncludePaths       []string
	WebIncludesEnabled bool
	WebIncludePrefix   string
	YyDebugLevel       int
	OnPreprocessLine   func([]xc.Token)
	SizeModel          Model
}

func CheckParseConfig(cfg *ParseConfig) error {
	if cfg == nil {
		return errors.New("no config specified")
	}
	if len(cfg.Paths) == 0 && len(cfg.Input) == 0 {
		return errors.New("no paths or input specified to parse")
	}
	if likelyIsURL(cfg.WebIncludePrefix) {
		if u, err := url.Parse(cfg.WebIncludePrefix); err == nil {
			cfg.WebIncludePrefix = u.String()
		} else {
			return fmt.Errorf("web include prefix specified but doesn't represent a valid URL: %v", err)
		}
	}
	fromSlashes(cfg.Paths)
	fromSlashes(cfg.IncludePaths)
	fromSlashes(cfg.SysIncludePaths)

	if cfg.SizeModel == nil {
		// some random 64-bit model taken from ccgo
		cfg.SizeModel = Model{
			Ptr:       {Size: 8, Align: 8, More: "__TODO_PTR"},
			Void:      {Size: 0, Align: 1, More: "__TODO_VOID"},
			Char:      {Size: 1, Align: 1, More: "int8"},
			UChar:     {Size: 1, Align: 1, More: "byte"},
			Short:     {Size: 2, Align: 2, More: "int16"},
			UShort:    {Size: 2, Align: 2, More: "uint16"},
			Int:       {Size: 4, Align: 4, More: "int32"},
			UInt:      {Size: 4, Align: 4, More: "uint32"},
			Long:      {Size: 8, Align: 8, More: "int64"},
			ULong:     {Size: 8, Align: 8, More: "uint64"},
			LongLong:  {Size: 8, Align: 8, More: "int64"},
			ULongLong: {Size: 8, Align: 8, More: "uint64"},
			Float:     {Size: 4, Align: 4, More: "float32"},
			Double:    {Size: 8, Align: 8, More: "float64"},
			Bool:      {Size: 1, Align: 1, More: "bool"},
			Complex:   {Size: 8, Align: 8, More: "complex128"},
		}
	}
	return nil
}

// Parse clears any existing macros and define any macros in predefine. Then
// Parse preprocesses and parses the translation unit consisting of files in
// paths. The m communicates the scalar types model and opts allow to amend the
// parser. Parse is not reentrant nor safe to call concurrently.
func Parse(cfg *ParseConfig) (*TranslationUnit, error) {
	if err := CheckParseConfig(cfg); err != nil {
		return nil, err
	}
	compilation.ClearErrors()

	Macros = make(map[int]*macro)
	if err := cfg.SizeModel.sanityCheck(); err != nil {
		compilation.Err(0, "%s", err.Error())
		return nil, compilation.Errors(true)
	}

	model = cfg.SizeModel
	maxAlignment = -1
	for _, v := range model {
		maxAlignment = mathutil.Max(maxAlignment, v.Align)
	}

	includePaths = cfg.IncludePaths
	sysIncludePaths = cfg.SysIncludePaths
	webIncludePrefix = cfg.WebIncludePrefix
	webIncludesEnabled = cfg.WebIncludesEnabled
	yyDebug = cfg.YyDebugLevel

	lx := newTULexer()
	lx.cpp = cfg.OnPreprocessLine
	if err := compilation.Errors(true); err != nil {
		return nil, err
	}
	lx.ch = make(chan []xc.Token, 1000)
	lx0 := newTULexer()
	ctx := newEvalCtx(lx0, lx.ch)

	go func() {
		defer close(lx.ch)

		predefines(cfg.Predefined, lx.ch)
		if len(cfg.Paths) > 0 {
			for _, path := range cfg.Paths {
				ppFileByPath(xc.Token{}, path).preprocess(ctx)
			}
		} else if len(cfg.Input) > 0 {
			ppFileBySrc(xc.Token{}, cfg.Input, "<input>").preprocess(ctx)
		}
	}()

	if err := compilation.Errors(true); err != nil {
		// Do not parse if preprocessing failed.
		return nil, err
	}

	lx.state = lsTranslationUnit0
	r := yyParse(lx)
	if err := compilation.Errors(true); err != nil {
		return nil, err
	}
	if r != 0 {
		return nil, errors.New("internal error")
	}

	return lx.tu, nil
}
