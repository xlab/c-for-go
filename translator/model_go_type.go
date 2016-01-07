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
	Raw      string
	Bits     uint16
}

func (spec *GoTypeSpec) splitPointers(ptrTip Tip, n uint8) {
	if n == 0 {
		return
	}
	switch ptrTip {
	case TipPtrRef:
		spec.Slices = spec.Slices + n - 1
		spec.Pointers++
	case TipPtrSRef:
		spec.Pointers += n
	case TipPtrArr:
		spec.Slices += n
	default: // TipPtrArr
		spec.Slices += n
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
	case "int", "byte", "rune", "float", "unsafe.Pointer", "bool":
		return true
	case "string":
		return false
	}
	return false
}

func (spec *GoTypeSpec) IsReference() bool {
	return len(spec.Arrays) > 0
}

func (spec *GoTypeSpec) GetName() string {
	if len(spec.Raw) > 0 {
		return spec.Raw
	}
	return spec.Base
}

func (spec GoTypeSpec) String() string {
	buf := new(bytes.Buffer)
	if len(spec.Arrays) > 0 {
		buf.WriteString(spec.Arrays)
	}
	if spec.Slices > 0 {
		buf.WriteString(strings.Repeat("[]", int(spec.Slices)))
	}

	var unsafePointer uint8
	if spec.Base == "unsafe.Pointer" && len(spec.Raw) == 0 {
		unsafePointer = 1
	}
	if spec.Pointers > 0 {
		buf.WriteString(ptrs(spec.Pointers - unsafePointer))
	}
	if len(spec.Raw) > 0 {
		buf.WriteString(spec.Raw)
		return buf.String()
	}
	if spec.Unsigned {
		switch spec.Base {
		case "char", "short", "long", "int":
			buf.WriteString("u")
		}
	}
	buf.WriteString(spec.Base)
	if spec.Bits > 0 {
		fmt.Fprintf(buf, "%d", int(spec.Bits))
	}
	return buf.String()
}

func (spec GoTypeSpec) UnderlyingString() string {
	buf := new(bytes.Buffer)
	if len(spec.Arrays) > 0 {
		buf.WriteString(spec.Arrays)
	}
	if spec.Slices > 0 {
		buf.WriteString(strings.Repeat("[]", int(spec.Slices)))
	}

	var unsafePointer uint8
	if spec.Base == "unsafe.Pointer" {
		unsafePointer = 1
	}
	if spec.Pointers > 0 {
		buf.WriteString(ptrs(spec.Pointers - unsafePointer))
	}
	if spec.Unsigned {
		buf.WriteString("u")
	}
	buf.WriteString(spec.Base)
	if spec.Bits > 0 {
		fmt.Fprintf(buf, "%d", int(spec.Bits))
	}
	return buf.String()
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

func (spec CGoSpec) String() string {
	buf := new(bytes.Buffer)
	for _, size := range spec.Arrays {
		buf.WriteString(size.String())
	}
	fmt.Fprintf(buf, "%s%s", ptrs(spec.Pointers), spec.Base)
	return buf.String()
}

func (spec *CGoSpec) PointersAtLevel(level uint8) uint8 {
	// var pointers uint8
	// if len(spec.Arrays) > 0 {
	// 	pointers++
	// }
	if int(level) > len(spec.Arrays) {
		if delta := int(spec.Pointers) + len(spec.Arrays) - int(level); delta > 0 {
			return uint8(delta)
		}
	}
	if level <= spec.Pointers {
		return spec.Pointers - level
	}
	return 0
}

func (spec *CGoSpec) AtLevel(level uint8) string {
	buf := new(bytes.Buffer)
	for i, size := range spec.Arrays {
		if i < int(level) {
			continue
		} else if i == 0 {
			fmt.Fprint(buf, "*")
			continue
		}
		buf.WriteString(size.String())
	}
	if int(level) > len(spec.Arrays) {
		if delta := int(spec.Pointers) + len(spec.Arrays) - int(level); delta > 0 {
			fmt.Fprint(buf, strings.Repeat("*", delta))
		}
	} else {
		fmt.Fprint(buf, ptrs(spec.Pointers))
	}
	fmt.Fprint(buf, spec.Base)
	return buf.String()
}

func (spec *CGoSpec) SpecAtLevel(level uint8) CGoSpec {
	buf := CGoSpec{
		Base: spec.Base,
	}
	for i, size := range spec.Arrays {
		if i < int(level) {
			continue
		} else if i == 0 {
			buf.Pointers = 1
			continue
		}
		buf.Arrays = append(buf.Arrays, size)
	}
	if int(level) > len(spec.Arrays) {
		if delta := int(spec.Pointers) + len(spec.Arrays) - int(level); delta > 0 {
			buf.Pointers += uint8(delta)
		}
	} else {
		buf.Pointers += spec.Pointers
	}
	return buf
}

func ptrs(n uint8) string {
	return strings.Repeat("*", int(n))
}
