package translator

import (
	"fmt"
	"strings"
)

type CStructSpec struct {
	Tag      string
	Union    bool
	Members  []CDecl
	Pointers uint8
}

func (c *CStructSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c CStructSpec) Kind() CTypeKind {
	return StructDef
}

func (c CStructSpec) Copy() CType {
	return &c
}

func (css CStructSpec) String() string {
	var members []string
	for _, m := range css.Members {
		members = append(members, m.String())
	}
	membersColumn := strings.Join(members, ", ")

	str := "struct"
	if css.Union {
		str = "union"
	}
	if len(css.Tag) > 0 {
		str = fmt.Sprintf("%s %s", str, css.Tag)
	}
	if len(members) > 0 {
		str = fmt.Sprintf("%s {%s}", str, membersColumn)
	}
	if css.Pointers > 0 {
		str += strings.Repeat("*", int(css.Pointers))
	}
	return str
}
