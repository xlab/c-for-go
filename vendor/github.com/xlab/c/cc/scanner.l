%{
// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.4.
%}

%yyc c
%yyn c = l.Next()
%yym l.Mark()
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

func (l *lexer) scan() (r int) {
        c := l.Enter()
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
identifier                      {identifier-nondigit}({identifier-nondigit}|{digit}|{ucn-digit}|"$")*
identifier-nondigit             {nondigit}|{universal-character-name}|{ucn-nondigit}
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
                                c = l.Rule0()

[ \t\f\v]+
"//".*

"/*"                            l.commentPos0 = l.First.Pos()
                                l.push(scCOMMENT)

<COMMENT>{comment-close}        l.pop()

<COMMENT>{eof}                  l.report.Err(l.commentPos0, commentNotClosed)
                                l.pop()
                                return rune2class(lex.RuneEOF)

<*>{eof}                        return rune2class(lex.RuneEOF)

"!="                            return NEQ
"%:"                            return '#'
"%="                            return MODASSIGN
"%>"                            return '}'
"&&"                            return ANDAND
"&="                            return ANDASSIGN
"*="                            return MULASSIGN
"++"                            return INC
"+="                            return ADDASSIGN
"--"                            return DEC
"-="                            return SUBASSIGN
"->"                            return ARROW
"..."                           return DDD
"/="                            return DIVASSIGN
":>"                            return ']'
"<%"                            return '{'
"<:"                            return '['
"<<"                            return LSH
"<<="                           return LSHASSIGN
"<="                            return LEQ
"=="                            return EQ
">="                            return GEQ
">>"                            return RSH
">>="                           return RSHASSIGN
"^="                            return XORASSIGN
"|="                            return ORASSIGN
"||"                            return OROR

"##"                            |
"#%:"                           |
"%:#"                           |
"%:%:"                          return PPPASTE


<DIRECTIVE>"define"             l.pop(); return PPDEFINE
<DIRECTIVE>"elif"               l.pop(); return PPELIF
<DIRECTIVE>"else"               l.pop(); return PPELSE
<DIRECTIVE>"endif"              l.pop(); return PPENDIF
<DIRECTIVE>"error"              l.pop(); return PPERROR
<DIRECTIVE>"if"                 l.pop(); return PPIF
<DIRECTIVE>"ifdef"              l.pop(); return PPIFDEF
<DIRECTIVE>"ifndef"             l.pop(); return PPIFNDEF
<DIRECTIVE>"include"            l.pop(); return PPINCLUDE
<DIRECTIVE>"line"               l.pop(); return PPLINE
<DIRECTIVE>"pragma"             l.pop(); return PPPRAGMA
<DIRECTIVE>"undef"              l.pop(); return PPUNDEF

<HEADER>{header-name}           l.sc = scINITIAL
                                return PPHEADER_NAME

L{character-constant}           return LONGCHARCONST
L{string-literal}               return LONGSTRINGLITERAL
{character-constant}            return CHARCONST
{identifier}                    return IDENTIFIER
<DEFINE>{identifier}"("         return IDENTIFIER_LPAREN
{integer-constant}              return INTCONST
{floating-constant}             return FLOATCONST
{pp-number}                     return PPNUMBER
{string-literal}                return STRINGLITERAL

%%
        if c, ok := l.Abort(); ok {
                return c
        }
        
        goto yyAction
}
