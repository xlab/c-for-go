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
