// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.4.
//

// Implements translation phases 1 and 2 of [0], 5.1.1.2.

package cc

import (
	"fmt"
)

const (
	_ = iota
	scTRIGRAPHS
)

func (t *trigraphsReader) scan() (r int) {
	c := t.Enter()

yystate0:
	yyrule := -1
	_ = yyrule
	c = t.Rule0()

	switch yyt := t.sc; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: TRIGRAPHS
		goto yystart5
	}

	goto yystate0 // silence unused label error
	goto yyAction // silence unused label error
yyAction:
	switch yyrule {
	case 1:
		goto yyrule1
	case 2:
		goto yyrule2
	case 3:
		goto yyrule3
	case 4:
		goto yyrule4
	case 5:
		goto yyrule5
	case 6:
		goto yyrule6
	case 7:
		goto yyrule7
	case 8:
		goto yyrule8
	case 9:
		goto yyrule9
	case 10:
		goto yyrule10
	case 11:
		goto yyrule11
	}
	goto yystate1 // silence unused label error
yystate1:
	c = t.Next()
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
	c = t.Next()
	yyrule = 10
	t.Mark()
	goto yyrule10

yystate3:
	c = t.Next()
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate2
	case c == '\r':
		goto yystate4
	}

yystate4:
	c = t.Next()
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate2
	}

	goto yystate5 // silence unused label error
yystate5:
	c = t.Next()
yystart5:
	switch {
	default:
		goto yyabort
	case c == '?':
		goto yystate6
	case c == '\\':
		goto yystate3
	case c == '\r':
		goto yystate2
	}

yystate6:
	c = t.Next()
	switch {
	default:
		goto yyabort
	case c == '?':
		goto yystate7
	}

yystate7:
	c = t.Next()
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate8
	case c == '(':
		goto yystate10
	case c == ')':
		goto yystate11
	case c == '-':
		goto yystate12
	case c == '/':
		goto yystate13
	case c == '<':
		goto yystate16
	case c == '=':
		goto yystate17
	case c == '>':
		goto yystate18
	case c == '\'':
		goto yystate9
	}

yystate8:
	c = t.Next()
	yyrule = 1
	t.Mark()
	goto yyrule1

yystate9:
	c = t.Next()
	yyrule = 2
	t.Mark()
	goto yyrule2

yystate10:
	c = t.Next()
	yyrule = 3
	t.Mark()
	goto yyrule3

yystate11:
	c = t.Next()
	yyrule = 4
	t.Mark()
	goto yyrule4

yystate12:
	c = t.Next()
	yyrule = 5
	t.Mark()
	goto yyrule5

yystate13:
	c = t.Next()
	yyrule = 6
	t.Mark()
	switch {
	default:
		goto yyrule6
	case c == '\n':
		goto yystate14
	case c == '\r':
		goto yystate15
	}

yystate14:
	c = t.Next()
	yyrule = 11
	t.Mark()
	goto yyrule11

yystate15:
	c = t.Next()
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate14
	}

yystate16:
	c = t.Next()
	yyrule = 7
	t.Mark()
	goto yyrule7

yystate17:
	c = t.Next()
	yyrule = 8
	t.Mark()
	goto yyrule8

yystate18:
	c = t.Next()
	yyrule = 9
	t.Mark()
	goto yyrule9

yyrule1: // "??!"
	{
		return '|'
	}
yyrule2: // "??'"
	{
		return '^'
	}
yyrule3: // "??("
	{
		return '['
	}
yyrule4: // "??)"
	{
		return ']'
	}
yyrule5: // "??-"
	{
		return '~'
	}
yyrule6: // "??/"
	{
		return '\\'
	}
yyrule7: // "??<"
	{
		return '{'
	}
yyrule8: // "??="
	{
		return '#'
	}
yyrule9: // "??>"
	{
		return '}'
	}
yyrule10: // \\\r?\n|\r
yyrule11: // "??/"\r?\n

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	if c, ok := t.Abort(); ok {
		return c
	}

	goto yyAction
}
