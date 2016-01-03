// Copyright 2015 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"fmt"
	"go/token"

	"github.com/cznic/mathutil"
	"github.com/xlab/c/xc"
)

type operand interface {
	eval(*Model) (value interface{}, typ Type)
	node
}

type node interface {
	Pos() token.Pos
}

// ---------------------------------------------------------- CompoundStatement

// Scope returns n's scope.
func (n *CompoundStatement) Scope() *Bindings { return n.scope }

// ---------------------------------------------------------------- Declaration

// Declarator returns a synthetic Declarator when n.InitDeclaratorListOpt is
// nil.
func (n *Declaration) Declarator() *Declarator { return n.declarator }

// ------------------------------------------------------ DeclarationSpecifiers

// isInline implements specifier.
func (n *DeclarationSpecifiers) isInline() bool {
	return n.attr&saInline != 0
}

// IsTypedef implements specifier.
func (n *DeclarationSpecifiers) IsTypedef() bool {
	return n.attr&saTypedef != 0
}

// isExtern implements specifier.
func (n *DeclarationSpecifiers) isExtern() bool {
	return n.attr&saExtern != 0
}

// IsStatic implements specifier.
func (n *DeclarationSpecifiers) IsStatic() bool {
	return n.attr&saStatic != 0
}

// isAuto implements specifier.
func (n *DeclarationSpecifiers) isAuto() bool {
	return n.attr&saAuto != 0
}

// isRegister implements specifier.
func (n *DeclarationSpecifiers) isRegister() bool {
	return n.attr&saRegister != 0
}

// isConst implements specifier.
func (n *DeclarationSpecifiers) IsConst() bool {
	return n.attr&saConst != 0
}

// isRestrict implements specifier.
func (n *DeclarationSpecifiers) isRestrict() bool {
	return n.attr&saRestrict != 0
}

// isVolatile implements specifier.
func (n *DeclarationSpecifiers) isVolatile() bool {
	return n.attr&saVolatile != 0
}

// kind implements specifier.
func (n *DeclarationSpecifiers) kind() Kind { return tsValid[n.typeSpecifiers()] }

// typeSpecifiers implements specifier.
func (n *DeclarationSpecifiers) typeSpecifiers() int {
	return n.typeSpecifier
}

// firstTypeSpecifier implements specifier.
func (n *DeclarationSpecifiers) firstTypeSpecifier() *TypeSpecifier {
	for n.Case != 1 { // TypeSpecifier DeclarationSpecifiersOpt
		o := n.DeclarationSpecifiersOpt
		if o == nil {
			return nil
		}

		n = o.DeclarationSpecifiers
	}
	return n.TypeSpecifier
}

// attrs implements specifier.
func (n *DeclarationSpecifiers) attrs() int { return n.attr }

// member implements specifier.
func (n *DeclarationSpecifiers) member(nm int) (*Member, error) {
	return n.firstTypeSpecifier().member(nm)
}

// str implements specifier.
func (n *DeclarationSpecifiers) str() string {
	return specifierString(n)
}

// TypedefName implements Specifier.
func (n *DeclarationSpecifiers) TypedefName() int {
	if n.kind() == TypedefName {
		return n.firstTypeSpecifier().Token.Val
	}
	return 0
}

// ----------------------------------------------------------------- Declarator

// Identifier returns the ID of the name declared by n and the scope the name
// is declared in.
func (n *Declarator) Identifier() (int, *Bindings) {
	dd := n.DirectDeclarator.bottom()
	if dd != nil {
		return dd.Token.Val, dd.DeclarationScope()
	}

	return 0, nil
}

// RawSpecifier returns the raw Specifier associated with n before expanding
// typedefs. The effective Specifier is accessible via the Type field of n.
func (n *Declarator) RawSpecifier() Specifier { return n.specifier }

func (n *Declarator) clone() *Declarator {
	m := *n
	return &m
}

func (n *Declarator) stars() int { return n.PointerOpt.stars() }

func (n *Declarator) isCompatible(m *Declarator) (r bool) {
	return n == m || n.Type.(*ctype).isCompatible(m.Type.(*ctype))
}

func (n *Declarator) setFull(lx *lexer) Type {
	d := n
	var dds []*DirectDeclarator
	for dd := d.DirectDeclarator; dd != nil; dd = dd.directDeclarator() {
		dds = append(dds, dd)
	}
	for i, j := 0, len(dds)-1; i < j; i, j = i+1, j-1 { // reverse
		dds[i], dds[j] = dds[j], dds[i]
	}

	resultAttr := 0
	mask := 0
	for d.specifier != nil && d.specifier.kind() == TypedefName {
		resultAttr |= d.specifier.attrs()
		ts := d.specifier.firstTypeSpecifier()
		dd := ts.scope.Lookup(NSIdentifiers, ts.Token.Val).Node.(*DirectDeclarator) // eg. typedef T dd, (*dd), dd(int), ...
		if dd.Case != 0 {                                                           // IDENTIFIER
			panic("internal error")
		}

		nd := dd.top().declarator
		mask = saTypedef // nd.specifier.isTypedef() == true
		dds2 := nd.Type.(*ctype).dds
		d2 := d.clone()
		d2.specifier = nil
		dd2 := &DirectDeclarator{
			Case:       1, //  '(' Declarator ')'
			Declarator: d2,
		}
		dds = append(dds, dd2)
		dds = append(dds, dds2[1:]...)
		d = nd
	}

	// Inner ((...)) -> (...)
	for {
		changed := false
		w := 0
		for r := 0; r < len(dds); {
			dd := dds[r]
			if r == len(dds)-1 || dd.Case != 1 { // '(' Declarator ')'
				dds[w] = dd
				w++
				r++
				continue
			}

			dd2 := dds[r+1]
			if dd2.Case != 1 {
				dds[w] = dd
				w++
				r++
				continue
			}

			d := dd.Declarator
			d2 := dd2.Declarator
			switch s, s2 := d.stars(), d2.stars(); {
			case s == 0 && s2 == 0:
				dds[w] = dd
				w++
				r += 2
				changed = true
			case s == 0 && s2 != 0:
				dds[w] = dd2
				w++
				r += 2
				changed = true
			case s != 0 && s2 == 0:
				dds[w] = dd
				w++
				r += 2
				changed = true
			case s != 0 && s2 != 0:
				d2 := d2.clone()
				var p *Pointer
				for i := 0; i < s+s2; i++ {
					p = &Pointer{Pointer: p}
				}
				d2.PointerOpt = &PointerOpt{Pointer: p}
				dd2 := dd2.clone()
				dd2.Declarator = d2
				dds[w] = dd2
				w++
				r += 2
				changed = true
			}

		}
		dds = dds[:w]
		if !changed {
			break
		}
	}

	// Outer (...) -> ...
	for {
		i := len(dds) - 1
		if dd := dds[i]; dd.Case == 1 /* '(' Declarator ')' */ && dd.Declarator.stars() == 0 {
			dds = dds[:i:i]
			continue
		}

		break
	}
	resultStars := 0
	i := len(dds) - 1
	if dd := dds[i]; dd.Case == 1 /* '(' Declarator ')' */ {
		resultStars = dd.Declarator.stars()
		dds = dds[:i:i]
	}

	stars := 0
	resultStars += d.stars()
	switch {
	case len(dds) == 1:
		if dds[0].Case != 0 { // IDENTIFIER
			panic("internal error")
		}

		stars, resultStars = resultStars, 0
	default:
		i := 1
	loop:
		for {
			switch dd := dds[i]; dd.Case {
			case 1: // '(' Declarator ')'
				if dds[i-1].Case == 0 { // IDENTIFIER
					stars = dd.Declarator.stars()
				} else {
					//dbg("", resultStars, stars, d.specifier.str(), ddsStr(dds))
					panic("TODO")
				}
				i++
			case
				2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
				6, // DirectDeclarator '(' ParameterTypeList ')'
				7: // DirectDeclarator '(' IdentifierListOpt ')'
				break loop
			default:
				//dbg("", position(n.Pos()), resultStars, stars, d.specifier.str(), ddsStr(dds))
				panic(dd.Case)
			}
		}
	}

	resultSpecifier := d.specifier
	resultAttr |= resultSpecifier.attrs()
	resultAttr &^= mask
	t := &ctype{
		dds:             dds,
		model:           lx.model,
		resultAttr:      resultAttr,
		resultSpecifier: resultSpecifier,
		resultStars:     resultStars,
		stars:           stars,
	}
	n.Type = t
	//dbg("setFull %v: %v: %v, %v %v", position(n.Pos()), t, t.Kind(), t.resultStars, t.stars)
	//dbg("", t.str())

	dd := dds[0]
	id := dd.Token.Val
	if id == 0 {
		return t
	}

	// Determine linkage

	scs := resultAttr & (saTypedef | saExtern | saStatic | saAuto | saRegister)
	sk := lx.scope.kind
	var prev *Declarator
	if b := dd.prev; b != nil {
		prev = b.Node.(*DirectDeclarator).TopDeclarator()
	}

	switch {
	case
		// [0]6.2.2, 6: The following identifiers have no linkage: an
		// identifier declared to be anything other than an object or a
		// function; an identifier declared to be a function parameter;
		// a block scope identifier for an object declared without the
		// storage-class specifier extern.
		resultAttr&saTypedef != 0,
		sk == ScopeParams,
		(sk == ScopeBlock || sk == ScopeMembers) && resultAttr&saExtern == 0:

		n.Linkage = None
	case
		// [0]6.2.2, 3: If the declaration of a file scope identifier
		// for an object or a function contains the storage-class
		// specifier static, the identifier has internal linkage.
		sk == ScopeFile && resultAttr&saStatic != 0:

		n.Linkage = Internal
	case
		// [0]6.2.2, 4: For an identifier declared with the
		// storage-class specifier extern in a scope in which a prior
		// declaration of that identifier is visible, if the prior
		// declaration specifies internal or external linkage, the
		// linkage of the identifier at the later declaration is the
		// same as the linkage specified at the prior declaration.
		resultAttr&saExtern != 0 && prev != nil && (prev.Linkage == Internal || prev.Linkage == External):

		n.Linkage = prev.Linkage
	case
		// [0]6.2.2, 4: If no prior declaration is visible, or if the
		// prior declaration specifies no linkage, then the identifier
		// has external linkage.
		resultAttr&saExtern != 0 && (prev == nil || prev.Linkage == None):

		n.Linkage = External
	case
		// [0]6.2.2, 5: If the declaration of an identifier for a
		// function has no storage-class specifier, its linkage is
		// determined exactly as if it were declared with the
		// storage-class specifier extern.
		t.Kind() == Function && scs == 0,
		// [0]6.2.2, 5: If the declaration of an identifier for an
		// object has file scope and no storage-class specifier, its
		// linkage is external.
		t.Kind() != Function && sk == ScopeFile && scs == 0:

		n.Linkage = External
	}

	if isGenerating {
		return t
	}

	switch n.Linkage {
	case External:
		// [0]6.2.2, 2: In the set of translation units and libraries
		// that constitutes an entire program, each declaration of a
		// particular identifier with external linkage denotes the same
		// object or function.
		if prev, ok := lx.externs[id]; ok && !n.isCompatible(prev) {
			lx.report.Err(n.Pos(),
				"conflicting types for %s '%s' with external linkage, previous declaration at %s '%s'",
				xc.Dict.S(id), n.Type, prev.Pos(), prev.Type)
			break
		}

		lx.externs[id] = n
	case Internal:
		// [0]6.2.2, 2: Within one translation unit, each declaration
		// of an identifier with internal linkage denotes the same
		// object or function.
		if prev != nil && !n.isCompatible(prev) {
			lx.report.Err(n.Pos(),
				"conflicting types for %s '%s' with internal linkage, previous declaration at %s '%s'",
				xc.Dict.S(id), n.Type, prev.Pos(), prev.Type)
		}
	case None:
		// [0]6.2.2, 2: Each declaration of an identifier with no
		// linkage denotes a unique entity.
		if prev != nil {
			lx.report.Err(n.Pos(),
				"redeclaration of %s '%s' with no linkage, previous declaration at %v '%s'",
				xc.Dict.S(id), n.Type, position(prev.Pos()), prev.Type)

		}
	default:
		panic("internal error")
	}

	return t
}

// ----------------------------------------------------------- DirectDeclarator

// DeclarationScope returns the scope a name declared by n is in. If n does not
// declare a name or n declares a name of a built in type, DeclarationScope
// returns nil.
func (n *DirectDeclarator) DeclarationScope() *Bindings {
	return n.idScope
}

// TopDeclarator returns the top level Declarator associated with n.
func (n *DirectDeclarator) TopDeclarator() *Declarator {
	return n.top().declarator
}

func (n *DirectDeclarator) top() *DirectDeclarator {
	for n.parent != nil {
		n = n.parent
	}
	return n
}

func (n *DirectDeclarator) bottom() *DirectDeclarator {
	for n.Case != 0 { // IDENTIFIER
		n = n.directDeclarator()
	}
	return n
}

func (n *DirectDeclarator) clone() *DirectDeclarator {
	m := *n
	return &m
}

func (n *DirectDeclarator) isCompatible(m *DirectDeclarator) (r bool) {
	if n == m {
		return true
	}

	if n.Case != m.Case {
		return false
	}

	switch n.Case {
	case 0: // IDENTIFIER
		return true
	case 1: // '(' Declarator ')'
		return true // Declarator checked before
	case 2: // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
		// [0]6.7.5.3 6: For two array types to be compatible, both
		// shall have compatible element types, and if both size
		// specifiers are present, and are integer constant
		// expressions, then both size specifiers shall have the same
		// constant value. If the two array types are used in a context
		// which requires them to be compatible, it is undefined
		// behavior if the two size specifiers evaluate to unequal
		// values.
		var nv, mv interface{}
		if o := n.ExpressionOpt; o != nil {
			nv = o.Expression.Value
		}
		if o := m.ExpressionOpt; o != nil {
			mv = o.Expression.Value
		}
		if nv != nil && mv != nil && nv != mv {
			return false
		}

		return true
	case 6: // DirectDeclarator '(' ParameterTypeList ')'
		return n.ParameterTypeList.isCompatible(m.ParameterTypeList)
	case 7: // DirectDeclarator '(' IdentifierListOpt ')'
		if o := n.IdentifierListOpt; o != nil {
			panic("TODO")
		}

		return true
	default:
		panic(n.Case)
	}
}

func (n *DirectDeclarator) directDeclarator() *DirectDeclarator {
	switch n.Case {
	case 0: // IDENTIFIER
		return nil
	case 1: // '(' Declarator ')'
		return n.Declarator.DirectDeclarator
	case
		2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
		3, // DirectDeclarator '[' "static" TypeQualifierListOpt Expression ']'
		4, // DirectDeclarator '[' TypeQualifierList "static" Expression ']'
		5, // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
		6, // DirectDeclarator '(' ParameterTypeList ')'
		7: // DirectDeclarator '(' IdentifierListOpt ')'
		return n.DirectDeclarator
	default:
		panic(n.Case)
	}
}

// ----------------------------------------------------------------- Expression

func (n *Expression) eval(m *Model) (interface{}, Type) {
	if n.Type == nil {
		n.Type = undefined
		switch n.Case {
		case 0: // IDENTIFIER
			b := n.scope.Lookup(NSIdentifiers, n.Token.Val)
			if b.Node == nil {
				m.report.ErrTok(n.Token, "undefined: %s", n.Token.S())
				break
			}

			dd := b.Node.(*DirectDeclarator)
			n.Type = dd.top().declarator.Type
			if dd.isEnumConst {
				n.Value = m.value(dd.enumVal, m.IntType)
			}
		case 1: // CHARCONST
			n.Value, n.Type = m.charConst(n.Token)
		case 2: // FLOATCONST
			n.Value, n.Type = m.floatConst(n.Token)
		case 3: // INTCONST
			n.Value, n.Type = m.intConst(n.Token)
		case 6: // STRINGLITERAL
			n.Value, n.Type = m.strConst(n.Token)
		case 7: //  '(' ExpressionList ')'
			n.Value, n.Type = n.ExpressionList.eval(m)
		case 8: // Expression '[' ExpressionList ']'
			_, t := n.Expression.eval(m)
			if k := t.Kind(); k != Ptr && k != Array {
				m.report.ErrTok(n.Token, "subscripted value is not a pointer (have '%s')", t)
				break
			}

			_, t2 := n.ExpressionList.eval(m)
			if !IsIntType(t2) {
				m.report.Err(n.ExpressionList.Pos(), "array subscript is not an integer (have '%s')", t2)
			}
			n.Type = t.Element()
		case 9: // Expression '(' ArgumentExpressionListOpt ')'
			_, t := n.Expression.eval(m)
			if t.Kind() == Ptr {
				t = t.Element()
			}
			if t.Kind() != Function {
				m.report.Err(n.Expression.Pos(), "called object is not a function or function pointer (have '%s')", t)
				break
			}

			n.Type = t.Result()
			params, isVariadic := t.Parameters()
			if params == nil {
				panic("internal error")
			}

			var args []*Expression
			var types []Type
			if o := n.ArgumentExpressionListOpt; o != nil {
				for l := o.ArgumentExpressionList; l != nil; l = l.ArgumentExpressionList {
					ex := l.Expression
					args = append(args, ex)
					_, t := ex.eval(m)
					types = append(types, t)
				}
			}

			if g, e := len(args), len(params); g < e {
				m.report.ErrTok(n.Token, "too few arguments to function (have %v, want %v)", g, e)
				break
			}

			if !isVariadic {
				if len(args) > len(params) {
					m.report.Err(n.ArgumentExpressionListOpt.Pos(), "too many arguments to function")
					break
				}
			}

			for i, param := range params {
				typ := types[i]
				if !typ.CanAssignTo(param.Type) {
					m.report.Err(args[i].Pos(), "expected '%s' but argument is of type '%s'", param.Type, typ)
				}
			}
		case 10: // Expression '.' IDENTIFIER
			_, t := n.Expression.eval(m)
			mb, err := t.Member(n.Token2.Val)
			if err != nil {
				m.report.Err(n.Token2.Pos(), "%v", err)
				break
			}

			n.Type = mb.Type
		case 11: // Expression "->" IDENTIFIER
			v, t := n.Expression.eval(m)
			if t.Kind() != Ptr {
				m.report.ErrTok(n.Token2, "invalid type argument of -> (have '%v')", t)
				break
			}

			t = t.Element()
			mb, err := t.Member(n.Token2.Val)
			if err != nil {
				m.report.Err(n.Token2.Pos(), "%v", err)
				break
			}

			n.Type = mb.Type
			switch x := v.(type) {
			case nil:
				// nop
			case uintptr:
				n.Value = x + uintptr(mb.OffsetOf)
			default:
				panic("internal error")
			}
		case 12: // Expression "++"
			n.Value, n.Type = n.Expression.eval(m)
		case 13: // Expression "--"
			n.Value, n.Type = n.Expression.eval(m)
		case 15: // "++" Expression
			n.Value, n.Type = n.Expression.eval(m)
		case 16: // "--" Expression
			n.Value, n.Type = n.Expression.eval(m)
		case 17: // '&' Expression
			var t Type
			n.Value, t = n.Expression.eval(m)
			n.Type = t.Pointer()
		case 18: // '*' Expression
			_, t := n.Expression.eval(m)
			if t.Kind() != Ptr {
				m.report.ErrTok(n.Token, "invalid type argument of unary * (have '%v')", t)
				break
			}

			n.Type = t.Element()
		case 19: // '+' Expression
			n.Value, n.Type = n.Expression.eval(m)
		case 20: // '-' Expression
			v, t := n.Expression.eval(m)
			n.Type = t
			switch x := v.(type) {
			case nil:
				// nop
			case int32:
				n.Value = -x
			case uint32:
				n.Value = -x
			case int64:
				n.Value = -x
			case float64:
				n.Value = -x
			default:
				panic(fmt.Errorf("internal error: %T", x))
			}
		case 21: // '~' Expression
			v, t := n.Expression.eval(m)
			n.Type = t
			switch x := v.(type) {
			case nil:
				// nop
			case int32:
				n.Value = ^x
			case uint32:
				n.Value = ^x
			case uint64:
				n.Value = ^x
			default:
				panic(fmt.Errorf("internal error: %T", x))
			}
		case 22: // '!' Expression
			v, _ := n.Expression.eval(m)
			n.Type = m.IntType
			if v == nil {
				break
			}

			n.Value = m.cBool(isZero(v))
		case 23: // "sizeof" Expression
			n.Type = m.getSizeType(n.Token)
			_, t := n.Expression.eval(m)
			n.Value = m.value(int32(t.SizeOf()), n.Type)
		case 24: // "sizeof" '(' TypeName ')'
			n.Type = m.getSizeType(n.Token)
			n.Value = m.value(int32(n.TypeName.declarator.Type.SizeOf()), n.Type)
		case 25: // '(' TypeName ')' Expression
			v, _ := n.Expression.eval(m)
			n.Type = n.TypeName.declarator.Type
			if v != nil {
				n.Value = m.value(v, n.Type)
			}
		case 26: // Expression '*' Expression
			var a, b interface{}
			a, b, n.Type = m.binOp(n.Expression, n.Expression2)

			switch x := a.(type) {
			case nil:
				// nop
			case int32:
				n.Value = x * b.(int32)
			case uint32:
				n.Value = x * b.(uint32)
			case int64:
				n.Value = x * b.(int64)
			case float64:
				n.Value = x * b.(float64)
			default:
				panic(fmt.Errorf("internal error: %T", x))
			}
		case 27: // Expression '/' Expression
			var a, b interface{}
			a, b, n.Type = m.binOp(n.Expression, n.Expression2)
			if b != nil && isZero(b) {
				m.report.Err(n.Expression2.Pos(), "division by zero")
				break
			}

			switch x := a.(type) {
			case nil:
				// nop
			case int32:
				n.Value = x / b.(int32)
			case uint32:
				n.Value = x / b.(uint32)
			case int64:
				n.Value = x / b.(int64)
			case float64:
				n.Value = x / b.(float64)
			default:
				panic(fmt.Errorf("internal error: %T", x))
			}
		case 28: // Expression '%' Expression
			var a, b interface{}
			a, b, n.Type = m.binOp(n.Expression, n.Expression2)
			if b != nil && isZero(b) {
				m.report.Err(n.Expression2.Pos(), "division by zero")
				break
			}

			switch x := a.(type) {
			case nil:
				// nop
			case int32:
				n.Value = x % b.(int32)
			case uint32:
				n.Value = x % b.(uint32)
			default:
				panic(fmt.Errorf("internal error: %T", x))
			}
		case 29: // Expression '+' Expression
			_, at := n.Expression.eval(m)
			_, bt := n.Expression2.eval(m)
			if at.Kind() > bt.Kind() {
				at, bt = bt, at
			}
			switch {
			case at.Kind() == Ptr:
				if IsIntType(bt) {
					n.Type = at
					break
				}

				m.report.ErrTok(n.Token, "incompatible types ('%s' + '%s')", at, bt)
			case IsArithmeticType(at):
				fallthrough
			default:
				var a, b interface{}
				a, b, n.Type = m.binOp(n.Expression, n.Expression2)
				switch x := a.(type) {
				case nil:
					// nop
				case int32:
					n.Value = x + b.(int32)
				case int64:
					n.Value = x + b.(int64)
				default:
					panic(fmt.Errorf("internal error: %T", x))
				}
			}
		case 30: // Expression '-' Expression
			_, at := n.Expression.eval(m)
			_, bt := n.Expression.eval(m)
			if at.Kind() == Ptr && bt.Kind() == Ptr {
				if !at.CanAssignTo(bt) {
					n.Type = undefined
					m.report.Err(n.Expression2.Pos(), "incompatible types ('%s' - '%s')", at, bt)
					break
				}

				n.Type = m.getPtrDiffType(n.Token)
				break
			}

			var a, b interface{}
			a, b, n.Type = m.binOp(n.Expression, n.Expression2)
			switch x := a.(type) {
			case nil:
				// nop
			case int32:
				n.Value = x - b.(int32)
			case uint32:
				n.Value = x - b.(uint32)
			case int64:
				n.Value = x - b.(int64)
			case uint64:
				n.Value = x - b.(uint64)
			default:
				panic(fmt.Errorf("internal error: %T", x))
			}
		case 31: // Expression "<<" Expression
			av, at := n.Expression.eval(m)
			bv, _ := n.Expression2.eval(m)
			n.Type = at
			if av == nil || bv == nil {
				break
			}

			switch x := av.(type) {
			case int32:
				switch y := bv.(type) {
				case int32:
					switch {
					case y > 0:
						n.Value = x << uint(y)
					case y < 0:
						n.Value = x >> uint(-y)
					default:
						n.Value = x
					}
				default:
					panic(fmt.Errorf("internal error: %T", y))
				}
			case uint32:
				switch y := bv.(type) {
				case int32:
					switch {
					case y > 0:
						n.Value = x << uint(y)
					case y < 0:
						n.Value = x >> uint(-y)
					default:
						n.Value = x
					}
				default:
					panic(fmt.Errorf("internal error: %T", y))
				}
			case int64:
				switch y := bv.(type) {
				case int32:
					switch {
					case y > 0:
						n.Value = x << uint(y)
					case y < 0:
						n.Value = x >> uint(-y)
					default:
						n.Value = x
					}
				default:
					panic(fmt.Errorf("internal error: %T", y))
				}
			case uint64:
				switch y := bv.(type) {
				case int32:
					switch {
					case y > 0:
						n.Value = x << uint(y)
					case y < 0:
						n.Value = x >> uint(-y)
					default:
						n.Value = x
					}
				default:
					panic(fmt.Errorf("internal error: %T", y))
				}
			default:
				panic(fmt.Errorf("internal error: %T", x))
			}
		case 32: // Expression ">>" Expression
			av, at := n.Expression.eval(m)
			bv, _ := n.Expression2.eval(m)
			n.Type = at
			if av == nil || bv == nil {
				break
			}

			switch x := av.(type) {
			case int32:
				switch y := bv.(type) {
				case int32:
					switch {
					case y > 0:
						n.Value = x >> uint(y)
					case y < 0:
						n.Value = x << uint(-y)
					default:
						n.Value = x
					}
				default:
					panic(fmt.Errorf("internal error: %T", y))
				}
			case uint32:
				switch y := bv.(type) {
				case int32:
					switch {
					case y > 0:
						n.Value = x >> uint(y)
					case y < 0:
						n.Value = x << uint(-y)
					default:
						n.Value = x
					}
				default:
					panic(fmt.Errorf("internal error: %T", y))
				}
			case int64:
				switch y := bv.(type) {
				case int32:
					switch {
					case y > 0:
						n.Value = x >> uint(y)
					case y < 0:
						n.Value = x << uint(-y)
					default:
						n.Value = x
					}
				default:
					panic(fmt.Errorf("internal error: %T", y))
				}
			default:
				panic(fmt.Errorf("internal error: %T", x))
			}
		case 33: // Expression '<' Expression
			n.Type = m.IntType
			_, at := n.Expression.eval(m)
			_, bt := n.Expression2.eval(m)
			a0, b0 := at, bt
			if at.Kind() > bt.Kind() {
				at, bt = bt, at
			}
			switch {
			case at.Kind() == Ptr:
				if !at.CanAssignTo(bt) {
					m.report.ErrTok(n.Token, "incompatible types ('%s' < '%s')", a0, b0)
				}
				break
			case IsArithmeticType(at):
				fallthrough
			default:
				n.Type = m.IntType
				var a, b interface{}
				a, b, _ = m.binOp(n.Expression, n.Expression2)
				switch x := a.(type) {
				case nil:
					// nop
				case int32:
					n.Value = m.cBool(x < b.(int32))
				case uint32:
					n.Value = m.cBool(x < b.(uint32))
				case int64:
					n.Value = m.cBool(x < b.(int64))
				default:
					panic(fmt.Errorf("internal error: %T", x))
				}
			}
		case 34: // Expression '>' Expression
			n.Type = m.IntType
			_, at := n.Expression.eval(m)
			_, bt := n.Expression2.eval(m)
			a0, b0 := at, bt
			if at.Kind() > bt.Kind() {
				at, bt = bt, at
			}
			switch {
			case at.Kind() == Ptr:
				if !at.CanAssignTo(bt) {
					m.report.ErrTok(n.Token, "incompatible types ('%s' > '%s')", a0, b0)
				}
				break
			case IsArithmeticType(at):
				fallthrough
			default:
				n.Type = m.IntType
				var a, b interface{}
				a, b, _ = m.binOp(n.Expression, n.Expression2)
				switch x := a.(type) {
				case nil:
					// nop
				case int32:
					n.Value = m.cBool(x > b.(int32))
				case uint32:
					n.Value = m.cBool(x > b.(uint32))
				default:
					panic(fmt.Errorf("internal error: %T", x))
				}
			}
		case 35: // Expression "<=" Expression
			n.Type = m.IntType
			_, at := n.Expression.eval(m)
			_, bt := n.Expression2.eval(m)
			a0, b0 := at, bt
			if at.Kind() > bt.Kind() {
				at, bt = bt, at
			}
			switch {
			case at.Kind() == Ptr:
				if !at.CanAssignTo(bt) {
					m.report.ErrTok(n.Token, "incompatible types ('%s' <= '%s')", a0, b0)
				}
				break
			case IsArithmeticType(at):
				fallthrough
			default:
				n.Type = m.IntType
				var a, b interface{}
				a, b, _ = m.binOp(n.Expression, n.Expression2)
				switch x := a.(type) {
				case nil:
					// nop
				case int32:
					n.Value = m.cBool(x <= b.(int32))
				case uint32:
					n.Value = m.cBool(x <= b.(uint32))
				case int64:
					n.Value = m.cBool(x <= b.(int64))
				default:
					panic(fmt.Errorf("internal error: %T", x))
				}
			}
		case 36: // Expression ">=" Expression
			n.Type = m.IntType
			_, at := n.Expression.eval(m)
			_, bt := n.Expression2.eval(m)
			a0, b0 := at, bt
			if at.Kind() > bt.Kind() {
				at, bt = bt, at
			}
			switch {
			case at.Kind() == Ptr:
				if !at.CanAssignTo(bt) {
					m.report.ErrTok(n.Token, "incompatible types ('%s' >= '%s')", a0, b0)
				}
				break
			case IsArithmeticType(at):
				fallthrough
			default:
				var a, b interface{}
				a, b, _ = m.binOp(n.Expression, n.Expression2)
				switch x := a.(type) {
				case nil:
					// nop
				case int32:
					n.Value = m.cBool(x >= b.(int32))
				case uint32:
					n.Value = m.cBool(x >= b.(uint32))
				case int64:
					n.Value = m.cBool(x >= b.(int64))
				default:
					panic(fmt.Errorf("internal error: %T", x))
				}
			}
		case 37: // Expression "==" Expression
			n.Type = m.IntType
			_, at := n.Expression.eval(m)
			_, bt := n.Expression2.eval(m)
			a0, b0 := at, bt
			if at.Kind() > bt.Kind() {
				at, bt = bt, at
			}
			switch {
			case at.Kind() == Ptr:
				if IsIntType(bt) {
					break
				}

				if bt.Kind() == Array {
					bt = bt.(*ctype).arrayDecay()
				}
				if bt.Kind() == Function && at.Element().Kind() == Function {
					bt = bt.Pointer()
				}
				if !at.CanAssignTo(bt) {
					m.report.ErrTok(n.Token, "incompatible types ('%s' == '%s')", a0, b0)
				}
				break
			case IsArithmeticType(at):
				fallthrough
			default:
				var a, b interface{}
				a, b, _ = m.binOp(n.Expression, n.Expression2)
				if a == nil {
					break
				}

				n.Value = m.cBool(a == b)
			}
		case 38: // Expression "!=" Expression
			n.Type = m.IntType
			_, at := n.Expression.eval(m)
			_, bt := n.Expression2.eval(m)
			if at.Kind() > bt.Kind() {
				at, bt = bt, at
			}
			switch {
			case at.Kind() == Ptr:
				if IsIntType(bt) {
					break
				}

				if bt.Kind() == Function && at.Element().Kind() == Function {
					bt = bt.Pointer()
				}
				if bt.Kind() == Array {
					bt = bt.(*ctype).arrayDecay()
				}
				if !at.CanAssignTo(bt) {
					m.report.ErrTok(n.Token, "incompatible types ('%s' != '%s')", at, bt)
				}
				break
			case IsArithmeticType(at):
				fallthrough
			default:
				var a, b interface{}
				a, b, _ = m.binOp(n.Expression, n.Expression2)
				if a == nil {
					break
				}

				n.Value = m.cBool(a != b)
			}
		case 39: // Expression '&' Expression
			var a, b interface{}
			a, b, n.Type = m.binOp(n.Expression, n.Expression2)
			switch x := a.(type) {
			case nil:
				// nop
			case int32:
				n.Value = x & b.(int32)
			case uint32:
				n.Value = x & b.(uint32)
			case int64:
				n.Value = x & b.(int64)
			default:
				panic(fmt.Errorf("internal error: %T", x))
			}
		case 40: // Expression '^' Expression
			var a, b interface{}
			a, b, n.Type = m.binOp(n.Expression, n.Expression2)
			switch x := a.(type) {
			case nil:
				// nop
			case int32:
				n.Value = x ^ b.(int32)
			case uint32:
				n.Value = x ^ b.(uint32)
			default:
				panic(fmt.Errorf("internal error: %T", x))
			}
		case 41: // Expression '|' Expression
			var a, b interface{}
			a, b, n.Type = m.binOp(n.Expression, n.Expression2)
			switch x := a.(type) {
			case nil:
				// nop
			case int32:
				n.Value = x | b.(int32)
			case uint32:
				n.Value = x | b.(uint32)
			case int64:
				n.Value = x | b.(int64)
			case uint64:
				n.Value = x | b.(uint64)
			default:
				panic(fmt.Sprintf("internal error: %T", x))
			}
		case 42: // Expression "&&" Expression
			n.Type = m.IntType
			a, _ := n.Expression.eval(m)
			if a != nil && isZero(a) {
				n.Value = m.cBool(false)
				break
			}

			b, _ := n.Expression2.eval(m)
			if a != nil && b != nil {
				if isZero(b) {
					n.Value = m.cBool(false)
					break
				}

				n.Value = m.cBool(true)
				break
			}
		case 43: // Expression "||" Expression
			n.Type = m.IntType
			av, _ := n.Expression.eval(m)
			if av != nil && isNonZero(av) {
				n.Value = m.cBool(true)
				break
			}

			bv, _ := n.Expression2.eval(m)
			if av != nil && bv != nil {
				n.Value = m.cBool(isNonZero(bv))
				break
			}
		case 44: // Expression '?' ExpressionList ':' Expression
			lv, _ := n.Expression.eval(m)
			if lv == nil {
				_, at := n.ExpressionList.eval(m)
				_, bt := n.Expression2.eval(m)
				if IsArithmeticType(at) && IsArithmeticType(bt) {
					n.Type = m.binOpType(at, bt)
					break
				}

				ak := at.Kind()
				bk := bt.Kind()
				if ak == Struct && bk == Struct ||
					ak == Union && bk == Union {
					if at.CanAssignTo(bt) {
						n.Type = at
						break
					}
				}

				if ak == Void && bk == Void {
					n.Type = at
					break
				}

				if ak == Array && bk == Array {
					if at.(*ctype).isCompatible(bt.(*ctype)) {
						n.Type = at
						break
					}

					at = at.(*ctype).arrayDecay()
					ak = at.Kind()
					bt = bt.(*ctype).arrayDecay()
					bk = bt.Kind()
				}

				if ak == Array && bk == Ptr && at.CanAssignTo(bt) {
					n.Type = bt
					break
				}

				if ak == Ptr && bk == Array && bt.CanAssignTo(at) {
					n.Type = at
					break
				}

				if ak == Ptr && bk == Ptr {
					if at.CanAssignTo(bt) {
						n.Type = at
						break
					}
				}

				if ak == Ptr && IsIntType(bt) {
					n.Type = at
					break
				}

				if bk == Ptr && IsIntType(at) {
					n.Type = bt
					break
				}

				if ak == Ptr && at.Element().Kind() == Void && bk == Ptr {
					n.Type = bt
					break
				}

				if bk == Ptr && bt.Element().Kind() == Void && ak == Ptr {
					n.Type = at
					break
				}

				m.report.ErrTok(n.Token2, "'%s'/'%s' mismatch in conditional expression", at, bt)
				break
			}

			if isNonZero(lv) {
				n.Value, n.Type = n.ExpressionList.eval(m)
				break
			}

			n.Value, n.Type = n.Expression2.eval(m)
		case 45: // Expression '=' Expression
			_, at := n.Expression.eval(m)
			_, bt := n.Expression2.eval(m)
			if bt.Kind() == Function {
				bt = bt.Pointer()
			}
			if !bt.CanAssignTo(at) {
				m.report.Err(n.Expression2.Pos(), "assignment from incompatible type ('%s' = '%s')", at, bt)
				break
			}

			n.Type = at
		case 46: // Expression "*=" Expression
			_, n.Type = n.Expression.eval(m)
			if _, _, t := m.binOp(n.Expression, n.Expression2); t.Kind() == Undefined {
				m.report.ErrTok(n.Token, "incompatible types") //TODO have ...
			}
		case
			47, // Expression "/=" Expression
			48: // Expression "%=" Expression
			m.checkArithmeticType(n.Expression, n.Expression2)
			n.Type = n.Expression.Type
			if v := n.Expression2.Value; v != nil && isZero(v) {
				m.report.Err(n.Expression2.Pos(), "division by zero")
				break
			}

			if _, _, t := m.binOp(n.Expression, n.Expression2); t.Kind() == Undefined {
				m.report.ErrTok(n.Token, "incompatible types") //TODO have ...
			}
		case
			49, // Expression "+=" Expression
			50: // Expression "-=" Expression
			_, at := n.Expression.eval(m)
			_, bt := n.Expression2.eval(m)
			n.Type = at
			switch {
			case at.Kind() == Ptr:
				if IsIntType(bt) {
					break
				}

				panic("TODO")
			case IsArithmeticType(at):
				fallthrough
			default:
				if _, _, n.Type = m.binOp(n.Expression, n.Expression2); n.Type.Kind() == Undefined {
					m.report.ErrTok(n.Token, "incompatible types") //TODO have ...
				}
			}
		case
			51, // Expression "<<=" Expression
			52, // Expression ">>=" Expression
			53, // Expression "&=" Expression
			55: // Expression "|=" Expression
			m.checkIntegerType(n.Expression, n.Expression2)
			n.Type = n.Expression.Type
		default:
			panic(n.Case)
		}
	}
	//ct := n.Type.(*ctype)
	//s := ""
	//if n.Value != nil {
	//	s = fmt.Sprintf("value: %T(%#v)", n.Value, n.Value)
	//}
	//dbg("tc %v %v %v %v %v: %v %v", position(n.Pos()), n.Case, ct.resultStars, ct.stars, ct, ct.Kind(), s)
	if m.adjustFnArgs {
		if n.Type.Kind() == Function {
			n.Type = n.Type.Pointer()
			//ct = n.Type.(*ctype)
			//dbg("\tadjusted")
			//dbg("\ttc %v: %v %v %v: %v", position(n.Pos()), ct.resultStars, ct.stars, ct, ct.Kind())
		}
	}
	return n.Value, n.Type
}

// ------------------------------------------------------------- ExpressionList

func (n *ExpressionList) eval(m *Model) (interface{}, Type) {
	if n.Type != nil {
		return n.Value, n.Type
	}

	n0 := n
	for ; n != nil; n = n.ExpressionList {
		n.Value, n.Type = n.Expression.eval(m)
		n0.Value, n0.Type = n.Value, n.Type
	}
	return n0.Value, n0.Type
}

// Len returns the number of items in n.
func (n *ExpressionList) Len() (r int) {
	for ; n != nil; n = n.ExpressionList {
		r++
	}
	return r
}

// ---------------------------------------------------------------- Initializer

func (n *Initializer) typeCheck(dt Type, mb []Member, i, limit int, lx *lexer) {
	if limit >= 0 && i > limit {
		lx.report.Err(n.Pos(), "excess elements in array/struct initializer")
		return
	}

	if dt == nil {
		dt = mb[i].Type
	}

	switch n.Case {
	case 0: // Expression
		st := n.Expression.Type
		if !st.CanAssignTo(dt) {
			lx.report.Err(n.Expression.Pos(), "incompatible types when initializing type '%s' using type ‘%s'", dt, st)
		}
	case 1: // '{' InitializerList CommaOpt '}'
		switch dk := dt.Kind(); dk {
		case Array:
			lim := dt.Elements()
			dt := dt.Element()
			i := 0
			for l := n.InitializerList; l != nil; l = l.InitializerList {
				l.Initializer.typeCheck(dt, nil, i, lim, lx)
				i++
			}
		case Struct, Union:
			mb, _ := dt.Members()
			if mb == nil {
				panic("internal error")
			}

			lim := len(mb)
			if dk == Union {
				lim = 1
			}
			i := 0
			for l := n.InitializerList; l != nil; l = l.InitializerList {
				l.Initializer.typeCheck(nil, mb, i, lim, lx)
				i++
			}
		default:
			panic(dk.String())
		}
	default:
		panic(n.Case)
	}
}

// ------------------------------------------------------------ InitializerList

// Len returns the number of items in n.
func (n *InitializerList) Len() (r int) {
	for ; n != nil; n = n.InitializerList {
		r++
	}
	return r
}

// ------------------------------------------------------- ParameterDeclaration

func (n *ParameterDeclaration) isCompatible(m *ParameterDeclaration) (r bool) {
	return n == m || n.declarator.isCompatible(m.declarator)
}

// -------------------------------------------------------------- ParameterList

func (n *ParameterList) isCompatible(m *ParameterList) (r bool) {
	if n == m {
		return true
	}

	for ; n != nil; n, m = n.ParameterList, m.ParameterList {
		if m == nil || !n.ParameterDeclaration.isCompatible(m.ParameterDeclaration) {
			return false
		}
	}
	return m == nil
}

// ---------------------------------------------------------- ParameterTypeList

func (n *ParameterTypeList) isCompatible(m *ParameterTypeList) (r bool) {
	return n == m || n.Case == m.Case && n.ParameterList.isCompatible(m.ParameterList)
}

func (n *ParameterTypeList) post() {
	for l := n.ParameterList; l != nil; l = l.ParameterList {
		d := l.ParameterDeclaration.declarator
		nm, _ := d.Identifier()
		t := d.Type
		//dbg("", position(d.Pos()), t)
		if t.Kind() == Function {
			t = t.Pointer()
		}
		if t.Kind() == Array {
			t = t.(*ctype).arrayDecay()
		}
		n.params = append(n.params, Parameter{
			Declarator: d,
			Name:       nm,
			Type:       t,
		})
	}
	if len(n.params) == 1 && n.params[0].Type.Kind() == Void {
		n.params = make([]Parameter, 0) // Must be non nil.
	}
}

// -------------------------------------------------------------------- Pointer

func (n *Pointer) stars() (r int) {
	for ; n != nil; n = n.Pointer {
		r++
	}
	return r
}

// ----------------------------------------------------------------- PointerOpt

func (n *PointerOpt) stars() int {
	if n == nil {
		return 0
	}

	return n.Pointer.stars()
}

// ----------------------------------------------------- SpecifierQualifierList

// isInline implements specifier.
func (n *SpecifierQualifierList) isInline() bool {
	return n.attr&saInline != 0
}

// IsTypedef implements specifier.
func (n *SpecifierQualifierList) IsTypedef() bool {
	return n.attr&saTypedef != 0
}

// isExtern implements specifier.
func (n *SpecifierQualifierList) isExtern() bool {
	return n.attr&saExtern != 0
}

// IsStatic implements specifier.
func (n *SpecifierQualifierList) IsStatic() bool {
	return n.attr&saStatic != 0
}

// isAuto implements specifier.
func (n *SpecifierQualifierList) isAuto() bool {
	return n.attr&saAuto != 0
}

// isRegister implements specifier.
func (n *SpecifierQualifierList) isRegister() bool {
	return n.attr&saRegister != 0
}

// isConst implements specifier.
func (n *SpecifierQualifierList) IsConst() bool {
	return n.attr&saConst != 0
}

// isRestrict implements specifier.
func (n *SpecifierQualifierList) isRestrict() bool {
	return n.attr&saRestrict != 0
}

// isVolatile implements specifier.
func (n *SpecifierQualifierList) isVolatile() bool {
	return n.attr&saVolatile != 0
}

// kind implements specifier.
func (n *SpecifierQualifierList) kind() Kind { return tsValid[n.typeSpecifiers()] }

// typeSpecifiers implements specifier.
func (n *SpecifierQualifierList) typeSpecifiers() int {
	return n.typeSpecifier
}

// firstTypeSpecifier implements specifier.
func (n *SpecifierQualifierList) firstTypeSpecifier() *TypeSpecifier {
	for n.Case != 0 { // TypeSpecifier SpecifierQualifierListOpt
		o := n.SpecifierQualifierListOpt
		if o == nil {
			return nil
		}

		n = o.SpecifierQualifierList
	}
	return n.TypeSpecifier
}

// attrs implements specifier.
func (n *SpecifierQualifierList) attrs() int { return n.attr }

// member implements specifier.
func (n *SpecifierQualifierList) member(nm int) (*Member, error) {
	return n.firstTypeSpecifier().member(nm)
}

// str implements specifier.
func (n *SpecifierQualifierList) str() string {
	return specifierString(n)
}

// TypedefName implements Specifier.
func (n *SpecifierQualifierList) TypedefName() int {
	if n.kind() == TypedefName {
		return n.firstTypeSpecifier().Token.Val
	}
	return 0
}

// ----------------------------------------------------------- StructDeclarator

func (n *StructDeclarator) post(lx *lexer) {
	sc := lx.scope
	switch n.Case {
	case 0: // Declarator
		if sc.bitOffset != 0 {
			finishBitField(lx)
		}

		t := n.Declarator.Type
		sz := t.SizeOf()
		al := t.AlignOf()
		switch {
		case sc.isUnion:
			// Track union size.
			sc.maxSize = mathutil.Max(sc.maxSize, sz)
		default:
			off := sc.offset
			sc.offset = align(sc.offset, al) // Bump offset if necessary.
			if pd := sc.prevStructDeclarator; pd != nil {
				pd.padding = sc.offset - off
			}
			n.Declarator.offsetOf = sc.offset
			sc.offset += sz // Allocate sz.
		}
		sc.maxAlign = mathutil.Max(sc.maxAlign, al)
		sc.prevStructDeclarator = n.Declarator
	case 1: // DeclaratorOpt ':' ConstantExpression
		t := lx.model.IntType
		if o := n.DeclaratorOpt; o != nil {
			o.Declarator.offsetOf = sc.offset
			o.Declarator.bitOffset = sc.bitOffset
			sc.prevStructDeclarator = o.Declarator
			t = o.Declarator.Type
			switch t.Kind() {
			case Int, UInt, Long, ULong, Short, UShort:
				// ok
			default:
				lx.report.Err(n.ConstantExpression.Pos(), "bit field has invalid type")
				t = lx.model.IntType
			}
		}
		var w int
		switch x := n.ConstantExpression.Value.(type) {
		case int32:
			w = int(x)
		default:
			panic("internal error")
		}
		if m := t.SizeOf() * 8; w > m {
			lx.report.Err(n.ConstantExpression.Pos(), "width of bit field exceeds its type")
			w = m
		}
		maxBits := lx.model.LongType.SizeOf() * 8
		switch {
		case sc.bitOffset+w > maxBits:
			panic("TODO")
		case sc.bitOffset+w == maxBits:
			panic("TODO")
		default:
			sc.bitOffset += w
			maxBits := lx.model.LongType.SizeOf() * 8
			if sc.bitOffset >= maxBits {
				finishBitField(lx)
			}
		}
	default:
		panic(n.Case)
	}
}

// -------------------------------------------------------------- StructOrUnion

func (n *StructOrUnion) typeSpecifiers() int {
	switch n.Token.Rune {
	case STRUCT:
		return tsStructSpecifier
	case UNION:
		return tsUnionSpecifier
	default:
		panic("internal error")
	}
}

func (n *StructOrUnion) isCompatible(m *StructOrUnion) (r bool) {
	return n == m || n.Case == m.Case
}

func (n *StructOrUnion) str() string {
	switch n.Token.Rune {
	case STRUCT:
		return "struct"
	case UNION:
		return "union"
	default:
		panic("internal error")
	}
}

// ----------------------------------------------------- StructOrUnionSpecifier

func (n *StructOrUnionSpecifier) typeSpecifiers() int { return n.StructOrUnion.typeSpecifiers() }

func (n *StructOrUnionSpecifier) isCompatible(m *StructOrUnionSpecifier) (r bool) {
	if n == m {
		return true
	}

	if !n.StructOrUnion.isCompatible(m.StructOrUnion) {
		return false
	}

	if n.Case > m.Case {
		n, m = m, n
	}
	switch n.Case {
	case 0: // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
		switch m.Case {
		case 1: // StructOrUnion IDENTIFIER
			if o := n.IdentifierOpt; o != nil {
				return o.Token.Val == m.Token.Val
			}

			panic("TODO")
		default:
			panic(m.Case)
		}
	case 1: // StructOrUnion IDENTIFIER
		switch m.Case {
		case 1: // StructOrUnion IDENTIFIER
			return n.Token.Val == m.Token.Val
		default:
			panic(m.Case)
		}
	default:
		panic(n.Case)
	}
}

func (n *StructOrUnionSpecifier) member(nm int) (*Member, error) {
	switch n.Case {
	case 0: // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
		b, s := n.scope.Lookup2(NSIdentifiers, nm)
		if s != n.scope {
			var t []byte
			if o := n.IdentifierOpt; o != nil {
				t = o.Token.S()
			}
			return nil, fmt.Errorf("%s %s has no member named %s", n.StructOrUnion.str(), t, xc.Dict.S(nm))
		}

		d := b.Node.(*DirectDeclarator).top().declarator
		return &Member{
			Bits:       d.bits,
			Declarator: d,
			Name:       nm,
			OffsetOf:   d.offsetOf,
			Type:       d.Type,
		}, nil
	case 1: // StructOrUnion IDENTIFIER
		b := n.scope.Lookup(NSTags, n.Token.Val)
		n2, def := b.Node.(*StructOrUnionSpecifier)
		if !def {
			return nil, fmt.Errorf("invalid use of undefined type '%s %s'", n.StructOrUnion.str(), n.Token.S())
		}

		return n2.member(nm)
	default:
		panic(n.Case)
	}
}

// -------------------------------------------------------------- TypeSpecifier

func (n *TypeSpecifier) member(nm int) (*Member, error) {
	switch n.Case {
	case 11: // StructOrUnionSpecifier
		return n.StructOrUnionSpecifier.member(nm)
	default:
		panic("internal error")
	}
}
