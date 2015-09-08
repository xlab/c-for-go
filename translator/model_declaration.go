package translator

import (
	"bytes"
	"fmt"
	"go/token"
)

type CTypeKind int

const (
	TypeKind CTypeKind = iota
	DefineKind
	StructKind
	UnionKind
	FunctionKind
	EnumKind
)

type CType interface {
	GetBase() string
	GetArrays() string
	GetVarArrays() uint8
	GetPointers() uint8
	SetPointers(uint8)
	AddArray(uint64)
	//
	IsConst() bool
	Kind() CTypeKind
	String() string
	Copy() CType
}

type (
	Value      interface{}
	Expression []byte
)

type CDecl struct {
	Spec       CType
	Name       string
	Value      Value
	Expression Expression
	IsStatic   bool
	IsTypedef  bool
	IsDefine   bool
	Pos        token.Pos
	Src        string
}

func (c *CDecl) IsTemplate() bool {
	switch spec := c.Spec.(type) {
	case *CStructSpec:
		return len(spec.Members) > 0
	case *CEnumSpec:
		return len(spec.Enumerators) > 0
	default:
		return true
	}
}

func (c *CDecl) Kind() CTypeKind {
	return c.Spec.Kind()
}

func (c *CDecl) IsConst() bool {
	return c.Spec.IsConst()
}

func (c *CDecl) SetPointers(n uint8) {
	c.Spec.SetPointers(n)
}

func (c *CDecl) AddArray(size uint64) {
	c.Spec.AddArray(size)
}

func (c CDecl) String() string {
	buf := new(bytes.Buffer)
	if len(c.Name) > 0 {
		fmt.Fprintf(buf, "%s %s", c.Spec, c.Name)
	} else {
		buf.WriteString(c.Spec.String())
	}
	if len(c.Expression) > 0 {
		fmt.Fprintf(buf, " = %s", string(c.Expression))
	}
	return buf.String()
}
