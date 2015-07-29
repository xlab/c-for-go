package translator

import (
	"bytes"
	"errors"
	"fmt"
	"go/token"
	"sort"
	"strconv"
)

type CDefKind int

const (
	TypeDef CDefKind = iota
	StructDef
	FunctionDef
)

type CDef interface {
	Kind() CDefKind
	String() string
}

type CTypeDecl struct {
	Spec CDef
	Name string
	Pos  token.Pos
}

func (c *CTypeDecl) String() string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf("%s %s", c.Spec.String(), c.Name)
}

type CFunctionSpec struct {
	Returns  CTypeDecl
	Members  []CTypeDecl
	Pointers uint8
}

func (c *CFunctionSpec) Kind() CDefKind {
	return FunctionDef
}

func (c *CFunctionSpec) String() string {
	return "// TODO: function spec"
}

type CStructSpec struct {
	Tag      string
	Union    bool
	Members  []CTypeDecl
	Pointers uint8
}

func (c *CStructSpec) Kind() CDefKind {
	return StructDef
}

func (c *CStructSpec) String() string {
	return "// TODO: struct spec"
}

type CTypeSpec struct {
	Base     string
	Const    bool
	Unsigned bool
	Short    bool
	Long     bool
	Struct   bool
	Pointers uint8
}

func (c *CTypeSpec) Kind() CDefKind {
	return TypeDef
}

func (cts *CTypeSpec) String() string {
	var str string
	if cts.Const {
		str += "const "
	}
	switch {
	case cts.Unsigned:
		str += "unsigned "
	case cts.Struct:
		str += "struct "
	}
	switch {
	case cts.Long:
		str += "long "
	case cts.Short:
		str += "short "
	}
	str += cts.Base
	for i := uint8(0); i < cts.Pointers; i++ {
		str += "*"
	}
	return str
}

func (cts *CTypeSpec) MarshalJSON() ([]byte, error) {
	if cts == nil {
		return nil, nil
	}
	if len(cts.Base) == 0 {
		return nil, errors.New("base type isn't specified")
	}
	return []byte(cts.String()), nil
}

type bytesSlice [][]byte

func (s bytesSlice) Len() int           { return len(s) }
func (s bytesSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s bytesSlice) Less(i, j int) bool { return bytes.Compare(s[i], s[j]) < 0 }

var (
	qualConst       = []byte("const")
	specStruct      = []byte("struct")
	specUnion       = []byte("union")
	specUnsigned    = []byte("unsigned")
	specSigned      = []byte("signed")
	specLong        = []byte("long")
	specShort       = []byte("short")
	ptrStr          = []byte("*")
	sliceStr        = []byte("[]")
	spaceStr        = []byte(" ")
	emptyStr        = []byte{}
	restrictedNames = bytes.Join([][]byte{
		qualConst, specStruct, specUnion, specUnsigned, specSigned, specShort,
	}, spaceStr)
)

func (cts *CTypeSpec) UnmarshalJSON(b []byte) error {
	parts := bytes.Split(b, spaceStr)
	if len(parts) == 0 {
		return errors.New("unexpected EOF")
	}
	ts := CTypeSpec{}
	sort.Reverse(bytesSlice(parts))

	// states:
	// 0 — pointers
	// 1 — base
	// 2 — qualifiers
	var state int
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		switch state {
		case 0:
			// read pointers count
			for bytes.HasSuffix(part, ptrStr) {
				ts.Pointers++
				part = part[:len(part)-1]
			}
			state = 1
		case 1:
			// read the base name
			if bytes.Contains(restrictedNames, part) {
				return errors.New("ctype: can't use keyword as a base type name: " + string(part))
			}
			ts.Base = string(part)
			state = 2
		case 2:
			// read specifiers and qualifiers
			switch {
			case bytes.Equal(part, specStruct), bytes.Equal(part, specUnion):
				ts.Struct = true
				if len(ts.Base) == 0 {
					return errors.New("ctype: no base type name specified")
				}
				*cts = ts
				return nil
			case bytes.Equal(part, specShort):
				ts.Short = true
			case bytes.Equal(part, specLong):
				ts.Long = true
			case bytes.Equal(part, specUnsigned):
				ts.Unsigned = true
			case bytes.Equal(part, qualConst):
				ts.Const = true
			}
		}
	}

	if len(ts.Base) == 0 {
		return errors.New("ctype: no base type name specified")
	}
	*cts = ts
	return nil
}

// func (cts CTypeSpec)  string {
// 	return cts.Base // TODO: lookup in cache
// }

type GoTypeSpec struct {
	Slices   uint8
	Pointers uint8
	Unsigned bool
	Base     string
	Inner    *GoTypeSpec
	Bits     uint16
}

func (gts *GoTypeSpec) Wrap(innerGTS GoTypeSpec) GoTypeSpec {
	return GoTypeSpec{
		Slices:   gts.Slices,
		Pointers: gts.Pointers,
		Inner:    &innerGTS,
	}
}

func (gts *GoTypeSpec) String() string {
	var str string
	for i := uint8(0); i < gts.Slices; i++ {
		str += "[]"
	}
	for i := uint8(0); i < gts.Pointers; i++ {
		str += "*"
	}
	if gts.Inner != nil {
		str += gts.Inner.String()
	} else {
		if gts.Unsigned {
			str += "u"
		}
		str += gts.Base
		if gts.Bits > 0 {
			str += strconv.Itoa(int(gts.Bits))
		}
	}
	return str
}

func (gts *GoTypeSpec) MarshalJSON() ([]byte, error) {
	if gts == nil {
		return []byte{}, nil
	}
	return []byte(gts.String()), nil
}

func (gts *GoTypeSpec) UnmarshalJSON(b []byte) error {
	// purge any spaces
	b = bytes.Replace(b, spaceStr, emptyStr, -1)

	// states:
	// 0 — beginning
	// 1 — processing slices
	// 2 — processing pointers
	// 3
	// 3 — base
	var state int
	// var ts = GoTypeSpec{}

	for {
		switch state {
		case 0:
			switch {
			case bytes.HasPrefix(b, sliceStr):
				state = 1
				continue
			case bytes.HasPrefix(b, ptrStr):
				state = 2
				continue
			}
		}
	}
	// TODO
	return nil
}

type CTypeMap map[CTypeSpec]GoTypeSpec
type GoTypeMap map[string]GoTypeSpec

var (
	BoolSpec        = GoTypeSpec{Base: "bool"}
	IntSpec         = GoTypeSpec{Base: "int"}
	UintSpec        = GoTypeSpec{Base: "int", Unsigned: true}
	Int8Spec        = GoTypeSpec{Base: "int", Bits: 8}
	Uint8Spec       = GoTypeSpec{Base: "int", Bits: 8, Unsigned: true}
	Int16Spec       = GoTypeSpec{Base: "int", Bits: 16}
	Uint16Spec      = GoTypeSpec{Base: "int", Bits: 16, Unsigned: true}
	Int32Spec       = GoTypeSpec{Base: "int", Bits: 32}
	Uint32Spec      = GoTypeSpec{Base: "int", Bits: 32, Unsigned: true}
	Int64Spec       = GoTypeSpec{Base: "int", Bits: 64}
	Uint64Spec      = GoTypeSpec{Base: "int", Bits: 64, Unsigned: true}
	RuneSpec        = GoTypeSpec{Base: "rune"}
	RuneSliceSpec   = GoTypeSpec{Base: "rune", Slices: 1}
	ByteSpec        = GoTypeSpec{Base: "byte"}
	ByteSliceSpec   = GoTypeSpec{Base: "byte", Slices: 1}
	StringSpec      = GoTypeSpec{Base: "string"}
	StringSliceSpec = GoTypeSpec{Base: "string", Slices: 1}
	Float32Spec     = GoTypeSpec{Base: "float", Bits: 32}
	Float64Spec     = GoTypeSpec{Base: "float", Bits: 64}
	PointerSpec     = GoTypeSpec{Base: "unsafe.Pointer"}
	UintptrSpec     = GoTypeSpec{Base: "uintptr"}
)

var builtinGoTypeMap = GoTypeMap{
	BoolSpec.String():        BoolSpec,
	IntSpec.String():         IntSpec,
	UintSpec.String():        UintSpec,
	Int8Spec.String():        Int8Spec,
	Int16Spec.String():       Int16Spec,
	Int32Spec.String():       Int32Spec,
	Int64Spec.String():       Int64Spec,
	Uint8Spec.String():       Uint8Spec,
	Uint16Spec.String():      Uint16Spec,
	Uint32Spec.String():      Uint32Spec,
	Uint64Spec.String():      Uint64Spec,
	RuneSpec.String():        RuneSpec,
	RuneSliceSpec.String():   RuneSliceSpec,
	ByteSliceSpec.String():   ByteSliceSpec,
	StringSpec.String():      StringSpec,
	StringSliceSpec.String(): StringSliceSpec,
	Float32Spec.String():     Float32Spec,
	Float64Spec.String():     Float64Spec,
	PointerSpec.String():     PointerSpec,
	UintptrSpec.String():     UintptrSpec,
}

var builtinCTypeMap = CTypeMap{
	// CHAR TYPES
	// ----------
	// char -> int8
	CTypeSpec{Base: "char"}: Int8Spec,
	// char* -> []byte
	CTypeSpec{Base: "char", Pointers: 1}: ByteSliceSpec,
	// const char* -> string
	CTypeSpec{Base: "char", Pointers: 1}: StringSpec,
	// unsigned char -> byte
	CTypeSpec{Base: "char", Unsigned: true}: ByteSpec,
	// unsigned char* -> []byte
	CTypeSpec{Base: "char", Unsigned: true, Pointers: 1}: ByteSliceSpec,
	// const unsigned char* -> string
	CTypeSpec{Base: "char", Const: true, Unsigned: true, Pointers: 1}: StringSpec,

	// SHORT TYPES
	// -----------
	// short -> int16
	CTypeSpec{Base: "short"}: Int16Spec,
	// unsigned short -> uint16
	CTypeSpec{Base: "short", Unsigned: true}: Uint16Spec,

	// LONG TYPES
	// ----------
	// long -> int64
	CTypeSpec{Base: "long"}: Int64Spec,
	// unsigned long -> uint64
	CTypeSpec{Base: "long", Unsigned: true}: Uint64Spec,
	// long long -> int64
	CTypeSpec{Base: "long", Long: true}: Int64Spec,
	// unsigned long long -> uint64
	CTypeSpec{Base: "long", Long: true, Unsigned: true}: Uint64Spec,

	// INT TYPES
	// ----------
	// int -> int32
	CTypeSpec{Base: "int"}: Int32Spec,
	// unsigned int -> uint32
	CTypeSpec{Base: "int", Unsigned: true}: Uint32Spec,
	// short int -> int16
	CTypeSpec{Base: "int", Short: true}: Int16Spec,
	// unsigned short int -> uint16
	CTypeSpec{Base: "int", Short: true, Unsigned: true}: Uint16Spec,

	// FLOAT TYPES
	// ----------
	// float -> float32
	CTypeSpec{Base: "float"}: Float32Spec,
	// double -> float64
	CTypeSpec{Base: "double"}: Float64Spec,

	// OTHER TYPES
	// ----------
	// void* -> unsafe.Pointer
	CTypeSpec{Base: "void", Pointers: 1}: PointerSpec,

	// DEFINED TYPES
	// ----------
	CTypeSpec{Base: "bool"}:      BoolSpec,
	CTypeSpec{Base: "_Bool"}:     BoolSpec, // C99
	CTypeSpec{Base: "ssize_t"}:   Int64Spec,
	CTypeSpec{Base: "size_t"}:    Uint64Spec,
	CTypeSpec{Base: "int_t"}:     IntSpec,
	CTypeSpec{Base: "uint_t"}:    UintSpec,
	CTypeSpec{Base: "int8_t"}:    Int8Spec,
	CTypeSpec{Base: "int16_t"}:   Int16Spec,
	CTypeSpec{Base: "int32_t"}:   Int32Spec,
	CTypeSpec{Base: "int64_t"}:   Int64Spec,
	CTypeSpec{Base: "uint8_t"}:   Uint8Spec,
	CTypeSpec{Base: "uint16_t"}:  Uint16Spec,
	CTypeSpec{Base: "uint32_t"}:  Uint32Spec,
	CTypeSpec{Base: "uint64_t"}:  Uint64Spec,
	CTypeSpec{Base: "intptr_t"}:  UintptrSpec,
	CTypeSpec{Base: "uintptr_t"}: UintptrSpec,
	// wchar_t -> rune
	CTypeSpec{Base: "wchar_t"}: RuneSpec,
	// wchar_t* -> []rune
	CTypeSpec{Base: "wchar_t", Pointers: 1}: RuneSliceSpec,
	// const wchar_t* -> string
	CTypeSpec{Base: "wchar_t", Const: true, Pointers: 1}: StringSpec,
}

// TODO consider:
// > const char** -> []string
// or should it be *string? how about:
// > const char*** -> []*string OR *[]string or even [][]string
// TODO: make some tests and CGO evaluations
