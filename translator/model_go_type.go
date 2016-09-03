package translator

import (
	"bytes"
	"fmt"
	"strings"
)

type GoTypeSpec struct {
	Slices   uint8
	Pointers uint8
	InnerArr ArraySpec
	OuterArr ArraySpec
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
	case PlainTypeKind, OpaqueStructKind, EnumKind, UnionKind:
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

func (spec *GoTypeSpec) PlainType() string {
	if len(spec.Raw) > 0 {
		return spec.Raw
	}
	buf := new(bytes.Buffer)
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

func (spec GoTypeSpec) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString(slcs(spec.Slices))
	buf.WriteString(arrs(spec.OuterArr))

	var unsafePointer uint8
	if spec.Base == "unsafe.Pointer" && len(spec.Raw) == 0 {
		unsafePointer = 1
	}
	if spec.Pointers > 0 {
		buf.WriteString(ptrs(spec.Pointers - unsafePointer))
	}
	buf.WriteString(arrs(spec.InnerArr))

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
	buf.WriteString(slcs(spec.Slices))
	buf.WriteString(arrs(spec.OuterArr))

	var unsafePointer uint8
	if spec.Base == "unsafe.Pointer" {
		unsafePointer = 1
	}
	if spec.Pointers > 0 {
		buf.WriteString(ptrs(spec.Pointers - unsafePointer))
	}
	buf.WriteString(arrs(spec.InnerArr))

	if spec.Unsigned {
		buf.WriteString("u")
	}
	buf.WriteString(spec.Base)
	if spec.Bits > 0 {
		fmt.Fprintf(buf, "%d", int(spec.Bits))
	}

	return buf.String()
}

type CGoSpec struct {
	Base     string
	Pointers uint8
	OuterArr ArraySpec
	InnerArr ArraySpec
}

func (spec CGoSpec) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString(arrs(spec.OuterArr))
	buf.WriteString(ptrs(spec.Pointers))
	buf.WriteString(arrs(spec.InnerArr))
	buf.WriteString(spec.Base)
	return buf.String()
}

func (spec *CGoSpec) PointersAtLevel(level uint8) uint8 {
	if int(level) > len(spec.OuterArr) {
		if delta := int(spec.Pointers) + len(spec.OuterArr.Sizes()) - int(level); delta > 0 {
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
	outerArrSizes := spec.OuterArr.Sizes()
	for i, size := range outerArrSizes {
		if i < int(level) {
			continue
		} else if i == 0 {
			buf.WriteRune('*')
			continue
		}
		fmt.Fprintf(buf, "[%d]", size)
	}
	if int(level) > len(outerArrSizes) {
		if delta := int(spec.Pointers) + len(outerArrSizes) - int(level); delta > 0 {
			buf.WriteString(ptrs(uint8(delta)))
		}
	} else {
		buf.WriteString(ptrs(spec.Pointers))
	}
	// drop inner arrays at levels
	// buf.WriteString(arrs(spec.InnerArr))
	buf.WriteString(spec.Base)

	return buf.String()
}

func (spec *CGoSpec) SpecAtLevel(level uint8) CGoSpec {
	buf := CGoSpec{
		// drop inner arrays at levels
		// InnerArr: spec.InnerArr,
		Base: spec.Base,
	}
	for i, size := range spec.OuterArr.Sizes() {
		if i < int(level) {
			continue
		} else if i == 0 {
			buf.Pointers = 1
			continue
		}
		buf.OuterArr.AddSized(size.N)
	}
	if int(level) > len(spec.OuterArr) {
		if delta := int(spec.Pointers) + len(spec.OuterArr.Sizes()) - int(level); delta > 0 {
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

func slcs(slcs uint8) string {
	return strings.Repeat("[]", int(slcs))
}

func arrs(spec ArraySpec) string {
	return string(spec)
}
