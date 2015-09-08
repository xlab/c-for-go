// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.4.
//
// CAUTION: Generated file (unless it's trigraphs.l) - DO NOT EDIT!

// Implements translation phases 1 and 2 of [0], 5.1.1.2.

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

yystate0:

	c = l.rule0()

	switch yyt := l.sc; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: TRIGRAPHS
		goto yystart4
	}

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = l.next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '\\':
		goto yystate3
	case c == '\r':
		goto yystate2
	}

yystate2:
	c = l.next()
	goto yyrule10

yystate3:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate2
	}

	goto yystate4 // silence unused label error
yystate4:
	c = l.next()
yystart4:
	switch {
	default:
		goto yyabort
	case c == '?':
		goto yystate5
	case c == '\\':
		goto yystate3
	case c == '\r':
		goto yystate2
	}

yystate5:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '?':
		goto yystate6
	}

yystate6:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate7
	case c == '(':
		goto yystate9
	case c == ')':
		goto yystate10
	case c == '-':
		goto yystate11
	case c == '/':
		goto yystate12
	case c == '<':
		goto yystate14
	case c == '=':
		goto yystate15
	case c == '>':
		goto yystate16
	case c == '\'':
		goto yystate8
	}

yystate7:
	c = l.next()
	goto yyrule1

yystate8:
	c = l.next()
	goto yyrule2

yystate9:
	c = l.next()
	goto yyrule3

yystate10:
	c = l.next()
	goto yyrule4

yystate11:
	c = l.next()
	goto yyrule5

yystate12:
	c = l.next()
	switch {
	default:
		goto yyrule6
	case c == '\n':
		goto yystate13
	}

yystate13:
	c = l.next()
	goto yyrule11

yystate14:
	c = l.next()
	goto yyrule7

yystate15:
	c = l.next()
	goto yyrule8

yystate16:
	c = l.next()
	goto yyrule9

yyrule1: // "??!"
	{
		return l.char('|')
	}
yyrule2: // "??'"
	{
		return l.char('^')
	}
yyrule3: // "??("
	{
		return l.char('[')
	}
yyrule4: // "??)"
	{
		return l.char(']')
	}
yyrule5: // "??-"
	{
		return l.char('~')
	}
yyrule6: // "??/"
	{
		return l.char('\\')
	}
yyrule7: // "??<"
	{
		return l.char('{')
	}
yyrule8: // "??="
	{
		return l.char('#')
	}
yyrule9: // "??>"
	{
		return l.char('}')
	}
yyrule10: // \\\n|\r
yyrule11: // "??/"\n

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	return l.abort()
}
