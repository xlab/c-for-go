package translator

import (
	"bytes"
	"fmt"
	"go/token"
)

type CTypeKind int

const (
	TypeKind CTypeKind = iota
	PlainTypeKind
	StructKind
	OpaqueStructKind
	UnionKind
	FunctionKind
	EnumKind
)

type CType interface {
	GetBase() string
	GetTag() string
	SetRaw(x string)
	CGoName() string
	GetPointers() uint8
	SetPointers(uint8)
	AddOuterArr(uint64)
	AddInnerArr(uint64)
	OuterArrays() ArraySpec
	InnerArrays() ArraySpec
	OuterArraySizes() []ArraySizeSpec
	InnerArraySizes() []ArraySizeSpec
	//
	IsConst() bool
	IsOpaque() bool
	IsComplete() bool
	Kind() CTypeKind
	String() string
	Copy() CType
}

type (
	Value interface{}
)

type CDecl struct {
	Spec       CType
	Name       string
	Value      Value
	Expression string
	IsStatic   bool
	IsTypedef  bool
	IsDefine   bool
	Pos        token.Pos
	Src        string
}

func (c CDecl) String() string {
	buf := new(bytes.Buffer)
	switch {
	case len(c.Name) > 0:
		fmt.Fprintf(buf, "%s %s", c.Spec, c.Name)
	default:
		buf.WriteString(c.Spec.String())
	}
	if len(c.Expression) > 0 {
		fmt.Fprintf(buf, " = %s", string(c.Expression))
	}
	return buf.String()
}
