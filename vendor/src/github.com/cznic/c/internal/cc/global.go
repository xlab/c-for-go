// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"encoding/binary"
	"fmt"
	"go/token"
	"reflect"
	"strconv"
	"sync"

	"github.com/cznic/c/internal/xc"
	"github.com/cznic/golex/lex"
	"github.com/cznic/mathutil"
	"github.com/cznic/strutil"
)

const (
	maxUvarint  = (64 + 7) / 7
	dbPageShift = 20 //DONE bench tune again using corpus
	dbPageSize  = 1 << dbPageShift
	dbPageMask  = dbPageSize - 1
)

var (
	compilation  = xc.Compilation
	db           = newTokDB()
	dict         = xc.Dict
	fileset      = xc.FileSet
	includePaths []string

	isExample bool

	// dict id: macro. No mutex needed, AST preprocessing is serial.

	// Macros records macros defined during preprocessing. Map key is the identifier ID.
	Macros = map[int]*macro{} //TODO unexport? no thanks

	maxAlignment    int
	model           Model
	printHooks      = strutil.PrettyPrintHooks{}
	sysIncludePaths []string

	webIncludePrefix  string
	enableWebIncludes bool
)

func init() {
	if n, m := dbPageSize, maxUvarint; n < m {
		panic(fmt.Errorf("dbPageSize too low: %v < %d", n, m))
	}

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
		f.Format("%s%v: %s"+suffix, prefix, fileset.Position(c.Pos()), s)
	}

	printHooks[lcRT] = lcH
	printHooks[reflect.TypeOf(xc.Token{})] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		t := v.(xc.Token)
		if !t.Pos().IsValid() {
			return
		}

		lcH(f, t.Char, prefix, "")
		if s := dict.S(t.Val); len(s) != 0 {
			f.Format(" %q", s)
		}
		f.Format(suffix)
		return
	}

	printHooks[reflect.TypeOf(PpTokenList(0))] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		x := v.(PpTokenList)
		if x == 0 {
			return
		}

		toks := db.tokens(x)
		strutil.PrettyPrint(f, toks, prefix, suffix, printHooks)
		return
	}
}

type tokDB struct {
	id    int
	mu    sync.RWMutex
	pages [][]byte
}

func newTokDB() *tokDB {
	r := &tokDB{}
	r.putInt(0) // Reserve id #0.
	return r
}

func (d *tokDB) cap() int {
	d.mu.RLock() // R-

	r := 0
	for _, v := range d.pages {
		r += cap(v)
	}

	d.mu.RUnlock() // R-

	return r
}

func (d *tokDB) int(id int) int {
	d.mu.RLock() // R+

	n, _ := binary.Uvarint(d.pages[id>>dbPageShift][id&dbPageMask:])

	d.mu.RUnlock() // R-

	return int(n)
}

func (d *tokDB) intpUnlocked(id *PpTokenList) int {
	n, l := binary.Uvarint(d.pages[*id>>dbPageShift][*id&dbPageMask:])
	off := (int(*id) & dbPageMask) + l
	if d := dbPageSize - off; d < maxUvarint {
		l += d
	}
	*id = *id + PpTokenList(l)
	return int(n)
}

func (d *tokDB) len() int {
	d.mu.RLock() // R+

	r := 0
	for _, v := range d.pages {
		r += len(v)
	}

	d.mu.RUnlock() // R-

	return r
}

func (d *tokDB) putInt(n int) int {
	d.mu.Lock() // W+

	pi := d.id >> dbPageShift
	if pi < len(d.pages) {
		p := d.pages[pi]
		if cap(p)-len(p) >= maxUvarint {
			l := binary.PutUvarint(p[len(p):cap(p)], uint64(n))
			p = p[:len(p)+l]
			d.pages[pi] = p
			id := d.id
			d.id += l

			d.mu.Unlock() // W-

			return id
		}

		pi++
	}

	p := make([]byte, mathutil.Max(dbPageSize, maxUvarint))
	p = p[:binary.PutUvarint(p, uint64(n))]
	d.pages = append(d.pages, p)
	id := pi << dbPageShift
	d.id = id + mathutil.Min(dbPageSize, len(p))

	d.mu.Unlock() // W-

	return id

}

func (d *tokDB) putIntUnlocked(n int) int {
	pi := d.id >> dbPageShift
	if pi < len(d.pages) {
		p := d.pages[pi]
		if cap(p)-len(p) >= maxUvarint {
			l := binary.PutUvarint(p[len(p):cap(p)], uint64(n))
			p = p[:len(p)+l]
			d.pages[pi] = p
			id := d.id
			d.id += l
			return id
		}

		pi++
	}

	p := make([]byte, mathutil.Max(dbPageSize, maxUvarint))
	p = p[:binary.PutUvarint(p, uint64(n))]
	d.pages = append(d.pages, p)
	id := pi << dbPageShift
	d.id = id + mathutil.Min(dbPageSize, len(p))
	return id
}

//func (d *tokDB) putToken(lastPos *token.Pos, pos token.Pos, r rune, idVal int) int {
//	d.mu.Lock() // W+
//
//	id := d.putIntUnlocked(int(r))
//	if r != '\n' {
//		d.putIntUnlocked(int(pos - *lastPos))
//		if idVal != 0 {
//			d.putIntUnlocked(int(idVal))
//		}
//	}
//
//	d.mu.Unlock() // W-
//
//	*lastPos = pos
//	return id
//}

func (d *tokDB) putTokens(ts []xc.Token) (id int) {
	d.mu.Lock() // W+

	var pos token.Pos
	for i, t := range ts {
		c := t.Char
		id0 := d.putIntUnlocked(int(c.Rune))
		if i == 0 {
			id = id0
		}

		d.putIntUnlocked(int(c.Pos() - pos))
		if t.Val != 0 {
			d.putIntUnlocked(int(t.Val))
		}

		pos = c.Pos()
	}
	d.putIntUnlocked('\n')

	d.mu.Unlock() // W-
	return id
}

//func (d *tokDB) token(id int) Token {
//	d.mu.RLock() // R+
//
//	c := rune(d.intpUnlocked(&id))
//	pos := token.Pos(d.intpUnlocked(&id))
//	var v int
//	if tokHasVal[c] {
//		v = d.intpUnlocked(&id)
//	}
//
//	d.mu.RUnlock() // R-
//
//	return Token{c: newChar(pos, c), val: v}
//}

func (d *tokDB) tokens(id PpTokenList) (r []xc.Token) {
	if id == 0 {
		return nil
	}

	d.mu.RLock() // R+

	var pos token.Pos
	for {
		c := rune(d.intpUnlocked(&id))
		if c == '\n' {

			d.mu.RUnlock() // R-

			return r
		}

		pos += token.Pos(d.intpUnlocked(&id))
		var v int
		if tokHasVal[c] {
			v = d.intpUnlocked(&id)
		}
		t := xc.Token{Char: lex.NewChar(pos, c), Val: v}
		r = append(r, t)
	}
}
