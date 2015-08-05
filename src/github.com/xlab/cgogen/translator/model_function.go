package translator

import (
	"fmt"
	"strings"
)

type CFunctionSpec struct {
	Returns   CDecl
	ParamList []CDecl
	Pointers  uint8
}

func (c *CFunctionSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c CFunctionSpec) Kind() CTypeKind {
	return FunctionDef
}

func (c CFunctionSpec) Copy() CType {
	return &c
}

func (c CFunctionSpec) String() string {
	var params []string
	for _, param := range c.ParamList {
		params = append(params, param.String())
	}
	return fmt.Sprintf("%s (%s)", c.Returns, strings.Join(params, ", "))
}
