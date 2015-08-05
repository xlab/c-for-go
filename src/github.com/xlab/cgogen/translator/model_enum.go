package translator

import (
	"fmt"
	"strings"
)

type CEnumSpec struct {
	Tag         string
	Enumerators []CDecl
	Pointers    uint8
	Type        CTypeSpec
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

func (c *CEnumSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c CEnumSpec) Kind() CTypeKind {
	return EnumKind
}

func (c CEnumSpec) Copy() CType {
	return &c
}

func (ces CEnumSpec) String() string {
	var members []string
	for _, m := range ces.Enumerators {
		members = append(members, m.String())
	}
	membersColumn := strings.Join(members, ", ")

	str := "enum"
	if len(ces.Tag) > 0 {
		str = fmt.Sprintf("%s %s", str, ces.Tag)
	}
	if len(members) > 0 {
		str = fmt.Sprintf("%s {%s}", str, membersColumn)
	}
	if ces.Pointers > 0 {
		str += strings.Repeat("*", int(ces.Pointers))
	}
	return str
}
