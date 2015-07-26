// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"sort"
	"time"

	"github.com/cznic/c/internal/xc"
	"github.com/cznic/mathutil"
	"github.com/cznic/sortutil"
)

const (
	scINITIAL = 0

	intBits  = mathutil.IntBits
	bitShift = intBits>>6 + 5
	bitMask  = intBits - 1

	// Invalid runes used for signaling.
	runeEOF   = -1
	runeError = -2

	// Character class is an 8 bit encoding of an Unicode rune for the
	// golex generated FSM.
	//
	// Every ASCII rune is its own class.  DO NOT change any of the
	// existing values. Adding new classes is OK.
	ccEOF         = 0x80
	ccError       = 0x81
	ccOther       = 0x82 // Any other rune.
	ccUCNDigit    = 0x83 // [0], Annex D, Universal character names for identifiers - digits.
	ccUCNNonDigit = 0x84 // [0], Annex D, Universal character names for identifiers - non digits.
)

// TypeSpecifier cases and combinations.
const (
	// TypeSpecifier.Case
	tsVoid          = iota // 0
	tsChar                 // 1
	tsShort                // 2
	tsInt                  // 3
	tsLong                 // 4
	tsFloat                // 5
	tsDouble               // 6
	tsSigned               // 7
	tsUnsigned             // 8
	tsBool                 // 9
	tsComplex              // 10
	tsStructOrUnion        // 11
	tsEnumSpecifier        // 12
	tsTypedefname          // 13

	// Combinations/other.
	tsUChar     // 14
	tsUShort    // 15
	tsUInt      // 16
	tsULong     // 17
	tsLongLong  // 18
	tsULongLong // 19
	tsStruct    // 20
	tsUnion     // 21
)

var (
	classStr = map[int]string{
		ccEOF:         "ccEOF",
		ccError:       "ccError",
		ccOther:       "ccOther",
		ccUCNDigit:    "ccUCNDigit",
		ccUCNNonDigit: "ccUCNNonDigit",
	}

	typeSums = map[int]int{
		typeSum4(tsVoid):                            tsVoid,
		typeSum4(tsChar):                            tsChar,
		typeSum4(tsSigned, tsChar):                  tsChar,
		typeSum4(tsUnsigned, tsChar):                tsUChar,
		typeSum4(tsShort):                           tsShort,
		typeSum4(tsShort, tsInt):                    tsShort,
		typeSum4(tsSigned, tsShort):                 tsShort,
		typeSum4(tsSigned, tsShort, tsInt):          tsShort,
		typeSum4(tsUnsigned, tsShort):               tsUShort,
		typeSum4(tsUnsigned, tsShort, tsInt):        tsUShort,
		typeSum4(tsInt):                             tsInt,
		typeSum4(tsSigned, tsInt):                   tsInt,
		typeSum4(tsUnsigned):                        tsUInt,
		typeSum4(tsUnsigned, tsInt):                 tsUInt,
		typeSum4(tsLong):                            tsLong,
		typeSum4(tsLong, tsInt):                     tsLong,
		typeSum4(tsSigned, tsLong):                  tsLong,
		typeSum4(tsSigned, tsLong, tsInt):           tsLong,
		typeSum4(tsUnsigned, tsLong):                tsULong,
		typeSum4(tsUnsigned, tsLong, tsInt):         tsULong,
		typeSum4(tsLong, tsLong):                    tsLongLong,
		typeSum4(tsLong, tsLong, tsInt):             tsLongLong,
		typeSum4(tsSigned, tsLong, tsLong):          tsLongLong,
		typeSum4(tsSigned, tsLong, tsLong, tsInt):   tsLongLong,
		typeSum4(tsUnsigned, tsLong, tsLong):        tsULongLong,
		typeSum4(tsUnsigned, tsLong, tsLong, tsInt): tsULongLong,
		typeSum4(tsFloat):                           tsFloat,
		typeSum4(tsDouble):                          tsDouble,
		typeSum4(tsLong, tsDouble):                  tsDouble,
		typeSum4(tsBool):                            tsBool,
		typeSum4(tsComplex):                         tsComplex,
		typeSum4(tsStruct):                          tsStruct,
		typeSum4(tsUnion):                           tsUnion,
		typeSum4(tsEnumSpecifier):                   tsEnumSpecifier,
		typeSum4(tsTypedefname):                     tsTypedefname,
	}

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

	idDefined = dict.ID([]byte("defined"))
	idVAARGS  = dict.SID("__VA_ARGS__")

	// idSTDCIEC559        = dict.SID("__STDC_IEC_559__")
	// idSTDCIEC559Complex = dict.SID("__STDC_IEC_559_COMPLEX__")
	// idSTDCISO10646      = dict.SID("__STDC_ISO_10646__")

	idBuiltinOffsetof  = dict.SID("__builtin_offsetof")
	idDate             = dict.SID("__DATE__")
	idFile             = dict.SID("__FILE__")
	idLine             = dict.SID("__LINE__")
	idSTDC             = dict.SID("__STDC__")
	idSTDCHosted       = dict.SID("__STDC_HOSTED__")
	idSTDCMBMightNeqWc = dict.SID("__STDC_MB_MIGHT_NEQ_WC__")
	idSTDCVersion      = dict.SID("__STDC_VERSION__")
	idTDate            = dict.SID(tuTime.Format("Jan _2 2006")) // The date of translation of the preprocessing translation unit.
	idTTtime           = dict.SID(tuTime.Format("15:04:05"))    // The time of translation of the preprocessing translation unit.
	idTime             = dict.SID("__TIME__")
	tuTime             = time.Now()

	tokHasVal = map[rune]bool{
		CHARCONST:         true,
		FLOATCONST:        true,
		IDENTIFIER:        true,
		IDENTIFIER_LPAREN: true,
		INTCONST:          true,
		LONGCHARCONST:     true,
		LONGSTRINGLITERAL: true,
		MACRO_ARG:         true,
		MACRO_ARG_EMPTY:   true,
		PPHEADER_NAME:     true,
		PPNUMBER:          true,
		STRINGLITERAL:     true,
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

	id1s          = dict.SID(`"1"`)
	idEmptyString = dict.SID(`""`)
	idUnion       = dict.SID("union")
)

func init() {
	if ccEOF != 128 || ccError != 129 || ccOther != 130 ||
		ccUCNDigit != 131 || ccUCNNonDigit != 132 {
		panic("invalid character class constant value")
	}

	for r := range tokConstVals {
		if tokHasVal[r] {
			panic("internal error 006")
		}
	}
}

func typeSum4(arg ...int) int {
	if len(arg) > 4 {
		panic("internal error")
	}

	sort.Ints(arg)
	r := 0
	for _, v := range arg {
		if v == tsStructOrUnion {
			panic("internal error")
		}

		r |= r<<8 | v
	}
	return r
}

func typeSum1(n int) int {
	var a [4]byte
	for i := 0; i < 4; i++ {
		v := byte(n)
		if v == tsStructOrUnion {
			panic("internal error")
		}

		a[i] = v
		n >>= 8
	}
	sort.Sort(sortutil.ByteSlice(a[:]))
	r := 0
	for _, v := range a {
		r |= r<<8 | int(v)
	}
	return r
}

// TokSrc returns t in its source form.
func TokSrc(t xc.Token) string {
	if x, ok := tokConstVals[t.Rune]; ok {
		return string(dict.S(x))
	}

	if tokHasVal[t.Rune] {
		return string(dict.S(t.Val))
	}

	return string(t.Rune)
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

	panic("internal error 008")
}

func isUCNDigit(r rune) bool {
	return int(r) < len(ucnDigits)<<bitShift && ucnDigits[uint(r)>>bitShift]&(1<<uint(r&bitMask)) != 0
}

func isUCNNonDigit(r rune) bool {
	return int(r) < len(ucnNonDigits)<<bitShift && ucnNonDigits[uint(r)>>bitShift]&(1<<uint(r&bitMask)) != 0
}

func rune2class(r rune) int {
	switch {
	case r >= 0 && r < 128:
		return int(r)
	case r == runeEOF:
		return ccEOF
	case r == runeError:
		return ccError
	case isUCNDigit(r):
		return ccUCNDigit
	case isUCNNonDigit(r):
		return ccUCNNonDigit
	default:
		return ccOther
	}
}

func typeSpecifier2TypeKind(ts int) Kind {
	switch ts {
	case tsVoid:
		return VoidType
	case tsChar:
		return CharType
	case tsShort:
		return ShortType
	case tsInt:
		return IntType
	case tsLong:
		return LongType
	case tsFloat:
		return FloatType
	case tsDouble:
		return DoubleType
	case tsSigned:
		return IntType
	case tsUnsigned:
		return UIntType
	case tsBool:
		return BoolType
	case tsComplex:
		return ComplexType
	case tsEnumSpecifier:
		return IntType
	case tsTypedefname:
		return NamedType
	case tsUChar:
		return UCharType
	case tsUShort:
		return UShortType
	case tsUInt:
		return UIntType
	case tsULong:
		return ULongType
	case tsLongLong:
		return LongLongType
	case tsULongLong:
		return ULongLongType
	case tsStruct:
		return StructType
	case tsUnion:
		return UnionType
	default:
		panic("internal error")
	}
}
