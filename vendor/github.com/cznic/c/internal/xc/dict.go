// Copyright 2015 The XC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xc

import (
	"bytes"
	"sync"
)

var (
	// R/O zero values
	dictZDE int
	dictZK  int
	dictZXE dictXE
)

const (
	dictKX = 64 //Done: benchmark tune.
	dictKD = 32 //Done: benchmark tune.
)

type dictD struct {
	c int
	d [2*dictKD + 1]int
}

func (l *dictD) mvL(r *dictD, c int) {
	copy(l.d[l.c:], r.d[:c])
	copy(r.d[:], r.d[c:r.c])
	l.c += c
	r.c -= c
}

func (l *dictD) mvR(r *dictD, c int) {
	copy(r.d[c:], r.d[:r.c])
	copy(r.d[:c], l.d[l.c-c:])
	r.c += c
	l.c -= c
}

type dictX struct {
	c int
	x [2*dictKX + 2]dictXE
}

func newDictX(ch0 interface{}) *dictX {
	r := &dictX{}
	r.x[0].ch = ch0
	return r
}

func (q *dictX) insert(i int, k int, ch interface{}) *dictX {
	c := q.c
	if i < c {
		q.x[c+1].ch = q.x[c].ch
		copy(q.x[i+2:], q.x[i+1:c])
		q.x[i+1].k = q.x[i].k
	}
	c++
	q.c = c
	q.x[i].k = k
	q.x[i+1].ch = ch
	return q
}

func (q *dictX) siblings(i int) (l, r *dictD) {
	if i >= 0 {
		if i > 0 {
			l = q.x[i-1].ch.(*dictD)
		}
		if i < q.c {
			r = q.x[i+1].ch.(*dictD)
		}
	}
	return
}

type dictXE struct {
	ch interface{}
	k  int
}

// Dictionary is a bijection of a set of byte slice values and their associated
// numeric identifiers. A zero length byte slice is guaranteed to have zero ID.
// Dictionary uses the global DB variable.
type Dictionary struct {
	mu sync.Mutex //TODO RWMutex
	r  interface{}
}

func newDictionary() *Dictionary {
	return &Dictionary{}
}

func (d *Dictionary) find(q interface{}, k []byte) (i int, ok bool) {
	l := 0
	switch x := q.(type) {
	case *dictX:
		h := x.c - 1
		for l <= h {
			m := (l + h) >> 1
			switch cmp := bytes.Compare(k, DB.bytesUnlocked(x.x[m].k)); {
			case cmp > 0:
				l = m + 1
			case cmp == 0:
				return m, true
			default:
				h = m - 1
			}
		}
	case *dictD:
		h := x.c - 1
		for l <= h {
			m := (l + h) >> 1
			switch cmp := bytes.Compare(k, DB.bytesUnlocked(x.d[m])); {
			case cmp > 0:
				l = m + 1
			case cmp == 0:
				return m, true
			default:
				h = m - 1
			}
		}
	}
	return l, false
}

// ID returns the identifier associated with s. If len(s) is 0, ID returns zero.
func (d *Dictionary) ID(s []byte) int {
	d.mu.Lock()
	DB.mu.RLock()
	i := d.put(s)
	d.mu.Unlock()
	return i
}

// SID is like ID but accepts a string instead of a byte slice.
func (d *Dictionary) SID(s string) int {
	return d.ID([]byte(s))
}

func (d *Dictionary) insert(q *dictD, i int, k int) *dictD {
	c := q.c
	if i < c {
		copy(q.d[i+1:], q.d[i:c])
	}
	c++
	q.c = c
	q.d[i] = k
	return q
}

func (d *Dictionary) overflow(p *dictX, q *dictD, pi, i int, k int) {
	l, r := p.siblings(pi)

	if l != nil && l.c < 2*dictKD {
		l.mvL(q, 1)
		d.insert(q, i-1, k)
		p.x[pi-1].k = q.d[0]
		return
	}

	if r != nil && r.c < 2*dictKD {
		if i < 2*dictKD {
			q.mvR(r, 1)
			d.insert(q, i, k)
			p.x[pi].k = r.d[0]
		} else {
			d.insert(r, 0, k)
			p.x[pi].k = k
		}
		return
	}

	d.split(p, q, pi, i, k)
}

func (d *Dictionary) put(k []byte) int {
	pi := -1
	var p *dictX
	q := d.r
	if q == nil {
		DB.mu.RUnlock()
		k := DB.PutBytes(k)
		z := d.insert(&dictD{}, 0, k)
		d.r = z
		return k
	}

	for {
		i, ok := d.find(q, k)
		if ok {
			switch x := q.(type) {
			case *dictX:
				if x.c > 2*dictKX {
					x, i = d.splitX(p, x, pi, i)
				}
				pi = i + 1
				p = x
				q = x.x[i+1].ch
				continue
			case *dictD:
				DB.mu.RUnlock()
				return x.d[i]
			}
		}

		switch x := q.(type) {
		case *dictX:
			if x.c > 2*dictKX {
				x, i = d.splitX(p, x, pi, i)
			}
			pi = i
			p = x
			q = x.x[i].ch
		case *dictD:
			DB.mu.RUnlock()
			k := DB.PutBytes(k)
			switch {
			case x.c < 2*dictKD:
				d.insert(x, i, k)
			default:
				d.overflow(p, x, pi, i, k)
			}
			return k
		}
	}
}

// S returns the byte slice associated with id.
//
// NOTE: The result refers to the dictionary backing storage directly without
// copying and is thus strictly R/O.
func (d *Dictionary) S(id int) []byte {
	return DB.Bytes(id)
}

func (d *Dictionary) split(p *dictX, q *dictD, pi, i int, k int) {
	r := &dictD{}
	copy(r.d[:], q.d[dictKD:2*dictKD])
	for i := range q.d[dictKD:] {
		q.d[dictKD+i] = dictZDE
	}
	q.c = dictKD
	r.c = dictKD
	var done bool
	if i > dictKD {
		done = true
		d.insert(r, i-dictKD, k)
	}
	if pi >= 0 {
		p.insert(pi, r.d[0], r)
	} else {
		d.r = newDictX(q).insert(0, r.d[0], r)
	}
	if done {
		return
	}

	d.insert(q, i, k)
}

func (d *Dictionary) splitX(p *dictX, q *dictX, pi int, i int) (*dictX, int) {
	r := &dictX{}
	copy(r.x[:], q.x[dictKX+1:])
	q.c = dictKX
	r.c = dictKX
	if pi >= 0 {
		p.insert(pi, q.x[dictKX].k, r)
		q.x[dictKX].k = dictZK
		for i := range q.x[dictKX+1:] {
			q.x[dictKX+i+1] = dictZXE
		}

		switch {
		case i < dictKX:
			return q, i
		case i == dictKX:
			return p, pi
		default:
			return r, i - dictKX - 1
		}
	}

	nr := newDictX(q).insert(0, q.x[dictKX].k, r)
	d.r = nr
	q.x[dictKX].k = dictZK
	for i := range q.x[dictKX+1:] {
		q.x[dictKX+i+1] = dictZXE
	}

	switch {
	case i < dictKX:
		return q, i
	case i == dictKX:
		return nr, 0
	default:
		return r, i - dictKX - 1
	}
}
