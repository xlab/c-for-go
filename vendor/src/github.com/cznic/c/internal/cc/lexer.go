// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//TODO recycle []Token
//TODO recycle lexer0, lexer, ppLexer

package cc

import (
	"bufio"
	"bytes"
	"go/token"
	"io"
	"sync"
	"unicode/utf8"

	"github.com/cznic/c/internal/xc"
	"github.com/cznic/golex/lex"
)

const (
	commentNotClosed = "comment not closed"
)

var (
	lexer0Pool = sync.Pool{New: func() interface{} { return &lexer0{} }}

	noTypedefNameAfter = map[rune]bool{
		BOOL:        true,
		CHAR:        true,
		COMPLEX:     true,
		DOUBLE:      true,
		ENUM:        true,
		FLOAT:       true,
		GOTO:        true,
		INT:         true,
		LONG:        true,
		SHORT:       true,
		SIGNED:      true,
		STRUCT:      true,
		TYPEDEFNAME: true,
		UNION:       true,
		UNSIGNED:    true,
		VOID:        true,
	}

	// R/O
	zLexer0 lexer0
)

func newToken(pos token.Pos, r rune, val int) xc.Token {
	return xc.Token{Char: lex.NewChar(pos, r), Val: val}
} //TODO use instead of Token{...}

func toC(t xc.Token) xc.Token {
	if t.Rune != IDENTIFIER {
		return t
	}

	if x, ok := cwords[t.Val]; ok {
		t.Rune = x
	}

	return t
}

type lexer0 struct {
	c           lex.Char    // Lookahead if .IsValid.
	example     interface{} // AST example hook.
	exampleRule int         // AST example hook.
	file        *token.File //
	in          []lex.Char  // Lexeme collector.
	last        lex.Char    //
	m           int         // Longest match marker.
	sc          int         // Starting condition.
	scs         []int       // SC stack.
	src         scanner     // Input pipeline.
	un          []lex.Char  // Unget buffer.
}

func newLexer0(src scanner) *lexer0 {
	lx := lexer0Pool.Get().(*lexer0)
	in, scs, un := lx.in[:0], lx.scs[:0], lx.un[:0]
	*lx = zLexer0
	lx.in, lx.scs, lx.un = in, scs, un
	lx.example = -1
	lx.src = src
	return lx
}

func (l *lexer0) abort() lex.Char {
	switch n := len(l.in); n {
	case 0: // [] z
		c := l.c
		l.next()
		return c
	case 1: // [a] z
		return l.in[0]
	default: // [a, b, ...], z
		c := l.in[0] // a
		l.unget(l.c) // z
		for i := n - 1; i > 1; i-- {
			l.unget(l.in[i]) // ...
		}
		l.c = l.in[1] // b
		l.in = l.in[:1]
		return c
	}
}

func (l *lexer0) char(r rune) lex.Char { return lex.NewChar(l.in[0].Pos(), r) }

func (l *lexer0) close() {
	if x, ok := l.src.(*ppLexer); ok {
		x.close()
	}
	lexer0Pool.Put(l)
}

func (l *lexer0) enter() int {
	if !l.c.IsValid() {
		l.next()
	}
	return rune2class(l.c.Rune)
}

// Implements yyLexer.
func (l *lexer0) Error(msg string) {
	compilation.Err(l.last.Pos(), "%s", msg)
}

// Implements scanner.
func (l *lexer0) error() error {
	return nil
}

func (l *lexer0) next() int {
	if c := l.c; c.IsValid() {
		l.in = append(l.in, c)
	}
	if n := len(l.un); n != 0 {
		l.c = l.un[n-1]
		l.un = l.un[:n-1]
		return rune2class(l.c.Rune)
	}

	l.c = l.src.scan()
	if rune2class(l.c.Rune) == ccError {
		compilation.Err(l.c.Pos(), l.src.error().Error())
	}
	return rune2class(l.c.Rune)
}

func (l *lexer0) push(sc int) {
	l.scs = append(l.scs, l.sc)
	l.sc = sc
}

func (l *lexer0) pop() {
	n := len(l.scs)
	l.sc = l.scs[n-1]
	l.scs = l.scs[:n-1]
}

func (l *lexer0) rule0() int {
	l.last = l.c
	l.m = -1
	l.in = l.in[:0]
	return rune2class(l.c.Rune)
}

func (l *lexer0) unget(c ...lex.Char) {
	for _, c := range c {
		l.un = append(l.un, c)
	}
	l.c = lex.Char{}
}

type lexer struct {
	ast              *PreprocessingFile  //
	ch               chan []xc.Token     //
	commentPos0      token.Pos           // /*-style comment start position.
	compoundStmt     int                 // Nesting level.
	constExpr        *ConstantExpression //
	constExprTok     int                 //
	constExprToks    []xc.Token          //
	cpp              func([]xc.Token)    //
	declaratorSerial int                 //
	encBuf           []byte              //
	fileScope        *Bindings           //
	lexer0                               //
	line             []xc.Token          //
	macroArg         []xc.Token          //
	macroArgs        [][]xc.Token        //
	macroArgsCommas  []xc.Token          //
	macroArgsLevel   int                 // Nesting.
	macroArgsTok     int                 //
	macroArgsToks    []xc.Token          //
	macroMore        int                 //
	pgl              **GroupList         // Multiline macro invocation hook.
	predefines       bool                //
	prev             rune                //
	runeEnc          []byte              //
	scope            *Bindings           //
	seenPrefix       bool                // Seen one of PREPROCESSING_FILE, CONSTANT_EXPRESSION, TRANSLATION_UNIT, MACRO_ARGS.
	state            int                 //
	tok2             xc.Token            //
	tok2r            int                 //
	tokValLen        int                 //
	tu               *TranslationUnit    //
	zipToks          []xc.Token          //
}

func newTULexer() *lexer {
	scope := newBindings(ScopeFile, nil)
	return &lexer{
		fileScope: scope,
		scope:     scope,
		tu:        &TranslationUnit{Declarations: scope},
	}
}

func newLexer(src scanner, trigraphs bool) *lexer {
	lx := newTULexer()
	lx.encBuf = make([]byte, 30)
	lx.lexer0 = *newLexer0(newPPLexer(src, trigraphs))
	lx.runeEnc = make([]byte, 4)
	return lx
}

func (l *lexer) pushScope(typ Scope) *Bindings {
	l.scope = newBindings(typ, l.scope)
	return l.scope
}

func (l *lexer) popScope(t xc.Token) *Bindings {
	if l.scope == nil {
		compilation.ErrTok(t, "cannot pop scope")
		return nil
	}

	p := l.scope.parent
	if p == nil {
		compilation.ErrTok(t, "cannot pop scope")
		return nil
	}

	old := l.scope
	l.scope = p
	return old
}

func (l *lexer) close() { l.lexer0.close() }

func (l *lexer) directive(v rune) lex.Char {
	l.pop()
	return l.char(v)
}

// Implements yyLexerEx
func (l *lexer) Reduced(rule, state int, lval *yySymType) (stop bool) {
	if n := l.exampleRule; n >= 0 && rule != n {
		return false
	}

	switch x := lval.item.(type) {
	case interface {
		fragment() interface{}
	}:
		l.example = x.fragment()
	default:
		l.example = x
	}
	return true
}

func (l *lexer) scan() lex.Char {
	c := l.scan0()
	cls := rune2class(c.Rune)
	if cls == ccEOF && l.prev != '\n' {
		c = lex.NewChar(c.Pos(), '\n')
	}
	l.prev = c.Rune
	return c
}

func (l *lexer) token(lval *yySymType) (y int) {
	var t xc.Token
	t.Char = l.scan()
	r := t.Rune
	l.tokValLen = 0
	if tokHasVal[r] {
		buf := l.encBuf[:0]
		for _, c := range l.in {
			if n, m := len(buf), cap(buf); m-n >= 4 {
				buf = buf[:n+utf8.EncodeRune(buf[n:m], c.Rune)]
				continue
			}

			n := utf8.EncodeRune(l.runeEnc, c.Rune)
			buf = append(buf, l.runeEnc[:n]...)
		}
		l.tokValLen = len(buf)
		if r == IDENTIFIER_LPAREN {
			buf = buf[:len(buf)-1] // Remove trailing '('
		}
		t.Val = dict.ID(buf)
		l.encBuf = buf[:0]
	}
	lval.Token = t
	if c := rune2class(t.Rune); c != ccEOF {
		if c == '\n' {
			l.state = lsPPBOL
		}
		return int(r)
	}

	return -1
}

const (
	lsPreprocess = iota //TODO rename to lsZero or something like that
	lsPPBOL
	lsPPDefine
	lsPPTokens
	lsPPVerbatim
	lsPPSeekRParen
	lsExample
	lsConstExpr0
	lsConstExpr
	lsMacroArgs0
	lsMacroArgs
	lsConstExprExample
	lsTranslationUnit0
	lsTranslationUnit
)

// Implements yyLexer
func (l *lexer) Lex(lval *yySymType) (r int) {
	switch l.state {
	case lsTranslationUnit0:
		l.state = lsTranslationUnit
		lval.Token = xc.Token{Char: lex.NewChar(1, TRANSLATION_UNIT)}
		return TRANSLATION_UNIT
	case lsTranslationUnit:
	readLine:
		if len(l.line) != 0 {
			t := toC(l.line[0])
			l.line = l.line[1:]
			if t.Rune == IDENTIFIER && !noTypedefNameAfter[l.last.Rune] {
				if l.scope.IsTypedefName(t.Val) {
					t.Char.Rune = TYPEDEFNAME
					lval.Token = t
				}
			}
			lval.Token = t
			l.last = t.Char
			return int(t.Rune)
		}

		var line []xc.Token
		for {
			ln, ok := <-l.ch
			if !ok {
				lval.Token = xc.Token{Char: lex.NewChar(l.last.Pos(), 0)}
				return -1
			}

			if len(ln) == 0 {
				continue
			}

			switch len(line) {
			case 0:
				line = ln
			default:
				line = append(line, ln...)
			}
			if r := line[len(line)-1].Rune; r != LONGSTRINGLITERAL && r != STRINGLITERAL {
				break
			}
		}

		// Preproc phase 6. Adjacent string literal tokens are concatenated.
		w := 0
		for r := 0; r < len(line); r++ {
			v := line[r]
			switch v.Rune {
			case STRINGLITERAL, LONGSTRINGLITERAL:
				to := r
				for to < len(line)-1 && line[to+1].Rune == STRINGLITERAL {
					to++
				}
				if to == r {
					line[w] = v
					w++
					break
				}

				var buf bytes.Buffer
				s := dict.S(v.Val)
				s = s[:len(s)-1] // Remove trailing "
				buf.Write(s)
				for i := r + 1; i <= to; i++ {
					s = dict.S(line[i].Val)
					s = s[1 : len(s)-1] // Remove leading and trailing "
					buf.Write(s)
				}
				r = to
				buf.WriteByte('"')
				v.Val = dict.ID(buf.Bytes())
				fallthrough
			default:
				line[w] = v
				w++
			}
		}
		l.line = line[:w]
		if l.cpp != nil {
			l.cpp(l.line)
		}
		goto readLine
	case lsConstExpr: // Argument of #{if,elif}.
		if i, s := l.constExprTok, l.constExprToks; i < len(s) {
			t := s[i]
			l.last = t.Char
			r := t.Rune
			l.constExprTok++
			lval.Token = t
			return int(r)
		}

		return -1
	case lsMacroArgs0:
		l.state = lsMacroArgs
		return MACRO_ARGS
	case lsConstExpr0:
		l.state = lsConstExpr
		return CONSTANT_EXPRESSION
	case lsPPVerbatim:
		return l.token(lval)
	case lsMacroArgs:
		for {
			if i, s := l.macroArgsTok, l.macroArgsToks; i < len(s) {
				t := s[i]
				l.last = t.Char
				r := t.Rune
				l.macroArgsTok++
				switch r {
				case '(':
					if l.macroArgsLevel != 0 {
						r = MACRO_ARG
					}
					l.macroArgsLevel++
				case ')':
					l.macroArgsLevel--
					if l.macroArgsLevel != 0 {
						r = MACRO_ARG
						break
					}

					l.macroArgsToks = nil
				case ',':
					l.macroArgsCommas = append(l.macroArgsCommas, t)
					if l.macroArgsLevel != 1 {
						r = MACRO_ARG
					}
				default:
					r = MACRO_ARG
				}
				lval.Token = t
				return int(r)
			}

			if l.macroArgsLevel == 0 || l.pgl == nil {
				return -1
			}

			g := *l.pgl
			nx := g.GroupList
			if nx == nil {
				return -1
			}

			*l.pgl = nx
			group := nx.GroupPart
			if group.PpTokenList == 0 {
				return -1
			}

			l.macroArgsToks = db.tokens(group.PpTokenList)
			l.macroMore += l.macroArgsTok
			l.macroArgsTok = 0
		}
	case lsPPDefine:
		r = l.token(lval)
		l.sc = scINITIAL
		switch r {
		case IDENTIFIER:
			l.state = lsPPTokens
			return r
		case IDENTIFIER_LPAREN:
			l.state = lsPPSeekRParen
			return r
		default:
			l.state = lsPPVerbatim
			return r
		}
	case lsPPTokens:
		r = l.token(lval)
		l.sc = scINITIAL
		switch r {
		case '\n':
			return r
		default:
			return PPOTHER
		}
	case lsPPSeekRParen:
		r = l.token(lval)
		switch r {
		case ')':
			l.state = lsPPTokens
			return r
		default:
			return r
		}
	case lsExample:
	again:
		r = l.token(lval)
		t := lval.Token
		if r == '\n' {
			goto again
		}

		if r != IDENTIFIER {
			return r
		}

		if t.Val == idDefined {
			t.Rune = DEFINED
			lval.Token = t
			return DEFINED
		}

		t = toC(t)
		lval.Token = t
		if t.Rune == IDENTIFIER && !noTypedefNameAfter[l.last.Rune] {
			if l.scope.IsTypedefName(t.Val) {
				t.Char.Rune = TYPEDEFNAME
				lval.Token = t
			}
		}
		return int(t.Rune)
	case lsPPBOL:
		switch r = l.token(lval); r {
		case '\n', -1:
			return r
		case '#':
			var lval2 yySymType
			l.push(scDIRECTIVE)
			switch r2 := l.token(&lval2); r2 {
			case PPASSERT, PPUNASSERT:
				fallthrough //TODO Make this obsolete gcc extension an option.
			case PPIDENT:
				fallthrough //TODO Make this unofficial gcc extension an option.
			case PPWARNING:
				fallthrough //TODO Make this gcc extension an option.
			case PPIF, PPERROR, PPELIF, PPPRAGMA, PPLINE:
				lval.Token = lval2.Token
				l.sc = scINITIAL
				l.state = lsPPTokens
				return r2
			case PPINCLUDE_NEXT, PPIMPORT:
				fallthrough //TODO Make this gcc extension an option.
			case PPINCLUDE:
				l.sc = scHEADER
				lval.Token = lval2.Token
				l.state = lsPPTokens
				return r2
			case PPDEFINE:
				l.sc = scDEFINE
				lval.Token = lval2.Token
				l.state = lsPPDefine
				return r2
			case PPIFDEF, PPIFNDEF, PPELSE, PPUNDEF:
				lval.Token = lval2.Token
				l.state = lsPPVerbatim
				return r2
			case PPENDIF:
				lval.Token = lval2.Token
				l.state = lsPPTokens //TODO Make this gcc extension an option.
				return r2
			case '\n':
				return PPHASH_NL
			default:
				lval.Token = lval2.Token
				l.sc = scINITIAL
				l.state = lsPPTokens
				return PPNONDIRECTIVE
			}
		default:
			l.state = lsPPTokens
			return PPOTHER
		}
	case lsPreprocess:
		switch r = l.token(lval); r {
		case CONSTANT_EXPRESSION:
			if !l.seenPrefix {
				l.state = lsExample
			}

			l.seenPrefix = true
			return r
		case PREPROCESSING_FILE:
			if !l.seenPrefix {
				l.state = lsPPBOL
			}

			l.seenPrefix = true
			return r
		case TRANSLATION_UNIT:
			if !l.seenPrefix {
				l.state = lsExample
			}

			l.seenPrefix = true
			return r
		default:
			panic("internal error")
		}
	default:
		panic("internal error")
	}
}

func (l *lexer) parseConstExpr(toks []xc.Token) (r interface{}, _ bool) {
	l.constExprToks = toks
	l.constExprTok = 0
	l.state = lsConstExpr0
	if yyParse(l) != 0 {
		return intT(0), false
	}

	return l.constExpr.eval(), true
}

func (l *lexer) parseMacroArgs(toks []xc.Token, pgl **GroupList) (int, [][]xc.Token) {
	l.macroArg = nil
	l.macroArgs = nil
	l.macroArgsCommas = nil
	l.macroArgsLevel = 0
	l.macroArgsTok = 0
	l.macroArgsToks = toks
	l.macroMore = 0
	l.pgl = pgl
	l.state = lsMacroArgs0
	if yyParse(l) != 0 {
		return l.macroArgsTok, nil
	}

	return l.macroArgsTok + l.macroMore, l.macroArgs
}

type ppLexer struct {
	lexer0
}

func newPPLexer(src scanner, trigraphs bool) *ppLexer {
	lx := &ppLexer{*newLexer0(src)}
	if trigraphs {
		lx.sc = scTRIGRAPHS
	}
	return lx
}

func (l *ppLexer) close() { l.lexer0.close() }

type scanner interface {
	error() error
	scan() lex.Char
}

type utf8src struct {
	err    error
	file   *token.File
	off    int
	src    io.RuneReader
	sticky lex.Char
}

func newUtf8src(src io.Reader, file *token.File) *utf8src {
	var rr io.RuneReader
	switch x := src.(type) {
	case io.RuneReader:
		rr = x
	default:
		rr = bufio.NewReader(src)
	}
	return &utf8src{
		file: file,
		src:  rr,
	}
}

// Implements scanner.
func (s *utf8src) error() error { return s.err }

// Implements scanner.
func (s *utf8src) scan() lex.Char {
	if c := s.sticky; c.IsValid() {
		return c
	}

	r, sz, err := s.src.ReadRune()
	pos := s.file.Pos(s.off)
	switch {
	case err == nil:
		s.off += sz
		if r == '\n' {
			s.file.AddLine(s.off)
		}
		return lex.NewChar(pos, r)
	case err == io.EOF:
		s.sticky = lex.NewChar(pos, runeEOF)
		return s.sticky
	default:
		s.err = err
		s.sticky = lex.NewChar(pos, runeEOF)
		return lex.NewChar(pos, runeError)
	}
}

type tuLexer struct {
	ppFiles []*PreprocessingFile
}
