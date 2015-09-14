package translator

type CTypeMap map[CTypeSpec]GoTypeSpec
type GoTypeMap map[string]GoTypeSpec

var (
	BoolSpec          = GoTypeSpec{Base: "bool"}
	IntSpec           = GoTypeSpec{Base: "int"}
	UintSpec          = GoTypeSpec{Base: "int", Unsigned: true}
	Int8Spec          = GoTypeSpec{Base: "int", Bits: 8}
	Uint8Spec         = GoTypeSpec{Base: "int", Bits: 8, Unsigned: true}
	Int16Spec         = GoTypeSpec{Base: "int", Bits: 16}
	Uint16Spec        = GoTypeSpec{Base: "int", Bits: 16, Unsigned: true}
	Int32Spec         = GoTypeSpec{Base: "int", Bits: 32}
	Uint32Spec        = GoTypeSpec{Base: "int", Bits: 32, Unsigned: true}
	Int64Spec         = GoTypeSpec{Base: "int", Bits: 64}
	Uint64Spec        = GoTypeSpec{Base: "int", Bits: 64, Unsigned: true}
	RuneSpec          = GoTypeSpec{Base: "rune"}
	ByteSpec          = GoTypeSpec{Base: "byte"}
	StringSpec        = GoTypeSpec{Base: "string"}
	UStringSpec       = GoTypeSpec{Base: "string", Unsigned: true}
	Float32Spec       = GoTypeSpec{Base: "float", Bits: 32}
	Float64Spec       = GoTypeSpec{Base: "float", Bits: 64}
	UnsafePointerSpec = GoTypeSpec{Base: "unsafe.Pointer"}
	VoidSpec          = GoTypeSpec{Base: "byte", Arrays: "[0]"}
	UintptrSpec       = GoTypeSpec{Base: "uintptr"}
	//
	InterfaceSliceSpec = GoTypeSpec{Base: "[]interface{}"}
)

var builtinCTypeMap = CTypeMap{
	// CHAR TYPES
	// ----------
	// char -> byte
	CTypeSpec{Base: "char"}: ByteSpec,

	// Sould work but it won't since C developers
	// sometimes may use char* buffers to read uint8_t from it. Well played.
	// char* -> string
	// CTypeSpec{Base: "char", Pointers: 1}: StringSpec,

	// const char* -> string
	CTypeSpec{Base: "char", Const: true, Pointers: 1}: StringSpec,
	// unsigned char -> byte
	CTypeSpec{Base: "char", Unsigned: true}: ByteSpec,
	// const unsigned char* -> string
	CTypeSpec{Base: "char", Const: true, Unsigned: true, Pointers: 1}: UStringSpec,

	// SHORT TYPES
	// -----------
	// short -> int16
	CTypeSpec{Base: "short"}: Int16Spec,
	// unsigned short -> uint16
	CTypeSpec{Base: "short", Unsigned: true}: Uint16Spec,

	// LONG TYPES
	// ----------
	// long -> int32
	CTypeSpec{Base: "long"}: Int32Spec,
	// unsigned long -> uint32
	CTypeSpec{Base: "long", Unsigned: true}: Uint32Spec,
	// long long -> int64
	CTypeSpec{Base: "long", Long: true}: Int64Spec,
	// unsigned long long -> uint64
	CTypeSpec{Base: "long", Long: true, Unsigned: true}: Uint64Spec,

	// INT TYPES
	// ----------
	// int -> int
	CTypeSpec{Base: "int"}: IntSpec,
	// unsigned int -> uint
	CTypeSpec{Base: "int", Unsigned: true}: UintSpec,
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
	CTypeSpec{Base: "void", Pointers: 1}: UnsafePointerSpec,
	// void -> [0]byte
	CTypeSpec{Base: "void"}: VoidSpec,

	// DEFINED TYPES
	// ----------
	CTypeSpec{Base: "bool"}:    BoolSpec,
	CTypeSpec{Base: "_Bool"}:   BoolSpec, // C99
	CTypeSpec{Base: "ssize_t"}: Int64Spec,
	CTypeSpec{Base: "size_t"}:  Uint64Spec,
	CTypeSpec{Base: "int_t"}:   IntSpec,
	CTypeSpec{Base: "uint_t"}:  UintSpec,

	CTypeSpec{Base: "int8_t", Const: true, Pointers: 1}:  StringSpec,
	CTypeSpec{Base: "uint8_t", Const: true, Pointers: 1}: UStringSpec,
	CTypeSpec{Base: "uint8_t"}:                           ByteSpec,

	CTypeSpec{Base: "int8_t"}:    Int8Spec,
	CTypeSpec{Base: "int16_t"}:   Int16Spec,
	CTypeSpec{Base: "int32_t"}:   Int32Spec,
	CTypeSpec{Base: "int64_t"}:   Int64Spec,
	CTypeSpec{Base: "uint16_t"}:  Uint16Spec,
	CTypeSpec{Base: "uint32_t"}:  Uint32Spec,
	CTypeSpec{Base: "uint64_t"}:  Uint64Spec,
	CTypeSpec{Base: "intptr_t"}:  UintptrSpec,
	CTypeSpec{Base: "uintptr_t"}: UintptrSpec,
	// wchar_t -> rune
	CTypeSpec{Base: "wchar_t"}: RuneSpec,
	// const wchar_t* -> string
	CTypeSpec{Base: "wchar_t", Const: true, Pointers: 1}: StringSpec,
	// tr1/cstdarg
	CTypeSpec{Base: "va_list"}: InterfaceSliceSpec,
}
