// Copyright 2015 The XC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xc

import (
	"encoding/binary"
	"sync"

	"github.com/cznic/mathutil"
)

const (
	maxUvarint  = (64 + 7) / 7
	dbPageShift = 20 //TODO bench tune.
	dbPageSize  = 1 << dbPageShift
	dbPageMask  = dbPageSize - 1
)

// MemDB stores data accessible using a numeric identifier.
type MemDB struct {
	b256   [256]byte
	mu     sync.RWMutex // Access guard.
	nextID int
	pages  [][]byte
}

func newMemDB() *MemDB {
	d := &MemDB{}
	for i := 0; i < 256; i++ {
		d.b256[i] = byte(i)
	}
	return d
}

// Bytes returns the byte slice associated with id. Passing id not obtained by
// PutBytes has undefined behavior. Bytes RLocks the DB to obtain the result.
//
// NOTE: The result refers to the DB backing storage directly without copying
// and is thus strictly R/O.
func (d *MemDB) Bytes(id int) []byte {
	if id == 0 {
		return nil
	}

	if id < 257 {
		return d.b256[id-1 : id:id]
	}

	id -= 257
	off := id & dbPageMask

	d.mu.RLock() // R+

	page := d.pages[id>>dbPageShift]
	n, l := binary.Uvarint(page[off:])
	off += l
	h := off + int(n)
	r := page[off:h:h]

	d.mu.RUnlock() // R-

	return r
}

// bytesUnlocked is like Bytes but it does not RLock the DB.
func (d *MemDB) bytesUnlocked(id int) []byte {
	if id == 0 {
		return nil
	}

	if id < 257 {
		return d.b256[id-1 : id:id]
	}

	id -= 257
	off := id & dbPageMask

	page := d.pages[id>>dbPageShift]
	n, l := binary.Uvarint(page[off:])
	off += l
	r := page[off : off+int(n)]

	return r
}

// Cap reports the current DB capacity.
func (d *MemDB) Cap() int {
	d.mu.RLock() // R+

	r := 0
	for _, v := range d.pages {
		r += cap(v)
	}

	d.mu.RUnlock() // R-

	return r
}

// Len reports the size of the DB used by data.
func (d *MemDB) Len() int {
	d.mu.RLock() // R+

	r := 0
	for _, v := range d.pages {
		r += len(v)
	}

	d.mu.RUnlock() // R-

	return r
}

// PutBytes stores b in the DB and returns its id.  Zero length byte slices are
// guaranteed to return zero ID. PutBytes Locks the DB before updating it.
func (d *MemDB) PutBytes(b []byte) int {
	if len(b) == 0 {
		return 0
	}

	if len(b) == 1 {
		return int(b[0]) + 1
	}

	d.mu.Lock() // W+

	id := d.nextID
	pi := id >> dbPageShift
	if pi < len(d.pages) {
		p := d.pages[pi]
		off := id & dbPageMask
		if n := cap(p) - off - maxUvarint; n >= len(b) {
			p = p[:cap(p)]
			l := binary.PutUvarint(p[off:], uint64(len(b)))
			copy(p[off+l:], b)
			n = l + len(b)
			d.pages[pi] = p[:off+n]
			d.nextID += n

			d.mu.Unlock() // W-

			return id + 257
		}

		pi++
	}

	p := make([]byte, mathutil.Max(dbPageSize, maxUvarint+len(b)))
	p = p[:binary.PutUvarint(p, uint64(len(b)))]
	p = append(p, b...)
	d.pages = append(d.pages, p)
	id = pi << dbPageShift
	d.nextID = id + mathutil.Min(dbPageSize, len(p))

	d.mu.Unlock() // W-

	return id + 257
}
