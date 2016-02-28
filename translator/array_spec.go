package translator

import (
	"fmt"
	"strconv"
	"strings"
)

type ArraySpec string

func (a ArraySpec) String() string {
	return string(a)
}

func (a *ArraySpec) AddSized(size uint64) {
	*a = ArraySpec(fmt.Sprintf("%s[%d]", a, size))
}

func (a *ArraySpec) Prepend(spec ArraySpec) {
	*a = spec + *a
}

type ArraySizeSpec struct {
	N   uint64
	Str string
}

func (a ArraySpec) Sizes() (sizes []ArraySizeSpec) {
	if len(a) == 0 {
		return
	}
	arr := string(a)
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
