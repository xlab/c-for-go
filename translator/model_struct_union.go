package translator

import (
	"bytes"
	"fmt"
	"strings"
)

type CStructMember struct {
	Name string
	Type CType
}

func (c CStructMember) String() string {
	return c.Name + " " + c.Type.String()
}

type CStructSpec struct {
	Tag       string
	IsUnion   bool
	Members   []CStructMember
	Arrays    string
	VarArrays uint8
	Pointers  uint8
}

func (c *CStructSpec) AddArray(size uint64) {
	if size > 0 {
		c.Arrays = fmt.Sprintf("%s[%d]", c.Arrays, size)
		return
	}
	c.VarArrays++
}

func (spec CStructSpec) String() string {
	var members []string
	for _, m := range spec.Members {
		members = append(members, m.Name)
	}
	membersColumn := strings.Join(members, ", ")

	buf := new(bytes.Buffer)
	if spec.IsUnion {
		buf.WriteString("union")
	} else {
		buf.WriteString("struct")
	}
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

func (c *CStructSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c *CStructSpec) Kind() CTypeKind {
	switch {
	case c.IsUnion:
		return UnionKind
	case len(c.Members) == 0:
		return OpaqueStructKind
	default:
		return StructKind
	}
}

func (c *CStructSpec) IsComplete() bool {
	return len(c.Members) > 0
}

func (c *CStructSpec) IsOpaque() bool {
	return len(c.Members) == 0
}

func (c CStructSpec) Copy() CType {
	return &c
}

func (c *CStructSpec) GetBase() string {
	return c.Tag
}

func (c *CStructSpec) GetArrays() string {
	return c.Arrays
}

func (c *CStructSpec) GetVarArrays() uint8 {
	return c.VarArrays
}

func (c *CStructSpec) GetPointers() uint8 {
	return c.Pointers
}

func (c *CStructSpec) IsConst() bool {
	return false
}
