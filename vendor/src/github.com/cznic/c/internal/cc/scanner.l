%{
// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.4.
%}

%yyc c
%yyn c = l.next()
%yym l.m = len(l.in)
%yyt l.sc

%x COMMENT
%s DEFINE DIRECTIVE HEADER

%{
package cc

import (
        "fmt"
	
	"github.com/cznic/golex/lex"
)

const (
        _ = iota
        scCOMMENT       // [`/*`, `*/`]
        scDEFINE        // [^#define, next token]
        scDIRECTIVE     // [^#, next token]
        scHEADER        // [`#include`, next token]
)

func (l *lexer) scan0() lex.Char {
        c := l.enter()
%}

binary-exponent-part            [pP]{sign}?{digit-sequence}
c-char                          [^'\n\x80\\]|{escape-sequence}
c-char-sequence                 {c-char}+
character-constant              '{c-char-sequence}'
comment-close                   ([^*\x80]|\*+[^*/\x80])*\*+\/
decimal-constant                {nonzero-digit}{digit}*
decimal-floating-constant       ({fractional-constant}{exponent-part}?|{digit-sequence}{exponent-part}){floating-suffix}?
digit                           [0-9]
digit-sequence                  {digit}+
eof                             \x80
escape-sequence                 {simple-sequence}|{octal-escape-sequence}|{hexadecimal-escape-sequence}|{universal-character-name}
exponent-part                   [eE]{sign}?{digit-sequence}
floating-constant               {decimal-floating-constant}|{hexadecimal-floating-constant}
floating-suffix                 [flFL]
fractional-constant             {digit-sequence}?\.{digit-sequence}|{digit-sequence}\.
h-char                          [^>\n\x80]
h-char-sequence                 {h-char}+
header-name                     <{h-char-sequence}>|\x22{q-char-sequence}\x22
hex-quad                        {hexadecimal-digit}{hexadecimal-digit}{hexadecimal-digit}{hexadecimal-digit}
hexadecimal-constant            {hexadecimal-prefix}{hexadecimal-digit}+
hexadecimal-digit               [0-9a-fA-F]
hexadecimal-digit-sequence      {hexadecimal-digit}+
hexadecimal-escape-sequence     \\x{hexadecimal-digit}+
hexadecimal-floating-constant   {hexadecimal-prefix}({hexadecimal-fractional-constant}|{hexadecimal-digit-sequence}){binary-exponent-part}{floating-suffix}?
hexadecimal-fractional-constant {hexadecimal-digit-sequence}?\.{hexadecimal-digit-sequence}|{hexadecimal-digit-sequence}\.
hexadecimal-prefix              0[xX]
identifier                      {identifier-nondigit}({identifier-nondigit}|{digit}|{ucn-digit})*
identifier-nondigit             {nondigit}|{universal-character-name}|{ucn-nondigit}|"$" //TODO Reject GCC extension $ when not enabled.
integer-constant                ({decimal-constant}|{octal-constant}|{hexadecimal-constant}){integer-suffix}?
integer-suffix                  {unsigned-suffix}({long-suffix}?|{long-long-suffix})|{long-suffix}{unsigned-suffix}?|{long-long-suffix}{unsigned-suffix}?
long-long-suffix                ll|LL
long-suffix                     [lL]
nondigit                        [_a-zA-Z]
nonzero-digit                   [1-9]
octal-constant                  0{octal-digit}*
octal-digit                     [0-7]
octal-escape-sequence           \\{octal-digit}{octal-digit}?{octal-digit}?
pp-number                       ({digit}|\.{digit})({digit}|{identifier-nondigit}|[eEpP]{sign}|\.)*
q-char                          [^\n\x22\x80]
q-char-sequence                 {q-char}+
s-char                          [^\x22\n\x80\\]|{escape-sequence}
s-char-sequence                 {s-char}+
sign                            [-+]
simple-sequence                 \\['\x22?\\abfnrtv]
string-literal                  \x22{s-char-sequence}?\x22
ucn-digit                       \x83
ucn-nondigit                    \x84
universal-character-name        \\u{hex-quad}|\\U{hex-quad}{hex-quad}
unsigned-suffix                 [uU]

%%
                                c = l.rule0()

[ \t\f\v]+
"//".*

"/*"                            l.commentPos0 = l.in[0].Pos()
                                l.push(scCOMMENT)

<COMMENT>{comment-close}        l.pop()

<COMMENT>{eof}                  compilation.Err(l.commentPos0, commentNotClosed)
                                l.pop()
                                return lex.NewChar(l.in[len(l.in)-1].Pos(), 0)

<*>{eof}                        return lex.NewChar(l.c.Pos(), runeEOF)

"!="                            return l.char(NEQ)
"%="                            return l.char(MODASSIGN)
"%>"                            return l.char('}')
"&&"                            return l.char(ANDAND)
"&="                            return l.char(ANDASSIGN)
"*="                            return l.char(MULASSIGN)
"++"                            return l.char(INC)
"+="                            return l.char(ADDASSIGN)
"--"                            return l.char(DEC)
"-="                            return l.char(SUBASSIGN)
"->"                            return l.char(ARROW)
"..."                           return l.char(DDD)
"/="                            return l.char(DIVASSIGN)
":>"                            return l.char(']')
"<%"                            return l.char('{')
"<:"                            return l.char('[')
"<<"                            return l.char(LSH)
"<<="                           return l.char(LSHASSIGN)
"<="                            return l.char(LEQ)
"=="                            return l.char(EQ)
">="                            return l.char(GEQ)
">>"                            return l.char(RSH)
">>="                           return l.char(RSHASSIGN)
"^="                            return l.char(XORASSIGN)
"|="                            return l.char(ORASSIGN)
"||"                            return l.char(OROR)

"##"                            |
"#%:"                           return l.char(PPPASTE)

"%:"                            // ['%', ':'], z
                                l.unget(l.c, lex.NewChar(l.in[0].Pos(), '#'))
                                l.next()

<DIRECTIVE>"assert"             return l.directive(PPASSERT)
<DIRECTIVE>"define"             return l.directive(PPDEFINE)
<DIRECTIVE>"elif"               return l.directive(PPELIF)
<DIRECTIVE>"else"               return l.directive(PPELSE)
<DIRECTIVE>"endif"              return l.directive(PPENDIF)
<DIRECTIVE>"error"              return l.directive(PPERROR)
<DIRECTIVE>"ident"              return l.directive(PPIDENT)
<DIRECTIVE>"if"                 return l.directive(PPIF)
<DIRECTIVE>"ifdef"              return l.directive(PPIFDEF)
<DIRECTIVE>"ifndef"             return l.directive(PPIFNDEF)
<DIRECTIVE>"import"             return l.directive(PPIMPORT)
<DIRECTIVE>"include"            return l.directive(PPINCLUDE)
<DIRECTIVE>"include_next"       return l.directive(PPINCLUDE_NEXT)
<DIRECTIVE>"line"               return l.directive(PPLINE)
<DIRECTIVE>"pragma"             return l.directive(PPPRAGMA)
<DIRECTIVE>"unassert"           return l.directive(PPUNASSERT)
<DIRECTIVE>"undef"              return l.directive(PPUNDEF)
<DIRECTIVE>"warning"            return l.directive(PPWARNING)

<HEADER>{header-name}           l.sc = scINITIAL
                                return l.char(PPHEADER_NAME)

L{character-constant}           return l.char(LONGCHARCONST)
L{string-literal}               return l.char(LONGSTRINGLITERAL)
{character-constant}            return l.char(CHARCONST)
{identifier}                    return l.char(IDENTIFIER)
<DEFINE>{identifier}"("         return l.char(IDENTIFIER_LPAREN)
{integer-constant}              return l.char(INTCONST)
{floating-constant}             return l.char(FLOATCONST)
{pp-number}                     return l.char(PPNUMBER)
{string-literal}                return l.char(STRINGLITERAL)

%%
        if l.m >= 0 {
                if len(l.in) > l.m {
                        l.unget(l.c)
                        for i := len(l.in) - 1; i >= l.m; i-- {
                                l.unget(l.in[i])
                        }
                        l.next()
                }
                l.in = l.in[:l.m]
                goto yyAction
        }

        return l.abort()
}
