package translator

import (
	"bytes"
	"fmt"
	"strings"
)

type CEnumSpec struct {
	Tag       string
	Members   []*CDecl
	Type      CTypeSpec
	Arrays    string
	VarArrays uint8
	Pointers  uint8
}

func (c *CEnumSpec) AddArray(size uint64) {
	if size > 0 {
		c.Arrays = fmt.Sprintf("%s[%d]", c.Arrays, size)
		return
	}
	c.VarArrays++
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

func (spec CEnumSpec) String() string {
	var members []string
	for _, m := range spec.Members {
		members = append(members, m.String())
	}
	membersColumn := strings.Join(members, ", ")

	buf := new(bytes.Buffer)
	fmt.Fprint(buf, "enum")
	if len(spec.Tag) > 0 {
		buf.WriteString(" " + spec.Tag)
	}
	if len(members) > 0 {
		fmt.Fprintf(buf, " {%s}", membersColumn)
	}
	buf.WriteString(strings.Repeat("*", int(spec.Pointers)))
	buf.WriteString(spec.Arrays)
	return buf.String()
}

func (c *CEnumSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c *CEnumSpec) Kind() CTypeKind {
	return EnumKind
}

func (c *CEnumSpec) IsComplete() bool {
	return len(c.Members) > 0
}

func (c *CEnumSpec) IsOpaque() bool {
	return len(c.Members) == 0
}

func (c CEnumSpec) Copy() CType {
	return &c
}

func (c *CEnumSpec) GetBase() string {
	return c.Tag
}

func (c *CEnumSpec) GetArrays() string {
	return c.Arrays
}

func (c *CEnumSpec) GetVarArrays() uint8 {
	return c.VarArrays
}

func (c *CEnumSpec) GetPointers() uint8 {
	return c.Pointers
}

func (c *CEnumSpec) IsConst() bool {
	// could be c.Const
	return false
}
