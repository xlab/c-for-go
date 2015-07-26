// Copyright 2015 The XC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xc provides cross language compiler support.
package xc

import (
	"fmt"
	"go/scanner"
	"go/token"
	"log"
	"os"
	"reflect"
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

	// ErrLimit limits the number of calls to Compilation.Err(). After the
	// limit is reached, all errors are reported using log.Print and the
	// log.Fatal() is called with a message about too many errors. To
	// disable error limit set ErrLimit to value less or equal zero.
	ErrLimit = 10

	// FileSet represents a set of source files.
	FileSet = token.NewFileSet()

	// IgnoreErrors is a testing hook allowing to ignore errors reported to
	// Compilation.Err.
	IgnoreErrors bool

	// PanicOnError is a testing hook allowing to fast fail by panicking on
	// first call of Compilation.Err.
	PanicOnError bool

	// TraceErrors enable printing errors to stderr as they come, which
	// means they may not be sorted. Intended for testing.
	TraceErrors bool

	// Compilation provides centralized source files management and error
	// handling. See also *Compiler methods.
	Compilation = newCompiler()

	// Default strutil.PrettyPrintHooks.
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
	o.v = v
	o.m.Unlock()
}

// Value returns the value of o. Value possibly waits for Set to be called
// before returning.
func (o *Once) Value() interface{} {
	o.m.Lock()
	r := o.v
	o.m.Unlock()
	return r
}

// Compiler provides centralized error collecting and source file
// handling.
type Compiler struct {
	errors   scanner.ErrorList
	errorsMu sync.Mutex
	onceMu   sync.RWMutex
	onces    map[string]*Once
}

func newCompiler() *Compiler {
	return &Compiler{onces: map[string]*Once{}}
}

// ErrPosition reports an error at position. ErrPosition is safe for concurrent
// use by multiple goroutines.
func (c *Compiler) ErrPosition(position token.Position, format string, arg ...interface{}) {
	if TraceErrors {
		fmt.Fprintf(os.Stderr, "%v: %s\n", position, fmt.Sprintf(format, arg...))
	}

	if IgnoreErrors {
		return
	}

	c.errorsMu.Lock() // X+

	c.errors.Add(position, fmt.Sprintf(format, arg...))

	if ErrLimit > 0 {
		ErrLimit--
		if ErrLimit == 0 {
			scanner.PrintError(os.Stderr, c.errors)
			log.Fatalf("too many errors")
		}
	}

	c.errorsMu.Unlock() // X-

	if PanicOnError {
		panic(c.Errors(true))
	}
}

// Err reports an error at pos. Err is safe for concurrent use by multiple
// goroutines.
func (c *Compiler) Err(pos token.Pos, format string, arg ...interface{}) {
	c.ErrPosition(FileSet.Position(pos), format, arg...)
}

// ErrChar reports an error for char ch. ErrChar is safe for concurrent use by
// multiple goroutines.
func (c *Compiler) ErrChar(ch lex.Char, format string, arg ...interface{}) {
	c.Err(ch.Pos(), format, arg...)
}

// ErrTok reports an error for token tok. ErrTok is safe for concurrent use by
// multiple goroutines.
func (c *Compiler) ErrTok(tok Token, format string, arg ...interface{}) {
	c.Err(tok.Pos(), format, arg...)
}

// ClearErrors clears any errors already reported to c and returnes them.
// ClearError is safe for concurrent use by multiple goroutines.
func (c *Compiler) ClearErrors() error {
	c.errorsMu.Lock()
	defer c.errorsMu.Unlock()

	r := c.errors
	c.errors = nil
	return r
}

// Errors returns a go/scanner.ErrorList or nil if there were no errors
// reported so far by the Err* methods. Error is safe for concurrent use by
// multiple goroutines.
func (c *Compiler) Errors(sorted bool) error {
	c.errorsMu.Lock()
	defer c.errorsMu.Unlock()

	if len(c.errors) == 0 {
		return nil
	}

	if sorted {
		c.errors.Sort()
	}
	return c.errors
}

// Once returns the *Once associated with id. Once is safe for concurrent use
// by multiple goroutines. If the value of the returned object needs to be set,
// Once starts set in a new goroutine to obtain the value.
func (c *Compiler) Once(id string, set func() interface{}) *Once {
	c.onceMu.Lock()
	v := c.onces[id]
	if v == nil {
		v = newOnce()
		c.onces[id] = v
		go func() {
			v.set(set())
		}()
	}
	c.onceMu.Unlock()
	return v
}

// Val represents an item of a Dict. It prettyPrints as a string Dict.S(theVal).
type Val int

// String implements fmt.Stringer.
func (v Val) String() string { return strutil.PrettyString(v, "", "", PrintHooks) }
