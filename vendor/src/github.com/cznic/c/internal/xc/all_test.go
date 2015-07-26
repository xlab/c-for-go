// Copyright 2015 The XC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xc

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/cznic/mathutil"
)

func caller(s string, va ...interface{}) {
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "\tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
}

func TODO(...interface{}) string {
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("TODO: %s:%d:\n", path.Base(fn), fl)
}

func use(...interface{}) {}

// ============================================================================

func newRng() *mathutil.FC32 {
	r, err := mathutil.NewFC32(0, math.MaxInt32, true)
	if err != nil {
		panic(err)
	}

	r.Seed(42)
	return r
}

func TestDBBytes(t *testing.T) {
	const n = 29e4
	rng := newRng()
	db := newMemDB()
	exp := map[int][]byte{}
	var buf []byte
	for i := 0; i < n; i++ {
		buf = buf[:0]
		for i, l := 0, rng.Next()%31; i < l; i++ {
			buf = append(buf, byte(rng.Next()))
		}
		id := db.PutBytes(buf)
		switch {
		case len(buf) == 0:
			if g, e := id, 0; g != e {
				t.Fatal(i, g, e)
			}
		case len(buf) == 1:
			if g, e := id, int(buf[0])+1; g != e {
				t.Fatal(i, g, e)
			}
		default:
			if _, ok := exp[id]; ok {
				t.Fatal(i, id)

			}
		}

		exp[id] = append([]byte(nil), buf...)
	}
	for id, e := range exp {
		g := db.Bytes(id)
		if id == 0 {
			if g != nil {
				t.Fatalf("|% x|", g)
			}
			continue
		}

		if !bytes.Equal(g, e) {
			t.Fatalf("%v |% x| |% x|", id, g, e)
		}
	}
	db.mu.RLock()
	defer db.mu.RUnlock()
	for id, e := range exp {
		g := db.bytesUnlocked(id)
		if id == 0 {
			if g != nil {
				t.Fatalf("|% x|", g)
			}
			continue
		}

		if !bytes.Equal(g, e) {
			t.Fatalf("%v |% x| |% x|", id, g, e)
		}
	}
	l, c := db.Len(), db.Cap()
	w := c - l
	t.Logf("len %v, cap %v, wasted %v (%.1f%%)", l, c, w, 100*float64(w)/float64(c))
}

func TestDict(t *testing.T) {
	const n = 15e4

	dict := newDictionary()

	rng := newRng()
	m := map[string]int{}
	for i := 0; i < n; i++ {
		x := rng.Next()
		s := fmt.Sprintf("%d.%d", x, x)
		m[s] = dict.ID([]byte(s))
	}
	for k, v := range m {
		b := dict.S(v)
		if g, e := string(b), k; g != e {
			t.Fatalf("%q %q", g, e)
		}
	}
	for k, v := range m {
		if g, e := dict.ID([]byte(k)), v; g != e {
			t.Fatal(g, e)
		}
	}
}
