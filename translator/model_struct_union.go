package translator

import (
	"bytes"
	"fmt"
	"strings"
)

type CStructSpec struct {
	Tag       string
	Typedef   string
	IsUnion   bool
	Members   []*CDecl
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
	buf := new(bytes.Buffer)
	writePrefix := func() {
		if spec.IsUnion {
			buf.WriteString("union ")
		} else {
			buf.WriteString("struct ")
		}
	}

	switch {
	case len(spec.Typedef) > 0:
		buf.WriteString(spec.Typedef)
	case len(spec.Tag) > 0:
		writePrefix()
		buf.WriteString(spec.Tag)
	case len(spec.Members) > 0:
		var members []string
		for _, m := range spec.Members {
			members = append(members, m.String())
		}
		membersColumn := strings.Join(members, ",\n")
		writePrefix()
		fmt.Fprintf(buf, " {%s}", membersColumn)
	default:
		writePrefix()
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
	if len(c.Typedef) > 0 {
		return c.Typedef
	}
	return c.Tag
}

func (c *CStructSpec) GetTag() string {
	return c.Tag
}

func (c *CStructSpec) SetRaw(x string) {
	c.Typedef = x
}

func (c *CStructSpec) CGoName() string {
	if len(c.Typedef) > 0 {
		return c.Typedef
	}
	if c.IsUnion {
		return "union_" + c.Tag
	}
	return "struct_" + c.Tag
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
