package translator

import (
	"bytes"
	"fmt"
	"strings"
)

type GoTypeSpec struct {
	Arrays   []uint64
	Slices   uint8
	Pointers uint8
	Unsigned bool
	Base     string
	Bits     uint16
}

func (gts *GoTypeSpec) splitPointers(n uint8) {
	for n > 0 {
		if n > 1 {
			gts.Slices++
		} else {
			gts.Pointers++
		}
		n--
	}
}

func (gts GoTypeSpec) String() string {
	buf := new(bytes.Buffer)
	for _, size := range gts.Arrays {
		fmt.Fprintf(buf, "[%d]", size)
	}
	if gts.Slices > 0 {
		buf.WriteString(strings.Repeat("[]", int(gts.Slices)))
	}
	if gts.Pointers > 0 {
		buf.WriteString(strings.Repeat("*", int(gts.Pointers)))
	}
	if gts.Unsigned {
		buf.WriteString("u")
	}
	buf.WriteString(gts.Base)
	if gts.Bits > 0 {
		fmt.Fprintf(buf, "%d", int(gts.Bits))
	}
	return buf.String()
}

type CGoSpec struct {
	Base     string
	Pointers uint8
	Arrays   []uint8
}

func (cgs CGoSpec) String() string {
	return fmt.Sprintf("%s%s", strings.Repeat("*", int(cgs.Pointers)), cgs.Base)
}
