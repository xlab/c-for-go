// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"bytes"
	"fmt"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/cznic/golex/lex"
	"github.com/cznic/mathutil"
	"github.com/cznic/strutil"
	"github.com/xlab/c/xc"
)

var (
	_ Specifier = (*DeclarationSpecifiers)(nil)
	_ Specifier = (*SpecifierQualifierList)(nil)
	_ Specifier = (*spec)(nil)

	_ Type = (*ctype)(nil)
)

var (
	noTypedefNameAfter = map[rune]bool{
		'.':         true,
		ARROW:       true,
		BOOL:        true,
		CHAR:        true,
		COMPLEX:     true,
		DOUBLE:      true,
		ENUM:        true,
		FLOAT:       true,
		GOTO:        true,
		INT:         true,
		LONG:        true,
		SHORT:       true,
		SIGNED:      true,
		STRUCT:      true,
		TYPEDEFNAME: true,
		UNION:       true,
		UNSIGNED:    true,
		VOID:        true,
	}

	undefined        = &ctype{}
	debugTypeStrings bool
)

// Specifier describes a combination of {Function,StorageClass,Type}Specifiers
// and TypeQualifiers.
type Specifier interface {
	IsStatic() bool                     // StorageClassSpecifier "static" present.
	IsTypedef() bool                    // StorageClassSpecifier "typedef" present.
	TypedefName() int                   // TypedefName returns the typedef name ID used, if any, zero otherwise.
	attrs() int                         // Encoded attributes.
	firstTypeSpecifier() *TypeSpecifier //
	isAuto() bool                       // StorageClassSpecifier "auto" present.
	isConst() bool                      // TypeQualifier "const" present.
	isExtern() bool                     // StorageClassSpecifier "extern" present.
	isInline() bool                     // FunctionSpecifier "inline" present.
	isRegister() bool                   // StorageClassSpecifier "register" present.
	isRestrict() bool                   // TypeQualifier "restrict" present.
	isVolatile() bool                   // TypeQualifier "volatile" present.
	kind() Kind                         //
	member(int) (*Member, error)        //
	str() string                        //
	typeSpecifiers() int                // Encoded TypeSpecifier combination.
}

// Type decribes properties of a C type.
type Type interface {
	// AlignOf returns the alignment in bytes of a value of this type when
	// allocated in memory. Incomplete struct types have no alignment and
	// the value returned will be < 0.
	AlignOf() int

	// CanAssignTo returns whether this type can be assigned to dst.
	CanAssignTo(dst Type) bool

	// Declarator returns the full Declarator which defined this type.
	Declarator() *Declarator

	// Element returns the type this Ptr type points to.
	Element() Type

	// Elements returns the number of elements an Array type has. The
	// returned value is < 0 if this type is not an Array or if the array
	// is not of a constant size.
	Elements() int

	// Kidn returns one of Ptr, Void, Int, ...
	Kind() Kind

	// Member returns the type of a member of this Struct or Union type,
	// having numeric name identifier nm.
	Member(nm int) (*Member, error)

	// Members returns the members of a Struct or Union type in declaration
	// order. Returned members are valid iff non nil.
	//
	// Note: Non nil members of length 0 means the struct/union has no
	// members or the type is incomplete, which is indicated by the
	// isIncomplete return value.
	//
	// Note 2: C99 standard does not allow empty structs/unions, but GCC
	// supports that as an extension.
	Members() (members []Member, isIncomplete bool)

	// Parameters returns the parameters of a Function type in declaration
	// order. Result is valid iff params is not nil.
	//
	// Note: len(params) == 0 is fine and just means the function has no
	// parameters.
	Parameters() (params []Parameter, isVariadic bool)

	// Pointer returns a type that points to this type.
	Pointer() Type

	// Result returns the result type of a Function type.
	Result() Type

	// Sizeof returns the number of bytes needed to store a value of this
	// type. Incomplete struct types have no alignment and the value
	// returned will be < 0.
	SizeOf() int

	// Specifier returns the Specifier of this type.
	Specifier() Specifier

	// String returns a C-like type specifier of this type.
	String() string

	// Tag returns the ID of a tag of a Struct or Union type, if any.
	// Otherwise the returned value is zero.
	Tag() int
}

// Member describes a member of a struct or union.
type Member struct {
	BitFieldType Type
	BitOffsetOf  int         // Bit field starting bit.
	Bits         int         // Size in bits for bit fields, 0 otherwise.
	Declarator   *Declarator // Possibly nil for bit fields.
	Name         int
	OffsetOf     int
	Padding      int // Number of unused bytes added to the end of the field to force proper alignment requirements.
	Type         Type
}

// Parameter describes a function argument.
type Parameter struct {
	Declarator *Declarator
	Name       int
	Type       Type
}

// PrettyString pretty prints things produced by this package.
func PrettyString(v interface{}) string {
	return strutil.PrettyString(v, "", "", printHooks)
}

func position(pos token.Pos) token.Position { return fset.Position(pos) }

// Binding records the declaration Node of a declared name.
type Binding struct {
	Node node
	enum bool
}

// Bindings record names declared in a scope.
type Bindings struct {
	identifiers map[int]Binding //
	kind        Scope           // ScopeFile, ...
	parent      *Bindings       //
	tags        map[int]Binding //

	// Scoped helpers.

	mergeScope *Bindings // Fn params.

	specifier Specifier // To store in full declarators.

	// Struct/union field handling.
	bitFieldTypes        []Type      //
	bitOffset            int         //
	isUnion              bool        //
	maxAlign             int         //
	maxSize              int         //
	offset               int         //
	prevStructDeclarator *Declarator //
}

func newBindings(parent *Bindings, kind Scope) *Bindings {
	return &Bindings{
		kind:   kind,
		parent: parent,
	}
}

// Scope retuns the kind of b.
func (b *Bindings) Scope() Scope { return b.kind }

func (b *Bindings) merge(c *Bindings) {
	if b.kind != ScopeBlock || len(b.identifiers) != 0 || c.kind != ScopeParams {
		panic("internal error")
	}

	b.boot(NSIdentifiers)
	for k, v := range c.identifiers {
		b.identifiers[k] = v
	}
}

func (b *Bindings) boot(ns Namespace) map[int]Binding {
	var m *map[int]Binding
	switch ns {
	case NSIdentifiers:
		m = &b.identifiers
	case NSTags:
		m = &b.tags
	default:
		panic(fmt.Errorf("internal error %v", ns))
	}

	mp := *m
	if mp == nil {
		mp = make(map[int]Binding)
		*m = mp
	}
	return mp
}

func (b *Bindings) root() *Bindings {
	for b.parent != nil {
		b = b.parent
	}
	return b
}

// Lookup returns the Binding of id in ns or any of its parents. If id is
// undeclared, the returned Binding has its Node field set to nil.
func (b *Bindings) Lookup(ns Namespace, id int) Binding {
	r, _ := b.Lookup2(ns, id)
	return r
}

// Lookup2 is like Lookup but addionally it returns also the scope in which id
// was found.
func (b *Bindings) Lookup2(ns Namespace, id int) (Binding, *Bindings) {
	if ns == NSTags {
		b = b.root()
	}
	for b != nil {
		m := b.boot(ns)
		if x, ok := m[id]; ok {
			return x, b
		}

		b = b.parent
	}

	return Binding{}, nil
}

func (b *Bindings) declareIdentifier(tok xc.Token, d *DirectDeclarator, report *xc.Report) {
	m := b.boot(NSIdentifiers)
	var p *Binding
	if ex, ok := m[tok.Val]; ok {
		p = &ex
	}

	d.prev = p
	m[tok.Val] = Binding{d, false}
}

func (b *Bindings) declareEnumTag(tok xc.Token, report *xc.Report) {
	b = b.root()
	m := b.boot(NSTags)
	if ex, ok := m[tok.Val]; ok {
		if !ex.enum {
			report.ErrTok(tok, "struct tag redeclared as enum tag, previous declaration/definition: %s", position(ex.Node.Pos()))
		}
		return
	}

	m[tok.Val] = Binding{tok, true}
}

func (b *Bindings) defineEnumTag(tok xc.Token, n node, report *xc.Report) {
	b = b.root()
	m := b.boot(NSTags)
	if ex, ok := m[tok.Val]; ok {
		if !ex.enum {
			report.ErrTok(tok, "struct tag redefined as enum tag, previous declaration/definition: %s", position(ex.Node.Pos()))
			return
		}

		if _, ok := ex.Node.(xc.Token); !ok {
			report.ErrTok(tok, "enum tag redefined, previous definition: %s", position(ex.Node.Pos()))
			return
		}
	}

	m[tok.Val] = Binding{n, true}
}

func (b *Bindings) defineEnumConst(lx *lexer, tok xc.Token, v int) {
	d := lx.model.makeDeclarator(0, tsInt)
	dd := d.DirectDeclarator
	dd.Token = tok
	dd.isEnumConst = true
	dd.enumVal = int32(v)
	d.setFull(lx)
	b.declareIdentifier(tok, dd, lx.report)
	lx.iota = v + 1
}

func (b *Bindings) declareStructTag(tok xc.Token, report *xc.Report) {
	b = b.root()
	m := b.boot(NSTags)
	if ex, ok := m[tok.Val]; ok {
		if ex.enum {
			report.ErrTok(tok, "enum tag redeclared as struct tag, previous declaration/definition: %s", position(ex.Node.Pos()))
		}
		return
	}

	m[tok.Val] = Binding{tok, false}
}

func (b *Bindings) defineStructTag(tok xc.Token, n node, report *xc.Report) {
	b = b.root()
	m := b.boot(NSTags)
	if ex, ok := m[tok.Val]; ok {
		if ex.enum {
			report.ErrTok(tok, "enum tag redefined as struct tag, previous declaration/definition: %s", position(ex.Node.Pos()))
			return
		}

		if _, ok := ex.Node.(xc.Token); !ok {
			report.ErrTok(tok, "struct tag redefined, previous definition: %s", position(ex.Node.Pos()))
			return
		}
	}

	m[tok.Val] = Binding{n, false}
}

func (b *Bindings) isTypedefName(id int) bool {
	x := b.Lookup(NSIdentifiers, id)
	if dd, ok := x.Node.(*DirectDeclarator); ok {
		return dd.specifier.IsTypedef()
	}

	return false
}

func (b *Bindings) lexerHack(tok, prev xc.Token) xc.Token { // https://en.wikipedia.org/wiki/The_lexer_hack
	if noTypedefNameAfter[prev.Rune] {
		return tok
	}

	if tok.Rune == IDENTIFIER && b.isTypedefName(tok.Val) {
		tok.Char = lex.NewChar(tok.Pos(), TYPEDEFNAME)
	}
	return tok
}

func errPos(a ...token.Pos) token.Pos {
	for _, v := range a {
		if v.IsValid() {
			return v
		}
	}

	return token.Pos(0)
}

func isZero(v interface{}) bool { return !isNonZero(v) }

func isNonZero(v interface{}) bool {
	switch x := v.(type) {
	case int32:
		return x != 0
	case int:
		return x != 0
	case uint32:
		return x != 0
	case int64:
		return x != 0
	case float64:
		return x != 0
	case StringLitID, LongStringLitID:
		return true
	default:
		panic(fmt.Errorf("internal error: %T", x))
	}
}

func fromSlashes(a []string) []string {
	for i, v := range a {
		a[i] = filepath.FromSlash(v)
	}
	return a
}

type ctype struct {
	dds             []*DirectDeclarator
	model           *Model
	resultAttr      int
	resultSpecifier Specifier
	resultStars     int
	stars           int
}

func (n *ctype) arrayDecay() *ctype {
	return n.setElements(-1)
}

func (n *ctype) setElements(elems int) *ctype {
	m := *n
	m.dds = append([]*DirectDeclarator(nil), n.dds...)
	for i, dd := range m.dds {
		switch dd.Case {
		case 0: // IDENTIFIER
			// nop
		case 2: // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
			dd := dd.clone()
			dd.elements = elems
			m.dds[i] = dd
			return &m
		default:
			//dbg("", position(dd.Pos()))
			panic(dd.Case)
		}
	}
	panic("internal error") // Not an Array
}

func (n *ctype) isCompatible(m *ctype) (r bool) {
	const ignore = saConst | saStatic | saExtern | saRegister

	if n == m {
		return true
	}

	if len(n.dds) != len(m.dds) || n.resultAttr&^ignore != m.resultAttr&^ignore ||
		n.resultStars != m.resultStars || n.stars != m.stars {
		return false
	}

	for i, n := range n.dds {
		if !n.isCompatible(m.dds[i]) {
			return false
		}
	}

	ns := n.resultSpecifier
	ms := m.resultSpecifier
	if ns == ms {
		return true
	}

	if n.Kind() != m.Kind() {
		return false
	}

	switch ns.kind() {
	case Array:
		panic("internal error")
	case Struct, Union:
		return n.structOrUnionSpecifier().isCompatible(m.structOrUnionSpecifier())
	case Enum:
		panic("internal error")
	case TypedefName:
		panic("internal error")
	default:
		return true
	}
}

func (n *ctype) index(d int) int { return len(n.dds) - 1 + d }

func (n *ctype) top(d int) *DirectDeclarator {
	return n.dds[n.index(d)]
}

// AlignOf implements Type.
func (n *ctype) AlignOf() int {
	if n == undefined {
		return 1
	}

	if n.Kind() == Array {
		return n.Element().AlignOf()
	}

	switch k := n.Kind(); k {
	case Enum:
		return 1
	case
		Void,
		Ptr,
		Char,
		SChar,
		UChar,
		Short,
		UShort,
		Int,
		UInt,
		Long,
		ULong,
		LongLong,
		ULongLong,
		Float,
		Double,
		LongDouble,
		Bool,
		FloatComplex,
		DoubleComplex,
		LongDoubleComplex:
		return n.model.Items[k].Align
	case Struct, Union:
		switch sus := n.structOrUnionSpecifier(); sus.Case {
		case 1: // StructOrUnion IDENTIFIER
			return -1 // Incomplete type
		case 0: // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
			return sus.alignOf
		default:
			panic(sus.Case)
		}
	default:
		panic(k.String())
	}
}

// CanAssignTo implements Type.
func (n *ctype) CanAssignTo(dst Type) bool {
	if n == undefined || dst.Kind() == Undefined {
		return false
	}

	if n.Kind() == Array && dst.Kind() == Ptr {
		n = n.arrayDecay()
	}

	if IsArithmeticType(n) && IsArithmeticType(dst) {
		return true
	}

	if n.Kind() == Ptr && dst.Kind() == Ptr && dst.Element().Kind() == Void {
		return true
	}

	if n.Kind() == Ptr && n.Element().Kind() == Void && dst.Kind() == Ptr {
		return true
	}

	if n.isCompatible(dst.(*ctype)) {
		return true
	}

	if n.Kind() == Ptr && dst.Kind() == Ptr {
		t := Type(n)
		u := dst
		for t.Kind() == Ptr && u.Kind() == Ptr {
			t = t.Element()
			u = u.Element()
		}
		if t.Kind() == Ptr || u.Kind() == Ptr {
			return false
		}

		return t.(*ctype).isCompatible(u.(*ctype))
	}

	if n.Kind() == Function && dst.Kind() == Ptr && dst.Element().Kind() == Function {
		return n.isCompatible(dst.Element().(*ctype))
	}

	if dst.Kind() == Ptr {
		if IsIntType(n) {
			return true
		}
	}

	return false
}

// Declarator implements Type.
func (n *ctype) Declarator() *Declarator {
	return n.dds[0].TopDeclarator()
}

// Element implements Type.
func (n *ctype) Element() Type {
	if n == undefined {
		return n
	}

	if n.Kind() != Ptr && n.Kind() != Array {
		return undefined
	}

	if len(n.dds) == 1 {
		m := *n
		m.stars--
		return &m
	}

	switch dd := n.dds[1]; dd.Case {
	case 1: // '(' Declarator ')'
		if n.stars == 1 {
			m := *n
			m.dds = append([]*DirectDeclarator{n.dds[0]}, n.dds[2:]...)
			m.stars--
			return &m
		}

		m := *n
		m.stars--
		return &m
	case 2: // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
		m := *n
		m.dds = append([]*DirectDeclarator{n.dds[0]}, n.dds[2:]...)
		if len(m.dds) == 1 {
			m.stars += m.resultStars
			m.resultStars = 0
		}
		return &m
	default:
		//dbg("", position(n.dds[0].Pos()), n, n.Kind())
		//dbg("", n.str())
		panic(dd.Case)
	}
}

// Kind implements Type.
func (n *ctype) Kind() Kind {
	if n == undefined {
		return Undefined
	}

	if n.stars > 0 {
		return Ptr
	}

	if len(n.dds) == 1 {
		return n.resultSpecifier.kind()
	}

	i := 1
	for {
		switch dd := n.dds[i]; dd.Case {
		case 2: // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
			if dd.elements < 0 {
				return Ptr
			}

			return Array
		case
			6, // DirectDeclarator '(' ParameterTypeList ')'
			7: // DirectDeclarator '(' IdentifierListOpt ')'
			return Function
		default:
			//dbg("", n)
			//dbg("", n.str())
			panic(dd.Case)
		}
	}
}

// Member implements Type.
func (n *ctype) Member(nm int) (*Member, error) {
	if n == undefined {
		return nil, fmt.Errorf("not a struct/union (have '%s')", n)
	}

	if n.Kind() == Array {
		panic("TODO")
	}

	if k := n.Kind(); k != Struct && k != Union {
		return nil, fmt.Errorf("request for member %s in something not a structure or union (have '%s')", xc.Dict.S(nm), n)
	}

	return n.resultSpecifier.member(nm)
}

func (n *ctype) structOrUnionSpecifier() *StructOrUnionSpecifier {
	if k := n.Kind(); k != Struct && k != Union {
		return nil
	}

	ts := n.resultSpecifier.firstTypeSpecifier()
	if ts.Case != 11 { // StructOrUnionSpecifier
		panic("internal error")
	}

	switch sus := ts.StructOrUnionSpecifier; sus.Case {
	case 0: // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
		return sus
	case 1: // StructOrUnion IDENTIFIER
		b := sus.scope.Lookup(NSTags, sus.Token.Val)
		switch x := b.Node.(type) {
		case nil:
			return sus
		case *StructOrUnionSpecifier:
			return x
		case xc.Token:
			return sus
		default:
			panic("internal error")
		}
	default:
		panic(sus.Case)
	}
}

// Members implements Type.
func (n *ctype) Members() (r []Member, isIncomplete bool) {
	if k := n.Kind(); k != Struct && k != Union {
		return nil, false
	}

	switch sus := n.structOrUnionSpecifier(); sus.Case {
	case 0: // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
		for l := sus.StructDeclarationList; l != nil; l = l.StructDeclarationList {
			for l := l.StructDeclaration.StructDeclaratorList; l != nil; l = l.StructDeclaratorList {
				var d *Declarator
				var bits int
				switch sd := l.StructDeclarator; sd.Case {
				case 0: // Declarator
					d = sd.Declarator
				case 1: // DeclaratorOpt ':' ConstantExpression
					if o := sd.DeclaratorOpt; o != nil {
						d = o.Declarator
					}
					switch x := sd.ConstantExpression.Value.(type) {
					case int32:
						bits = int(x)
					default:
						panic("internal error")
					}
				default:
					panic(sd.Case)
				}
				var id, off, pad, bitoff int
				t := n.model.IntType
				var bt Type
				if d != nil {
					id, _ = d.Identifier()
					t = d.Type
					off = d.offsetOf
					pad = d.padding
					bitoff = d.bitOffset
					bt = d.bitFieldType
				}
				r = append(r, Member{
					BitFieldType: bt,
					BitOffsetOf:  bitoff,
					Bits:         bits,
					Declarator:   d,
					Name:         id,
					OffsetOf:     off,
					Padding:      pad,
					Type:         t,
				})
			}
		}
		return r, false
	case 1: // StructOrUnion IDENTIFIER
		return []Member{}, true
	default:
		panic(sus.Case)
	}
}

// Parameters implements Type.
func (n *ctype) Parameters() ([]Parameter, bool) {
	if n == undefined || n.Kind() != Function {
		return nil, false
	}

	switch dd := n.dds[1]; dd.Case {
	case 6: // DirectDeclarator '(' ParameterTypeList ')'
		l := dd.ParameterTypeList
		return l.params, l.Case == 1 // ParameterList ',' "..."
	case 7: // DirectDeclarator '(' IdentifierListOpt ')'
		if o := dd.IdentifierListOpt; o == nil {
			return make([]Parameter, 0), false
		}

		panic("TODO")
	default:
		//dbg("", dd.Case)
		panic("internal error")
	}
}

// Pointer implements Type.
func (n *ctype) Pointer() Type {
	if n == undefined {
		return n
	}

	if len(n.dds) == 1 {
		m := *n
		m.stars++
		return &m
	}

	if n.Kind() == Array {
		m := *n
		m.stars++
		return &m
	}

	switch dd := n.dds[1]; dd.Case {
	case
		6, // DirectDeclarator '(' ParameterTypeList ')'
		7: // DirectDeclarator '(' IdentifierListOpt ')'
		dd := &DirectDeclarator{
			Case: 1, // '(' Declarator ')'
			Declarator: &Declarator{
				DirectDeclarator: &DirectDeclarator{},
				PointerOpt: &PointerOpt{
					Pointer: &Pointer{},
				},
			},
		}
		m := *n
		m.dds = append(append([]*DirectDeclarator{n.dds[0]}, dd), n.dds[1:]...)
		m.stars++
		return &m
	default:
		m := *n
		m.stars++
		return &m
	}
}

// Result implements Type.
func (n *ctype) Result() Type {
	if n == undefined {
		return n
	}

	if n.Kind() != Function {
		//dbg("", n, n.Kind())
		//dbg("", n.str())
		panic("TODO")
	}

	i := 1
	for {
		switch dd := n.dds[i]; dd.Case {
		case
			6, // DirectDeclarator '(' ParameterTypeList ')'
			7: // DirectDeclarator '(' IdentifierListOpt ')'
			if i == len(n.dds)-1 { // Outermost function.
				if i == 1 {
					m := *n
					m.dds = m.dds[:1:1]
					m.stars += m.resultStars
					m.resultStars = 0
					return &m
				}

				//dbg("", n)
				//dbg("", n.str())
				panic("TODO")
			}

			m := *n
			m.dds = append([]*DirectDeclarator{n.dds[0]}, n.dds[i+1:]...)
			if dd := m.dds[1]; dd.Case == 1 { // '(' Declarator ')'
				m.stars = dd.Declarator.stars()
			}
			return &m
		default:
			//dbg("", position(n.dds[0].Pos()), n)
			//dbg("", n.str())
			panic(dd.Case)
		}

	}
}

// Elements implements Type.
func (n *ctype) Elements() int {
	for _, dd := range n.dds {
		switch dd.Case {
		case 0: // IDENTIFIER
		case
			2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
			3, // DirectDeclarator '[' "static" TypeQualifierListOpt Expression ']'
			4, // DirectDeclarator '[' TypeQualifierList "static" Expression ']'
			5: // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
			return dd.elements
		default:
			//dbg("", position(n.dds[0].Pos()))
			panic(dd.Case)
		}
	}
	return -1
}

// SizeOf implements Type.
func (n *ctype) SizeOf() int {
	if n == undefined {
		return 1
	}

	if n.Kind() == Array {
		switch nelem := n.Elements(); {
		case nelem < 0:
			return n.model.Items[Ptr].Size
		default:
			return nelem * n.Element().SizeOf()
		}
	}

	switch k := n.Kind(); k {
	case Enum:
		return 1
	case
		Void,
		Ptr,
		Char,
		SChar,
		UChar,
		Short,
		UShort,
		Int,
		UInt,
		Long,
		ULong,
		LongLong,
		ULongLong,
		Float,
		Double,
		LongDouble,
		Bool,
		FloatComplex,
		DoubleComplex,
		LongDoubleComplex:
		return n.model.Items[k].Size
	case Struct, Union:
		switch sus := n.structOrUnionSpecifier(); sus.Case {
		case 1: // StructOrUnion IDENTIFIER
			return -1 // Incomplete type
		case 0: // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
			return sus.sizeOf
		default:
			panic(sus.Case)
		}
	default:
		panic(n.Kind().String())
	}
}

// Specifier implements Type.
func (n *ctype) Specifier() Specifier { return &spec{n.resultAttr, n.resultSpecifier.typeSpecifiers()} }

// String implements Type.
func (n *ctype) String() string {
	if n == undefined {
		return "<undefined>"
	}

	var buf bytes.Buffer
	s := attrString(n.resultAttr)
	buf.WriteString(s)
	if s != "" {
		buf.WriteString(" ")
	}
	s = specifierString(n.resultSpecifier)
	buf.WriteString(s)
	buf.WriteString(strings.Repeat("*", n.resultStars))

	pd := func(n *ParameterDeclaration) {
		t := n.declarator.Type
		buf.WriteString(t.String())
	}

	pl := func(n *ParameterList) {
		first := true
		for ; n != nil; n = n.ParameterList {
			if !first {
				buf.WriteString(",")
			}
			pd(n.ParameterDeclaration)
			first = false
		}
	}

	ptl := func(n *ParameterTypeList) {
		pl(n.ParameterList)
		if n.Case == 1 { // ParameterList ',' "..."
			buf.WriteString(",...")
		}
	}

	il := func(n *IdentifierList) {
		first := true
		for ; n != nil; n = n.IdentifierList {
			if !first {
				buf.WriteString(",")
			}
			first = false
			switch n.Case {
			case 0: // IDENTIFIER
				fmt.Fprintf(&buf, "%s", n.Token.S())
			case 1: // IdentifierList ',' IDENTIFIER
				fmt.Fprintf(&buf, "%s", n.Token2.S())
			}
		}
	}

	ilo := func(n *IdentifierListOpt) {
		if n != nil {
			il(n.IdentifierList)
		}
	}

	var f func(int)
	starsWritten := false
	f = func(x int) {
		switch dd := n.top(x); dd.Case {
		case 0: // IDENTIFIER
			if debugTypeStrings {
				id := dd.Token.Val
				if id == 0 {
					id = idID
				}
				fmt.Fprintf(&buf, "<%s>", xc.Dict.S(id))
			}
			if !starsWritten {
				buf.WriteString(strings.Repeat("*", n.stars))
			}
		case 1: // '(' Declarator ')'
			buf.WriteString("(")
			s := 0
			switch dd2 := n.top(x - 1); dd2.Case {
			case 0: // IDENTIFIER
				s = n.stars
				starsWritten = true
			default:
				s = dd.Declarator.stars()
			}
			buf.WriteString(strings.Repeat("*", s))
			f(x - 1)
			buf.WriteString(")")
		case 2: // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
			f(x - 1)
			buf.WriteString("[")
			sep := ""
			if o := dd.TypeQualifierListOpt; o != nil {
				buf.WriteString(attrString(o.TypeQualifierList.attr))
				sep = " "
			}
			if e := dd.elements; e > 0 {
				buf.WriteString(sep)
				fmt.Fprint(&buf, e)
			}
			buf.WriteString("]")
		case 6: // DirectDeclarator '(' ParameterTypeList ')'
			f(x - 1)
			buf.WriteString("(")
			ptl(dd.ParameterTypeList)
			buf.WriteString(")")
		case 7: // DirectDeclarator '(' IdentifierListOpt ')'
			f(x - 1)
			buf.WriteString("(")
			ilo(dd.IdentifierListOpt)
			buf.WriteString(")")
		default:
			panic(dd.Case)
		}
	}
	f(0)
	return buf.String()
}

// Tag implements Type.
func (n *ctype) Tag() int {
	if k := n.Kind(); k != Struct && k != Union {
		return 0
	}

	switch sus := n.structOrUnionSpecifier(); sus.Case {
	case 0: // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
		if o := sus.IdentifierOpt; o != nil {
			return o.Token.Val
		}

		return 0
	case 1: // StructOrUnion IDENTIFIER
		return sus.Token.Val
	default:
		panic(sus.Case)
	}
}

type spec struct {
	attr int
	ts   int
}

func (s *spec) IsStatic() bool                     { return s.attr&saStatic != 0 }
func (s *spec) IsTypedef() bool                    { return s.attr&saTypedef != 0 }
func (s *spec) TypedefName() int                   { return 0 }
func (s *spec) attrs() int                         { return s.attr }
func (s *spec) firstTypeSpecifier() *TypeSpecifier { panic("TODO") }
func (s *spec) isAuto() bool                       { return s.attr&saAuto != 0 }
func (s *spec) isConst() bool                      { return s.attr&saConst != 0 }
func (s *spec) isExtern() bool                     { return s.attr&saExtern != 0 }
func (s *spec) isInline() bool                     { return s.attr&saInline != 0 }
func (s *spec) isRegister() bool                   { return s.attr&saRegister != 0 }
func (s *spec) isRestrict() bool                   { return s.attr&saRestrict != 0 }
func (s *spec) isVolatile() bool                   { return s.attr&saVolatile != 0 }
func (s *spec) kind() Kind                         { return tsValid[s.ts] }
func (s *spec) member(int) (*Member, error)        { panic("TODO") }
func (s *spec) str() string                        { return specifierString(s) }
func (s *spec) typeSpecifiers() int                { return s.ts }

func specifierString(sp Specifier) string {
	if sp == nil {
		return ""
	}

	var buf bytes.Buffer
	switch k := sp.kind(); k {
	case Enum:
		panic("TODO Enum")
	case Function:
		panic("TODO Function")
	case Struct, Union:
		switch ts := sp.firstTypeSpecifier(); ts.Case {
		case 11: // StructOrUnionSpecifier
			sus := ts.StructOrUnionSpecifier
			buf.WriteString(sus.StructOrUnion.str())
			switch sus.Case {
			case 0: // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
				if o := sus.IdentifierOpt; o != nil {
					buf.WriteString(" ")
					buf.Write(o.Token.S())
					break
				}

				buf.WriteString("{")
				outerFirst := true
				for l := sus.StructDeclarationList; l != nil; l = l.StructDeclarationList {
					if !outerFirst {
						buf.WriteString("; ")
					}
					outerFirst = false
					first := true
					for l := l.StructDeclaration.StructDeclaratorList; l != nil; l = l.StructDeclaratorList {
						if !first {
							buf.WriteString(", ")
						}
						first = false
						switch sd := l.StructDeclarator; sd.Case {
						case 0: // Declarator
							buf.WriteString(sd.Declarator.Type.String())
						default:
							fmt.Fprintf(&buf, "specifierString_TODO%v", sd.Case)
						}
					}
				}
				buf.WriteString(";}")
			case 1: // StructOrUnion IDENTIFIER
				buf.WriteString(" ")
				buf.Write(sus.Token.S())
			default:
				panic(sus.Case)
			}
		default:
			panic(ts.Case)
		}
	default:
		buf.WriteString(k.CString())
	}
	return buf.String()
}

func align(off, algn int) int {
	r := off % algn
	if r != 0 {
		off += algn - r
	}
	return off
}

func finishBitField(lx *lexer) {
	sc := lx.scope
	maxBits := lx.model.LongType.SizeOf() * 8
	bits := sc.bitOffset
	if bits > maxBits || bits == 0 {
		panic("internal error")
	}

	var bytes, al int
	for _, k := range []Kind{Char, Short, Int, Long} {
		bytes = lx.model.Items[k].Size
		al = lx.model.Items[k].Align
		if bytes*8 >= bits {
			var t Type
			switch k {
			case Char:
				t = lx.model.CharType
			case Short:
				t = lx.model.ShortType
			case Int:
				t = lx.model.IntType
			case Long:
				t = lx.model.LongType
			default:
				panic("internal error")
			}
			sc.bitFieldTypes = append(sc.bitFieldTypes, t)
			break
		}
	}
	switch {
	case sc.isUnion:
		panic("TODO")
	default:
		off := sc.offset
		sc.offset = align(sc.offset, al)
		if pd := sc.prevStructDeclarator; pd != nil {
			pd.padding = sc.offset - off
			pd.bitOffset = sc.bitOffset
			pd.offsetOf = sc.offset
		}
		sc.offset += bytes
		sc.bitOffset = 0
	}
	sc.maxAlign = mathutil.Max(sc.maxAlign, al)
}

// IsArithmeticType reports wheter t.Kind() is one of Char, SChar, UChar,
// Short, UShort, Int, UInt, Long, ULong, LongLong, ULongLong, Float, Double,
// LongDouble, FloatComplex, DoubleComplex or LongDoubleComplex.
func IsArithmeticType(t Type) bool {
	switch t.Kind() {
	case
		Char,
		SChar,
		UChar,
		Short,
		UShort,
		Int,
		UInt,
		Long,
		ULong,
		LongLong,
		ULongLong,
		Float,
		Double,
		LongDouble,
		FloatComplex,
		DoubleComplex,
		LongDoubleComplex:
		return true
	default:
		return false
	}
}

// IsIntType reports t.Kind() is one of Char, SChar, UChar, Short, UShort, Int,
// UInt, Long, ULong, LongLong or ULongLong.
func IsIntType(t Type) bool {
	switch t.Kind() {
	case
		Char,
		SChar,
		UChar,
		Short,
		UShort,
		Int,
		UInt,
		Long,
		ULong,
		LongLong,
		ULongLong:
		return true
	default:
		return false
	}
}

func elements(v interface{}) (int, error) {
	r, err := toInt(v)
	if err != nil {
		return -1, err
	}

	if r <= 0 {
		return -1, fmt.Errorf("array size must be positive: %v", v)
	}

	return r, nil
}

func toInt(v interface{}) (int, error) {
	switch x := v.(type) {
	case int8:
		return int(x), nil
	case byte:
		return int(x), nil
	case int16:
		return int(x), nil
	case uint16:
		return int(x), nil
	case int32:
		return int(x), nil
	case uint32:
		return int(x), nil
	case int64:
		if x < mathutil.MinInt || x > mathutil.MaxInt {
			return 0, fmt.Errorf("value out of bounds: %v", x)
		}

		return int(x), nil
	case uint64:
		if x > mathutil.MaxInt {
			return 0, fmt.Errorf("value out of bounds: %v", x)
		}

		return int(x), nil
	case int:
		return x, nil
	default:
		return -1, fmt.Errorf("not a constant integer expression: %v", x)
	}
}
