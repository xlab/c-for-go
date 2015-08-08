package translator

import (
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
	SetPointers(n uint8)
	Kind() CTypeKind
	String() string
	Copy() CType
}

type (
	Value         interface{}
	Expression    []byte
	ArraySizeSpec []byte
)

type CDecl struct {
	Spec       CType
	Name       string
	Value      Value
	Expression Expression
	IsTypedef  bool
	IsDefine   bool
	Arrays     []ArraySizeSpec
	Pos        token.Pos
	Src        string
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
