// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"bufio"
	"bytes"
	"fmt"
	goscanner "go/scanner"
	"go/token"
	"hash/fnv"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/cznic/c/internal/xc"
	"github.com/cznic/golex/lex"
	"github.com/cznic/mathutil"
)

const (
	testCorpusDir   = "testdata"
	testCorpusName  = "corpus"
	testFileList    = "files"
	testLexiconName = "lexicon"
)

var (
	onceCorpus    sync.Once
	testCorpus    []byte
	testCorpusNLs []int
)

func corpus() ([]byte, []int, error) {
	var err error
	onceCorpus.Do(func() {
		var buf []byte
		if buf, err = ioutil.ReadFile(filepath.Join(testCorpusDir, testCorpusName)); err != nil {
			return
		}

		testCorpus = buf
		testCorpusNLs = make([]int, 0, len(buf)/3)
		for i, v := range testCorpus {
			if v == '\n' {
				testCorpusNLs = append(testCorpusNLs, i)
			}
		}
	})
	return testCorpus, testCorpusNLs, err
}

type errStop struct{}

func (errStop) Error() string { return "limit reached" }

func TestDevMakeCorpus(t *testing.T) {
	if !*oDev {
		t.Log("not enabled")
		return
	}

	limit := *oMaxToks

	f0, err := os.OpenFile(filepath.Join(testCorpusDir, testCorpusName), os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}

	f := bufio.NewWriter(f0)

	defer func() {
		if err := f.Flush(); err != nil {
			t.Error(err)
		}
		if err := f0.Close(); err != nil {
			t.Error(err)
		}
	}()

	lst0, err := os.Create(filepath.Join(testCorpusDir, testFileList))
	if err != nil {
		t.Fatal(err)
	}

	lst := bufio.NewWriter(lst0)

	defer func() {
		if err := lst.Flush(); err != nil {
			t.Error(err)
		}
		if err := lst0.Close(); err != nil {
			t.Error(err)
		}
	}()

	var files, used, tokens, fileSize, usedSize, size, nls, bigSz int
	var biggestSz string
	m := map[uint32]struct{}{}
	dict := map[string]int{}
	if err := filepath.Walk(testCorpusDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
		}

		if info.IsDir() {
			return nil
		}

		if ext := filepath.Ext(path); ext != ".c" && ext != ".h" {
			return nil
		}

		files++
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		fileSize += len(buf)
		h := fnv.New32a()
		h.Write(buf)
		sum := h.Sum32()
		if _, ok := m[sum]; ok {
			return nil
		}

		used++
		usedSize += len(buf)
		m[sum] = struct{}{}
		if _, err := lst.WriteString(path + "\n"); err != nil {
			return err
		}

		if sz := len(buf); sz > bigSz {
			bigSz = sz
			biggestSz = path
		}
		file := fileset.AddFile(path, -1, len(buf))
		lx := newLexer(newUtf8src(bytes.NewBuffer(buf), file), true)
		lx.unget(lex.NewChar(0, PREPROCESSING_FILE))
		var lval yySymType
		var runes []rune
		for {
			c := lx.Lex(&lval)
			if c <= 0 {
				err := lx.error()
				if err != nil {
					if strings.HasSuffix(err.Error(), commentNotClosed) { // gcc tests
						err = nil
					}
				}
				return err
			}

			runes = runes[:0]
			for _, c := range lx.in {
				runes = append(runes, c.Rune)
			}
			s := string(runes)
			if s == "\n" {
				nls++
				continue
			}

			dict[s]++
			tokens++
			size += len(s)
			if _, err := f.WriteString(s); err != nil {
				return err
			}

			if err := f.WriteByte('\n'); err != nil {
				return err
			}

			limit--
			if limit == 0 {
				return errStop{}
			}
		}
	}); err != nil {
		switch x := err.(type) {
		case goscanner.ErrorList:
			for _, v := range x {
				t.Error(v)
			}
		case errStop:
			// nop
		default:
			t.Fatal(err)
		}
	}
	a := make([]string, 0, len(dict))
	dsz := 0
	for k := range dict {
		dsz += len(k)
		a = append(a, k)
	}
	sort.Strings(a)

	g0, err := os.Create(filepath.Join(testCorpusDir, testLexiconName))
	if err != nil {
		t.Fatal(err)
	}

	g := bufio.NewWriter(g0)
	defer func() {
		if err := g.Flush(); err != nil {
			t.Error(err)
		}
		if err := g0.Close(); err != nil {
			t.Error(err)
		}
	}()
	for _, v := range a {
		if _, err := g.WriteString(v); err != nil {
			t.Fatal(err)
		}
		if err := g.WriteByte('\n'); err != nil {
			t.Fatal(err)
		}
	}

	t.Logf(
		`files %v (%v bytes), used %v (%v bytes)
tokens %v, bytes %v
dictionary tokens %v, bytes %v
biggest %d: %s`,
		files, fileSize, used, usedSize, tokens, size, len(a), dsz, bigSz, biggestSz,
	)

	a = a[:0]
	for k, v := range dict {
		if len(k) == 1 {
			a = append(a, fmt.Sprintf("%020d%s", v, k))
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
	rt := 0
	ft := float64(tokens)
	t.Logf("Top %d (len == 1)", len(a))
	for i, v := range a {
		n, err := strconv.Atoi(v[:20])
		if err != nil {
			t.Fatal(err)
		}

		rt += n
		frt := float64(rt)
		t.Logf("%4d: %10dx (%4.1f%%, run %10d, %4.1f%%) %v", i+1, n, 100*float64(n)/ft, rt, 100*frt/ft, v[20:])
	}

	a = a[:0]
	for k, v := range dict {
		if len(k) != 1 {
			a = append(a, fmt.Sprintf("%020d%s", v, k))
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
	rt = 0
	a = a[:mathutil.Min(256, len(a))]
	t.Logf("Top %d (len > 1)", len(a))
	for i, v := range a {
		n, err := strconv.Atoi(v[:20])
		if err != nil {
			t.Fatal(err)
		}

		rt += n
		frt := float64(rt)
		t.Logf("%4d: %10dx (%4.1f%%, run %10d, %4.1f%%) %v", i+1, n, 100*float64(n)/ft, rt, 100*frt/ft, v[20:])
	}

	a = a[:0]
	for k, v := range dict {
		a = append(a, fmt.Sprintf("%020d%s", v, k))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
	rt = 0
	a = a[:mathutil.Min(256, len(a))]
	t.Logf("Top %d (any len)", len(a))
	for i, v := range a {
		n, err := strconv.Atoi(v[:20])
		if err != nil {
			t.Fatal(err)
		}

		rt += n
		frt := float64(rt)
		t.Logf("%4d: %10dx (%4.1f%%, run %10d, %4.1f%%) %v", i+1, n, 100*float64(n)/ft, rt, 100*frt/ft, v[20:])
	}
	t.Logf("New line tokens (line count): %v", nls)
	return
}

func isSeqPosition(lpos *token.Position, np token.Position) bool {
	if !lpos.IsValid() || lpos.Filename != np.Filename || np.Line-lpos.Line > 1 {
		*lpos = np
		return false
	}

	*lpos = np
	return true
}

func testDevParse(t *testing.T, predefine string, src []string, opts ...Opt) *TranslationUnit {
	sysIncludePaths = nil
	includePaths = nil
	tu, err := Parse(predefine, src, model, opts...)
	if err != nil {
		t.Fatal(errString(err))
	}

	if err := compilation.Errors(true); err != nil {
		t.Fatal(errString(err))
	}
	return tu
}

func TestDevParseTmp(t *testing.T) {
	if !*oDev {
		t.Log("not enabled")
		return
	}

	logcpp, err := os.Create("logcpp.c")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Fprintf(logcpp, "// +build ignore\n")
	var lpos token.Position
	tu := testDevParse(
		t,
		`
#define __STDC_HOSTED__ 1
#define __STDC_VERSION__ 199901L
#define __STDC__ 1
		`,
		[]string{"tmp.c"},
		[]Opt{
			Cpp(func(l []xc.Token) {
				if len(l) == 0 {
					return
				}

				p0 := xc.FileSet.Position(l[0].Pos())
				isSeq := isSeqPosition(&lpos, p0)
				if !isSeq {
					fmt.Fprintln(logcpp)
				}
				for _, v := range l {
					fmt.Fprintf(logcpp, "%s ", TokSrc(v))
				}
				if !isSeq {
					fmt.Fprintf(logcpp, "\t// %s", p0)
				}
				fmt.Fprintln(logcpp)
			}),
		}...,
	)
	d := tu.Declarations
	tu.Declarations = nil
	t.Log(PrettyString(tu))
	t.Log(PrettyString(d))
	b, _ := d.Lookup(NSIdentifiers, dict.SID("f"))
	if b != nil {
		t.Log(PrettyString(b.Node))
		d := b.Node.(*Declarator)
		t.Log(PrettyString(d.DirectDeclarator.specifier))
	}
}
