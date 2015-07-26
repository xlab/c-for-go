// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// stringer'ed enumerations.

package cc

// Kind describes a type category.
type Kind int

// Kind values
const (
	_ Kind = iota
	ArrayType
	BoolType
	CharType
	ComplexType
	DoubleType
	EnumType
	FloatType
	FunctionType
	IntType
	LongLongType
	LongType
	NamedType
	PtrType
	ShortType
	StructType
	UCharType
	UIntType
	ULongLongType
	ULongType
	UShortType
	UnionType
	VoidType
)

// ScalarType describes primitive types.
type ScalarType int

// ScalarType values.
const (
	_ ScalarType = iota
	Ptr
	Void
	Char
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
	Bool
	Complex
	Enum
)

// Alignof reports the required alignment of i.
func (i ScalarType) Alignof() int {
	switch i {
	case Ptr, Void, Char, UChar, Short, UShort, Int, UInt, Long, ULong, LongLong, ULongLong, Float, Double, Bool, Complex:
		return model[i].Align
	case Enum:
		return model[Int].Align
	default:
		panic("internal error")
	}
}

// Sizeof reports the size of i.
func (i ScalarType) Sizeof() int {
	switch i {
	case Ptr, Void, Char, UChar, Short, UShort, Int, UInt, Long, ULong, LongLong, ULongLong, Float, Double, Bool, Complex:
		return model[i].Size
	case Enum:
		return model[Int].Size
	default:
		panic("internal error")
	}
}

// Namespace represents C name spaces. ([0]6.2.3)
type Namespace int

// Namespace values.
const (
	_ Namespace = iota
	NSLabels
	NSTags
	NSMembers
	NSIdentifiers
)

// Scope represents the scope of Bindings.
type Scope int

// Scope values. ([0]6.2.1)
const (
	_ Scope = iota
	ScopeFile
	ScopeFunction // == also function prototype scope.
	ScopeBlock
	ScopeMembers // struct, union
	ScopeFnParams
)
