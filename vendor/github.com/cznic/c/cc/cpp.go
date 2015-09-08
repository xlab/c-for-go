// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

//DONE
// The invocation of the macro need not be restricted to a single logical
// line—it can cross as many lines in the source file as you wish.
//
//	https://gcc.gnu.org/onlinedocs/cpp/Macro-Arguments.html

import (
	"bytes"
	"fmt"
	"go/token"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cznic/c/xc"
	"github.com/cznic/golex/lex"
)

//TODO 6.10.8 Predefined macro names

/*

5.1.1.2 Translation phases

1 The precedence among the syntax rules of translation is specified by the
following phases. (*5)

    1. Physical source file multibyte characters are mapped, in an
    implementation defined manner, to the source character set (introducing
    new-line characters for end-of-line indicators) if necessary. Trigraph
    sequences are replaced by corresponding single-character internal
    representations.

    2. Each instance of a backslash character (\) immediately followed by a
    new-line character is deleted, splicing physical source lines to form
    logical source lines.  Only the last backslash on any physical source line
    shall be eligible for being part of such a splice. A source file that is
    not empty shall end in a new-line character, which shall not be immediately
    preceded by a backslash character before any such splicing takes place.

    3. The source file is decomposed into preprocessing tokens (*6) and
    sequences of white-space characters (including comments). A source file
    shall not end in a partial preprocessing token or in a partial comment.
    Each comment is replaced by one space character. New-line characters are
    retained. Whether each nonempty sequence of white-space characters other
    than new-line is retained or replaced by one space character is
    implementation-defined.

    4. Preprocessing directives are executed, macro invocations are expanded,
    and _Pragma unary operator expressions are executed. If a character
    sequence that matches the syntax of a universal character name is produced
    by token concatenation (6.10.3.3), the behavior is undefined. A #include
    preprocessing directive causes the named header or source file to be
    processed from phase 1 through phase 4, recursively. All preprocessing
    directives are then deleted.

    5. Each source character set member and escape sequence in character
    constants and string literals is converted to the corresponding member of
    the execution character set; if there is no corresponding member, it is
    converted to an implementation-defined member other than the null (wide)
    character. (*7)

    6. Adjacent string literal tokens are concatenated.

    7. White-space characters separating tokens are no longer significant. Each
    preprocessing token is converted into a token. The resulting tokens are
    syntactically and semantically analyzed and translated as a translation
    unit.

    8. All external object and function references are resolved. Library
    components are linked to satisfy external references to functions and
    objects not defined in the current translation. All such translator output
    is collected into a program image which contains information needed for
    execution in its execution environment.

(*5) Implementations shall behave as if these separate phases occur, even
though many are typically folded together in practice. Source files,
translation units, and translated translation units need not necessarily be
stored as files, nor need there be any one-to-one correspondence between these
entities and any external representation. The description is conceptual only,
and does not specify any particular implementation.

(*6) The process of dividing a source file’s characters into preprocessing
tokens is context-dependent.  For example, see the handling of < within a
#include preprocessing directive.

(*7) White-space characters separating tokens are no longer significant. Each
preprocessing token is converted into a token. The resulting tokens are
syntactically and semantically analyzed and translated as a translation unit.

*/

var (
	protectedMacros = map[int]bool{
		idDate:             true,
		idDefined:          true,
		idFile:             true,
		idLine:             true,
		idSTDC:             true,
		idSTDCHosted:       true,
		idSTDCMBMightNeqWc: true,
		idSTDCVersion:      true,
		idTime:             true,
		idVAARGS:           true,
	}

	macrosToImplement = map[int]bool{ //TODO
		idDate: true,
		// idSTDCIEC559:        true,
		// idSTDCIEC559Complex: true,
		// idSTDCISO10646:      true,
		idSTDCMBMightNeqWc: true,
	}
)

type macro struct {
	args    []int
	ddd     bool
	dddName int
	defTok  xc.Token
	repl    PpTokenList
}

// ExpandDefine returns (pos, expanded-list, unexpanded-list, true) if id
// refers to a object like definition of a token list. Otherwise (zero-value,
// nil, nil, false) is returned.
func ExpandDefine(id int) (token.Pos, []xc.Token, []xc.Token, bool) {
	m, ok := Macros[id]
	if !ok || m.isFunctionLike() {
		return token.Pos(0), nil, nil, false
	}

	unexpandedToks := db.tokens(m.repl)
	tok := xc.Token{Char: lex.NewChar(0, IDENTIFIER), Val: id}
	toks := []xc.Token{tok}
	lx := newLexer(nil, false)
	ctx := newEvalCtx(lx, nil)
	r := cppExpand0(ctx, 0, toks, nil, nil)

	// If the replacement list is a well formed constant expression,
	// evaluate it and return the value instead.
	if len(r) > 1 {
		v, ok := lx.parseConstExpr(r)
		if ok {
			switch x := v.(type) {
			case int32:
				t := xc.Token{
					Char: lex.NewChar(r[0].Pos(), INTCONST),
					Val:  dict.SID(fmt.Sprint(x)),
				}
				return m.defTok.Pos(), []xc.Token{t}, unexpandedToks, true
			default:
				panic(fmt.Sprintf("TODO: %T", x))
			}
		}
	}

	return m.defTok.Pos(), r, unexpandedToks, true
}

func newMacro(defTok xc.Token, repl PpTokenList) *macro { return &macro{defTok: defTok, repl: repl} }

func newMacro2(defTok xc.Token, args *IdentifierListOpt, repl PpTokenList) *macro {
	a := make([]int, 0, 1) // Must be non nil even for zero args.
	if args != nil {
		first := true
		for l := args.IdentifierList; l != nil; l = l.IdentifierList {
			switch first {
			case true:
				a = append(a, l.Token.Val)
				first = false
			default:
				a = append(a, l.Token2.Val)
			}
		}
	}
	return &macro{args: a, defTok: defTok, repl: repl}
}

func newMacro3(defTok xc.Token, args *IdentifierListOpt, repl PpTokenList) *macro {
	r := newMacro2(defTok, args, repl)
	r.ddd = true
	return r
}

func (m *macro) isObjectLike() bool   { return m.args == nil }
func (m *macro) isFunctionLike() bool { return m.args != nil }

func (m *macro) whitespace() []byte {
	return whitespace(db.tokens(m.repl))
}

func cppDefineFailed(nmTok xc.Token, msg string) {
	compilation.Err(nmTok.Pos(), "#define %s: %s", dict.S(nmTok.Val), msg)
}

func cppDefine(ctx *evalCtx, defTok, nmTok xc.Token, m *macro) {
	nm := nmTok.Val
	if protectedMacros[nm] && !ctx.lx.predefines {
		cppDefineFailed(nmTok, "cannot redefine predeclared macro")
		return
	}

	ex, ok := Macros[nm]
	if !ok {
		Macros[nm] = m
		return
	}

	if ex.isObjectLike() != m.isObjectLike() {
		cppDefineFailed(nmTok, "cannot redefine an object-like macro using a function-like macro or vice versa")
		return
	}

	if m.repl == 0 { // Treat redefining a macro w/o replacement list as being "effectively the same".
		return
	}

	switch {
	case ex.isFunctionLike():
		if g, e := len(m.args), len(ex.args); g != e {
			cppDefineFailed(nmTok, "cannot redefine macro using an argument list of different length")
			return
		}

		for i, g := range m.args {
			if e := ex.args[i]; g != e {
				//dbg("%v %s %s", i, dict.S(e), dict.S(g))
				cppDefineFailed(nmTok, "cannot redefine macro using a different argument list")
				return
			}
		}

		fallthrough
	default:
		mt := db.tokens(m.repl)
		ext := db.tokens(ex.repl)
		if len(mt) != len(ext) {
			//dbg("", PrettyString(mt))
			//dbg("", PrettyString(ext))
			cppDefineFailed(nmTok, "cannot redefine macro using a replacement list of different length")
			return
		}

		for i, mv := range mt {
			exv := ext[i]
			if mv.Rune != exv.Rune || mv.Val != exv.Val {
				cppDefineFailed(nmTok, "cannot redefine macro using a different replacement list")
				return
			}
		}

		wm := m.whitespace()
		wex := ex.whitespace()
		if len(wm) != len(wex) {
			panic("internal error 002")
		}

		if !bytes.Equal(wm, wex) {
			cppDefineFailed(nmTok, "cannot redefine macro, white space differs")
		}
	}
}

func cppUndef(undefTok, nmTok xc.Token) {
	nm := nmTok.Val
	if protectedMacros[nm] {
		compilation.Err(nmTok.Pos(), "#undef %s: forbidden", dict.S(nm))
		return
	}

	delete(Macros, nm)
}

func cppExpand0(ctx *evalCtx, pos token.Pos, toks []xc.Token, nonRepl []int, pgl **GroupList) (r []xc.Token) {
	var a []string
	for _, v := range nonRepl {
		a = append(a, string(dict.S(v)))
	}
next:
	for p := 0; p < len(toks); {
		t := toks[p]
		if t.Rune != IDENTIFIER {
			p++
			continue
		}

		id := t.Val
		for _, v := range nonRepl {
			if id == v {
				p++
				continue next
			}
		}

		if t.Val == idDefined {
			t.Char = lex.NewChar(t.Pos(), DEFINED)
			toks[p] = t
			p++
			continue
		}

		if p > 0 && toks[p-1].Rune == DEFINED || p > 1 && toks[p-2].Rune == DEFINED && toks[p-1].Rune == '(' {
			// Must not expand argument of `defined sym` or `defined(sym)`.
			p++
			continue
		}

		switch id {
		case idLine:
			toks[p] = newToken(toks[p].Pos(), INTCONST, dict.ID([]byte(strconv.Itoa(fileset.Position(pos).Line))))
			p++
			continue
		case idFile:
			toks[p] = newToken(toks[p].Pos(), STRINGLITERAL, dict.ID([]byte(fmt.Sprintf("%q", fileset.Position(pos).Filename))))
			p++
			continue
		case idTime:
			toks[p] = newToken(toks[p].Pos(), STRINGLITERAL, idTTtime)
			p++
			continue
		case idDate:
			toks[p] = newToken(toks[p].Pos(), STRINGLITERAL, idTDate)
			p++
			continue
		default:
			if macrosToImplement[id] {
				panic("TODO")
			}

			macro, ok := Macros[id]
			if !ok {
				p++
				continue
			}

			if macro.isFunctionLike() { // Must consume and replace args.
				if len(toks) == p+1 || toks[p+1].Rune != '(' {
					p++
					continue
				}

				var g0 *GroupList
				if pgl != nil {
					g0 = *pgl
				}
				n, args := ctx.lx.parseMacroArgs(toks[p+1:], pgl)
				if g, e := len(args), len(macro.args); !macro.ddd && g != e {
					compilation.Err(t.Pos(), "number of actual and formal arguments of a macro differ: %v != %v", g, e)
					return nil
				}

				// Fix arg list
				for i, v := range args {
					if len(v) > 1 && v[0].Rune == MACRO_ARG_EMPTY {
						args[i] = v[1:]
					}
				}

				if g0 != nil && g0 != *pgl {
					for g := g0.GroupList; ; g = g.GroupList {
						toks = append(toks, db.tokens(g.GroupPart.PpTokenList)...)
						if g == *pgl {
							break
						}
					}
				}

				replList := db.tokens(macro.repl)
				for i := len(replList) - 1; i >= 0; i-- {
					switch replList[i].Rune {
					case '#':
						y := replList[i+1]
						for iarg, v := range macro.args {
							if y.Val == v {
								arg := args[iarg]
								replList[i] = stringify(arg, false)
								copy(replList[i+1:], replList[i+2:])
								replList = replList[:len(replList)-1]
								i--
								break
							}
						}
					case PPPASTE:
						if i == 0 || i == len(replList)-1 {
							continue
						}

						x := replList[i-1]
						y := replList[i+1]

						for iarg, v := range macro.args {
							if x.Val == v {
								arg := args[iarg]
								if len(arg) == 1 {
									x = arg[0]
									continue
								}

								compilation.Err(arg[0].Pos(), "multitoken ## not supported")
								return nil
							}
							if y.Val == v {
								arg := args[iarg]
								if len(arg) == 1 {
									y = arg[0]
									continue
								}

								compilation.Err(arg[0].Pos(), "multitoken ## not supported")
								return nil
							}
						}

						replList[i-1].Val = dict.ID(append(dict.S(x.Val), dict.S(y.Val)...))
						copy(replList[i:], replList[i+2:])
						replList = replList[:len(replList)-2]
					}
				}

				for i, arg := range args {
					args[i] = cppExpand0(ctx, pos, arg, append(nonRepl, id), pgl)
				}
				for i := len(replList) - 1; i >= 0; i-- {
					t := replList[i]
					if t.Rune != IDENTIFIER {
						continue
					}

					if t.Val == idVAARGS {
						if !macro.ddd {
							compilation.Err(t.Pos(), "invalid __VA_ARGS__ within non variadic macro")
							return nil
						}

						//TODO this supports only macro(...), not for example macro(arg, ...)
						var vaToks []xc.Token
						commas := ctx.lx.macroArgsCommas
						for _, v := range args {
							vaToks = append(vaToks, v...)
							if len(commas) != 0 {
								vaToks = append(vaToks, commas[0])
								commas = commas[1:]
							}
						}
						if len(vaToks) == 1 { // Fast path
							replList[i] = vaToks[0]
							continue
						}

						replList = append(replList[:i], append(vaToks, replList[i+1:]...)...)
						continue
					}

					// Check if t is named arg
					for ia, v := range macro.args {
						if v == t.Val {
							arg := args[ia]
							if len(arg) == 1 { // fast path
								replList[i] = arg[0]
								continue
							}

							replList = append(replList[:i], append(arg, replList[i+1:]...)...)
							continue
						}
					}
				}

				ns := cppExpand0(ctx, pos, replList, append(nonRepl, id), pgl)
				toks = append(toks[:p], append(ns, toks[p+1+n:]...)...)
				p += len(ns)
				continue
			}

			if macro.repl == 0 { // Empty replacement list.
				copy(toks[p:], toks[p+1:])
				toks = toks[:len(toks)-1]
				continue
			}

			replList := db.tokens(macro.repl)
			if len(replList) == 1 && replList[0].Rune != IDENTIFIER {
				toks[p] = replList[0]
				p++
				continue
			}

			ns := cppExpand0(ctx, pos, replList, append(nonRepl, id), pgl)
			if len(ns) == 1 {
				toks[p] = ns[0]
				p++
				continue
			}

			toks = append(toks[:p], append(ns, toks[p+1:]...)...)
			p += len(ns)
		}
	}
	return toks
}

// pgl != nil -> expanding TextLine
func cppExpand(ctx *evalCtx, list PpTokenList, pgl **GroupList) (r []xc.Token) {
	toks := db.tokens(list)
	return cppExpand0(ctx, toks[0].Pos(), toks, nil, pgl)
}

func whitespace(toks []xc.Token) []byte {
	if len(toks) < 2 {
		return nil
	}

	r := make([]byte, 0, len(toks)-1)
	ltok := toks[0]
	for _, tok := range toks[1:] {
		pos0 := int(ltok.Pos()) + len(dict.S(tokVal(ltok)))
		pos1 := int(tok.Pos())
		d := byte(0)
		if pos0 != pos1 {
			d = 1
		}
		r = append(r, d)
		ltok = tok
	}
	return r
}

func unquote(b []byte) []byte {
	if len(b) != 0 && b[0] == '"' {
		b = b[1:]
	}
	if n := len(b); n != 0 && b[n-1] == '"' {
		b = b[:n-1]
	}
	return b
}

func stringify(toks []xc.Token, paste bool) xc.Token {
	if len(toks) == 0 {
		return newToken(0, STRINGLITERAL, idEmptyString)
	}

	v := tokVal(toks[0])
	w := whitespace(toks)
	s := append([]byte{'"'}, unquote(dict.S(v))...)
	for i, t := range toks[1:] {
		v = tokVal(t)
		if !paste && w[i] != 0 {
			s = append(s, ' ')
		}
		s = append(s, unquote(dict.S(v))...)
	}
	s = append(s, '"')
	return newToken(toks[0].Pos(), STRINGLITERAL, dict.ID(s))
}

// --------------------------------------------------- Manual AST types/methods

type evalCtx struct {
	ch chan []xc.Token //
	lx *lexer          // For constant expression parse and eval
	// for #pragma once
	onceMap map[string]struct{}
}

func newEvalCtx(lx *lexer, ch chan []xc.Token) *evalCtx {
	return &evalCtx{
		lx: lx,
		ch: ch,
		//
		onceMap: make(map[string]struct{}),
	}
}

type PpTokenList int

// String implements fmt.Stringer.
func (p PpTokenList) String() string {
	return PrettyString(p)
}

// GroupList represents data reduced by production(s):
//
//	GroupList:
//		GroupPart
//	|	GroupList GroupPart  // Case 1
type GroupList struct {
	GroupList *GroupList
	GroupPart *GroupPart
}

func (g *GroupList) reverse() *GroupList {
	if g == nil {
		return nil
	}

	na := g
	nb := na.GroupList
	for nb != nil {
		nc := nb.GroupList
		nb.GroupList = na
		na = nb
		nb = nc
	}
	g.GroupList = nil
	return na
}

func (g *GroupList) fragment() interface{} { return g.reverse() }

// String implements fmt.Stringer.
func (g *GroupList) String() string {
	return PrettyString(g)
}

func (g *GroupList) preprocess(ctx *evalCtx) {
	for ; g != nil; g = g.GroupList {
		g.GroupPart.preprocess(ctx, &g)
	}
}

// GroupPart represents data reduced by production(s):
//
//	GroupPart:
//		ControlLine
//	|	IfSection                   // Case 1
//	|	PPNONDIRECTIVE PpTokenList  // Case 2
//	|	TextLine                    // Case 3
type GroupPart struct {
	ControlLine *ControlLine
	IfSection   *IfSection
	PpTokenList PpTokenList
	Token       xc.Token
}

func (g *GroupPart) fragment() interface{} { return g }

// String implements fmt.Stringer.
func (g *GroupPart) String() string {
	return PrettyString(g)
}

func (g *GroupPart) preprocess(ctx *evalCtx, pgl **GroupList) {
	if x := g.ControlLine; x != nil { // ControlLine
		x.preprocess(ctx)
		return
	}

	if x := g.IfSection; x != nil { // IfSection
		x.preprocess(ctx)
		return
	}

	if !g.Token.IsValid() { // Textline
		toks := cppExpand(ctx, g.PpTokenList, pgl)
		if ctx.ch == nil || len(toks) == 0 {
			return
		}

		ctx.ch <- toks
	}
}

func (p *PreprocessingFile) preprocess(ctx *evalCtx) {
	if p != nil {
		p.GroupList.preprocess(ctx)
	}
}

func (c *ControlLine) preprocess(ctx *evalCtx) {
	switch c.Case {
	case 0: // PPDEFINE IDENTIFIER ReplacementList
		cppDefine(ctx, c.Token, c.Token2, newMacro(c.Token, c.ReplacementList))
	case 1: // PPDEFINE IDENTIFIER_LPAREN DDD ')' ReplacementList
		cppDefine(ctx, c.Token, c.Token2, newMacro3(c.Token, nil, c.ReplacementList))
	case 2: //TODO PPDEFINE IDENTIFIER_LPAREN IdentifierList ',' DDD ')' ReplacementList
		// cppDefine(ctx, c.Token, c.Token2, newMacro3(c.Token, c.IdentifierListOpt, c.ReplacementList))
	case 3: // PPDEFINE IDENTIFIER_LPAREN IdentifierListOpt ')' ReplacementList
		cppDefine(ctx, c.Token, c.Token2, newMacro2(c.Token, c.IdentifierListOpt, c.ReplacementList))
	case 4: // PPERROR PpTokenListOpt
		s := stringify(db.tokens(c.PpTokenListOpt), false)
		compilation.Err(c.Token.Pos(), "%s", dict.S(s.Val))
	case 5: // PPHASH_NL
		// nop
	case 6: // PPINCLUDE PpTokenList
		var refTok xc.Token
		if tokList := cppExpand(ctx, c.PpTokenList, nil); len(tokList) == 0 {
			compilation.Err(c.Token.Pos(), "invalid #include")
			return
		} else {
			refTok = tokList[0]
		}

		var refValue string
		if refValue = string(dict.S(refTok.Val)); len(refValue) == 0 {
			compilation.Err(refTok.Pos(), "invalid #include: %s", refValue)
			return
		}

		var isSys bool
		switch {
		case strings.HasPrefix(refValue, `"`) && strings.HasSuffix(refValue, `"`):
		case strings.HasPrefix(refValue, `<`) && strings.HasSuffix(refValue, `>`):
			isSys = true
		default:
			compilation.Err(refTok.Pos(), "invalid #include: %s", refValue)
			return
		}

		refPath := strings.Trim(refValue, `"<>`)
		if likelyIsURL(refPath) {
			if !webIncludesEnabled {
				compilation.Err(refTok.Pos(), "forbidden #include: %s", refPath)
				return
			}
			u, err := url.Parse(refPath)
			if err != nil {
				compilation.Err(refTok.Pos(), "malformed #include: %s", refPath)
				return
			}
			if ppFile := ppFileByURL(refTok, u); ppFile != nil {
				ppFile.preprocess(ctx)
			}
			return
		}

		searchSys := func(refTok xc.Token, refPath string) (found bool) {
			for _, sysPath := range sysIncludePaths {
				fn := filepath.Join(sysPath, refPath)
				if _, ok := ctx.onceMap[fn]; ok {
					// skip this file
					return true
				}
				if _, err := os.Stat(fn); err != nil && !os.IsNotExist(err) {
					compilation.Err(refTok.Pos(), err.Error())
					return true
				} else if os.IsNotExist(err) {
					continue
				}
				if ppFile := ppFileByPath(refTok, fn); ppFile != nil {
					ppFile.preprocess(ctx)
				}
				return true
			}
			// try to fetch this file from web
			if webIncludesEnabled && len(webIncludePrefix) > 0 {
				u, _ := url.Parse(webIncludePrefix)
				u.Path = filepath.Join(u.Path, refPath)
				if ppFile := ppFileByURL(refTok, u); ppFile != nil {
					ppFile.preprocess(ctx)
				}
				return true
			}
			return false
		}

		if isSys {
			if searchSys(refTok, refPath) {
				return
			}
			compilation.Err(refTok.Pos(), "include file not found: %s", refPath)
			return
		}

		// TODO: consider case when Filename is actually an URL (when parsing PPINCLUDE in web-included file)
		// also consider a way to fallback to some default URL if local file wasn't found.

		srcDir := filepath.Dir(fileset.Position(c.Token.Pos()).Filename)
		fn := filepath.Join(srcDir, refPath)
		if _, ok := ctx.onceMap[fn]; ok {
			// skip this file
			return
		}
		if _, err := os.Stat(fn); err != nil && !os.IsNotExist(err) {
			compilation.Err(refTok.Pos(), err.Error())
			return
		} else if err == nil {
			if ppFile := ppFileByPath(refTok, fn); ppFile != nil {
				ppFile.preprocess(ctx)
			}
			return
		}

		for _, relPath := range includePaths {
			fn := filepath.Join(srcDir, relPath, refPath)
			if _, ok := ctx.onceMap[fn]; ok {
				// skip this file
				return
			}
			if _, err := os.Stat(fn); err != nil && !os.IsNotExist(err) {
				compilation.Err(refTok.Pos(), err.Error())
				return
			} else if os.IsNotExist(err) {
				continue
			}
			if ppFile := ppFileByPath(refTok, fn); ppFile != nil {
				ppFile.preprocess(ctx)
			}
			return
		}
		// try to search in sys include paths
		if searchSys(refTok, refPath) {
			return
		}
		compilation.Err(refTok.Pos(), "include file not found: %s", refPath)

	case 7: //TODO PPLINE PpTokenList
	case 8: // PPPRAGMA PpTokenListOpt
		var once bool
		if toks := db.tokens(c.PpTokenListOpt); len(toks) > 0 {
			once = bytes.Equal(toks[0].S(), []byte("once"))
		}
		if once {
			name := fileset.Position(c.Token.Pos()).Filename
			ctx.onceMap[name] = struct{}{}
		}
	case 9: // PPUNDEF IDENTIFIER '\n'
		cppUndef(c.Token, c.Token2)
	case 10: //TODO PPASSERT PpTokenList
	case 11: //TODO PPDEFINE IDENTIFIER_LPAREN IDENTIFIER DDD ')' ReplacementList
	case 12: //TODO PPDEFINE IDENTIFIER_LPAREN IdentifierList ',' IDENTIFIER DDD ')' ReplacementList
	case 13: //TODO PPIDENT PpTokenList
	case 14: //TODO PPIMPORT PpTokenList
	case 15: //TODO PPINCLUDE_NEXT PpTokenList
	case 16: //TODO PPUNASSERT PpTokenList
	case 17: //TODO PPWARNING PpTokenList
	default:
		panic("TODO")
	}
}

func (i *IfSection) preprocess(ctx *evalCtx) {
	if i.IfGroup.preprocess(ctx) {
		return
	}

	if x := i.ElifGroupListOpt; x != nil && x.ElifGroupList.preprocess(ctx) {
		return
	}

	if x := i.ElseGroupOpt; x != nil {
		x.ElseGroup.preprocess(ctx)
	}
}

func (e *ElseGroup) preprocess(ctx *evalCtx) {
	if l := e.GroupListOpt; l != nil {
		l.GroupList.preprocess(ctx)
	}
}

func (i *IfGroup) preprocess(ctx *evalCtx) bool {
	switch i.Case {
	case 0: // PPIF PpTokenList GroupListOpt
		v, _ := ctx.lx.parseConstExpr(cppExpand(ctx, i.PpTokenList, nil))
		if isZero(v) {
			return false
		}

		i.GroupListOpt.preprocess(ctx)
		return true
	case 1: // PPIFDEF IDENTIFIER '\n' GroupListOpt
		if _, ok := Macros[i.Token2.Val]; !ok {
			return false
		}

		i.GroupListOpt.preprocess(ctx)
		return true
	default: // 2: PPIFNDEF IDENTIFIER '\n' GroupListOpt
		if _, ok := Macros[i.Token2.Val]; ok {
			return false
		}

		i.GroupListOpt.preprocess(ctx)
		return true
	}
}

func (g *GroupListOpt) preprocess(ctx *evalCtx) {
	if g == nil {
		return
	}

	if x := g.GroupList; x != nil {
		x.preprocess(ctx)
	}
}

func (c *ConstantExpression) eval() interface{} {
	return c.ConditionalExpression.eval()
}

func (c *ConditionalExpression) eval() (v interface{}) {
	switch c.Case {
	case 0: // LogicalOrExpression
		return c.LogicalOrExpression.eval()
	default: // 1: LogicalOrExpression '?' ExpressionList ':' ConditionalExpression
		if isZero(c.LogicalOrExpression.eval()) {
			return c.ConditionalExpression.eval()
		}

		return c.ExpressionList.eval()
	}
}

func (l *LogicalOrExpression) eval() (v interface{}) {
	switch l.Case {
	case 0: // LogicalAndExpression
		return l.LogicalAndExpression.eval()
	case 1: // LogicalOrExpression "||" LogicalAndExpression
		v = l.LogicalOrExpression.eval()
		if !isZero(v) {
			return v
		}

		return l.LogicalAndExpression.eval()
	default:
		panic("internal error")
	}
}

func (l *LogicalAndExpression) eval() (v interface{}) {
	switch l.Case {
	case 0: // InclusiveOrExpression
		return l.InclusiveOrExpression.eval()
	case 1: // LogicalAndExpression "&&" InclusiveOrExpression
		v = l.LogicalAndExpression.eval()
		if isZero(v) {
			return v
		}

		if !isZero(l.InclusiveOrExpression.eval()) {
			return boolT(true)
		}

		return boolT(false)
	default:
		panic("internal error")
	}
}

func (i *InclusiveOrExpression) eval() (v interface{}) {
	switch i.Case {
	case 0: // ExclusiveOrExpression
		return i.ExclusiveOrExpression.eval()
	case 1: // InclusiveOrExpression '|' ExclusiveOrExpression
		v = i.InclusiveOrExpression.eval()
		v2 := i.ExclusiveOrExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x | y
			case int64:
				v = int64(x) | y
			case uint32:
				v = uint32(x) | y
			case uint64:
				v = uint64(x) | y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x | int64(y)
			case int64:
				v = x | y
			case uint32:
				v = uint64(x) | uint64(y)
			case uint64:
				v = uint64(x) | y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x | uint32(y)
			case int64:
				v = uint64(x) | uint64(y)
			case uint32:
				v = uint32(x) | y
			case uint64:
				v = uint64(x) | y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x | uint64(y)
			case int64:
				v = x | uint64(y)
			case uint32:
				v = x | uint64(y)
			case uint64:
				v = x | y
			}
		}
	default:
		panic("internal error")
	}
	return v
}

func (e *ExclusiveOrExpression) eval() (v interface{}) {
	switch e.Case {
	case 0: // AndExpression
		return e.AndExpression.eval()
	case 1: // ExclusiveOrExpression '^' AndExpression
		v = e.ExclusiveOrExpression.eval()
		v2 := e.AndExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x ^ y
			case int64:
				v = int64(x) ^ y
			case uint32:
				v = uint32(x) ^ y
			case uint64:
				v = uint64(x) ^ y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x ^ int64(y)
			case int64:
				v = x ^ y
			case uint32:
				v = uint64(x) ^ uint64(y)
			case uint64:
				v = uint64(x) ^ y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x ^ uint32(y)
			case int64:
				v = uint64(x) ^ uint64(y)
			case uint32:
				v = x ^ y
			case uint64:
				v = uint64(x) ^ y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x ^ uint64(y)
			case int64:
				v = x ^ uint64(y)
			case uint32:
				v = x ^ uint64(y)
			case uint64:
				v = x ^ y
			}
		}
	default:
		panic("internal error")
	}
	return v
}

func (a *AndExpression) eval() (v interface{}) {
	switch a.Case {
	case 0: // EqualityExpression
		return a.EqualityExpression.eval()
	case 1: // AndExpression '&' EqualityExpression
		v = a.AndExpression.eval()
		v2 := a.EqualityExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x & y
			case int64:
				v = int64(x) & y
			case uint32:
				v = uint32(x) & y
			case uint64:
				v = uint64(x) & y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x & int64(y)
			case int64:
				v = x & y
			case uint32:
				v = uint64(x) & uint64(y)
			case uint64:
				v = uint64(x) & y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x & uint32(y)
			case int64:
				v = uint64(x) & uint64(y)
			case uint32:
				v = x & y
			case uint64:
				v = uint64(x) & y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x & uint64(y)
			case int64:
				v = x & uint64(y)
			case uint32:
				v = x & uint64(y)
			case uint64:
				v = x & y
			}
		}
	default:
		panic("internal error")
	}
	return v
}

func (e *EqualityExpression) eval() (v interface{}) {
	switch e.Case {
	case 0: // RelationalExpression
		return e.RelationalExpression.eval()
	case 1: // EqualityExpression "==" RelationalExpression
		v = e.EqualityExpression.eval()
		v2 := e.RelationalExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x == y)
			case int64:
				v = boolT(int64(x) == y)
			case uint32:
				v = boolT(uint32(x) == y)
			case uint64:
				v = boolT(uint64(x) == y)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x == int64(y))
			case int64:
				v = boolT(x == y)
			case uint32:
				v = boolT(uint64(x) == uint64(y))
			case uint64:
				v = boolT(uint64(x) == y)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x == uint32(y))
			case int64:
				v = boolT(uint64(x) == uint64(y))
			case uint32:
				v = boolT(x == y)
			case uint64:
				v = boolT(uint64(x) == y)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x == uint64(y))
			case int64:
				v = boolT(x == uint64(y))
			case uint32:
				v = boolT(x == uint64(y))
			case uint64:
				v = boolT(x == y)
			}
		}
	case 2: // EqualityExpression "!=" RelationalExpression
		v = e.EqualityExpression.eval()
		v2 := e.RelationalExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x != y)
			case int64:
				v = boolT(int64(x) != y)
			case uint32:
				v = boolT(uint32(x) != y)
			case uint64:
				v = boolT(uint64(x) != y)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x != int64(y))
			case int64:
				v = boolT(x != y)
			case uint32:
				v = boolT(uint64(x) != uint64(y))
			case uint64:
				v = boolT(uint64(x) != y)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x != uint32(y))
			case int64:
				v = boolT(uint64(x) != uint64(y))
			case uint32:
				v = boolT(x != y)
			case uint64:
				v = boolT(uint64(x) != y)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x != uint64(y))
			case int64:
				v = boolT(x != uint64(y))
			case uint32:
				v = boolT(x != uint64(y))
			case uint64:
				v = boolT(x != y)
			}
		}
	default:
		panic("internal error")
	}
	return v
}

func (r *RelationalExpression) eval() (v interface{}) {
	switch r.Case {
	case 0: // ShiftExpression
		return r.ShiftExpression.eval()
	case 1: // RelationalExpression '<' ShiftExpression
		v = r.RelationalExpression.eval()
		v2 := r.ShiftExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x < y)
			case int64:
				v = boolT(int64(x) < y)
			case uint32:
				v = boolT(uint32(x) < y)
			case uint64:
				v = boolT(uint64(x) < y)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x < int64(y))
			case int64:
				v = boolT(x < y)
			case uint32:
				v = boolT(uint64(x) < uint64(y))
			case uint64:
				v = boolT(uint64(x) < y)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x < uint32(y))
			case int64:
				v = boolT(uint64(x) < uint64(y))
			case uint32:
				v = boolT(x < y)
			case uint64:
				v = boolT(uint64(x) < y)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x < uint64(y))
			case int64:
				v = boolT(x < uint64(y))
			case uint32:
				v = boolT(x < uint64(y))
			case uint64:
				v = boolT(x < y)
			}
		}
	case 2: // RelationalExpression '>' ShiftExpression
		v = r.RelationalExpression.eval()
		v2 := r.ShiftExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x > y)
			case int64:
				v = boolT(int64(x) > y)
			case uint32:
				v = boolT(uint32(x) > y)
			case uint64:
				v = boolT(uint64(x) > y)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x > int64(y))
			case int64:
				v = boolT(x > y)
			case uint32:
				v = boolT(uint64(x) > uint64(y))
			case uint64:
				v = boolT(uint64(x) > y)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x > uint32(y))
			case int64:
				v = boolT(uint64(x) > uint64(y))
			case uint32:
				v = boolT(x > y)
			case uint64:
				v = boolT(uint64(x) > y)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x > uint64(y))
			case int64:
				v = boolT(x > uint64(y))
			case uint32:
				v = boolT(x > uint64(y))
			case uint64:
				v = boolT(x > y)
			}
		}
	case 3: // RelationalExpression "<=" ShiftExpression
		v = r.RelationalExpression.eval()
		v2 := r.ShiftExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x <= y)
			case int64:
				v = boolT(int64(x) <= y)
			case uint32:
				v = boolT(uint32(x) <= y)
			case uint64:
				v = boolT(uint64(x) <= y)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x <= int64(y))
			case int64:
				v = boolT(x <= y)
			case uint32:
				v = boolT(uint64(x) <= uint64(y))
			case uint64:
				v = boolT(uint64(x) <= y)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x <= uint32(y))
			case int64:
				v = boolT(uint64(x) <= uint64(y))
			case uint32:
				v = boolT(x <= y)
			case uint64:
				v = boolT(uint64(x) <= y)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x <= uint64(y))
			case int64:
				v = boolT(x <= uint64(y))
			case uint32:
				v = boolT(x <= uint64(y))
			case uint64:
				v = boolT(x <= y)
			}
		}
	case 4: // RelationalExpression ">=" ShiftExpression
		v = r.RelationalExpression.eval()
		v2 := r.ShiftExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x >= y)
			case int64:
				v = boolT(int64(x) >= y)
			case uint32:
				v = boolT(uint32(x) >= y)
			case uint64:
				v = boolT(uint64(x) >= y)
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x >= int64(y))
			case int64:
				v = boolT(x >= y)
			case uint32:
				v = boolT(uint64(x) >= uint64(y))
			case uint64:
				v = boolT(uint64(x) >= y)
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = boolT(x >= uint32(y))
			case int64:
				v = boolT(uint64(x) >= uint64(y))
			case uint32:
				v = boolT(x >= y)
			case uint64:
				v = boolT(uint64(x) >= y)
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = boolT(x >= uint64(y))
			case int64:
				v = boolT(x >= uint64(y))
			case uint32:
				v = boolT(x >= uint64(y))
			case uint64:
				v = boolT(x >= y)
			}
		}
	default:
		panic("internal error")
	}
	return v
}

func (s *ShiftExpression) eval() (v interface{}) {
	switch s.Case {
	case 0: // AdditiveExpression
		return s.AdditiveExpression.eval()
	case 1: // ShiftExpression "<<" AdditiveExpression
		v = s.ShiftExpression.eval()
		v2 := s.AdditiveExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x << uint(y)
				case y < 0:
					v = x >> uint(-y)
				}
			}
		}
	case 2: // ShiftExpression ">>" AdditiveExpression
		v = s.ShiftExpression.eval()
		v2 := s.AdditiveExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case int64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint32:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			case uint64:
				switch {
				case y > 0:
					v = x >> uint(y)
				case y < 0:
					v = x << uint(-y)
				}
			}
		}
	default:
		panic("internal error")
	}
	return v
}

func (a *AdditiveExpression) eval() (v interface{}) {
	switch a.Case {
	case 0: // MultiplicativeExpression
		return a.MultiplicativeExpression.eval()
	case 1: // AdditiveExpression '+' MultiplicativeExpression
		v = a.AdditiveExpression.eval()
		v2 := a.MultiplicativeExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x + y
			case int64:
				v = int64(x) + y
			case uint32:
				v = uint32(x) + y
			case uint64:
				v = uint64(x) + y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x + int64(y)
			case int64:
				v = x + y
			case uint32:
				v = uint64(x) + uint64(y)
			case uint64:
				v = uint64(x) + y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x + uint32(y)
			case int64:
				v = uint64(x) + uint64(y)
			case uint32:
				v = x + y
			case uint64:
				v = uint64(x) + y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x + uint64(y)
			case int64:
				v = x + uint64(y)
			case uint32:
				v = x + uint64(y)
			case uint64:
				v = x + y
			}
		}
	case 2: // AdditiveExpression '-' MultiplicativeExpression
		v = a.AdditiveExpression.eval()
		v2 := a.MultiplicativeExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x - y
			case int64:
				v = int64(x) - y
			case uint32:
				v = uint32(x) - y
			case uint64:
				v = uint64(x) - y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x - int64(y)
			case int64:
				v = x - y
			case uint32:
				v = uint64(x) - uint64(y)
			case uint64:
				v = uint64(x) - y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x - uint32(y)
			case int64:
				v = uint64(x) - uint64(y)
			case uint32:
				v = x - y
			case uint64:
				v = uint64(x) - y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x - uint64(y)
			case int64:
				v = x - uint64(y)
			case uint32:
				v = x - uint64(y)
			case uint64:
				v = x - y
			}
		}
	default:
		panic("internal error")
	}
	return v
}

func (m *MultiplicativeExpression) eval() (v interface{}) {
	switch m.Case {
	case 0: // CastExpression
		return m.CastExpression.eval()
	case 1: // MultiplicativeExpression '*' CastExpression
		v = m.MultiplicativeExpression.eval()
		v2 := m.CastExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				v = x * y
			case int64:
				v = int64(x) * y
			case uint32:
				v = uint32(x) * y
			case uint64:
				v = uint64(x) * y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				v = x * int64(y)
			case int64:
				v = x * y
			case uint32:
				v = uint64(x) * uint64(y)
			case uint64:
				v = uint64(x) * y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				v = x * uint32(y)
			case int64:
				v = uint64(x) * uint64(y)
			case uint32:
				v = x * y
			case uint64:
				v = uint64(x) * y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				v = x * uint64(y)
			case int64:
				v = x * uint64(y)
			case uint32:
				v = x * uint64(y)
			case uint64:
				v = x * y
			}
		}
	case 2: // MultiplicativeExpression '/' CastExpression
		v = m.MultiplicativeExpression.eval()
		v2 := m.CastExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = int32(0)
					break
				}

				v = x / y
			case int64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = int64(0)
					break
				}

				v = int64(x) / y
			case uint32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint32(0)
					break
				}

				v = uint32(x) / y
			case uint64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = uint64(x) / y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = int64(0)
					break
				}

				v = x / int64(y)
			case int64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = int64(0)
					break
				}

				v = x / y
			case uint32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = uint64(x) / uint64(y)
			case uint64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = uint64(x) / y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint32(0)
					break
				}

				v = x / uint32(y)
			case int64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = uint64(x) / uint64(y)
			case uint32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint32(0)
					break
				}

				v = x / y
			case uint64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = uint64(x) / y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = x / uint64(y)
			case int64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = x / uint64(y)
			case uint32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = x / uint64(y)
			case uint64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = x / y
			}
		}
	case 3: // MultiplicativeExpression '%' CastExpression
		v = m.MultiplicativeExpression.eval()
		v2 := m.CastExpression.eval()
		switch x := v.(type) {
		case int32:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = int32(0)
					break
				}

				v = x % y
			case int64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = int64(0)
					break
				}

				v = int64(x) % y
			case uint32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint32(0)
					break
				}

				v = uint32(x) % y
			case uint64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = uint64(x) % y
			}
		case int64:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = int64(0)
					break
				}

				v = x % int64(y)
			case int64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = int64(0)
					break
				}

				v = x % y
			case uint32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = uint64(x) % uint64(y)
			case uint64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = uint64(x) % y
			}
		case uint32:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint32(0)
					break
				}

				v = x % uint32(y)
			case int64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = uint64(x) % uint64(y)
			case uint32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint32(0)
					break
				}

				v = x % y
			case uint64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = uint64(x) % y
			}
		case uint64:
			switch y := v2.(type) {
			case int32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = x % uint64(y)
			case int64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = x % uint64(y)
			case uint32:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = x % uint64(y)
			case uint64:
				if y == 0 {
					compilation.Err(m.Token.Pos(), "division by zero")
					v = uint64(0)
					break
				}

				v = x % y
			}
		}
	default:
		panic("internal error")
	}
	return v
}

func (c *CastExpression) eval() interface{} {
	switch c.Case {
	case 0: // UnaryExpression
		return c.UnaryExpression.eval()
	case 1: // '(' TypeName ')' CastExpression
		return c.CastExpression.eval()
	default:
		panic("TODO")
	}
}

func (u *UnaryExpression) eval() interface{} {
	switch u.Case {
	case 0: // PostfixExpression
		return u.PostfixExpression.eval()
	case 3: // UnaryOperator CastExpression
		x := u.CastExpression.eval()
		switch u.UnaryOperator.Token.Rune {
		case '!':
			return boolT(isZero(x))
		case '-':
			switch x := x.(type) {
			case int32:
				return -x
			case int64:
				return -x
			case uint32:
				return -x
			case uint64:
				return -x
			}
		case '+':
			switch x := x.(type) {
			case int32:
				return +x
			case int64:
				return +x
			case uint32:
				return +x
			case uint64:
				return +x
			}
		case '~':
			switch x := x.(type) {
			case int32:
				return ^x
			case int64:
				return ^x
			case uint32:
				return ^x
			case uint64:
				return ^x
			}
		case '&':
			compilation.Err(u.UnaryOperator.Token.Pos(), "cannot take address of a constant expression")
			return intT(0)
		default:
			panic("TODO")
		}
	case 5: // "sizeof" '(' TypeName ')'
		return intT(u.TypeName.Type().Sizeof())
	case 6: // DEFINED IDENTIFIER
		return boolT(Macros[u.Token2.Val] != nil)
	case 7: // DEFINED '(' IDENTIFIER ')'
		return boolT(Macros[u.Token3.Val] != nil)
	default:
		println(PrettyString(u)) //TODO-
		panic("TODO")
	}
	panic("unreachable")
}

func (e *ElifGroupList) preprocess(ctx *evalCtx) bool {
	for ; e != nil; e = e.ElifGroupList {
		if e.ElifGroup.preprocess(ctx) {
			return true
		}
	}

	return false
}

func (e *ElifGroup) preprocess(ctx *evalCtx) bool {
	v, _ := ctx.lx.parseConstExpr(cppExpand(ctx, e.PpTokenList, nil))
	if isZero(v) {
		return false
	}

	e.GroupListOpt.preprocess(ctx)
	return true
}

func (p *PostfixExpression) eval() (v interface{}) {
	switch p.Case {
	case 0: // PrimaryExpression
		return p.PrimaryExpression.eval()
	default:
		panic(PrettyString(p))
	}
	//for {
	//	switch p.Case {
	//	case 0: // PrimaryExpression
	//		v = p.PrimaryExpression.eval()
	//	case 2: // PostfixExpression '(' ArgumentExpressionListOpt ')'
	//		compilation.Err(p.Token.Pos(), "function call not allowed in constant expression")
	//		return intT(0)
	//	default:
	//		panic("TODO")
	//	}

	//	if p = p.PostfixExpression; p == nil {
	//		return v
	//	}
	//}
}

func (p *PrimaryExpression) eval() (r interface{}) {
	switch p.Case {
	case 0: // IDENTIFIER
		return intT(0)
	case 1: // Constant
		return p.Constant.eval()
	case 2: // '(' ExpressionList ')'
		return p.ExpressionList.eval()
	default:
		panic("TODO")
	}
}

func (c *Constant) eval() interface{} {
	switch c.Token.Rune {
	case CHARCONST, LONGCHARCONST:
		return charConst(c.Token)
	case INTCONST:
		return intConst(c.Token)
	default:
		compilation.Err(c.Token.Pos(), "invalid constant")
		return intT(0)
	}
}

func (e *ExpressionList) eval() (v interface{}) {
	for ; e != nil; e = e.ExpressionList {
		v = e.AssignmentExpression.eval()
	}
	return v
}

func (a *AssignmentExpression) eval() interface{} {
	switch a.Case {
	case 0:
		return a.ConditionalExpression.eval()
	default:
		panic("TODO")
	}
}

func (a *AssignmentExpression) eval2() int {
	switch n := a.eval().(type) {
	case int32:
		return int(n)
	default:
		panic(fmt.Sprintf("%T", n))
	}
}
