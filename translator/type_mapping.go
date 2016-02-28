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
	Complex64Spec     = GoTypeSpec{Base: "complex", Bits: 64}
	Complex128Spec    = GoTypeSpec{Base: "complex", Bits: 128}
	UnsafePointerSpec = GoTypeSpec{Base: "unsafe.Pointer", Pointers: 1}
	VoidSpec          = GoTypeSpec{Base: "byte", OuterArr: "[0]"}
	//
	InterfaceSliceSpec = GoTypeSpec{Base: "[]interface{}"}
)

// https://en.wikipedia.org/wiki/C_data_types
var builtinCTypeMap = CTypeMap{
	// char -> byte
	CTypeSpec{Base: "char"}: ByteSpec,
	// const char* -> string
	CTypeSpec{Base: "char", Const: true, Pointers: 1}: StringSpec,
	// unsigned char -> byte
	CTypeSpec{Base: "char", Unsigned: true}: ByteSpec,
	// const unsigned char* -> string
	CTypeSpec{Base: "char", Const: true, Unsigned: true, Pointers: 1}: UStringSpec,
	// short -> int16
	CTypeSpec{Base: "short"}: Int16Spec,
	// unsigned short -> uint16
	CTypeSpec{Base: "short", Unsigned: true}: Uint16Spec,
	// long -> int
	CTypeSpec{Base: "long"}: IntSpec,
	// unsigned long -> uint
	CTypeSpec{Base: "long", Unsigned: true}: UintSpec,
	// signed long -> int
	CTypeSpec{Base: "long", Signed: true}: IntSpec,
	// long long -> int64
	CTypeSpec{Base: "long", Long: true}: Int64Spec,
	// unsigned long long -> uint64
	CTypeSpec{Base: "long", Long: true, Unsigned: true}: Uint64Spec,
	// signed long long -> int64
	CTypeSpec{Base: "long", Long: true, Signed: true}: Int64Spec,
	// int -> int32
	CTypeSpec{Base: "int"}: Int32Spec,
	// unsigned int -> uint32
	CTypeSpec{Base: "int", Unsigned: true}: Uint32Spec,
	// signed int -> int32
	CTypeSpec{Base: "int", Signed: true}: Int32Spec,
	// short int -> int16
	CTypeSpec{Base: "int", Short: true}: Int16Spec,
	// unsigned short int -> uint16
	CTypeSpec{Base: "int", Short: true, Unsigned: true}: Uint16Spec,
	// signed short int -> uint16
	CTypeSpec{Base: "int", Short: true, Signed: true}: Int16Spec,
	// long int -> int
	CTypeSpec{Base: "int", Long: true}: IntSpec,
	// unsigned long int -> uint
	CTypeSpec{Base: "int", Long: true, Unsigned: true}: UintSpec,
	// signed long int -> uint
	CTypeSpec{Base: "int", Long: true, Signed: true}: IntSpec,
	// float -> float32
	CTypeSpec{Base: "float"}: Float32Spec,
	// double -> float64
	CTypeSpec{Base: "double"}: Float64Spec,
	// long double -> float64
	CTypeSpec{Base: "double", Long: true}: Float64Spec,
	// complex float -> complex164
	CTypeSpec{Base: "float", Complex: true}: Complex64Spec,
	// complex double -> complex128
	CTypeSpec{Base: "double", Complex: true}: Complex128Spec,
	// long complex double -> complex128
	CTypeSpec{Base: "double", Long: true, Complex: true}: Complex128Spec,
	// void* -> unsafe.Pointer
	CTypeSpec{Base: "void", Pointers: 1}: UnsafePointerSpec,
	// void -> [0]byte
	CTypeSpec{Base: "void"}: VoidSpec,
	// _Bool -> bool
	CTypeSpec{Base: "_Bool"}: BoolSpec,
}
