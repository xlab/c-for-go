package translator

import (
	"bytes"
	"fmt"
	"strings"
)

type CStructSpec struct {
	Tag       string
	IsUnion   bool
	Members   []CDecl
	Arrays    string
	VarArrays uint8
	Pointers  uint8
}

func (c *CStructSpec) AddArray(size uint64) {
	if size > 0 {
		c.Arrays = fmt.Sprintf("%s[%d]", size, c.Arrays)
		return
	}
	c.VarArrays++
}

func (css CStructSpec) String() string {
	var members []string
	for _, m := range css.Members {
		members = append(members, m.String())
	}
	membersColumn := strings.Join(members, ", ")

	buf := new(bytes.Buffer)
	if css.IsUnion {
		buf.WriteString("union")
	} else {
		buf.WriteString("struct")
	}
	if len(css.Tag) > 0 {
		buf.WriteString(" " + css.Tag)
	}
	if len(members) > 0 {
		fmt.Fprintf(buf, " {%s}", membersColumn)
	}
	buf.WriteString(strings.Repeat("*", int(css.Pointers)))
	buf.WriteString(css.Arrays)
	return buf.String()
}

func (c *CStructSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c CStructSpec) Kind() CTypeKind {
	if c.IsUnion {
		return UnionKind
	}
	return StructKind
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
