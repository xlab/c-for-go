package translator

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type CTypeSpec struct {
	Base      string
	Const     bool
	Unsigned  bool
	Short     bool
	Long      bool
	Arrays    string
	VarArrays uint8
	Pointers  uint8
}

func (c *CTypeSpec) AddArray(size uint32) {
	if size > 0 {
		c.Arrays = fmt.Sprintf("[%d]%s", size, c.Arrays)
		return
	}
	c.VarArrays++
}

func (cts CTypeSpec) String() string {
	buf := new(bytes.Buffer)
	if cts.Const {
		buf.WriteString("const ")
	}
	if cts.Unsigned {
		buf.WriteString("unsigned ")
	}
	switch {
	case cts.Long:
		buf.WriteString("long ")
	case cts.Short:
		buf.WriteString("short ")
	}
	fmt.Fprint(buf, cts.Base)
	buf.WriteString(strings.Repeat("*", int(cts.Pointers)))
	buf.WriteString(cts.Arrays)
	return buf.String()
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

func (c *CTypeSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c CTypeSpec) Kind() CTypeKind {
	return TypeKind
}

func (c CTypeSpec) Copy() CType {
	return &c
}

func (c *CTypeSpec) GetBase() string {
	return c.Base
}

func (c *CTypeSpec) GetArrays() string {
	return c.Arrays
}

func (c *CTypeSpec) GetVarArrays() uint8 {
	return c.VarArrays
}

func (c *CTypeSpec) GetPointers() uint8 {
	return c.Pointers
}

func (c *CTypeSpec) IsConst() bool {
	return c.Const
}
