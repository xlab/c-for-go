// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/cznic/c/xc"
	"github.com/cznic/golex/lex"
)

func predefines(src string, ch chan []xc.Token) {
	fileToken := fileset.AddFile("<built-in>", -1, len(src))
	scanner := newUtf8src(strings.NewReader(src), fileToken)
	lx := newLexer(scanner, false)
	defer lx.close()

	lx.predefines = true
	r := parsePreprocessingFile(lx)
	if r == nil {
		return
	}

	r.preprocess(newEvalCtx(lx, ch))
}

func ppFileByPath(refTok xc.Token, path string) *PreprocessingFile {
	once := compilation.Once(path, func() interface{} {
		f, err := os.Open(path)
		if err != nil {
			compilation.ErrTok(refTok, err.Error())
			return nil
		}
		defer f.Close()

		info, err := f.Stat()
		if err != nil {
			compilation.ErrTok(refTok, err.Error())
			return nil
		}

		fileToken := fileset.AddFile(path, -1, int(info.Size()))
		scanner := newUtf8src(f, fileToken)
		lx := newLexer(scanner, false)
		defer lx.close()

		return parsePreprocessingFile(lx)
	})
	//
	switch x := once.Value().(type) {
	case nil:
		return nil
	case *PreprocessingFile:
		return x
	default:
		compilation.ErrTok(refTok, "parsing error: %s", path)
		return nil
	}
}

func ppFileBySrc(refTok xc.Token, src []byte, name string) *PreprocessingFile {
	once := compilation.Once(name, func() interface{} {
		if len(src) == 0 {
			compilation.ErrTok(refTok, "unexpeceted EOF: empty source")
			return nil
		}
		fileToken := fileset.AddFile(name, -1, len(src))
		scanner := newUtf8src(bytes.NewReader(src), fileToken)
		lx := newLexer(scanner, false)
		defer lx.close()

		return parsePreprocessingFile(lx)
	})
	//
	switch x := once.Value().(type) {
	case nil:
		return nil
	case *PreprocessingFile:
		return x
	default:
		compilation.ErrTok(refTok, "parsing error: %s", name)
		return nil
	}
}

func ppFileByURL(refTok xc.Token, u *url.URL) *PreprocessingFile {
	resp, err := http.Get(u.String())
	if err != nil || resp.StatusCode != http.StatusOK {
		if err != nil {
			compilation.ErrTok(refTok, err.Error())
			return nil
		}
		compilation.ErrTok(refTok, fmt.Sprintf("unable to fetch include: %s", resp.Status))
		defer resp.Body.Close()

		base := filepath.Base(u.Path)
		u.Path = filepath.Dir(u.Path)
		if resp, err := http.Get(u.String()); err != nil || resp.StatusCode != http.StatusOK {
			return nil
		} else {
			u = resp.Request.URL
			u.Path = filepath.Join(u.Path, base)
		}
	}
	resp, err = http.Get(u.String())
	if err != nil {
		compilation.ErrTok(refTok, err.Error())
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		compilation.ErrTok(refTok, fmt.Sprintf("unable to fetch include: %s", resp.Status))
		return nil
	}
	compilation.ClearErrors()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		compilation.ErrTok(refTok, err.Error())
		return nil
	}
	return ppFileBySrc(refTok, body, u.String())
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
