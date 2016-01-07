// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"fmt"
	"go/token"
	"io"

	"github.com/cznic/golex/lex"
	"github.com/xlab/c/xc"
)

// Lexer state
const (
	lsZero             = iota
	lsBOL              // Preprocessor: Beginning of line.
	lsDefine           // Preprocessor: Seen ^#define.
	lsSeekRParen       // Preprocessor: Seen ^#define identifier(
	lsTokens           // Preprocessor: Convert anything to PPOTHER until EOL.
	lsUndef            // Preprocessor: Seen ^#undef.
	lsConstExpr0       // Preprocessor: Parsing constant expression.
	lsConstExpr        // Preprocessor: Parsing constant expression.
	lsTranslationUnit0 //
	lsTranslationUnit  //
)

type trigraphsReader struct {
	*lex.Lexer           //
	pos0       token.Pos //
	sc         int       // Start condition.
}

func (t *trigraphsReader) ReadRune() (rune, int, error) { return lex.RuneEOF, 0, io.EOF }

func (t *trigraphsReader) ReadChar() (c lex.Char, size int, err error) {
	r := rune(t.scan())
	pos0 := t.pos0
	pos := t.Lookahead().Pos()
	t.pos0 = pos
	c = lex.NewChar(t.First.Pos(), r)
	return c, int(pos - pos0), nil
}

type lexer struct {
	*lex.Lexer                             //
	ch                 chan []xc.Token     //
	commentPos0        token.Pos           //
	constantExpression *ConstantExpression //
	cpp                func([]xc.Token)    //
	ddBuf              []*DirectDeclarator //
	encBuf             []byte              // PPTokens
	encBuf1            [30]byte            // Rune, position, optional value ID.
	encPos             token.Pos           // For delta pos encoding
	eof                lex.Char            //
	example            interface{}         //
	exampleRule        int                 //
	externs            map[int]*Declarator //
	file               *token.File         //
	finalNLInjected    bool                //
	includePaths       []string            //
	iota               int                 //
	isPreprocessing    bool                //
	last               xc.Token            //
	model              *Model              //
	preprocessingFile  *PreprocessingFile  //
	report             *xc.Report          //
	sc                 int                 // Start condition.
	scope              *Bindings           //
	scs                int                 // Start condition stack.
	state              int                 // Lexer state.
	sysIncludePaths    []string            //
	t                  *trigraphsReader    //
	textLine           []xc.Token          //
	toC                bool                // Whether to translate preprocessor identifiers to reserved C words.
	tokLast            xc.Token            //
	tokPrev            xc.Token            //
	toks               []xc.Token          // Parsing preprocessor constant expression.
	translationUnit    *TranslationUnit    //
	tweaks             *tweaks             //
}

func newLexer(nm string, sz int, r io.RuneReader, report *xc.Report, tweaks *tweaks, opts ...lex.Option) (*lexer, error) {
	file := fset.AddFile(nm, -1, sz)
	t := &trigraphsReader{}
	lx, err := lex.New(
		file,
		r,
		lex.ErrorFunc(func(pos token.Pos, msg string) {
			report.Err(pos, msg)
		}),
		lex.RuneClass(func(r rune) int { return int(r) }),
	)
	if err != nil {
		return nil, err
	}

	t.Lexer = lx
	t.pos0 = lx.Lookahead().Pos()
	if tweaks.enableTrigraphs {
		t.sc = scTRIGRAPHS
	}
	r = t

	scope := newBindings(nil, ScopeFile)
	lexer := &lexer{
		externs: map[int]*Declarator{},
		file:    file,
		report:  report,
		scope:   scope,
		scs:     -1, // Stack empty
		t:       t,
		tweaks:  tweaks,
	}
	if lexer.Lexer, err = lex.New(
		file,
		r,
		append(opts, lex.RuneClass(rune2class))...,
	); err != nil {
		return nil, err
	}

	return lexer, nil
}

func newSimpleLexer(cpp func([]xc.Token), report *xc.Report, tweaks *tweaks) *lexer {
	return &lexer{
		cpp:     cpp,
		externs: map[int]*Declarator{},
		report:  report,
		scope:   newBindings(nil, ScopeFile),
		tweaks:  tweaks,
	}
}

func (l *lexer) push(sc int) {
	if l.scs >= 0 { // Stack overflow.
		panic("internal error")
	}

	l.scs = l.sc
	l.sc = sc
}

func (l *lexer) pop() {
	if l.scs < 0 { // Stack underflow
		panic("internal error")
	}
	l.sc = l.scs
	l.scs = -1 // Stack empty.
}

func (l *lexer) pushScope(kind Scope) (old *Bindings) {
	old = l.scope
	l.scope = newBindings(old, kind)
	return old
}

func (l *lexer) popScope(tok xc.Token) (old, new *Bindings) {
	old = l.scope
	new = l.scope.parent
	if new == nil {
		l.report.ErrTok(tok, "cannot pop scope")
		return nil, old
	}

	l.scope = new
	return old, new
}

var dlr = []byte{'$'}

func (l *lexer) scanChar() (c lex.Char) {
	r := rune(l.scan())
	switch r {
	case '\n':
		l.state = lsBOL
		l.sc = scINITIAL
		l.scs = -1 // Stack empty
	case PREPROCESSING_FILE:
		l.state = lsBOL
		l.isPreprocessing = true
	case CONSTANT_EXPRESSION, TRANSLATION_UNIT:
		l.toC = true
	}

	return lex.NewChar(l.First.Pos(), r)
}

func (l *lexer) scanToken() (tok xc.Token) {
	switch l.state {
	case lsConstExpr0:
		tok = xc.Token{Char: lex.NewChar(0, CONSTANT_EXPRESSION)}
		l.state = lsConstExpr
	case lsConstExpr:
		if len(l.toks) == 0 {
			tok = xc.Token{Char: lex.NewChar(l.tokLast.Pos(), lex.RuneEOF)}
			break
		}

		tok = l.toks[0]
		l.toks = l.toks[1:]
	case lsTranslationUnit0:
		tok = xc.Token{Char: lex.NewChar(0, TRANSLATION_UNIT)}
		l.state = lsTranslationUnit
		l.toC = true
	case lsTranslationUnit:
		if len(l.textLine) == 0 {
			var ok bool
			if l.textLine, ok = <-l.ch; !ok {
				return xc.Token{Char: lex.NewChar(l.tokLast.Pos(), lex.RuneEOF)}
			}

			if l.cpp != nil {
				l.cpp(l.textLine)
			}
		}
		tok = l.scope.lexerHack(l.textLine[0], l.tokLast)
		l.textLine = l.textLine[1:]
	default:
		c := l.scanChar()
		if c.Rune == ccEOF {
			c = lex.NewChar(c.Pos(), lex.RuneEOF)
			if l.isPreprocessing && l.last.Rune != '\n' && !l.finalNLInjected {
				l.finalNLInjected = true
				l.eof = c
				c.Rune = '\n'
				l.state = lsBOL
				return xc.Token{Char: c}
			}

			return xc.Token{Char: c}
		}

		val := 0
		if tokHasVal[c.Rune] {
			b := l.TokenBytes(nil)
			val = dict.ID(b)
			//TODO handle ID UCNs
			//TODO- chars := l.Token()
			//TODO- switch c.Rune {
			//TODO- case IDENTIFIER, IDENTIFIER_LPAREN:
			//TODO- 	b := l.TokenBytes(func(buf *bytes.Buffer) {
			//TODO- 		for i := 0; i < len(chars); {
			//TODO- 			switch c := chars[i]; {
			//TODO- 			case c.Rune == '$' && !l.tweaks.enableDlrInIdentifiers:
			//TODO- 				l.report.Err(c.Pos(), "identifier character set extension '$' not enabled")
			//TODO- 				i++
			//TODO- 			case c.Rune == '\\':
			//TODO- 				r, n := decodeUCN(chars[i:])
			//TODO- 				buf.WriteRune(r)
			//TODO- 				i += n
			//TODO- 			case c.Rune < 0x80: // ASCII
			//TODO- 				buf.WriteByte(byte(c.Rune))
			//TODO- 				i++
			//TODO- 			default:
			//TODO- 				panic("internal error")
			//TODO- 			}
			//TODO- 		}
			//TODO- 	})
			//TODO- 	val = dict.ID(b)
			//TODO- default:
			//TODO- 	panic("internal error: " + yySymName(int(c.Rune)))
			//TODO- }
		}
		tok = xc.Token{Char: c, Val: val}
		if !l.isPreprocessing {
			tok = l.scope.lexerHack(tok, l.tokLast)
		}
	}
	if l.toC {
		tok = toC(tok)
	}
	l.tokPrev = l.tokLast
	l.tokLast = tok
	return tok
}

// Lex implements yyLexer
func (l *lexer) Lex(lval *yySymType) int {
	tok := l.scanToken()
	l.last = tok
	if tok.Rune == lex.RuneEOF {
		lval.Token = tok
		return 0
	}

	switch l.state {
	case lsBOL:
		switch tok.Rune {
		case PREPROCESSING_FILE, '\n':
			// nop
		case '#':
			l.push(scDIRECTIVE)
			tok = l.scanToken()
			switch tok.Rune {
			case '\n':
				tok.Char = lex.NewChar(tok.Pos(), PPHASH_NL)
			case PPDEFINE:
				l.push(scDEFINE)
				l.state = lsDefine
			case PPELIF, PPENDIF, PPERROR, PPIF, PPLINE, PPPRAGMA:
				l.sc = scINITIAL
				l.state = lsTokens
			case PPELSE, PPIFDEF, PPIFNDEF:
				l.state = lsZero
			case PPUNDEF:
				l.state = lsUndef
			case PPINCLUDE:
				l.sc = scHEADER
				l.state = lsTokens
			default:
				l.state = lsTokens
				tok.Char = lex.NewChar(tok.Pos(), PPNONDIRECTIVE)
				l.pop()
			}
		default:
			l.encodeToken(tok)
			tok.Char = lex.NewChar(tok.Pos(), PPOTHER)
			l.state = lsTokens
		}
	case lsDefine:
		l.pop()
		switch tok.Rune {
		case IDENTIFIER:
			l.state = lsTokens
		case IDENTIFIER_LPAREN:
			l.state = lsSeekRParen
		default:
			l.state = lsZero
		}
	case lsSeekRParen:
		if tok.Rune == ')' {
			l.state = lsTokens
		}
	case lsTokens:
		l.encodeToken(tok)
		tok.Char = lex.NewChar(tok.Pos(), PPOTHER)
	case lsUndef:
		l.state = lsTokens
	}

	lval.Token = tok
	return int(tok.Char.Rune)
}

// Error Implements yyLexer.
func (l *lexer) Error(msg string) {
	if isTesting {
		msg = fmt.Sprintf("%s (last lexer token: %s)", msg, PrettyString(l.tokLast))
	}
	l.report.Err(errPos(l.tokLast.Pos()), "%s", msg)
}

// Reduced implements yyLexerEx
func (l *lexer) Reduced(rule, state int, lval *yySymType) (stop bool) {
	if n := l.exampleRule; n >= 0 && rule != n {
		return false
	}

	switch x := lval.node.(type) {
	case interface {
		fragment() interface{}
	}:
		l.example = x.fragment()
	default:
		l.example = x
	}
	return true
}

func (l *lexer) parsePPConstExpr(list PPTokenList, p *pp) bool {
	l.toks = l.toks[:0]
	p.expand(&tokenBuf{decodeTokens(list, nil)}, true, func(toks []xc.Token) {
		l.toks = append(l.toks, toks...)
	})
	for i, tok := range l.toks {
		if tok.Rune == IDENTIFIER {
			if p.macros.m[tok.Val] != nil {
				l.report.ErrTok(tok, "expected constant expression")
				return false
			}

			tok.Rune = INTCONST
			tok.Val = id0
			l.toks[i] = tok
		}
	}
	l.state = lsConstExpr0
	yyParse(l)
	if v := l.constantExpression.Value; v != nil {
		return isNonZero(v)
	}

	return false
}
