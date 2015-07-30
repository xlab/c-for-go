// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"fmt"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"

	"github.com/cznic/c/internal/xc"
	"github.com/cznic/mathutil"
	"github.com/cznic/strutil"
)

type Type interface {
	// Alignof return the required alignment of a type in bytes. Calling
	// Alignof of a FunctionType will panic.
	Alignof() int

	// BaseType returns the type a PtrType points to.  Calling BaseType for
	// other type kinds will panic.
	BaseType() Type

	// ElementType returns the element type of an ArrayType. Calling
	// ElementType for other type kinds will panic.
	ElementType() Type

	// One of PtrType, VoidType, IntType, ...
	Kind() Kind

	// Len returns a pair (declared size, true) of an ArrayType if the size
	// declaration is a constant expression or (0, false) otherwise.
	// Calling Len for other type kinds will panic.
	Len() (int, bool)

	// Name returns the typedef name of a NamedType.  Calling Name for
	// other type kinds will panic.
	Name() xc.Token

	// ParameterTypeList returns the ParameterTypeList of a FunctionType.
	// Calling ParameterTypeList for other type kinds will panic.
	ParameterTypeList() *ParameterTypeList

	// ResultType returns the result type of a function type. Calling ResultType
	// for other type kinds will panic.
	ResultType() Type

	// Sizeof return the size of a type in bytes. Calling Sizeof of a
	// FunctionType will panic.
	Sizeof() int

	// StructOrUnionType returns the struct or union specification of a
	// StructType or a UnionType. Calling StructOrUnionType for other type
	// kinds will panic.
	StructOrUnionType() *StructOrUnionSpecifier
}

//TODO
//TODO // UnderlyingType returns the underlying type of t. The underlying type of a
//TODO // NamedType T is the underlying type of the type T refers to in the typedef.
//TODO // Otherwise the underlying type of T is T itself.
//TODO func UnderlyingType(t Type) Type {
//TODO 	for t.Type() == NamedType {
//TODO 		t = t.Typedef()
//TODO 	}
//TODO 	return t
//TODO }

// PrettyString returns, if possible, pretty formatted string representation of v.
func PrettyString(v interface{}) string {
	return strutil.PrettyString(v, "", "", printHooks)
}

type nodeStat struct { //TODO -> all_test.go
	instances int
	sizeof    int
}

func astSum(m map[string]*nodeStat, v interface{}) { //TODO -> all_test.go
	if v == nil {
		return
	}

	r := func(nm string, words int) {
		st := m[nm]
		if st == nil {
			m[nm] = &nodeStat{instances: 1, sizeof: words * mathutil.UintPtrBits / 8}
			return
		}

		st.instances++
	}

	var g func(interface{}) int
	g = func(v interface{}) int {
		if v == nil {
			return 1
		}

		switch x := v.(type) {
		case xc.Token:
			r(reflect.TypeOf(x).Name(), 2)
			return 2
		case PpTokenList:
			r(reflect.TypeOf(x).Name(), 1)
			return 1
		}

		rt := reflect.TypeOf(v)
		rv := reflect.ValueOf(v)
		switch rt.Kind() {
		case reflect.Struct:
			w := 0
			for i := 0; i < rt.NumField(); i++ {
				f := rv.Field(i)
				if !f.CanInterface() {
					continue
				}

				w += g(f.Interface())
			}
			r(rt.Name(), w)
			return w
		case reflect.Ptr:
			if rv.IsNil() {
				return 1
			}

			g(rv.Elem().Interface())
			return 1
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
			r(rt.Name(), 1)
			return 1
		default:
			panic("internal error 004")
		}
	}
	g(v)
	return
}

func intT(i interface{}) interface{} { // C int
	switch x := i.(type) {
	case int:
		return int32(x)
	case int32:
		return x
	case int64:
		return int32(x)
	default:
		panic("TODO")
	}
}

func uintT(i interface{}) interface{} { // C unsigned
	switch x := i.(type) {
	case uint32:
		return x
	case uint64:
		return uint32(x)
	default:
		panic("internal error 009")
	}
}

func longT(i interface{}) interface{} { // C long
	switch x := i.(type) {
	case int64:
		return x
	case int:
		return int64(x)
	default:
		panic("TODO")
	}
}

func ulongT(i interface{}) interface{} { // C unsigned long
	switch x := i.(type) {
	case uint64:
		return x
	default:
		panic("TODO")
	}
}

func isZero(i interface{}) bool {
	switch x := i.(type) {
	case int32:
		return x == 0
	case int64:
		return x == 0
	case uint32:
		return x == 0
	case uint64:
		return x == 0
	}
	panic("unreachable")
}

func strConst(t xc.Token) interface{} {
	s := dict.S(t.Val)
	return string(s[1 : len(s)-1]) //TODO This is incomplete
}

func intConst(t xc.Token) interface{} {
	const (
		l = 1 << iota
		ll
		u
	)
	k := 0
	s := dict.S(t.Val)
	i := len(s) - 1
more:
	switch c := s[i]; c {
	case 'u', 'U':
		k |= u
		i--
		goto more
	case 'l', 'L':
		if i > 0 && (s[i-1] == 'l' || s[i-1] == 'L') {
			k |= ll
			i -= 2
			goto more
		}

		k |= l
		i--
		goto more
	}
	n, err := strconv.ParseUint(string(s[:i+1]), 0, 64)
	if err != nil {
		compilation.Err(t.Pos(), "invalid integer constant: %s", s)
	}

	switch {
	case k == 0:
		return intT(int(n))
	case k == l, k == ll:
		return longT(int64(n))
	case k == u:
		return uintT(n)
	case k == u|l, k == u|ll:
		return ulongT(n)
	default:
		panic("TODO")
	}
}

func intConst2(t xc.Token) int {
	switch x := intConst(t).(type) {
	case int32:
		return int(x)
	default:
		panic("TODO")
	}
}

func charConst(t xc.Token) interface{} {
	s := dict.S(t.Val)
	long := false
	if s[0] == 'L' {
		long = true
		s = s[1:]
	}
	if len(s) == 3 { // 'A'
		if long {
			return longT(int(s[1]))
		}

		return intT(int(s[1]))
	}

	switch s[2] {
	case '0', '1', '2', '3', '4', '5', '6', '7':
		v := 0
		i := 0
		for _, c := range s[2:] {
			if c < '0' || c > '7' {
				break
			}

			v = 8*v + int(c) - '0'
			i++
			if i == 3 {
				break
			}
		}
		if long {
			return longT(v)
		}

		return intT(int(v))
	case '\'':
		return intT('\'')
	case '"':
		return intT('"')
	case '?':
		return intT('?')
	case '\\':
		return intT('\\')
	case 'a':
		return intT('\a')
	case 'b':
		return intT('\b')
	case 'f':
		return intT('\f')
	case 'n':
		return intT('\n')
	case 'r':
		return intT('\r')
	case 't':
		return intT('\t')
	case 'v':
		return intT('\v')
	default:
		compilation.Err(t.Pos(), "invalid character constant")
		return intT(0)
	}
}

func boolT(v bool) interface{} {
	if v {
		return intT(1)
	}

	return intT(0)
}

type binding struct {
	Token xc.Token
	Node  interface{}
}

func newBinding(tok xc.Token, node interface{}) *binding { return &binding{Token: tok, Node: node} }

// Bindings maps identifiers to declarations or definitions.
type Bindings struct {
	SUSpecifier0 *StructOrUnionSpecifier0
	Type         Scope
	specifier    Type
	identifiers  map[int]*binding
	isTypedef    bool
	labels       map[int]*binding
	members      map[int]*binding
	parent       *Bindings
	tags         map[int]*binding
}

func newBindings(typ Scope, parent *Bindings) *Bindings {
	return &Bindings{Type: typ, parent: parent}
}

func (s *Bindings) Parent() *Bindings             { return s.parent }
func (s *Bindings) Identifiers() map[int]*binding { return s.identifiers }
func (s *Bindings) Tags() map[int]*binding        { return s.tags }

func (s *Bindings) boot0(p *map[int]*binding) map[int]*binding {
	if m := *p; m != nil {
		return m
	}

	m := map[int]*binding{}
	*p = m
	return m
}

func (s *Bindings) boot(ns Namespace) map[int]*binding {
	switch ns {
	case NSLabels:
		return s.boot0(&s.labels)
	case NSTags:
		return s.boot0(&s.tags)
	case NSMembers:
		return s.boot0(&s.members)
	case NSIdentifiers:
		return s.boot0(&s.identifiers)
	default:
		panic("internal error")
	}
}

func (s *Bindings) insert(ns Namespace, tok xc.Token, node interface{}) {
	if tok.Rune != IDENTIFIER || ns == NSTags && s.Type != ScopeFile {
		panic("internal error")
	}

	m := s.boot(ns)
	nm := tok.Val
	if ex, ok := m[nm]; ok {
		switch exn := ex.Node.(type) {
		case *Declarator:
			switch n := node.(type) {
			case *FunctionDefinition:
				m[nm] = newBinding(tok, node)
				return
			case *Declarator:
				switch {
				case !exn.IsDefinition && n.IsDefinition:
					m[nm] = newBinding(tok, node)
					return
				default:
					panic("internal error")
				}
			default:
				panic("internal error")
			}
		case *StructOrUnionSpecifier:
			switch n := node.(type) {
			case *StructOrUnionSpecifier:
				switch exn.Case {
				case 1: // StructOrUnion IDENTIFIER
					switch n.Case {
					case 0: // StructOrUnionSpecifier0 '{' StructDeclarationList '}'
						m[nm] = newBinding(tok, node)
						return
					case 1: // StructOrUnion IDENTIFIER
						return
					default:
						panic(n.Case)
					}
				case 0: // StructOrUnionSpecifier0 '{' StructDeclarationList '}'
					switch n.Case {
					case 1: // StructOrUnion IDENTIFIER
						return
					default:
						panic("internal error")
					}
				default:
					panic("internal error")
				}
			default:
				panic("internal error")
			}
		default:
			panic("internal error")
		}

		compilation.ErrTok(tok, "%s redeclared, previous declaration at %v", tok.S(), fileset.Position(ex.Token.Pos()))
		return
	}

	m[nm] = newBinding(tok, node)
}

// Lookup attempts to find the binding of nm in ns.
func (s *Bindings) Lookup(ns Namespace, nm int) (*binding, *Bindings) {
	var m map[int]*binding
	for s != nil {
		switch ns {
		case NSLabels:
			m = s.labels
		case NSTags:
			m = s.tags
		case NSMembers:
			m = s.members
		case NSIdentifiers:
			m = s.identifiers
		default:
			panic("internal error")
		}

		if b, ok := m[nm]; ok {
			return b, s
		}

		s = s.parent
	}
	return nil, nil
}

func (s *Bindings) IsTypedefName(nm int) bool {
	b, _ := s.Lookup(NSIdentifiers, nm)
	if b == nil {
		return false
	}

	switch x := b.Node.(type) {
	case *Declarator:
		return x.IsTypedef
	case *Declaration: //TODO-
		return x.IsTypedef
	case *FunctionDefinition:
		return false
	default:
		//dbg("", PrettyString(b))
		panic("TODO")
		//return false
	}
}

// ModelItem is a single item of a model.
type ModelItem struct {
	Size  int
	Align int
	More  interface{}
}

// Model describes size and align requirements of predeclared types.
type Model map[ScalarType]ModelItem

// sanityCheck reports model errors, if any.
func (m Model) sanityCheck() error {
	if len(m) == 0 {
		return fmt.Errorf("model has no items")
	}

	tab := map[ScalarType]struct {
		minSize, maxSize   int
		minAlign, maxAlign int
	}{
		Ptr:       {4, 8, 4, 8},
		Void:      {0, 0, 1, 1},
		Char:      {1, 1, 1, 1},
		UChar:     {1, 1, 1, 1},
		Short:     {2, 2, 2, 2},
		UShort:    {2, 2, 2, 2},
		Int:       {4, 4, 4, 4},
		UInt:      {4, 4, 4, 4},
		Long:      {4, 8, 4, 8},
		ULong:     {4, 8, 4, 8},
		LongLong:  {8, 8, 8, 8},
		ULongLong: {8, 8, 8, 8},
		Float:     {4, 4, 4, 4},
		Double:    {8, 8, 8, 8},
		Bool:      {1, 1, 1, 1},
		Complex:   {8, 16, 8, 16},
	}
	a := []int{}
	required := map[ScalarType]bool{}
	seen := map[ScalarType]bool{}
	for k := range tab {
		required[k] = true
		a = append(a, int(k))
	}
	sort.Ints(a)
	for k, v := range m {
		if seen[k] {
			return fmt.Errorf("model has duplicate item: %s", k)
		}

		seen[k] = true
		if !required[k] {
			return fmt.Errorf("model has invalid type: %s: %#v", k, v)
		}

		for typ, t := range tab {
			if typ == k {
				if v.Size < t.minSize {
					return fmt.Errorf("size %d too small: %s", v.Size, k)
				}

				if v.Size > t.maxSize {
					return fmt.Errorf("size %d too big: %s", v.Size, k)
				}

				if v.Size != 0 && mathutil.PopCount(v.Size) != 1 {
					return fmt.Errorf("size %d is not a power of two: %s", v.Size, k)
				}

				if v.Align < t.minAlign {
					return fmt.Errorf("align %d too small: %s", v.Align, k)
				}

				if v.Align > t.maxAlign {
					return fmt.Errorf("align %d too big: %s", v.Align, k)
				}

				if v.Align < v.Size {
					return fmt.Errorf("align is smaller than size: %s", k)
				}

				if mathutil.PopCount(v.Align) != 1 {
					return fmt.Errorf("align %d is not a power of two: %s", v.Align, k)
				}

				break
			}
		}
	}
	for _, typ := range a {
		if !seen[ScalarType(typ)] {
			return fmt.Errorf("model has no item for type %s", ScalarType(typ))
		}
	}
	return nil
}

func fromSlashes(a []string) []string {
	for i, v := range a {
		a[i] = filepath.FromSlash(v)
	}
	return a
}

type indirectType struct {
	specifier Type
	n         int
}

func newIndirectType(specifier Type, ind int) Type {
	if specifier.Kind() == PtrType || ind < 0 {
		panic("internal error")
	}

	if ind == 0 {
		return specifier
	}

	return &indirectType{specifier, ind}
}

// Alignof implements Type.
func (i *indirectType) Alignof() int { panic("TODO") }

// BaseType implements Type.
func (i *indirectType) BaseType() Type {
	if i.n == 1 {
		return i.specifier
	}

	return &indirectType{i.specifier, i.n - 1}
}

// ElementType implements Type.
func (i *indirectType) ElementType() Type { panic("internal error") }

// Kind implements Type.
func (i *indirectType) Kind() Kind { return PtrType }

// Len implements Type.
func (i *indirectType) Len() (int, bool) { panic("internal error") }

// Name implements Type.
func (i *indirectType) Name() xc.Token { panic("internal error") }

// ParameterTypeList implements Type.
func (i *indirectType) ParameterTypeList() *ParameterTypeList { panic("internal error") }

// ResultType implements Type.
func (i *indirectType) ResultType() Type { panic("internal error") }

// Sizeof implements Type.
func (i *indirectType) Sizeof() int { panic("TODO") }

// StructOrUnionType implements Type.
func (i *indirectType) StructOrUnionType() *StructOrUnionSpecifier { panic("internal error") }
