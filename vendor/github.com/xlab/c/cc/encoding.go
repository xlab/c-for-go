// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"encoding/binary"
	"go/token"
	"sort"
	"strings"
	"time"

	"github.com/cznic/golex/lex"
	"github.com/cznic/mathutil"
	"github.com/xlab/c/xc"
)

const (
	intBits  = mathutil.IntBits
	bitShift = intBits>>6 + 5
	bitMask  = intBits - 1

	scINITIAL = 0 // Start condition (shared value).
)

const (
	// Character class is an 8 bit encoding of an Unicode rune for the
	// golex generated FSM.
	//
	// Every ASCII rune is its own class.  DO NOT change any of the
	// existing values. Adding new classes is OK.
	ccEOF         = iota + 0x80
	_             // ccError
	ccOther       // Any other rune.
	ccUCNDigit    // [0], Annex D, Universal character names for identifiers - digits.
	ccUCNNonDigit // [0], Annex D, Universal character names for identifiers - non digits.
)

// Values of TypeSpecifier.Case.
const (
	tsVoid            = iota // "void"
	tsChar                   // "char"
	tsShort                  // "short"
	tsInt                    // "int"
	tsLong                   // "long"
	tsFloat                  // "float"
	tsDouble                 // "double"
	tsSigned                 // "signed"
	tsUnsigned               // "unsigned"
	tsBool                   // "_Bool"
	tsComplex                // "_Complex"
	tsStructSpecifier        // StructOrUnionSpecifier: struct
	tsUnionSpecifier         // StructOrUnionSpecifier: union
	tsEnumSpecifier          // EnumSpecifier
	tsTypedefName            // TYPEDEFNAME
)

const (
	tsBits = 4 // Values [0, 13]
	tsMask = 1<<tsBits - 1
)

// specifier attributes.
const (
	saInline = 1 << iota
	saTypedef
	saExtern
	saStatic
	saAuto
	saRegister
	saConst
	saRestrict
	saVolatile
)

func attrString(attr int) string {
	if attr == 0 {
		return ""
	}

	a := []string{}
	if attr&saAuto != 0 {
		a = append(a, "auto")
	}
	if attr&saConst != 0 {
		a = append(a, "const")
	}
	if attr&saExtern != 0 {
		a = append(a, "extern")
	}
	if attr&saInline != 0 {
		a = append(a, "inline")
	}
	if attr&saRegister != 0 {
		a = append(a, "register")
	}
	if attr&saRestrict != 0 {
		a = append(a, "restrict")
	}
	if attr&saStatic != 0 {
		a = append(a, "static")
	}
	if attr&saTypedef != 0 {
		a = append(a, "typedef")
	}
	if attr&saVolatile != 0 {
		a = append(a, "volatile")
	}
	return strings.Join(a, " ")
}

// PPTokenList represents a sequence of tokens.
type PPTokenList int

func (p PPTokenList) Pos() token.Pos {
	if p == 0 {
		return 0
	}

	return decodeTokens(p, nil)[0].Pos()
}

// Linkage is a C linkage kind ([0], 6.2.2, p. 30)
type Linkage int

// Values of type Linkage.
const (
	None Linkage = iota
	Internal
	External
)

// Kind is a type category. Kind formally implements Type the only method
// returning a non nil value is Kind.
type Kind int

// Values of type Kind.
const (
	Undefined Kind = iota
	Void
	Ptr
	Char
	SChar
	UChar
	Short
	UShort
	Int
	UInt
	Long
	ULong
	LongLong
	ULongLong
	Float
	Double
	LongDouble
	Bool
	FloatComplex
	DoubleComplex
	LongDoubleComplex
	Struct
	Union
	Enum
	TypedefName
	Function
	Array
)

func (k Kind) CString() string {
	switch k {
	case Undefined:
		return "undefined"
	case Void:
		return "void"
	case Ptr:
		return "pointer"
	case Char:
		return "char"
	case SChar:
		return "signed char"
	case UChar:
		return "unsigned char"
	case Short:
		return "short"
	case UShort:
		return "unsigned short"
	case Int:
		return "int"
	case UInt:
		return "unsigned"
	case Long:
		return "long"
	case ULong:
		return "unsigned long"
	case LongLong:
		return "long long"
	case ULongLong:
		return "unsigned long long"
	case Float:
		return "float"
	case Double:
		return "double"
	case LongDouble:
		return "long double"
	case Bool:
		return "bool"
	case FloatComplex:
		return "float complex"
	case DoubleComplex:
		return "double complex"
	case LongDoubleComplex:
		return "long double complex"
	case Struct:
		return "struct"
	case Union:
		return "union"
	case Enum:
		return "enum"
	case TypedefName:
		return "typedefname"
	case Function:
		return "function"
	case Array:
		return "array"
	default:
		panic("internal error")
	}
}

// Scope is a bindings category.
type Scope int

// Values of type Scope
const (
	ScopeFile Scope = iota
	ScopeBlock
	ScopeMembers
	ScopeParams
)

// Namespace is a binding category.
type Namespace int

// Values of type Namespace.
const (
	NSIdentifiers Namespace = iota
	NSTags
)

var (
	cwords = map[int]rune{
		dict.SID("auto"):     AUTO,
		dict.SID("_Bool"):    BOOL,
		dict.SID("break"):    BREAK,
		dict.SID("case"):     CASE,
		dict.SID("char"):     CHAR,
		dict.SID("_Complex"): COMPLEX,
		dict.SID("const"):    CONST,
		dict.SID("continue"): CONTINUE,
		dict.SID("default"):  DEFAULT,
		dict.SID("do"):       DO,
		dict.SID("double"):   DOUBLE,
		dict.SID("else"):     ELSE,
		dict.SID("enum"):     ENUM,
		dict.SID("extern"):   EXTERN,
		dict.SID("float"):    FLOAT,
		dict.SID("for"):      FOR,
		dict.SID("goto"):     GOTO,
		dict.SID("if"):       IF,
		dict.SID("inline"):   INLINE,
		dict.SID("int"):      INT,
		dict.SID("long"):     LONG,
		dict.SID("register"): REGISTER,
		dict.SID("restrict"): RESTRICT,
		dict.SID("return"):   RETURN,
		dict.SID("short"):    SHORT,
		dict.SID("signed"):   SIGNED,
		dict.SID("sizeof"):   SIZEOF,
		dict.SID("static"):   STATIC,
		dict.SID("struct"):   STRUCT,
		dict.SID("switch"):   SWITCH,
		dict.SID("typedef"):  TYPEDEF,
		dict.SID("union"):    UNION,
		dict.SID("unsigned"): UNSIGNED,
		dict.SID("void"):     VOID,
		dict.SID("volatile"): VOLATILE,
		dict.SID("while"):    WHILE,
	}

	tokConstVals = map[rune]int{
		ADDASSIGN: dict.SID("+="),
		ANDAND:    dict.SID("&&"),
		ANDASSIGN: dict.SID("&="),
		ARROW:     dict.SID("->"),
		AUTO:      dict.SID("auto"),
		BOOL:      dict.SID("_Bool"),
		BREAK:     dict.SID("break"),
		CASE:      dict.SID("case"),
		CHAR:      dict.SID("char"),
		COMPLEX:   dict.SID("_Complex"),
		CONST:     dict.SID("const"),
		CONTINUE:  dict.SID("continue"),
		DDD:       dict.SID("..."),
		DEC:       dict.SID("--"),
		DEFAULT:   dict.SID("default"),
		DIVASSIGN: dict.SID("/="),
		DO:        dict.SID("do"),
		DOUBLE:    dict.SID("double"),
		ELSE:      dict.SID("else"),
		ENUM:      dict.SID("enum"),
		EQ:        dict.SID("=="),
		EXTERN:    dict.SID("extrn"),
		FLOAT:     dict.SID("float"),
		FOR:       dict.SID("for"),
		GEQ:       dict.SID(">="),
		GOTO:      dict.SID("goto"),
		IF:        dict.SID("if"),
		INC:       dict.SID("++"),
		INLINE:    dict.SID("inline"),
		INT:       dict.SID("int"),
		LEQ:       dict.SID("<="),
		LONG:      dict.SID("long"),
		LSH:       dict.SID("<<"),
		LSHASSIGN: dict.SID("<<="),
		MODASSIGN: dict.SID("%="),
		MULASSIGN: dict.SID("*="),
		NEQ:       dict.SID("!="),
		ORASSIGN:  dict.SID("|="),
		OROR:      dict.SID("||"),
		PPPASTE:   dict.SID("##"),
		REGISTER:  dict.SID("register"),
		RESTRICT:  dict.SID("restrict"),
		RETURN:    dict.SID("return"),
		RSH:       dict.SID(">>"),
		RSHASSIGN: dict.SID(">>="),
		SHORT:     dict.SID("short"),
		SIGNED:    dict.SID("signed"),
		SIZEOF:    dict.SID("sizeof"),
		STATIC:    dict.SID("static"),
		STRUCT:    dict.SID("struct"),
		SUBASSIGN: dict.SID("-="),
		SWITCH:    dict.SID("switch"),
		TYPEDEF:   dict.SID("typedef"),
		UNION:     dict.SID("union"),
		UNSIGNED:  dict.SID("unsigned"),
		VOID:      dict.SID("void"),
		VOLATILE:  dict.SID("volatile"),
		WHILE:     dict.SID("while"),
		XORASSIGN: dict.SID("^="),
	}

	id0                = dict.SID("0")
	id1                = dict.SID("1")
	idDate             = dict.SID("__DATE__")
	idDefined          = dict.SID("defined")
	idEmptyString      = dict.SID(`""`)
	idFile             = dict.SID("__FILE__")
	idID               = dict.SID("ID")
	idLine             = dict.SID("__LINE__")
	idSTDC             = dict.SID("__STDC__")
	idSTDCHosted       = dict.SID("__STDC_HOSTED__")
	idSTDCMBMightNeqWc = dict.SID("__STDC_MB_MIGHT_NEQ_WC__")
	idSTDCVersion      = dict.SID("__STDC_VERSION__")
	idTDate            = dict.SID(tuTime.Format("Jan _2 2006")) // The date of translation of the preprocessing translation unit.
	idTTtime           = dict.SID(tuTime.Format("15:04:05"))    // The time of translation of the preprocessing translation unit.
	idTime             = dict.SID("__TIME__")
	idVAARGS           = dict.SID("__VA_ARGS__")
	tuTime             = time.Now()

	tokHasVal = map[rune]bool{
		CHARCONST:         true,
		FLOATCONST:        true,
		IDENTIFIER:        true,
		IDENTIFIER_LPAREN: true,
		INTCONST:          true,
		LONGCHARCONST:     true,
		LONGSTRINGLITERAL: true,
		PPHEADER_NAME:     true,
		PPNUMBER:          true,
		STRINGLITERAL:     true,
	}

	// Valid combinations of TypeSpecifier.Case ([0], 6.7.2, 2)
	tsValid = map[int]Kind{
		tsEncode(tsBool):                            Bool,              // _Bool
		tsEncode(tsChar):                            Char,              // char
		tsEncode(tsDouble):                          Double,            // double
		tsEncode(tsDouble, tsComplex):               DoubleComplex,     // double _Complex
		tsEncode(tsEnumSpecifier):                   Enum,              // enum specifier
		tsEncode(tsFloat):                           Float,             // float
		tsEncode(tsFloat, tsComplex):                FloatComplex,      // float _Complex
		tsEncode(tsInt):                             Int,               // int
		tsEncode(tsInt):                             Int,               // signed int
		tsEncode(tsLong):                            Long,              // long
		tsEncode(tsLong, tsDouble):                  LongDouble,        // long double
		tsEncode(tsLong, tsDouble, tsComplex):       LongDoubleComplex, // long double _Complex
		tsEncode(tsLong, tsInt):                     Long,              // long int
		tsEncode(tsLong, tsLong):                    LongLong,          // long long
		tsEncode(tsLong, tsLong, tsInt):             LongLong,          // long long int
		tsEncode(tsShort):                           Short,             // short
		tsEncode(tsShort, tsInt):                    Short,             // short int
		tsEncode(tsSigned):                          Int,               // signed
		tsEncode(tsSigned, tsChar):                  SChar,             // signed char
		tsEncode(tsSigned, tsLong):                  Long,              // signed long
		tsEncode(tsSigned, tsLong, tsInt):           Long,              // signed long int
		tsEncode(tsSigned, tsLong, tsLong):          LongLong,          // signed long long
		tsEncode(tsSigned, tsLong, tsLong, tsInt):   LongLong,          // signed long long int
		tsEncode(tsSigned, tsShort):                 Short,             // signed short
		tsEncode(tsSigned, tsShort, tsInt):          Short,             // signed short int
		tsEncode(tsStructSpecifier):                 Struct,            // struct
		tsEncode(tsTypedefName):                     TypedefName,       // typedef name
		tsEncode(tsUnionSpecifier):                  Union,             // union
		tsEncode(tsUnsigned):                        UInt,              // unsigned
		tsEncode(tsUnsigned, tsChar):                UChar,             // unsigned char
		tsEncode(tsUnsigned, tsInt):                 UInt,              // unsigned int
		tsEncode(tsUnsigned, tsLong):                ULong,             // unsigned long
		tsEncode(tsUnsigned, tsLong, tsInt):         ULong,             // unsigned long int
		tsEncode(tsUnsigned, tsLong, tsLong):        ULongLong,         // unsigned long long
		tsEncode(tsUnsigned, tsLong, tsLong, tsInt): ULongLong,         // unsigned long long int
		tsEncode(tsUnsigned, tsShort):               UShort,            // unsigned short
		tsEncode(tsUnsigned, tsShort, tsInt):        UShort,            // unsigned short int
		tsEncode(tsVoid):                            Void,              // void
	}
)

func isUCNDigit(r rune) bool {
	return int(r) < len(ucnDigits)<<bitShift && ucnDigits[uint(r)>>bitShift]&(1<<uint(r&bitMask)) != 0
}

func isUCNNonDigit(r rune) bool {
	return int(r) < len(ucnNonDigits)<<bitShift && ucnNonDigits[uint(r)>>bitShift]&(1<<uint(r&bitMask)) != 0
}

func rune2class(r rune) (c int) {
	switch {
	case r == lex.RuneEOF:
		return ccEOF
	case r < 128:
		return int(r)
	case isUCNDigit(r):
		return ccUCNDigit
	case isUCNNonDigit(r):
		return ccUCNNonDigit
	default:
		return ccOther
	}
}

func toC(t xc.Token) xc.Token {
	if t.Rune != IDENTIFIER {
		return t
	}

	if x, ok := cwords[t.Val]; ok {
		t.Rune = x
	}

	return t
}

func tsEncode(a ...int) (r int) {
	sort.Ints(a)
	for _, v := range a {
		r = r<<tsBits | v
	}
	return r<<1 | 1 // Bit 0 set: value is valid.
}

func tsDecode(n int) (r []int) {
	if n == 0 {
		return nil
	}

	n >>= 1 // Remove value is valid bit.
	for n != 0 {
		r = append(r, n&tsMask)
		n >>= tsBits
	}
	return r
}

func tsHasTypedefname(n int) bool {
	if n == 0 {
		return false
	}

	n >>= 1 // Remove value is valid bit.
	for n != 0 {
		if n&tsMask == tsTypedefName {
			return true
		}

		n >>= tsBits
	}
	return false
}

func (l *lexer) encodeToken(tok xc.Token) {
	n := binary.PutUvarint(l.encBuf1[:], uint64(tok.Rune))
	pos := tok.Pos()
	n += binary.PutUvarint(l.encBuf1[n:], uint64(pos-l.encPos))
	l.encPos = pos
	if tokHasVal[tok.Rune] {
		n += binary.PutUvarint(l.encBuf1[n:], uint64(tok.Val))
	}
	l.encBuf = append(l.encBuf, l.encBuf1[:n]...)
}

func decodeToken(p *[]byte, pos *token.Pos) xc.Token {
	b := *p
	r, n := binary.Uvarint(b)
	b = b[n:]
	d, n := binary.Uvarint(b)
	b = b[n:]
	np := *pos + token.Pos(d)
	*pos = np
	c := lex.NewChar(np, rune(r))
	var v uint64
	if tokHasVal[c.Rune] {
		v, n = binary.Uvarint(b)
		b = b[n:]
	}
	*p = b
	return xc.Token{Char: c, Val: int(v)}
}

func decodeTokens(id PPTokenList, r []xc.Token) []xc.Token {
	b := dict.S(int(id))
	var pos token.Pos
	r = r[:0]
	for len(b) != 0 {
		r = append(r, decodeToken(&b, &pos))
	}
	return r
}

func tokVal(t xc.Token) int {
	r := t.Rune
	if tokHasVal[r] {
		return t.Val
	}

	if r != 0 && r < 0x80 {
		return int(r) + 1
	}

	if i, ok := tokConstVals[r]; ok {
		return i
	}

	panic("internal error")
}

// TokSrc returns t in its source form.
func TokSrc(t xc.Token) string {
	if x, ok := tokConstVals[t.Rune]; ok {
		return string(dict.S(x))
	}

	if tokHasVal[t.Rune] {
		return string(t.S())
	}

	return string(t.Rune)
}

// universal-character-name	\\u{hex-quad}|\\U{hex-quad}{hex-quad}
func decodeUCN(runes []rune) (rune, int) {
	if runes[0] != '\\' {
		panic("internal error")
	}

	runes = runes[1:]
	switch runes[0] {
	case 'u':
		return rune(decodeHexQuad(runes[1:])), 6
	case 'U':
		return rune(decodeHexQuad(runes[1:])<<16 | decodeHexQuad(runes[5:])), 10
	default:
		panic("internal error")
	}
}

// hex-quad	{hexadecimal-digit}{hexadecimal-digit}{hexadecimal-digit}{hexadecimal-digit}
func decodeHexQuad(runes []rune) int {
	n := 0
	for _, r := range runes[:4] {
		n = n<<4 | decodeHex(r)
	}
	return n
}

func decodeHex(r rune) int {
	switch {
	case r >= '0' && r <= '9':
		return int(r) - '0'
	default:
		x := int(r) &^ 0x20
		return x - 'A' + 10
	}
}

// escape-sequence		{simple-sequence}|{octal-escape-sequence}|{hexadecimal-escape-sequence}|{universal-character-name}
// simple-sequence		\\['\x22?\\abfnrtv]
// octal-escape-sequence	\\{octal-digit}{octal-digit}?{octal-digit}?
// hexadecimal-escape-sequence	\\x{hexadecimal-digit}+
func decodeEscapeSequence(runes []rune) (rune, int) {
	if runes[0] != '\\' {
		panic("internal error")
	}

	r := runes[1]
	switch r {
	case '\'', '"', '?', '\\':
		return r, 2
	case 'a':
		return 7, 2
	case 'b':
		return 8, 2
	case 'f':
		return 12, 2
	case 'n':
		return 10, 2
	case 'r':
		return 13, 2
	case 't':
		return 9, 2
	case 'v':
		return 11, 2
	case 'x':
		v, n := 0, 2
		for _, r := range runes[2:] {
			switch {
			case r >= '0' && r <= '9', r >= 'a' && r <= 'f', r >= 'A' && r <= 'F':
				v = v<<4 | decodeHex(r)
				n++
			default:
				break
			}
		}
		return rune(v & 0xff), n
	case 'u', 'U':
		return decodeUCN(runes)
	}

	if r < '0' || r > '7' {
		panic("internal error")
	}

	v, n := 0, 1
	for _, r := range runes[1:] {
		switch {
		case r >= '0' && r <= '7':
			v = v<<3 | int(r) - '0'
			n++
		default:
			break
		}
	}
	return rune(n & 0xff), n
}
