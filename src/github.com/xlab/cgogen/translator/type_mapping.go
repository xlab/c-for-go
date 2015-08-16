package translator

type CGOTypeMap map[string]string
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
	//
	InterfaceSliceSpec = GoTypeSpec{Base: "[]interface{}"}
)

var builtinCGOTypeMap = CGOTypeMap{
	IntSpec.String():         "C.int",
	UintSpec.String():        "C.uint",
	Int8Spec.String():        "C.char",
	Int16Spec.String():       "C.short",
	Int32Spec.String():       "C.long",
	Int64Spec.String():       "C.int",
	Uint8Spec.String():       "C.uchar",
	Uint16Spec.String():      "C.ushort",
	Uint32Spec.String():      "C.uint",
	Uint64Spec.String():      "C.ulong",
	RuneSpec.String():        "C.int",
	RuneSliceSpec.String():   "*C.int",
	ByteSliceSpec.String():   "*C.char",
	StringSpec.String():      "*C.char",
	StringSliceSpec.String(): "**C.char",
	Float32Spec.String():     "C.float",
	Float64Spec.String():     "C.double",
	PointerSpec.String():     "unsafe.Pointer",
	UintptrSpec.String():     "unsafe.Pointer",
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
	// tr1/cstdarg
	CTypeSpec{Base: "va_list"}: InterfaceSliceSpec,
}

// TODO consider:
// > const char** -> []string
// or should it be *string? how about:
// > const char*** -> []*string OR *[]string or even [][]string
// TODO: make some tests and CGO evaluations
