package translator

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type CTypeSpec struct {
	Base      string
	Const     bool
	Signed    bool
	Unsigned  bool
	Short     bool
	Long      bool
	Arrays    string
	VarArrays uint8
	Pointers  uint8
}

func (c *CTypeSpec) AddArray(size uint64) {
	if size > 0 {
		c.Arrays += fmt.Sprintf("[%d]", size)
		return
	}
	c.VarArrays++
}

type ArraySizeSpec struct {
	N   uint64
	Str string
}

func GetArraySizes(arr string) (sizes []ArraySizeSpec) {
	if len(arr) == 0 {
		return
	}
	for len(arr) > 0 {
		// get "n" from "[k][l][m][n]"
		p1 := strings.LastIndexByte(arr, '[')
		p2 := strings.LastIndexByte(arr, ']')
		part := arr[p1+1 : p2]
		// and try to convert uint64
		if u, err := strconv.ParseUint(part, 10, 64); err != nil || u == 0 {
			// use size spec as-is (i.e. unsafe.Sizeof(x))
			sizes = append(sizes, ArraySizeSpec{Str: part})
		} else {
			sizes = append(sizes, ArraySizeSpec{N: u})
		}
		arr = arr[:p1]
	}
	return sizes
}

func (cts CTypeSpec) String() string {
	buf := new(bytes.Buffer)
	if cts.Unsigned {
		buf.WriteString("unsigned ")
	} else if cts.Signed {
		buf.WriteString("signed ")
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

func (c *CTypeSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c *CTypeSpec) IsOpaque() bool {
	return true
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
