package translator

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
)

type CTypeSpec struct {
	Base     string
	Const    bool
	Unsigned bool
	Short    bool
	Long     bool
	Pointers uint8
}

func (c *CTypeSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c CTypeSpec) Kind() CTypeKind {
	return TypeDef
}

func (cts CTypeSpec) String() string {
	var str string
	if cts.Const {
		str += "const "
	}
	if cts.Unsigned {
		str += "unsigned "
	}
	switch {
	case cts.Long:
		str += "long "
	case cts.Short:
		str += "short "
	}
	str += cts.Base
	for i := uint8(0); i < cts.Pointers; i++ {
		str += "*"
	}
	return str
}

func (c CTypeSpec) Copy() CType {
	return &c
}

func (cts *CTypeSpec) MarshalJSON() ([]byte, error) {
	if cts == nil {
		return nil, nil
	}
	if len(cts.Base) == 0 {
		return nil, errors.New("base type isn't specified")
	}
	return []byte(cts.String()), nil
}

func (cts *CTypeSpec) UnmarshalJSON(b []byte) error {
	parts := bytes.Split(b, spaceStr)
	if len(parts) == 0 {
		return errors.New("unexpected EOF")
	}
	ts := CTypeSpec{}
	sort.Reverse(bytesSlice(parts))

	// states:
	// 0 — pointers
	// 1 — base
	// 2 — qualifiers
	var state int
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		switch state {
		case 0:
			// read pointers count
			for bytes.HasSuffix(part, ptrStr) {
				ts.Pointers++
				part = part[:len(part)-1]
			}
			state = 1
		case 1:
			// read the base name
			if isRestrictedBase(part) {
				return errors.New("ctype: can't use keyword as a base type name: " + string(part))
			}
			ts.Base = string(part)
			state = 2
		case 2:
			// read specifiers and qualifiers
			switch {
			case bytes.Equal(part, specStruct), bytes.Equal(part, specUnion):
				return errors.New("struct is not a simple C type")
			case bytes.Equal(part, specShort):
				ts.Short = true
			case bytes.Equal(part, specLong):
				ts.Long = true
			case bytes.Equal(part, specUnsigned):
				ts.Unsigned = true
			case bytes.Equal(part, qualConst):
				ts.Const = true
			}
		}
	}

	if len(ts.Base) == 0 {
		return errors.New("ctype: no base type name specified")
	}
	*cts = ts
	return nil
}

func CTypeOf(v interface{}) (*CTypeSpec, error) {
	switch x := v.(type) {
	case int32:
		return &CTypeSpec{Base: "int"}, nil
	case int64:
		return &CTypeSpec{Base: "long"}, nil
	case uint32:
		return &CTypeSpec{Base: "int", Unsigned: true}, nil
	case uint64:
		return &CTypeSpec{Base: "long", Unsigned: true}, nil
	case float32:
		return &CTypeSpec{Base: "float"}, nil
	case float64:
		return &CTypeSpec{Base: "double"}, nil
	case string:
		return &CTypeSpec{Base: "char", Pointers: 1, Const: true}, nil
	default:
		return nil, errors.New(fmt.Sprintf("cannot resolve type %T", x))
	}
}
