package translator

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type GoTypeSpec struct {
	Slices   uint8
	Pointers uint8
	Arrays   []uint64
	InnerCGO string
	Inner    *GoTypeSpec
	Unsigned bool
	Base     string
	Bits     uint16
}

// func (gts GoTypeSpec) Wrap(innerGTS GoTypeSpec) GoTypeSpec {
// 	return GoTypeSpec{
// 		Slices:   gts.Slices,
// 		Pointers: gts.Pointers,
// 		Arrays:   gts.Arrays,
// 		Inner:    &innerGTS,
// 	}
// }

func (gts GoTypeSpec) String() string {
	var str string
	str += strings.Repeat("[]", int(gts.Slices))
	str += strings.Repeat("*", int(gts.Pointers))
	for _, size := range gts.Arrays {
		str += fmt.Sprintf("[%d]", size)
	}
	if gts.Inner != nil {
		str += gts.Inner.String()
	} else {
		if gts.Unsigned {
			str += "u"
		}
		str += gts.Base
		if gts.Bits > 0 {
			str += strconv.Itoa(int(gts.Bits))
		}
	}
	return str
}

func (gts GoTypeSpec) MarshalJSON() ([]byte, error) {
	return []byte(gts.String()), nil
}

func (gts *GoTypeSpec) UnmarshalJSON(b []byte) error {
	// purge any spaces
	b = bytes.Replace(b, spaceStr, emptyStr, -1)

	// states:
	// 0 — beginning
	// 1 — processing slices
	// 2 — processing pointers
	// 3
	// 3 — base
	var state int
	// var ts = GoTypeSpec{}

	for {
		switch state {
		case 0:
			switch {
			case bytes.HasPrefix(b, sliceStr):
				state = 1
				continue
			case bytes.HasPrefix(b, ptrStr):
				state = 2
				continue
			}
		}
	}
	// TODO
	return nil
}
