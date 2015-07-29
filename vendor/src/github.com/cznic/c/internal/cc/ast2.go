// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"go/token"

	"github.com/cznic/c/internal/xc"
)

var (
	_ Type = (*declarationSpecifiers)(nil)
	_ Type = (*directAbstractDeclarator)(nil)
	_ Type = (*directDeclarator)(nil)
	_ Type = (*functionDefintion)(nil)
	_ Type = (*indirectType)(nil)
	_ Type = (*specifierQualifierList)(nil)
	_ Type = (*structOrUnionSpecifier)(nil)
	_ Type = (*typeSpecifier)(nil)
)

type align struct {
	pos token.Pos
	val int
}

func (a *align) set(n int) int { a.val = n + 1; return n }
func (a *align) get() int      { return a.val - 1 }

// Alignof reports the alignment requirements of a variable, struct/union field
// or a type.
func (a *align) Alignof() int {
	if v := a.get(); v > 0 { // Alignment cannot be zero.
		return v
	}

	compilation.Err(a.pos, "alignment undetermined")
	return a.set(model[Char].Align)
}

type offset struct {
	pos token.Pos
	val int
}

func (o *offset) set(n int) int { o.val = n + 1; return n }
func (o *offset) get() int      { return o.val - 1 }

// Offsetof reports the relative address of a field within a struct or union.
func (o *offset) Offsetof() int {
	if v := o.get(); v >= 0 {
		return v
	}

	compilation.Err(o.pos, "offset undetermined")
	return o.set(0)
}

type size struct {
	pos token.Pos
	val int
}

func (s *size) set(n int) int { s.val = n + 1; return n }
func (s *size) get() int      { return s.val - 1 }

// Sizeof reports the size of a variable, struct/union field or a type.
func (s *size) Sizeof() int {
	if v := s.get(); v >= 0 {
		return v
	}

	compilation.Err(s.pos, "alignment undetermined")
	return s.set(model[Char].Size)
}

func (a *AbstractDeclarator) Type() Type {
	switch a.Case {
	case 0: // Pointer
		return newIndirectType(a.specifier, a.Pointer.indirection)
	default:
		panic(a.Case) //TODO-
		panic("internal error")
	}
}

// FindInitDeclarator returns the init declarator associated with nm or nil if
// no such init declarator exists.
func (d *Declaration) FindInitDeclarator(nm int) *InitDeclarator {
	return d.InitDeclaratorListOpt.FindInitDeclarator(nm)
}

type declarationSpecifiers DeclarationSpecifiers

func (d *declarationSpecifiers) Type() Type { return d }

// BaseType implements Type.
func (d *declarationSpecifiers) BaseType() Type { panic("internal error") }

// Kind implements Type.
func (d *declarationSpecifiers) Kind() Kind {
	return typeSpecifier2TypeKind(d.typ)
}

// Name implements Type.
func (d *declarationSpecifiers) Name() xc.Token { return (*typeSpecifier)(d.typeSpecification()).Name() }

// ParameterTypeList implements Type.
func (d *declarationSpecifiers) ParameterTypeList() *ParameterTypeList { panic("internal error") }

// ResultType implements Type.
func (d *declarationSpecifiers) ResultType() Type { panic("internal error") }

// StructOrUnionType implements Type.
func (d *declarationSpecifiers) StructOrUnionType() *StructOrUnionSpecifier {
	return (*typeSpecifier)(d.typeSpecification()).StructOrUnionType()
}

// typeSpecification returns the first TypeSpecifier of d, if any, or nil
// otherwise.
func (d *declarationSpecifiers) typeSpecification() *TypeSpecifier {
	for {
		if t := d.TypeSpecifier; t != nil {
			return t
		}

		o := d.DeclarationSpecifiersOpt
		if o == nil {
			return nil
		}

		d = (*declarationSpecifiers)(o.DeclarationSpecifiers)
	}
}

// Type returns the type of d.
func (d *DeclarationSpecifiers) Type() Type { return (*declarationSpecifiers)(d) }

func (d *DeclarationSpecifiers) sum(e *DeclarationSpecifiersOpt) {
	if e == nil {
		return
	}

	f := e.DeclarationSpecifiers
	d.IsAuto = d.IsAuto || f.IsAuto
	d.IsConst = d.IsConst || f.IsConst
	d.IsExtern = d.IsExtern || f.IsExtern
	d.IsInline = d.IsInline || f.IsInline
	d.IsRegister = d.IsRegister || f.IsRegister
	d.IsRestrict = d.IsRestrict || f.IsRestrict
	d.IsStatic = d.IsStatic || f.IsStatic
	d.IsTypedef = d.IsTypedef || f.IsTypedef
	d.IsVolatile = d.IsVolatile || f.IsVolatile
	d.typeSpecifiers = f.typeSpecifiers
	d.typ = f.typ
}

func (d *DeclarationSpecifiers) typeSum() {
	x, ok := typeSums[typeSum1(d.typeSpecifiers)]
	if !ok {
		compilation.ErrTok((*declarationSpecifiers)(d).typeSpecification().Tok(), "invalid type specification")
		x = tsVoid
	}

	d.typ = x
	return
}

// Type returns typeof d.
func (d *Declarator) Type() Type {
	return d.DirectDeclarator.spec0()
}

// Ident returns the name token of d.
func (d *Declarator) Ident() xc.Token {
	return d.DirectDeclarator.ident()
}

// SUSpecifier returns the StructOrUnionSpecifier of a field d declares.
// SUSpecifier will panic if d does not declare a field.
func (d *Declarator) SUSpecifier() *StructOrUnionSpecifier { return d.SUSpecifier0.SUSpecifier }

func (d *Declarator) insert(sc *Bindings, ns Namespace, isDefinition bool) {
	dd := d.DirectDeclarator
	for {
		switch dd.Case {
		case 0: // IDENTIFIER
			d.IsDefinition = isDefinition
			sc.insert(NSIdentifiers, dd.Token, d)
			return
		case 1: //  '(' Declarator ')'
			dd.Declarator.insert(sc, ns, isDefinition)
			return
		case
			2, // DirectDeclarator '[' TypeQualifierListOpt AssignmentExpressionOpt ']'
			3, // DirectDeclarator '[' "static" TypeQualifierListOpt AssignmentExpression ']'
			4, // DirectDeclarator '[' TypeQualifierList "static" AssignmentExpression ']'
			5, // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
			6: // DirectDeclarator '(' DirectDeclarator2
			dd = dd.DirectDeclarator
		default:
			panic("internal error")
		}
	}
}

type directAbstractDeclarator DirectAbstractDeclarator

// BaseType implements Type.
func (d *directAbstractDeclarator) BaseType() Type {
	panic("internal error")
}

// Kind implements Type.
func (d *directAbstractDeclarator) Kind() Kind {
	panic("internal error")
}

// Name implements Type.
func (d *directAbstractDeclarator) Name() xc.Token {
	panic("internal error")
}

// ParameterTypeList implements Type.
func (d *directAbstractDeclarator) ParameterTypeList() *ParameterTypeList {
	panic("internal error")
}

// ResultType implements Type.
func (d *directAbstractDeclarator) ResultType() Type {
	panic("internal error")
}

// StructOrUnionType implements Type.
func (d *directAbstractDeclarator) StructOrUnionType() *StructOrUnionSpecifier {
	panic("internal error")
}

type directDeclarator DirectDeclarator

// BaseType implements Type.
func (d *directDeclarator) BaseType() Type {
	return (*DirectDeclarator)(d).spec().BaseType()
}

// Kind implements Type.
func (d *directDeclarator) Kind() Kind {
	switch d.Case {
	case 0: // IDENTIFIER
		return (*DirectDeclarator)(d).spec().Kind()
	// case 1: // '(' Declarator ')'
	case
		2, // DirectDeclarator '[' TypeQualifierListOpt AssignmentExpressionOpt ']'
		3, // DirectDeclarator '[' "static" TypeQualifierListOpt AssignmentExpression ']'
		4, // DirectDeclarator '[' TypeQualifierList "static" AssignmentExpression ']'
		5: // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
		return ArrayType
	case 6: // DirectDeclarator '(' DirectDeclarator2
		return FunctionType
	default:
		panic("internal error")
	}
}

// Name implements Type.
func (d *directDeclarator) Name() xc.Token { return (*DirectDeclarator)(d).spec().Name() }

// ParameterTypeList implements Type.
func (d *directDeclarator) ParameterTypeList() *ParameterTypeList {
	switch d.Case {
	case 6: // DirectDeclarator '(' DirectDeclarator2
		return d.DirectDeclarator2.ParameterTypeList
	default:
		panic(d.Case)
	}
}

// ResultType implements Type.
func (d *directDeclarator) ResultType() Type {
	switch d.Case {
	case 6: // DirectDeclarator '(' DirectDeclarator2
		return (*DirectDeclarator)(d).spec()
	default:
		panic(d.Case)
	}
}

// StructOrUnionType implements Type.
func (d *directDeclarator) StructOrUnionType() *StructOrUnionSpecifier {
	return (*DirectDeclarator)(d).spec().StructOrUnionType()
}

func (d *DirectDeclarator) postProc(sc *Bindings) {
	dd := d.DirectDeclarator
	dd.specifier = (*directDeclarator)(d)
	if dd.Case == 1 {
		dd.Declarator.DirectDeclarator.specifier = (*directDeclarator)(d)
	}
	d.specifier = sc.specifier
}

func (d *DirectDeclarator) spec() Type { return newIndirectType(d.specifier, d.indirection) }

func (d *DirectDeclarator) ident() xc.Token {
	switch d.Case {
	case 0: // IDENTIFIER
		return d.Token
	case 1: // '(' Declarator ')'
		return d.Declarator.Ident()
	default:
		return d.DirectDeclarator.ident()
	}
}

func (d *DirectDeclarator) spec0() Type {
	switch d.Case {
	case 0: // IDENTIFIER
		return d.spec()
	case 1: // '(' Declarator ')'
		return d.Declarator.DirectDeclarator.spec0()
	default:
		return d.DirectDeclarator.spec0()
	}
}

// Tag returns (tag name of e, true) or (zero-value, false) if e has no tag
// name.
func (e *EnumSpecifier) Tag() (xc.Token, bool) {
	switch e.Case {
	case
		0, // EnumSpecifier0 '{' EnumeratorList '}'
		1: // EnumSpecifier0 '{' EnumeratorList ',' '}'
		return e.EnumSpecifier0.Tag()
	case 2: // "enum" IDENTIFIER
		return e.Token2, true
	default:
		panic("internal error")
	}
}

// Tok returns the effective token of e.
func (e *EnumSpecifier) Tok() xc.Token {
	switch e.Case {
	case
		0, // EnumSpecifier0 '{' EnumeratorList '}'
		1: // EnumSpecifier0 '{' EnumeratorList ',' '}'
		return e.EnumSpecifier0.Tok()
	case 2: // "enum" IDENTIFIER
		return e.Token2
	default:
		panic("internal error")
	}
}

// Tok returns the effective token of e.
func (e *EnumSpecifier0) Tok() xc.Token {
	if o := e.IdentifierOpt; o != nil {
		return o.Token
	}

	return e.Token
}

// Tag returns (tag name of e, true) or (zero-value, false) if e has no tag
// name.
func (e *EnumSpecifier0) Tag() (xc.Token, bool) {
	o := e.IdentifierOpt
	if o == nil {
		return xc.Token{}, false
	}

	return o.Token, true
}

// Type returns the type of f.
func (f *FunctionDefinition) Type() Type { return (*functionDefintion)(f) }

type functionDefintion FunctionDefinition

// BaseType implements Type.
func (f *functionDefintion) BaseType() Type { panic("internal error") }

// Kind implements Type.
func (f *functionDefintion) Kind() Kind { return FunctionType }

// Name implements Type.
func (f *functionDefintion) Name() xc.Token { panic("internal error") }

// ParameterTypeList implements Type.
func (f *functionDefintion) ParameterTypeList() *ParameterTypeList {
	return (*FunctionDefinition)(f).Declarator.DirectDeclarator.DirectDeclarator2.ParameterTypeList
}

// ResultType implements Type.
func (f *functionDefintion) ResultType() Type {
	return newIndirectType(f.DeclarationSpecifiers.Type(), f.Declarator.DirectDeclarator.indirection)
}

// StructOrUnionType implements Type.
func (f *functionDefintion) StructOrUnionType() *StructOrUnionSpecifier { panic("internal error") }

// FindInitDeclarator returns the init declarator associated with nm or nil if
// no such init declarator exists.
func (i *InitDeclaratorList) FindInitDeclarator(nm int) *InitDeclarator {
	for ; i != nil; i = i.InitDeclaratorList {
		if id := i.InitDeclarator; id.Declarator.Ident().Val == nm {
			return id
		}
	}
	return nil
}

// FindInitDeclarator returns the init declarator associated with nm or nil if
// no such init declarator exists.
func (i *InitDeclaratorListOpt) FindInitDeclarator(nm int) *InitDeclarator {
	if i == nil {
		return nil
	}

	return i.InitDeclaratorList.FindInitDeclarator(nm)
}

func (p *PointerOpt) indirection() int {
	if p == nil {
		return 0
	}

	return p.Pointer.indirection
}

type specifierQualifierList SpecifierQualifierList

// BaseType implements Type.
func (s *specifierQualifierList) BaseType() Type { panic("internal error") }

// Kind implements Type.
func (s *specifierQualifierList) Kind() Kind {
	return typeSpecifier2TypeKind(s.typ)
}

// Name implements Type.
func (s *specifierQualifierList) Name() xc.Token {
	return (*typeSpecifier)(s.typeSpecification()).Name()
}

// ParameterTypeList implements Type.
func (s *specifierQualifierList) ParameterTypeList() *ParameterTypeList { panic("internal error") }

// ResultType implements Type.
func (s *specifierQualifierList) ResultType() Type { panic("internal error") }

// StructOrUnionType implements Type.
func (s *specifierQualifierList) StructOrUnionType() *StructOrUnionSpecifier {
	return (*typeSpecifier)(s.typeSpecification()).StructOrUnionType()
}

// typeSpecification returns the first TypeSpecifier of s, if any, or nil
// otherwise.
func (s *specifierQualifierList) typeSpecification() *TypeSpecifier {
	for {
		if t := s.TypeSpecifier; t != nil {
			return t
		}

		o := s.SpecifierQualifierListOpt
		if o == nil {
			return nil
		}

		s = (*specifierQualifierList)(o.SpecifierQualifierList)
	}
}

func (s *SpecifierQualifierList) sum(e *SpecifierQualifierListOpt) {
	if e == nil {
		return
	}

	f := e.SpecifierQualifierList
	s.IsConst = s.IsConst || f.IsConst
	s.IsRestrict = s.IsRestrict || f.IsRestrict
	s.IsVolatile = s.IsVolatile || f.IsVolatile
	s.typeSpecifiers = f.typeSpecifiers
	s.typ = f.typ
}

func (s *SpecifierQualifierList) typeSum() {
	x, ok := typeSums[typeSum1(s.typeSpecifiers)]
	if !ok {
		compilation.ErrTok((*specifierQualifierList)(s).typeSpecification().Tok(), "invalid type specification")
		x = tsVoid
	}

	s.typ = x
	return
}

func (s *StructOrUnionSpecifier) Type() Type { return (*structOrUnionSpecifier)(s) }

// Tag returns (tag name of s, true) or (zero-value, false) if s has no
// tag name.
func (s *StructOrUnionSpecifier) Tag() (xc.Token, bool) {
	switch s.Case {
	case 0: // StructOrUnionSpecifier0 '{' StructDeclarationList '}'
		return s.StructOrUnionSpecifier0.Tag()
	case 1: // StructOrUnion IDENTIFIER
		return s.Token, true
	default:
		panic("internal error")
	}
}

// Tok returns the effective token of s.
func (s *StructOrUnionSpecifier) Tok() xc.Token {
	switch s.Case {
	case 0: // StructOrUnionSpecifier0 '{' StructDeclarationList '}'
		return s.StructOrUnionSpecifier0.Tok()
	case 1: // StructOrUnion IDENTIFIER
		return s.Token
	default:
		panic("internal error")
	}
}

type structOrUnionSpecifier StructOrUnionSpecifier

// BaseType implements Type.
func (s *structOrUnionSpecifier) BaseType() Type { panic("internal error") }

// Kind implements Type.
func (s *structOrUnionSpecifier) Kind() Kind {
	if s.isUnion {
		return UnionType
	}

	return StructType
}

// Name implements Type.
func (s *structOrUnionSpecifier) Name() xc.Token { panic("internal error") }

// ParameterTypeList implements Type.
func (s *structOrUnionSpecifier) ParameterTypeList() *ParameterTypeList { panic("internal error") }

// ResultType implements Type.
func (s *structOrUnionSpecifier) ResultType() Type { panic("internal error") }

// StructOrUnionType implements Type.
func (s *structOrUnionSpecifier) StructOrUnionType() *StructOrUnionSpecifier {
	return (*StructOrUnionSpecifier)(s)
}

// StructOrUnionType implements Type.
func (s *StructOrUnionSpecifier0) StructOrUnionType() *StructOrUnionSpecifier { return s.SUSpecifier }

// Tok returns the effective token of s.
func (s *StructOrUnionSpecifier0) Tok() xc.Token {
	if o := s.IdentifierOpt; o != nil {
		return o.Token
	}

	return s.StructOrUnion.Token
}

// Tag returns (tag name of s, true) or (zero-value, false) if s has no tag
// name.
func (s *StructOrUnionSpecifier0) Tag() (xc.Token, bool) {
	o := s.IdentifierOpt
	if o == nil {
		return xc.Token{}, false
	}

	return o.Token, true
}

// Tok returns the effective token of t.
func (t *TypeSpecifier) Tok() xc.Token {
	switch t.Case {
	case
		0,  // "void"
		1,  // "char"
		2,  // "short"
		3,  // "int"
		4,  // "long"
		5,  // "float"
		6,  // "double"
		7,  // "signed"
		8,  // "unsigned"
		9,  // "_Bool"
		10, // "_Complex"
		13: // TYPEDEFNAME
		return t.Token
	case 11: // StructOrUnionSpecifier
		return t.StructOrUnionSpecifier.Tok()
	case 12: // EnumSpecifier
		return t.EnumSpecifier.Tok()
	default:
		panic("internal error")
	}
}

type typeSpecifier TypeSpecifier

// BaseType implements Type.
func (t *typeSpecifier) BaseType() Type { panic("internal error") }

// Kind implements Type.
func (t *typeSpecifier) Kind() Kind {
	switch t.Case {
	case 11: // StructOrUnionSpecifier
		return typeSpecifier2TypeKind(t.case2)
	default:
		return typeSpecifier2TypeKind(t.Case)
	}
}

// Name implements Type.
func (t *typeSpecifier) Name() xc.Token {
	switch t.Case {
	case 13: // TYPEDEFNAME
		return t.Token
	default:
		panic("internal error")
	}
}

// ParameterTypeList implements Type.
func (t *typeSpecifier) ParameterTypeList() *ParameterTypeList { panic("internal error") }

// ResultType implements Type.
func (t *typeSpecifier) ResultType() Type { panic("internal error") }

// StructOrUnionType implements Type.
func (t *typeSpecifier) StructOrUnionType() *StructOrUnionSpecifier {
	switch t.Case {
	case 11: // StructOrUnionSpecifier
		return t.StructOrUnionSpecifier
	default:
		panic(PrettyString(t))
	}
}
