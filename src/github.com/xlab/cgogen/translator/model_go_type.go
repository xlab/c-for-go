package translator

import (
	"bytes"
	"fmt"
	"strings"
)

type GoTypeSpec struct {
	Arrays   string
	Slices   uint8
	Pointers uint8
	Unsigned bool
	Base     string
	Bits     uint16
}

func (gts *GoTypeSpec) splitPointers(spec PointerSpec, n uint8) {
	if n == 0 {
		return
	}
	switch spec {
	case PointerArr:
		if n > 1 {
			gts.Slices += n
		} else {
			gts.Slices++
		}
	case PointerRef:
		if n > 1 {
			gts.Slices += n - 1
			gts.Pointers++
		} else {
			gts.Pointers++
		}
	}
}

func (gts GoTypeSpec) IsPlain() bool {
	switch gts.Base {
	case "int", "byte", "rune", "float32", "float64", "void", "unsafe.Pointer", "bool":
		return true
	case "string":
		return false
	}
	return false
}

func (gts *GoTypeSpec) IsReference() bool {
	return len(gts.Arrays) > 0
}

func (gts GoTypeSpec) String() string {
	buf := new(bytes.Buffer)
	if len(gts.Arrays) > 0 {
		buf.WriteRune('*')
		buf.WriteString(gts.Arrays)
	}
	if gts.Slices > 0 {
		buf.WriteString(strings.Repeat("[]", int(gts.Slices)))
	}
	if gts.Pointers > 0 {
		buf.WriteString(ptrs(gts.Pointers))
	}
	if gts.Unsigned && gts.Base == "int" {
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
	Arrays   []uint64
}

func (cgs CGoSpec) String() string {
	buf := new(bytes.Buffer)
	for _, size := range cgs.Arrays {
		fmt.Fprintf(buf, "[%d]", size)
	}
	fmt.Fprintf(buf, "%s%s", ptrs(cgs.Pointers), cgs.Base)
	return buf.String()
}

func (cgs *CGoSpec) AtLevel(level uint8) string {
	buf := new(bytes.Buffer)
	for i, size := range cgs.Arrays {
		if i < int(level) {
			continue
		} else if i == 0 {
			fmt.Fprint(buf, "*")
			continue
		}
		fmt.Fprintf(buf, "[%d]", size)
	}
	if int(level) > len(cgs.Arrays) {
		if delta := int(cgs.Pointers) + len(cgs.Arrays) - int(level); delta > 0 {
			fmt.Fprint(buf, strings.Repeat("*", delta))
		}
	} else {
		fmt.Fprint(buf, ptrs(cgs.Pointers))
	}
	fmt.Fprint(buf, cgs.Base)
	return buf.String()
}

func ptrs(n uint8) string {
	return strings.Repeat("*", int(n))
}
