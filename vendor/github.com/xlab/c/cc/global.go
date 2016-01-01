// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/cznic/golex/lex"
	"github.com/cznic/strutil"
	"github.com/xlab/c/xc"
)

const (
	commentNotClosed = "comment not closed"
)

var (
	dict         = xc.Dict
	fset         = xc.FileSet
	isTesting    bool // Test hook.
	isGenerating bool // go generate hook.
	printHooks   = strutil.PrettyPrintHooks{}
)

func init() {
	for k, v := range xc.PrintHooks {
		printHooks[k] = v
	}
	lcRT := reflect.TypeOf(lex.Char{})
	lcH := func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		c := v.(lex.Char)
		r := c.Rune
		s := yySymName(int(r))
		if x := s[0]; x >= '0' && x <= '9' {
			s = strconv.QuoteRune(r)
		}
		f.Format("%s%v: %s"+suffix, prefix, fset.Position(c.Pos()), s)
	}

	printHooks[lcRT] = lcH
	printHooks[reflect.TypeOf(xc.Token{})] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		t := v.(xc.Token)
		if !t.Pos().IsValid() {
			return
		}

		lcH(f, t.Char, prefix, "")
		if s := t.S(); len(s) != 0 {
			f.Format(" %q", s)
		}
		f.Format(suffix)
	}

	printHooks[reflect.TypeOf(PPTokenList(0))] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		x := v.(PPTokenList)
		if x == 0 {
			return
		}

		a := strings.Split(prefix+PrettyString(decodeTokens(x, nil))+",", "\n")
		for _, v := range a {
			f.Format("%s\n", v)
		}

	}

	printHooks[reflect.TypeOf((*ctype)(nil))] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		f.Format(prefix)
		f.Format("%v", v.(Type).String())
		f.Format(suffix)
	}

	printHooks[reflect.TypeOf(Kind(0))] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		f.Format(prefix)
		f.Format("%v", v.(Kind))
		f.Format(suffix)
	}

	printHooks[reflect.TypeOf(Linkage(0))] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		f.Format(prefix)
		f.Format("%v", v.(Linkage))
		f.Format(suffix)
	}
}
