// Copyright 2015 The XC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xc provides cross language compiler support/utility stuff.
package xc

import (
	"fmt"
	"go/scanner"
	"go/token"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"

	"github.com/cznic/golex/lex"
	"github.com/cznic/strutil"
)

var (
	// DB keeps parser/compiler data shared between stages. See also *MemDB
	// methods.
	DB = newMemDB()

	// Dict collects unique []byte values using the global DB variable.
	// See also *Dictionary methods.
	Dict = newDictionary()

	// FileSet represents a set of source files.
	FileSet = token.NewFileSet()

	// Files is a ready to use FileCentral.
	Files = NewFileCentral()

	// PrintHooks define default strutil.PrettyPrintHooks.
	PrintHooks = strutil.PrettyPrintHooks{
		reflect.TypeOf(lex.Char{}): func(f strutil.Formatter, v interface{}, prefix, suffix string) {
			c := v.(lex.Char)
			suffix = strings.Replace(suffix, "%", "%%", -1)
			f.Format("%s%v: %q"+suffix, prefix, FileSet.Position(c.Pos()), string(c.Rune))
		},
		reflect.TypeOf(token.Pos(0)): func(f strutil.Formatter, v interface{}, prefix, suffix string) {
			p := v.(token.Pos)
			suffix = strings.Replace(suffix, "%", "%%", -1)
			f.Format("%s%v"+suffix, prefix, FileSet.Position(p))
		},
		reflect.TypeOf(Val(0)): func(f strutil.Formatter, v interface{}, prefix, suffix string) {
			suffix = strings.Replace(suffix, "%", "%%", -1)
			f.Format("%s%s"+suffix, prefix, Dict.S(int(v.(Val))))
		},
	}

	workLimiter = newLimiter(2 * runtime.NumCPU())
)

func init() {
	PrintHooks[reflect.TypeOf(Token{})] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		t := v.(Token)
		if !t.Pos().IsValid() {
			return
		}

		h := PrintHooks[reflect.TypeOf(lex.Char{})]
		h(f, t.Char, prefix, "")
		if s := Dict.S(t.Val); len(s) != 0 {
			f.Format(" %q", s)
		}
		f.Format(suffix)
		return
	}
}

// Once manages a single value instantiated only once.
type Once struct {
	m sync.Mutex
	v interface{}
}

func newOnce() *Once {
	o := &Once{}
	o.m.Lock()
	return o
}

func (o *Once) set(v interface{}) {
	if v == nil {
		panic("internal error")
	}

	o.v = v
	o.m.Unlock()
}

// Value returns the value of o. Value may block if o's value is not yet set.
func (o *Once) Value() interface{} {
	o.m.Lock()
	r := o.v
	o.m.Unlock()
	return r
}

// FileCentral provides centralized source file handling. Methods of
// FileCentral are synchonized; multiple goroutines may invoke them
// concurrently.
type FileCentral struct {
	onceMu sync.RWMutex
	onces  map[string]*Once
}

// NewFileCentral returns a newly created FileCentral.
func NewFileCentral() *FileCentral { return &FileCentral{onces: map[string]*Once{}} }

// Once returns the *Once associated with id. If the value of the returned
// object needs to be set, Once starts the set function in a new goroutine to
// obtain the value. Once will panic is the set function returns nil.
func (f *FileCentral) Once(id string, set func() interface{}) *Once {
	f.onceMu.Lock()
	v := f.onces[id]
	if v == nil {
		v = newOnce()
		f.onces[id] = v
		go func() {
			workLimiter.wait()
			defer workLimiter.signal()

			v.set(set())
		}()
	}
	f.onceMu.Unlock()
	return v
}

// Report provides centralized error collecting. Methods of Report are
// synchonized; multiple goroutines may invoke them concurrently.
type Report struct {
	// ErrLimit limits the number of calls to the error reporting methods.
	// After the limit is reached, all errors are reported using log.Print
	// and then log.Fatal() is called with a message about too many errors.
	// To disable error limit set ErrLimit to value less or equal zero.
	// Default value is 10.
	ErrLimit int

	// IgnoreErrors is a testing hook allowing to ignore errors.
	IgnoreErrors bool

	// PanicOnError is a testing hook allowing to fast fail by panicking on
	// first call to any of the error reporting methods.
	PanicOnError bool

	// TraceErrors enable printing errors to stderr as they come, which
	// means they may not be sorted. Intended for testing.
	TraceErrors bool

	errors   scanner.ErrorList
	errorsMu sync.Mutex
}

// NewReport returns a newly created Report.
func NewReport() *Report {
	return &Report{ErrLimit: 10}
}

// ErrPosition reports an error at position.
func (c *Report) ErrPosition(position token.Position, format string, arg ...interface{}) {
	if c.TraceErrors {
		fmt.Fprintf(os.Stderr, "%v: %s\n", position, fmt.Sprintf(format, arg...))
	}

	if c.IgnoreErrors {
		return
	}

	c.errorsMu.Lock()         // X+
	defer c.errorsMu.Unlock() // X-

	if c.PanicOnError {
		panic(fmt.Errorf("%s: %v", position, fmt.Sprintf(format, arg...)))
	}

	c.errors.Add(position, fmt.Sprintf(format, arg...))

	if c.ErrLimit > 0 {
		c.ErrLimit--
		if c.ErrLimit == 0 {
			scanner.PrintError(os.Stderr, c.errors)
			log.Fatalf("too many errors")
		}
	}

}

// Err reports an error at pos.
// goroutines.
func (c *Report) Err(pos token.Pos, format string, arg ...interface{}) {
	c.ErrPosition(FileSet.Position(pos), format, arg...)
}

// ErrChar reports an error for char ch. ErrChar is safe for concurrent use by
// multiple goroutines.
func (c *Report) ErrChar(ch lex.Char, format string, arg ...interface{}) {
	c.Err(ch.Pos(), format, arg...)
}

// ErrTok reports an error for token tok.
func (c *Report) ErrTok(tok Token, format string, arg ...interface{}) {
	c.Err(tok.Pos(), format, arg...)
}

// ClearErrors clears any errors already reported to c and returnes them.
func (c *Report) ClearErrors() error {
	c.errorsMu.Lock()
	defer c.errorsMu.Unlock()

	r := c.errors
	r.Sort()
	c.errors = nil
	return r
}

// Errors returns a go/scanner.ErrorList or nil if there were no errors
// reported so far by the Err* methods.
func (c *Report) Errors(sorted bool) error {
	c.errorsMu.Lock()
	defer c.errorsMu.Unlock()

	if sorted {
		c.errors.Sort()
	}
	if len(c.errors) == 0 {
		return nil
	}

	return c.errors
}

// Val represents an item of a Dict. It prettyPrints as a string Dict.S(theVal).
type Val int

// String implements fmt.Stringer.
func (v Val) String() string { return strutil.PrettyString(v, "", "", PrintHooks) }

type limiter chan struct{}

func newLimiter(n int) limiter { return make(limiter, n) }

func (l limiter) wait() { l <- struct{}{} }

func (l limiter) signal() { <-l }
