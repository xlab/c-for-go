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
	Kind     CTypeKind
	Base     string
	Bits     uint16
}

func (spec *GoTypeSpec) splitPointers(ptrTip Tip, n uint8) {
	if n == 0 {
		return
	}
	switch ptrTip {
	case TipPtrArr:
		if n > 1 {
			spec.Slices += n
		} else {
			spec.Slices++
		}
	case TipPtrRef:
		if n > 1 {
			spec.Slices += n - 1
			spec.Pointers++
		} else {
			spec.Pointers++
		}
	}
}

func (spec GoTypeSpec) IsPlainKind() bool {
	switch spec.Kind {
	case PlainTypeKind, EnumKind, OpaqueStructKind, UnionKind:
		return true
	}
	return false
}

func (spec GoTypeSpec) IsPlain() bool {
	switch spec.Base {
	case "int", "byte", "rune", "float", "void", "unsafe.Pointer", "bool":
		return true
	case "string":
		return false
	}
	return false
}

func (spec *GoTypeSpec) IsReference() bool {
	return len(spec.Arrays) > 0
}

func (spec GoTypeSpec) String() string {
	return string(spec.Bytes())
}

func (spec GoTypeSpec) Bytes() []byte {
	buf := new(bytes.Buffer)
	if len(spec.Arrays) > 0 {
		buf.WriteString(spec.Arrays)
	}
	if spec.Slices > 0 {
		buf.WriteString(strings.Repeat("[]", int(spec.Slices)))
	}
	if spec.Pointers > 0 {
		buf.WriteString(ptrs(spec.Pointers))
	}
	if spec.Unsigned && spec.Base == "int" {
		buf.WriteString("u")
	}
	buf.WriteString(spec.Base)
	if spec.Bits > 0 {
		fmt.Fprintf(buf, "%d", int(spec.Bits))
	}
	return buf.Bytes()
}

func (a ArraySizeSpec) String() string {
	if len(a.Str) > 0 {
		return fmt.Sprintf("[%s]", a.Str)
	} else {
		return fmt.Sprintf("[%d]", a.N)
	}
}

type CGoSpec struct {
	Base     string
	Pointers uint8
	Arrays   []ArraySizeSpec
}

func (cgs CGoSpec) String() string {
	buf := new(bytes.Buffer)
	for _, size := range cgs.Arrays {
		buf.WriteString(size.String())
	}
	fmt.Fprintf(buf, "%s%s", ptrs(cgs.Pointers), cgs.Base)
	return buf.String()
}

func (cgs *CGoSpec) PointersAtLevel(level uint8) uint8 {
	// var pointers uint8
	// if len(cgs.Arrays) > 0 {
	// 	pointers++
	// }
	if int(level) > len(cgs.Arrays) {
		if delta := int(cgs.Pointers) + len(cgs.Arrays) - int(level); delta > 0 {
			return uint8(delta)
		}
	}
	if level <= cgs.Pointers {
		return cgs.Pointers - level
	}
	return 0
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
		buf.WriteString(size.String())
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
