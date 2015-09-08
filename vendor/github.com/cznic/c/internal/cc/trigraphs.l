%{
// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.4.
//
// CAUTION: Generated file (unless it's trigraphs.l) - DO NOT EDIT!

// Implements translation phases 1 and 2 of [0], 5.1.1.2.
%}

%yyc c
%yyn c = l.next()
%yyt l.sc

%s TRIGRAPHS

%{
package cc

import (
        "fmt"

	"github.com/cznic/golex/lex"
)

const (
        _ = iota
        scTRIGRAPHS
)

func (l *ppLexer) scan() lex.Char {
        c := l.enter()
%}

%%
        c = l.rule0()

<TRIGRAPHS>"??!"        return l.char('|')
<TRIGRAPHS>"??'"        return l.char('^')
<TRIGRAPHS>"??("        return l.char('[')
<TRIGRAPHS>"??)"        return l.char(']')
<TRIGRAPHS>"??-"        return l.char('~')
<TRIGRAPHS>"??/"        return l.char('\\')
<TRIGRAPHS>"??<"        return l.char('{')
<TRIGRAPHS>"??="        return l.char('#')
<TRIGRAPHS>"??>"        return l.char('}')

\\\n|\r                 |
<TRIGRAPHS>"??/"\n

%%
        return l.abort()
}
