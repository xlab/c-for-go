package translator

import (
	"bytes"
	"fmt"
)

type CTypeSpec struct {
	Raw      string
	Base     string
	Const    bool
	Signed   bool
	Unsigned bool
	Short    bool
	Long     bool
	Complex  bool
	Opaque   bool
	Pointers uint8
	InnerArr ArraySpec
	OuterArr ArraySpec
}

func (spec CTypeSpec) String() string {
	buf := new(bytes.Buffer)
	if spec.Unsigned {
		buf.WriteString("unsigned ")
	} else if spec.Signed {
		buf.WriteString("signed ")
	}
	switch {
	case spec.Long:
		buf.WriteString("long ")
	case spec.Short:
		buf.WriteString("short ")
	case spec.Complex:
		buf.WriteString("complex ")
	}
	fmt.Fprint(buf, spec.Base)

	var unsafePointer uint8
	if spec.Base == "unsafe.Pointer" {
		unsafePointer = 1
	}

	buf.WriteString(arrs(spec.OuterArr))
	buf.WriteString(ptrs(spec.Pointers - unsafePointer))
	buf.WriteString(arrs(spec.InnerArr))
	return buf.String()
}

func (c *CTypeSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c *CTypeSpec) IsComplete() bool {
	return true
}

func (c *CTypeSpec) IsOpaque() bool {
	return c.Opaque
}

func (c *CTypeSpec) Kind() CTypeKind {
	return TypeKind
}

func (c CTypeSpec) Copy() CType {
	return &c
}

func (c *CTypeSpec) GetBase() string {
	return c.Base
}

func (c *CTypeSpec) GetTag() string {
	return ""
}

func (c *CTypeSpec) SetRaw(x string) {
	c.Raw = x
}

func (c *CTypeSpec) CGoName() (name string) {
	if len(c.Raw) > 0 {
		return c.Raw
	}
	switch c.Base {
	case "int", "short", "long", "char":
		if c.Unsigned {
			name += "u"
		} else if c.Signed {
			name += "s"
		}
		switch {
		case c.Long:
			name += "long"
			if c.Base == "long" {
				name += "long"
			}
		case c.Short:
			name += "short"
		default:
			name += c.Base
		}
	default:
		name = c.Base
	}
	return
}

func (c *CTypeSpec) AddOuterArr(size uint64) {
	c.OuterArr.AddSized(size)
}

func (c *CTypeSpec) AddInnerArr(size uint64) {
	c.InnerArr.AddSized(size)
}

func (c *CTypeSpec) OuterArraySizes() []ArraySizeSpec {
	return c.OuterArr.Sizes()
}

func (c *CTypeSpec) InnerArraySizes() []ArraySizeSpec {
	return c.InnerArr.Sizes()
}

func (c *CTypeSpec) OuterArrays() ArraySpec {
	return c.OuterArr
}

func (c *CTypeSpec) InnerArrays() ArraySpec {
	return c.InnerArr
}

func (c *CTypeSpec) GetPointers() uint8 {
	return c.Pointers
}

func (c *CTypeSpec) IsConst() bool {
	return c.Const
}
