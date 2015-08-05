package translator

import (
	"bytes"
	"errors"
	"fmt"
	"go/token"
	"sort"
	"strconv"
	"strings"
)

type CTypeKind int

const (
	TypeDef CTypeKind = iota
	StructDef
	FunctionDef
	EnumDef
)

type CType interface {
	SetPointers(n uint8)
	Kind() CTypeKind
	String() string
	Copy() CType
}

type ArraySizeSpec []byte

type CDecl struct {
	Spec       CType
	Name       string
	Value      Value
	Expression Expression
	Arrays     []ArraySizeSpec
	Pos        token.Pos
}

func (c CDecl) String() string {
	var str string
	if len(c.Name) > 0 {
		str = c.Spec.String() + " " + c.Name
	} else {
		str = c.Spec.String()
	}
	for _, size := range c.Arrays {
		if size != nil {
			str += fmt.Sprintf("[%s]", size)
		} else {
			str += "[]"
		}
	}
	if len(c.Expression) > 0 {
		str += " = " + string(c.Expression)
	}
	return str
}

func (c *CDecl) SetPointers(n uint8) {
	c.Spec.SetPointers(n)
}

func (c *CDecl) AddArray(size []byte) {
	c.Arrays = append(c.Arrays, size)
}

type CFunctionSpec struct {
	Returns   CDecl
	ParamList []CDecl
	Pointers  uint8
}

func (c *CFunctionSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c CFunctionSpec) Kind() CTypeKind {
	return FunctionDef
}

func (c CFunctionSpec) Copy() CType {
	return &c
}

func (c CFunctionSpec) String() string {
	var params []string
	for _, param := range c.ParamList {
		params = append(params, param.String())
	}
	return fmt.Sprintf("%s (%s)", c.Returns, strings.Join(params, ", "))
}

type CEnumSpec struct {
	Tag         string
	Enumerators []CDecl
	Pointers    uint8
	Type        CTypeSpec
}

func (c *CEnumSpec) PromoteType(v Value) *CTypeSpec {
	var (
		uint32Spec = CTypeSpec{Base: "int", Unsigned: true}
		int32Spec  = CTypeSpec{Base: "int"}
		uint64Spec = CTypeSpec{Base: "long", Unsigned: true}
		int64Spec  = CTypeSpec{Base: "long"}
	)
	switch c.Type {
	case uint32Spec:
		switch v := v.(type) {
		case int32:
			if v < 0 {
				c.Type = int32Spec
			}
		case uint64:
			c.Type = uint64Spec
		case int64:
			if v < 0 {
				c.Type = int64Spec
			} else {
				c.Type = uint64Spec
			}
		}
	case int32Spec:
		switch v := v.(type) {
		case uint64:
			c.Type = uint64Spec
		case int64:
			if v < 0 {
				c.Type = int64Spec
			} else {
				c.Type = uint64Spec
			}
		}
	case uint64Spec:
		switch v := v.(type) {
		case int64:
			if v < 0 {
				c.Type = int64Spec
			}
		}
	default:
		switch v := v.(type) {
		case uint32:
			c.Type = uint32Spec
		case int32:
			if v < 0 {
				c.Type = int32Spec
			} else {
				c.Type = uint32Spec
			}
		case uint64:
			c.Type = uint64Spec
		case int64:
			if v < 0 {
				c.Type = int64Spec
			} else {
				c.Type = uint64Spec
			}
		}
	}
	return &c.Type
}

func (c *CEnumSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c CEnumSpec) Kind() CTypeKind {
	return EnumDef
}

func (c CEnumSpec) Copy() CType {
	return &c
}

func (ces CEnumSpec) String() string {
	var members []string
	for _, m := range ces.Enumerators {
		members = append(members, m.String())
	}
	membersColumn := strings.Join(members, ", ")

	str := "enum"
	if len(ces.Tag) > 0 {
		str = fmt.Sprintf("%s %s", str, ces.Tag)
	}
	if len(members) > 0 {
		str = fmt.Sprintf("%s {%s}", str, membersColumn)
	}
	if ces.Pointers > 0 {
		str += strings.Repeat("*", int(ces.Pointers))
	}
	return str
}

type CStructSpec struct {
	Tag      string
	Union    bool
	Members  []CDecl
	Pointers uint8
}

func (c *CStructSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c CStructSpec) Kind() CTypeKind {
	return StructDef
}

func (c CStructSpec) Copy() CType {
	return &c
}

func (css CStructSpec) String() string {
	var members []string
	for _, m := range css.Members {
		members = append(members, m.String())
	}
	membersColumn := strings.Join(members, ", ")

	str := "struct"
	if css.Union {
		str = "union"
	}
	if len(css.Tag) > 0 {
		str = fmt.Sprintf("%s %s", str, css.Tag)
	}
	if len(members) > 0 {
		str = fmt.Sprintf("%s {%s}", str, membersColumn)
	}
	if css.Pointers > 0 {
		str += strings.Repeat("*", int(css.Pointers))
	}
	return str
}

type CTypeSpec struct {
	Base     string
	Const    bool
	Unsigned bool
	Short    bool
	Long     bool
	Pointers uint8
}

func (c *CTypeSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c CTypeSpec) Kind() CTypeKind {
	return TypeDef
}

func (cts CTypeSpec) String() string {
	var str string
	if cts.Const {
		str += "const "
	}
	if cts.Unsigned {
		str += "unsigned "
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

func (c CTypeSpec) Copy() CType {
	return &c
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
			if isRestrictedBase(part) {
				return errors.New("ctype: can't use keyword as a base type name: " + string(part))
			}
			ts.Base = string(part)
			state = 2
		case 2:
			// read specifiers and qualifiers
			switch {
			case bytes.Equal(part, specStruct), bytes.Equal(part, specUnion):
				return errors.New("struct is not a simple C type")
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

type GoTypeSpec struct {
	Slices   uint8
	Pointers uint8
	Arrays   []uint64
	Inner    *GoTypeSpec
	Unsigned bool
	Base     string
	Bits     uint16
}

func (gts *GoTypeSpec) Wrap(innerGTS GoTypeSpec) GoTypeSpec {
	return GoTypeSpec{
		Slices:   gts.Slices,
		Pointers: gts.Pointers,
		Arrays:   gts.Arrays,
		Inner:    &innerGTS,
	}
}

func (gts *GoTypeSpec) String() string {
	var str string
	str += strings.Repeat("[]", int(gts.Slices))
	str += strings.Repeat("*", int(gts.Pointers))
	for _, size := range gts.Arrays {
		str += fmt.Sprintf("[%d]", size)
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

func CTypeOf(v interface{}) (*CTypeSpec, error) {
	switch x := v.(type) {
	case int32:
		return &CTypeSpec{Base: "int"}, nil
	case int64:
		return &CTypeSpec{Base: "long"}, nil
	case uint32:
		return &CTypeSpec{Base: "int", Unsigned: true}, nil
	case uint64:
		return &CTypeSpec{Base: "long", Unsigned: true}, nil
	case float32:
		return &CTypeSpec{Base: "float"}, nil
	case float64:
		return &CTypeSpec{Base: "double"}, nil
	case string:
		return &CTypeSpec{Base: "char", Pointers: 1, Const: true}, nil
	default:
		return nil, errors.New(fmt.Sprintf("cannot resolve type %T", x))
	}
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
