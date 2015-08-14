package translator

import (
	"bytes"
	"fmt"
	"strings"
)

type GoTypeSpec struct {
	Slices   uint8
	Pointers uint8
	Arrays   string
	InnerCGO string
	Inner    *GoTypeSpec
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
	buf.WriteString(gts.Arrays)
	if gts.Slices > 0 {
		buf.WriteString(strings.Repeat("[]", int(gts.Slices)))
	}
	if gts.Pointers > 0 {
		buf.WriteString(strings.Repeat("*", int(gts.Pointers)))
	}
	if gts.Inner != nil {
		buf.WriteString(gts.Inner.String())
	} else {
		if gts.Unsigned {
			buf.WriteString("u")
		}
		buf.WriteString(gts.Base)
		if gts.Bits > 0 {
			fmt.Fprintf(buf, "%d", int(gts.Bits))
		}
	}
	return buf.String()
}
