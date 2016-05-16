package translator

import (
	"bytes"
	"fmt"
	"strings"
)

type CStructSpec struct {
	Tag      string
	Typedef  string
	IsUnion  bool
	Members  []*CDecl
	Pointers uint8
	InnerArr ArraySpec
	OuterArr ArraySpec
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

	buf.WriteString(arrs(spec.OuterArr))
	buf.WriteString(ptrs(spec.Pointers))
	buf.WriteString(arrs(spec.InnerArr))
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

func (c *CStructSpec) AddOuterArr(size uint64) {
	c.OuterArr.AddSized(size)
}

func (c *CStructSpec) AddInnerArr(size uint64) {
	c.InnerArr.AddSized(size)
}

func (c *CStructSpec) OuterArraySizes() []ArraySizeSpec {
	return c.OuterArr.Sizes()
}

func (c *CStructSpec) InnerArraySizes() []ArraySizeSpec {
	return c.InnerArr.Sizes()
}

func (c *CStructSpec) OuterArrays() ArraySpec {
	return c.OuterArr
}

func (c *CStructSpec) InnerArrays() ArraySpec {
	return c.InnerArr
}

func (c *CStructSpec) GetPointers() uint8 {
	return c.Pointers
}

func (c *CStructSpec) IsConst() bool {
	return false
}

func (c CStructSpec) AtLevel(level int) CType {
	spec := c
	var outerArrSpec ArraySpec
	for i, size := range spec.OuterArr.Sizes() {
		if i < int(level) {
			continue
		} else if i == 0 {
			spec.Pointers = 1
			continue
		}
		outerArrSpec.AddSized(size.N)
	}
	if int(level) > len(spec.OuterArr) {
		if delta := int(spec.Pointers) + len(spec.OuterArr.Sizes()) - int(level); delta > 0 {
			spec.Pointers = uint8(delta)
		}
	}
	spec.OuterArr = outerArrSpec
	spec.InnerArr = ArraySpec("")
	return &spec
}
