// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.4.

package cc

import (
	"fmt"

	"github.com/cznic/golex/lex"
)

const (
	_           = iota
	scCOMMENT   // [`/*`, `*/`]
	scDEFINE    // [^#define, next token]
	scDIRECTIVE // [^#, next token]
	scHEADER    // [`#include`, next token]
)

func (l *lexer) scan0() lex.Char {
	c := l.enter()

yystate0:
	yyrule := -1
	_ = yyrule
	c = l.rule0()

	switch yyt := l.sc; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: COMMENT
		goto yystart142
	case 2: // start condition: DEFINE
		goto yystart147
	case 3: // start condition: DIRECTIVE
		goto yystart160
	case 4: // start condition: HEADER
		goto yystart244
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
	case 12:
		goto yyrule12
	case 13:
		goto yyrule13
	case 14:
		goto yyrule14
	case 15:
		goto yyrule15
	case 16:
		goto yyrule16
	case 17:
		goto yyrule17
	case 18:
		goto yyrule18
	case 19:
		goto yyrule19
	case 20:
		goto yyrule20
	case 21:
		goto yyrule21
	case 22:
		goto yyrule22
	case 23:
		goto yyrule23
	case 24:
		goto yyrule24
	case 25:
		goto yyrule25
	case 26:
		goto yyrule26
	case 27:
		goto yyrule27
	case 28:
		goto yyrule28
	case 29:
		goto yyrule29
	case 30:
		goto yyrule30
	case 31:
		goto yyrule31
	case 32:
		goto yyrule32
	case 33:
		goto yyrule33
	case 34:
		goto yyrule34
	case 35:
		goto yyrule35
	case 36:
		goto yyrule36
	case 37:
		goto yyrule37
	case 38:
		goto yyrule38
	case 39:
		goto yyrule39
	case 40:
		goto yyrule40
	case 41:
		goto yyrule41
	case 42:
		goto yyrule42
	case 43:
		goto yyrule43
	case 44:
		goto yyrule44
	case 45:
		goto yyrule45
	case 46:
		goto yyrule46
	case 47:
		goto yyrule47
	case 48:
		goto yyrule48
	case 49:
		goto yyrule49
	case 50:
		goto yyrule50
	case 51:
		goto yyrule51
	case 52:
		goto yyrule52
	case 53:
		goto yyrule53
	case 54:
		goto yyrule54
	case 55:
		goto yyrule55
	case 56:
		goto yyrule56
	case 57:
		goto yyrule57
	case 58:
		goto yyrule58
	case 59:
		goto yyrule59
	case 60:
		goto yyrule60
	case 61:
		goto yyrule61
	case 62:
		goto yyrule62
	case 63:
		goto yyrule63
	}
	goto yystate1 // silence unused label error
yystate1:
	c = l.next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '#':
		goto yystate16
	case c == '$' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0084':
		goto yystate20
	case c == '%':
		goto yystate30
	case c == '&':
		goto yystate34
	case c == '*':
		goto yystate49
	case c == '+':
		goto yystate51
	case c == '-':
		goto yystate54
	case c == '.':
		goto yystate58
	case c == '/':
		goto yystate77
	case c == '0':
		goto yystate81
	case c == ':':
		goto yystate98
	case c == '<':
		goto yystate100
	case c == '=':
		goto yystate106
	case c == '>':
		goto yystate108
	case c == 'L':
		goto yystate112
	case c == '\'':
		goto yystate37
	case c == '\\':
		goto yystate21
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate2
	case c == '\u0080':
		goto yystate141
	case c == '^':
		goto yystate136
	case c == '|':
		goto yystate138
	case c >= '1' && c <= '9':
		goto yystate97
	}

yystate2:
	c = l.next()
	yyrule = 1
	l.m = len(l.in)
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate2
	}

yystate3:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate4
	}

yystate4:
	c = l.next()
	yyrule = 7
	l.m = len(l.in)
	goto yyrule7

yystate5:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate5
	}

yystate6:
	c = l.next()
	yyrule = 63
	l.m = len(l.in)
	goto yyrule63

yystate7:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"' || c == '\'' || c >= '0' && c <= '7' || c == '?' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate5
	case c == 'U':
		goto yystate8
	case c == 'u':
		goto yystate12
	case c == 'x':
		goto yystate15
	}

yystate8:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate9
	}

yystate9:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate10
	}

yystate10:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate11
	}

yystate11:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate12
	}

yystate12:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate13
	}

yystate13:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate14
	}

yystate14:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate15
	}

yystate15:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate5
	}

yystate16:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '#':
		goto yystate17
	case c == '%':
		goto yystate18
	}

yystate17:
	c = l.next()
	yyrule = 33
	l.m = len(l.in)
	goto yyrule33

yystate18:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == ':':
		goto yystate19
	}

yystate19:
	c = l.next()
	yyrule = 34
	l.m = len(l.in)
	goto yyrule34

yystate20:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate21:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'U':
		goto yystate22
	case c == 'u':
		goto yystate26
	}

yystate22:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate23
	}

yystate23:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate24
	}

yystate24:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate25
	}

yystate25:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate26
	}

yystate26:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate27
	}

yystate27:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate28
	}

yystate28:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate29
	}

yystate29:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate20
	}

yystate30:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == ':':
		goto yystate31
	case c == '=':
		goto yystate32
	case c == '>':
		goto yystate33
	}

yystate31:
	c = l.next()
	yyrule = 35
	l.m = len(l.in)
	goto yyrule35

yystate32:
	c = l.next()
	yyrule = 8
	l.m = len(l.in)
	goto yyrule8

yystate33:
	c = l.next()
	yyrule = 9
	l.m = len(l.in)
	goto yyrule9

yystate34:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '&':
		goto yystate35
	case c == '=':
		goto yystate36
	}

yystate35:
	c = l.next()
	yyrule = 10
	l.m = len(l.in)
	goto yyrule10

yystate36:
	c = l.next()
	yyrule = 11
	l.m = len(l.in)
	goto yyrule11

yystate37:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\\':
		goto yystate40
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate38
	}

yystate38:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate39
	case c == '\\':
		goto yystate40
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate38
	}

yystate39:
	c = l.next()
	yyrule = 57
	l.m = len(l.in)
	goto yyrule57

yystate40:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"' || c == '\'' || c >= '0' && c <= '7' || c == '?' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate38
	case c == 'U':
		goto yystate41
	case c == 'u':
		goto yystate45
	case c == 'x':
		goto yystate48
	}

yystate41:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate42
	}

yystate42:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate43
	}

yystate43:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate44
	}

yystate44:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate45
	}

yystate45:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate46
	}

yystate46:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate47
	}

yystate47:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate48
	}

yystate48:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate38
	}

yystate49:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate50
	}

yystate50:
	c = l.next()
	yyrule = 12
	l.m = len(l.in)
	goto yyrule12

yystate51:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '+':
		goto yystate52
	case c == '=':
		goto yystate53
	}

yystate52:
	c = l.next()
	yyrule = 13
	l.m = len(l.in)
	goto yyrule13

yystate53:
	c = l.next()
	yyrule = 14
	l.m = len(l.in)
	goto yyrule14

yystate54:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '-':
		goto yystate55
	case c == '=':
		goto yystate56
	case c == '>':
		goto yystate57
	}

yystate55:
	c = l.next()
	yyrule = 15
	l.m = len(l.in)
	goto yyrule15

yystate56:
	c = l.next()
	yyrule = 16
	l.m = len(l.in)
	goto yyrule16

yystate57:
	c = l.next()
	yyrule = 17
	l.m = len(l.in)
	goto yyrule17

yystate58:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate59
	case c >= '0' && c <= '9':
		goto yystate61
	}

yystate59:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate60
	}

yystate60:
	c = l.next()
	yyrule = 18
	l.m = len(l.in)
	goto yyrule18

yystate61:
	c = l.next()
	yyrule = 61
	l.m = len(l.in)
	switch {
	default:
		goto yyrule61
	case c == '$' || c == '.' || c >= 'A' && c <= 'D' || c >= 'G' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'g' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'e':
		goto yystate73
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate76
	case c == 'P' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	case c >= '0' && c <= '9':
		goto yystate61
	}

yystate62:
	c = l.next()
	yyrule = 62
	l.m = len(l.in)
	switch {
	default:
		goto yyrule62
	case c == '$' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	}

yystate63:
	c = l.next()
	yyrule = 62
	l.m = len(l.in)
	switch {
	default:
		goto yyrule62
	case c == '$' || c == '+' || c == '-' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	}

yystate64:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'U':
		goto yystate65
	case c == 'u':
		goto yystate69
	}

yystate65:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate66
	}

yystate66:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate67
	}

yystate67:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate68
	}

yystate68:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate69
	}

yystate69:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate70
	}

yystate70:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate71
	}

yystate71:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate72
	}

yystate72:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate62
	}

yystate73:
	c = l.next()
	yyrule = 62
	l.m = len(l.in)
	switch {
	default:
		goto yyrule62
	case c == '$' || c == '.' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == '+' || c == '-':
		goto yystate74
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	case c >= '0' && c <= '9':
		goto yystate75
	}

yystate74:
	c = l.next()
	yyrule = 62
	l.m = len(l.in)
	switch {
	default:
		goto yyrule62
	case c == '$' || c == '.' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	case c >= '0' && c <= '9':
		goto yystate75
	}

yystate75:
	c = l.next()
	yyrule = 61
	l.m = len(l.in)
	switch {
	default:
		goto yyrule61
	case c == '$' || c == '.' || c >= 'A' && c <= 'D' || c >= 'G' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'g' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate76
	case c == '\\':
		goto yystate64
	case c >= '0' && c <= '9':
		goto yystate75
	}

yystate76:
	c = l.next()
	yyrule = 61
	l.m = len(l.in)
	switch {
	default:
		goto yyrule61
	case c == '$' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	}

yystate77:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate78
	case c == '/':
		goto yystate79
	case c == '=':
		goto yystate80
	}

yystate78:
	c = l.next()
	yyrule = 3
	l.m = len(l.in)
	goto yyrule3

yystate79:
	c = l.next()
	yyrule = 2
	l.m = len(l.in)
	switch {
	default:
		goto yyrule2
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate79
	}

yystate80:
	c = l.next()
	yyrule = 19
	l.m = len(l.in)
	goto yyrule19

yystate81:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c == 'V' || c == 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c == 'v' || c == 'w' || c == 'y' || c == 'z' || c == '\u0084':
		goto yystate62
	case c == '.':
		goto yystate61
	case c == '8' || c == '9':
		goto yystate83
	case c == 'E' || c == 'e':
		goto yystate73
	case c == 'L':
		goto yystate84
	case c == 'P' || c == 'p':
		goto yystate63
	case c == 'U' || c == 'u':
		goto yystate87
	case c == 'X' || c == 'x':
		goto yystate91
	case c == '\\':
		goto yystate64
	case c == 'l':
		goto yystate90
	case c >= '0' && c <= '7':
		goto yystate82
	}

yystate82:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == '.':
		goto yystate61
	case c == '8' || c == '9':
		goto yystate83
	case c == 'E' || c == 'e':
		goto yystate73
	case c == 'L':
		goto yystate84
	case c == 'P' || c == 'p':
		goto yystate63
	case c == 'U' || c == 'u':
		goto yystate87
	case c == '\\':
		goto yystate64
	case c == 'l':
		goto yystate90
	case c >= '0' && c <= '7':
		goto yystate82
	}

yystate83:
	c = l.next()
	yyrule = 62
	l.m = len(l.in)
	switch {
	default:
		goto yyrule62
	case c == '$' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == '.':
		goto yystate61
	case c == 'E' || c == 'e':
		goto yystate73
	case c == 'P' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	case c >= '0' && c <= '9':
		goto yystate83
	}

yystate84:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == 'L':
		goto yystate85
	case c == 'U' || c == 'u':
		goto yystate86
	case c == '\\':
		goto yystate64
	}

yystate85:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == 'U' || c == 'u':
		goto yystate86
	case c == '\\':
		goto yystate64
	}

yystate86:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	}

yystate87:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == 'L':
		goto yystate88
	case c == '\\':
		goto yystate64
	case c == 'l':
		goto yystate89
	}

yystate88:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == 'L':
		goto yystate86
	case c == '\\':
		goto yystate64
	}

yystate89:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	case c == 'l':
		goto yystate86
	}

yystate90:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'P' || c == 'e' || c == 'p':
		goto yystate63
	case c == 'U' || c == 'u':
		goto yystate86
	case c == '\\':
		goto yystate64
	case c == 'l':
		goto yystate85
	}

yystate91:
	c = l.next()
	yyrule = 62
	l.m = len(l.in)
	switch {
	default:
		goto yyrule62
	case c == '$' || c >= 'G' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == '.':
		goto yystate92
	case c == 'E' || c == 'e':
		goto yystate96
	case c == 'P' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate95
	}

yystate92:
	c = l.next()
	yyrule = 62
	l.m = len(l.in)
	switch {
	default:
		goto yyrule62
	case c == '$' || c == '.' || c >= 'G' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'e':
		goto yystate94
	case c == 'P' || c == 'p':
		goto yystate63
	case c == '\\':
		goto yystate64
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate93
	}

yystate93:
	c = l.next()
	yyrule = 62
	l.m = len(l.in)
	switch {
	default:
		goto yyrule62
	case c == '$' || c == '.' || c >= 'G' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'e':
		goto yystate94
	case c == 'P' || c == 'p':
		goto yystate73
	case c == '\\':
		goto yystate64
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate93
	}

yystate94:
	c = l.next()
	yyrule = 62
	l.m = len(l.in)
	switch {
	default:
		goto yyrule62
	case c == '$' || c == '+' || c == '-' || c == '.' || c >= 'G' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == 'E' || c == 'e':
		goto yystate94
	case c == 'P' || c == 'p':
		goto yystate73
	case c == '\\':
		goto yystate64
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate93
	}

yystate95:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c >= 'G' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == '.':
		goto yystate93
	case c == 'E' || c == 'e':
		goto yystate96
	case c == 'L':
		goto yystate84
	case c == 'P' || c == 'p':
		goto yystate73
	case c == 'U' || c == 'u':
		goto yystate87
	case c == '\\':
		goto yystate64
	case c == 'l':
		goto yystate90
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate95
	}

yystate96:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c == '+' || c == '-' || c >= 'G' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'g' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == '.':
		goto yystate93
	case c == 'E' || c == 'e':
		goto yystate96
	case c == 'L':
		goto yystate84
	case c == 'P' || c == 'p':
		goto yystate73
	case c == 'U' || c == 'u':
		goto yystate87
	case c == '\\':
		goto yystate64
	case c == 'l':
		goto yystate90
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c == 'F' || c >= 'a' && c <= 'd' || c == 'f':
		goto yystate95
	}

yystate97:
	c = l.next()
	yyrule = 60
	l.m = len(l.in)
	switch {
	default:
		goto yyrule60
	case c == '$' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'O' || c >= 'Q' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0084':
		goto yystate62
	case c == '.':
		goto yystate61
	case c == 'E' || c == 'e':
		goto yystate73
	case c == 'L':
		goto yystate84
	case c == 'P' || c == 'p':
		goto yystate63
	case c == 'U' || c == 'u':
		goto yystate87
	case c == '\\':
		goto yystate64
	case c == 'l':
		goto yystate90
	case c >= '0' && c <= '9':
		goto yystate97
	}

yystate98:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '>':
		goto yystate99
	}

yystate99:
	c = l.next()
	yyrule = 20
	l.m = len(l.in)
	goto yyrule20

yystate100:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '%':
		goto yystate101
	case c == ':':
		goto yystate102
	case c == '<':
		goto yystate103
	case c == '=':
		goto yystate105
	}

yystate101:
	c = l.next()
	yyrule = 21
	l.m = len(l.in)
	goto yyrule21

yystate102:
	c = l.next()
	yyrule = 22
	l.m = len(l.in)
	goto yyrule22

yystate103:
	c = l.next()
	yyrule = 23
	l.m = len(l.in)
	switch {
	default:
		goto yyrule23
	case c == '=':
		goto yystate104
	}

yystate104:
	c = l.next()
	yyrule = 24
	l.m = len(l.in)
	goto yyrule24

yystate105:
	c = l.next()
	yyrule = 25
	l.m = len(l.in)
	goto yyrule25

yystate106:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate107
	}

yystate107:
	c = l.next()
	yyrule = 26
	l.m = len(l.in)
	goto yyrule26

yystate108:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate109
	case c == '>':
		goto yystate110
	}

yystate109:
	c = l.next()
	yyrule = 27
	l.m = len(l.in)
	goto yyrule27

yystate110:
	c = l.next()
	yyrule = 28
	l.m = len(l.in)
	switch {
	default:
		goto yyrule28
	case c == '=':
		goto yystate111
	}

yystate111:
	c = l.next()
	yyrule = 29
	l.m = len(l.in)
	goto yyrule29

yystate112:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '"':
		goto yystate113
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\'':
		goto yystate124
	case c == '\\':
		goto yystate21
	}

yystate113:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate114
	case c == '\\':
		goto yystate115
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate113
	}

yystate114:
	c = l.next()
	yyrule = 56
	l.m = len(l.in)
	goto yyrule56

yystate115:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"' || c == '\'' || c >= '0' && c <= '7' || c == '?' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate113
	case c == 'U':
		goto yystate116
	case c == 'u':
		goto yystate120
	case c == 'x':
		goto yystate123
	}

yystate116:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate117
	}

yystate117:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate118
	}

yystate118:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate119
	}

yystate119:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate120
	}

yystate120:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate121
	}

yystate121:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate122
	}

yystate122:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate123
	}

yystate123:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate113
	}

yystate124:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\\':
		goto yystate127
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate125
	}

yystate125:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate126
	case c == '\\':
		goto yystate127
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate125
	}

yystate126:
	c = l.next()
	yyrule = 55
	l.m = len(l.in)
	goto yyrule55

yystate127:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"' || c == '\'' || c >= '0' && c <= '7' || c == '?' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate125
	case c == 'U':
		goto yystate128
	case c == 'u':
		goto yystate132
	case c == 'x':
		goto yystate135
	}

yystate128:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate129
	}

yystate129:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate130
	}

yystate130:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate131
	}

yystate131:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate132
	}

yystate132:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate133
	}

yystate133:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate134
	}

yystate134:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate135
	}

yystate135:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate125
	}

yystate136:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate137
	}

yystate137:
	c = l.next()
	yyrule = 30
	l.m = len(l.in)
	goto yyrule30

yystate138:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate139
	case c == '|':
		goto yystate140
	}

yystate139:
	c = l.next()
	yyrule = 31
	l.m = len(l.in)
	goto yyrule31

yystate140:
	c = l.next()
	yyrule = 32
	l.m = len(l.in)
	goto yyrule32

yystate141:
	c = l.next()
	yyrule = 6
	l.m = len(l.in)
	goto yyrule6

	goto yystate142 // silence unused label error
yystate142:
	c = l.next()
yystart142:
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate144
	case c == '\u0080':
		goto yystate146
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate143
	}

yystate143:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate144
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate143
	}

yystate144:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate144
	case c == '/':
		goto yystate145
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '.' || c >= '0' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate143
	}

yystate145:
	c = l.next()
	yyrule = 4
	l.m = len(l.in)
	goto yyrule4

yystate146:
	c = l.next()
	yyrule = 5
	l.m = len(l.in)
	goto yyrule5

	goto yystate147 // silence unused label error
yystate147:
	c = l.next()
yystart147:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '#':
		goto yystate16
	case c == '$' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0084':
		goto yystate148
	case c == '%':
		goto yystate30
	case c == '&':
		goto yystate34
	case c == '*':
		goto yystate49
	case c == '+':
		goto yystate51
	case c == '-':
		goto yystate54
	case c == '.':
		goto yystate58
	case c == '/':
		goto yystate77
	case c == '0':
		goto yystate81
	case c == ':':
		goto yystate98
	case c == '<':
		goto yystate100
	case c == '=':
		goto yystate106
	case c == '>':
		goto yystate108
	case c == 'L':
		goto yystate159
	case c == '\'':
		goto yystate37
	case c == '\\':
		goto yystate150
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate2
	case c == '\u0080':
		goto yystate141
	case c == '^':
		goto yystate136
	case c == '|':
		goto yystate138
	case c >= '1' && c <= '9':
		goto yystate97
	}

yystate148:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate148
	case c == '(':
		goto yystate149
	case c == '\\':
		goto yystate150
	}

yystate149:
	c = l.next()
	yyrule = 59
	l.m = len(l.in)
	goto yyrule59

yystate150:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'U':
		goto yystate151
	case c == 'u':
		goto yystate155
	}

yystate151:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate152
	}

yystate152:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate153
	}

yystate153:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate154
	}

yystate154:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate155
	}

yystate155:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate156
	}

yystate156:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate157
	}

yystate157:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate158
	}

yystate158:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate148
	}

yystate159:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '"':
		goto yystate113
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate148
	case c == '(':
		goto yystate149
	case c == '\'':
		goto yystate124
	case c == '\\':
		goto yystate150
	}

	goto yystate160 // silence unused label error
yystate160:
	c = l.next()
yystart160:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '#':
		goto yystate16
	case c == '$' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c == 'b' || c == 'c' || c >= 'f' && c <= 'h' || c == 'j' || c == 'k' || c >= 'm' && c <= 'o' || c >= 'q' && c <= 't' || c == 'v' || c >= 'x' && c <= 'z' || c == '\u0084':
		goto yystate20
	case c == '%':
		goto yystate30
	case c == '&':
		goto yystate34
	case c == '*':
		goto yystate49
	case c == '+':
		goto yystate51
	case c == '-':
		goto yystate54
	case c == '.':
		goto yystate58
	case c == '/':
		goto yystate77
	case c == '0':
		goto yystate81
	case c == ':':
		goto yystate98
	case c == '<':
		goto yystate100
	case c == '=':
		goto yystate106
	case c == '>':
		goto yystate108
	case c == 'L':
		goto yystate112
	case c == '\'':
		goto yystate37
	case c == '\\':
		goto yystate21
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate2
	case c == '\u0080':
		goto yystate141
	case c == '^':
		goto yystate136
	case c == 'a':
		goto yystate161
	case c == 'd':
		goto yystate167
	case c == 'e':
		goto yystate173
	case c == 'i':
		goto yystate187
	case c == 'l':
		goto yystate216
	case c == 'p':
		goto yystate220
	case c == 'u':
		goto yystate226
	case c == 'w':
		goto yystate237
	case c == '|':
		goto yystate138
	case c >= '1' && c <= '9':
		goto yystate97
	}

yystate161:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 's':
		goto yystate162
	}

yystate162:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 's':
		goto yystate163
	}

yystate163:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate164
	}

yystate164:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'r':
		goto yystate165
	}

yystate165:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 't':
		goto yystate166
	}

yystate166:
	c = l.next()
	yyrule = 36
	l.m = len(l.in)
	switch {
	default:
		goto yyrule36
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate167:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate168
	}

yystate168:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'f':
		goto yystate169
	}

yystate169:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'i':
		goto yystate170
	}

yystate170:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'n':
		goto yystate171
	}

yystate171:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate172
	}

yystate172:
	c = l.next()
	yyrule = 37
	l.m = len(l.in)
	switch {
	default:
		goto yyrule37
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate173:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c == 'm' || c >= 'o' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'l':
		goto yystate174
	case c == 'n':
		goto yystate179
	case c == 'r':
		goto yystate183
	}

yystate174:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'i':
		goto yystate175
	case c == 's':
		goto yystate177
	}

yystate175:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'f':
		goto yystate176
	}

yystate176:
	c = l.next()
	yyrule = 38
	l.m = len(l.in)
	switch {
	default:
		goto yyrule38
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate177:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate178
	}

yystate178:
	c = l.next()
	yyrule = 39
	l.m = len(l.in)
	switch {
	default:
		goto yyrule39
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate179:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'd':
		goto yystate180
	}

yystate180:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'i':
		goto yystate181
	}

yystate181:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'f':
		goto yystate182
	}

yystate182:
	c = l.next()
	yyrule = 40
	l.m = len(l.in)
	switch {
	default:
		goto yyrule40
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate183:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'r':
		goto yystate184
	}

yystate184:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'o':
		goto yystate185
	}

yystate185:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'r':
		goto yystate186
	}

yystate186:
	c = l.next()
	yyrule = 41
	l.m = len(l.in)
	switch {
	default:
		goto yyrule41
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate187:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c == 'e' || c >= 'g' && c <= 'l' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'd':
		goto yystate188
	case c == 'f':
		goto yystate192
	case c == 'm':
		goto yystate200
	case c == 'n':
		goto yystate205
	}

yystate188:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate189
	}

yystate189:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'n':
		goto yystate190
	}

yystate190:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 't':
		goto yystate191
	}

yystate191:
	c = l.next()
	yyrule = 42
	l.m = len(l.in)
	switch {
	default:
		goto yyrule42
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate192:
	c = l.next()
	yyrule = 43
	l.m = len(l.in)
	switch {
	default:
		goto yyrule43
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'd':
		goto yystate193
	case c == 'n':
		goto yystate196
	}

yystate193:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate194
	}

yystate194:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'f':
		goto yystate195
	}

yystate195:
	c = l.next()
	yyrule = 44
	l.m = len(l.in)
	switch {
	default:
		goto yyrule44
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate196:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'd':
		goto yystate197
	}

yystate197:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate198
	}

yystate198:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'f':
		goto yystate199
	}

yystate199:
	c = l.next()
	yyrule = 45
	l.m = len(l.in)
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate200:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'p':
		goto yystate201
	}

yystate201:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'o':
		goto yystate202
	}

yystate202:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'r':
		goto yystate203
	}

yystate203:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 't':
		goto yystate204
	}

yystate204:
	c = l.next()
	yyrule = 46
	l.m = len(l.in)
	switch {
	default:
		goto yyrule46
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate205:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'c':
		goto yystate206
	}

yystate206:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'l':
		goto yystate207
	}

yystate207:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'u':
		goto yystate208
	}

yystate208:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'd':
		goto yystate209
	}

yystate209:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate210
	}

yystate210:
	c = l.next()
	yyrule = 47
	l.m = len(l.in)
	switch {
	default:
		goto yyrule47
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == '_':
		goto yystate211
	}

yystate211:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'n':
		goto yystate212
	}

yystate212:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate213
	}

yystate213:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'x':
		goto yystate214
	}

yystate214:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 't':
		goto yystate215
	}

yystate215:
	c = l.next()
	yyrule = 48
	l.m = len(l.in)
	switch {
	default:
		goto yyrule48
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate216:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'i':
		goto yystate217
	}

yystate217:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'n':
		goto yystate218
	}

yystate218:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate219
	}

yystate219:
	c = l.next()
	yyrule = 49
	l.m = len(l.in)
	switch {
	default:
		goto yyrule49
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate220:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'r':
		goto yystate221
	}

yystate221:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'a':
		goto yystate222
	}

yystate222:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'g':
		goto yystate223
	}

yystate223:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'm':
		goto yystate224
	}

yystate224:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'a':
		goto yystate225
	}

yystate225:
	c = l.next()
	yyrule = 50
	l.m = len(l.in)
	switch {
	default:
		goto yyrule50
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate226:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'n':
		goto yystate227
	}

yystate227:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'b' || c == 'c' || c >= 'e' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'a':
		goto yystate228
	case c == 'd':
		goto yystate234
	}

yystate228:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 's':
		goto yystate229
	}

yystate229:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 's':
		goto yystate230
	}

yystate230:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate231
	}

yystate231:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'r':
		goto yystate232
	}

yystate232:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 't':
		goto yystate233
	}

yystate233:
	c = l.next()
	yyrule = 51
	l.m = len(l.in)
	switch {
	default:
		goto yyrule51
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate234:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'e':
		goto yystate235
	}

yystate235:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'f':
		goto yystate236
	}

yystate236:
	c = l.next()
	yyrule = 52
	l.m = len(l.in)
	switch {
	default:
		goto yyrule52
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

yystate237:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'a':
		goto yystate238
	}

yystate238:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'r':
		goto yystate239
	}

yystate239:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'n':
		goto yystate240
	}

yystate240:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'i':
		goto yystate241
	}

yystate241:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'n':
		goto yystate242
	}

yystate242:
	c = l.next()
	yyrule = 58
	l.m = len(l.in)
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c == 'g':
		goto yystate243
	}

yystate243:
	c = l.next()
	yyrule = 53
	l.m = len(l.in)
	switch {
	default:
		goto yyrule53
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0083' || c == '\u0084':
		goto yystate20
	case c == '\\':
		goto yystate21
	}

	goto yystate244 // silence unused label error
yystate244:
	c = l.next()
yystart244:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate245
	case c == '#':
		goto yystate16
	case c == '$' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0084':
		goto yystate20
	case c == '%':
		goto yystate30
	case c == '&':
		goto yystate34
	case c == '*':
		goto yystate49
	case c == '+':
		goto yystate51
	case c == '-':
		goto yystate54
	case c == '.':
		goto yystate58
	case c == '/':
		goto yystate77
	case c == '0':
		goto yystate81
	case c == ':':
		goto yystate98
	case c == '<':
		goto yystate260
	case c == '=':
		goto yystate106
	case c == '>':
		goto yystate108
	case c == 'L':
		goto yystate112
	case c == '\'':
		goto yystate37
	case c == '\\':
		goto yystate21
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate2
	case c == '\u0080':
		goto yystate141
	case c == '^':
		goto yystate136
	case c == '|':
		goto yystate138
	case c >= '1' && c <= '9':
		goto yystate97
	}

yystate245:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate248
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate246
	}

yystate246:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate247
	case c == '\\':
		goto yystate248
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate246
	}

yystate247:
	c = l.next()
	yyrule = 54
	l.m = len(l.in)
	goto yyrule54

yystate248:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate251
	case c == 'U':
		goto yystate252
	case c == '\'' || c >= '0' && c <= '7' || c == '?' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate246
	case c == 'u':
		goto yystate256
	case c == 'x':
		goto yystate259
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '&' || c >= '(' && c <= '/' || c >= '8' && c <= '>' || c >= '@' && c <= 'T' || c >= 'V' && c <= '[' || c >= ']' && c <= '`' || c >= 'c' && c <= 'e' || c >= 'g' && c <= 'm' || c >= 'o' && c <= 'q' || c == 's' || c == 'w' || c >= 'y' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate249
	}

yystate249:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate250
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate249
	}

yystate250:
	c = l.next()
	yyrule = 54
	l.m = len(l.in)
	goto yyrule54

yystate251:
	c = l.next()
	yyrule = 54
	l.m = len(l.in)
	switch {
	default:
		goto yyrule54
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate5
	}

yystate252:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate250
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate253
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate249
	}

yystate253:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate250
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate254
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate249
	}

yystate254:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate250
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate255
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate249
	}

yystate255:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate250
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate256
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate249
	}

yystate256:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate250
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate257
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate249
	}

yystate257:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate250
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate258
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate249
	}

yystate258:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate250
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate259
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate249
	}

yystate259:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate250
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate246
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '/' || c >= ':' && c <= '@' || c >= 'G' && c <= '`' || c >= 'g' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate249
	}

yystate260:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '%':
		goto yystate262
	case c == ':':
		goto yystate263
	case c == '<':
		goto yystate264
	case c == '=':
		goto yystate266
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '$' || c >= '&' && c <= '9' || c == ';' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate261
	}

yystate261:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '>':
		goto yystate250
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '=' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate261
	}

yystate262:
	c = l.next()
	yyrule = 21
	l.m = len(l.in)
	switch {
	default:
		goto yyrule21
	case c == '>':
		goto yystate250
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '=' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate261
	}

yystate263:
	c = l.next()
	yyrule = 22
	l.m = len(l.in)
	switch {
	default:
		goto yyrule22
	case c == '>':
		goto yystate250
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '=' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate261
	}

yystate264:
	c = l.next()
	yyrule = 23
	l.m = len(l.in)
	switch {
	default:
		goto yyrule23
	case c == '=':
		goto yystate265
	case c == '>':
		goto yystate250
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '<' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate261
	}

yystate265:
	c = l.next()
	yyrule = 24
	l.m = len(l.in)
	switch {
	default:
		goto yyrule24
	case c == '>':
		goto yystate250
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '=' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate261
	}

yystate266:
	c = l.next()
	yyrule = 25
	l.m = len(l.in)
	switch {
	default:
		goto yyrule25
	case c == '>':
		goto yystate250
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '=' || c >= '?' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate261
	}

yyrule1: // [ \t\f\v]+

	goto yystate0
yyrule2: // "//".*

	goto yystate0
yyrule3: // "/*"
	{
		l.commentPos0 = l.in[0].Pos()
		l.push(scCOMMENT)
		goto yystate0
	}
yyrule4: // {comment-close}
	{
		l.pop()
		goto yystate0
	}
yyrule5: // {eof}
	{
		compilation.Err(l.commentPos0, commentNotClosed)
		l.pop()
		return lex.NewChar(l.in[len(l.in)-1].Pos(), 0)
	}
yyrule6: // {eof}
	{
		return lex.NewChar(l.c.Pos(), runeEOF)
	}
yyrule7: // "!="
	{
		return l.char(NEQ)
	}
yyrule8: // "%="
	{
		return l.char(MODASSIGN)
	}
yyrule9: // "%>"
	{
		return l.char('}')
	}
yyrule10: // "&&"
	{
		return l.char(ANDAND)
	}
yyrule11: // "&="
	{
		return l.char(ANDASSIGN)
	}
yyrule12: // "*="
	{
		return l.char(MULASSIGN)
	}
yyrule13: // "++"
	{
		return l.char(INC)
	}
yyrule14: // "+="
	{
		return l.char(ADDASSIGN)
	}
yyrule15: // "--"
	{
		return l.char(DEC)
	}
yyrule16: // "-="
	{
		return l.char(SUBASSIGN)
	}
yyrule17: // "->"
	{
		return l.char(ARROW)
	}
yyrule18: // "..."
	{
		return l.char(DDD)
	}
yyrule19: // "/="
	{
		return l.char(DIVASSIGN)
	}
yyrule20: // ":>"
	{
		return l.char(']')
	}
yyrule21: // "<%"
	{
		return l.char('{')
	}
yyrule22: // "<:"
	{
		return l.char('[')
	}
yyrule23: // "<<"
	{
		return l.char(LSH)
	}
yyrule24: // "<<="
	{
		return l.char(LSHASSIGN)
	}
yyrule25: // "<="
	{
		return l.char(LEQ)
	}
yyrule26: // "=="
	{
		return l.char(EQ)
	}
yyrule27: // ">="
	{
		return l.char(GEQ)
	}
yyrule28: // ">>"
	{
		return l.char(RSH)
	}
yyrule29: // ">>="
	{
		return l.char(RSHASSIGN)
	}
yyrule30: // "^="
	{
		return l.char(XORASSIGN)
	}
yyrule31: // "|="
	{
		return l.char(ORASSIGN)
	}
yyrule32: // "||"
	{
		return l.char(OROR)
	}
yyrule33: // "##"
yyrule34: // "#%:"
	{
		return l.char(PPPASTE)
	}
yyrule35: // "%:"
	{
		// ['%', ':'], z
		l.unget(l.c, lex.NewChar(l.in[0].Pos(), '#'))
		l.next()
		goto yystate0
	}
yyrule36: // "assert"
	{
		return l.directive(PPASSERT)
	}
yyrule37: // "define"
	{
		return l.directive(PPDEFINE)
	}
yyrule38: // "elif"
	{
		return l.directive(PPELIF)
	}
yyrule39: // "else"
	{
		return l.directive(PPELSE)
	}
yyrule40: // "endif"
	{
		return l.directive(PPENDIF)
	}
yyrule41: // "error"
	{
		return l.directive(PPERROR)
	}
yyrule42: // "ident"
	{
		return l.directive(PPIDENT)
	}
yyrule43: // "if"
	{
		return l.directive(PPIF)
	}
yyrule44: // "ifdef"
	{
		return l.directive(PPIFDEF)
	}
yyrule45: // "ifndef"
	{
		return l.directive(PPIFNDEF)
	}
yyrule46: // "import"
	{
		return l.directive(PPIMPORT)
	}
yyrule47: // "include"
	{
		return l.directive(PPINCLUDE)
	}
yyrule48: // "include_next"
	{
		return l.directive(PPINCLUDE_NEXT)
	}
yyrule49: // "line"
	{
		return l.directive(PPLINE)
	}
yyrule50: // "pragma"
	{
		return l.directive(PPPRAGMA)
	}
yyrule51: // "unassert"
	{
		return l.directive(PPUNASSERT)
	}
yyrule52: // "undef"
	{
		return l.directive(PPUNDEF)
	}
yyrule53: // "warning"
	{
		return l.directive(PPWARNING)
	}
yyrule54: // {header-name}
	{
		l.sc = scINITIAL
		return l.char(PPHEADER_NAME)
	}
yyrule55: // L{character-constant}
	{
		return l.char(LONGCHARCONST)
	}
yyrule56: // L{string-literal}
	{
		return l.char(LONGSTRINGLITERAL)
	}
yyrule57: // {character-constant}
	{
		return l.char(CHARCONST)
	}
yyrule58: // {identifier}
	{
		return l.char(IDENTIFIER)
	}
yyrule59: // {identifier}"("
	{
		return l.char(IDENTIFIER_LPAREN)
	}
yyrule60: // {integer-constant}
	{
		return l.char(INTCONST)
	}
yyrule61: // {floating-constant}
	{
		return l.char(FLOATCONST)
	}
yyrule62: // {pp-number}
	{
		return l.char(PPNUMBER)
	}
yyrule63: // {string-literal}
	{
		return l.char(STRINGLITERAL)
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	if l.m >= 0 {
		if len(l.in) > l.m {
			l.unget(l.c)
			for i := len(l.in) - 1; i >= l.m; i-- {
				l.unget(l.in[i])
			}
			l.next()
		}
		l.in = l.in[:l.m]
		goto yyAction
	}

	return l.abort()
}
