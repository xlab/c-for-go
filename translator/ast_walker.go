package translator

import (
	"fmt"

	"github.com/xlab/c/cc"
	"github.com/xlab/c/xc"
)

func (t *Translator) walkTranslationUnit(unit *cc.TranslationUnit) {
	t.fileScope = unit.Declarations
	for unit != nil {
		t.walkExternalDeclaration(unit.ExternalDeclaration)
		unit = unit.TranslationUnit
	}
}

func (t *Translator) walkExternalDeclaration(d *cc.ExternalDeclaration) {
	switch d.Case {
	case 0: // FunctionDefinition
		// (not an API definition)
	case 1: // Declaration
		declares := t.walkDeclaration(d.Declaration)
		for _, decl := range declares {
			if decl.IsStatic {
				continue
			} else if decl.IsTypedef {
				t.typedefs = append(t.typedefs, decl)
				t.typedefsSet[decl.Name] = struct{}{}
				continue
			}
			t.declares = append(t.declares, decl)
		}
	}
}

func (t *Translator) walkDeclaration(d *cc.Declaration) (declared []*CDecl) {
	if d.InitDeclaratorListOpt != nil {
		list := d.InitDeclaratorListOpt.InitDeclaratorList
		for list != nil {
			decl := t.declarator(list.InitDeclarator.Declarator)
			init := list.InitDeclarator.Initializer
			if init != nil {
				decl.Value = init.Expression.Value
				decl.Expression = blessName(init.Expression.Token.S())
			}
			t.registerTagsOf(decl)
			declared = append(declared, decl)
			if init != nil {
				t.valueMap[decl.Name] = decl.Value
				t.exprMap[decl.Name] = decl.Expression
			}
			list = list.InitDeclaratorList
		}
	} else if declr := d.Declarator(); declr != nil {
		decl := t.declarator(declr)
		t.registerTagsOf(decl)
		declared = append(declared, decl)
	}
	return
}

func (t *Translator) declarator(d *cc.Declarator) *CDecl {
	specifier := d.RawSpecifier()
	decl := &CDecl{
		Spec:      t.typeSpec(d.Type, false, false),
		Name:      identifierOf(d.DirectDeclarator),
		IsTypedef: specifier.IsTypedef(),
		IsStatic:  specifier.IsStatic(),
		Pos:       d.Pos(),
	}
	return decl
}

func (t *Translator) getStructTag(typ cc.Type) (tag string) {
	b := t.fileScope.Lookup(cc.NSTags, typ.Tag())
	sus, ok := b.Node.(*cc.StructOrUnionSpecifier)
	if !ok {
		return
	}
	if sus.IdentifierOpt != nil {
		return blessName(sus.IdentifierOpt.Token.S())
	}
	return
}

func (t *Translator) getEnumTag(typ cc.Type) (tag string) {
	b := t.fileScope.Lookup(cc.NSTags, typ.Tag())
	es, ok := b.Node.(*cc.EnumSpecifier)
	if !ok {
		return
	}
	if es.IdentifierOpt != nil {
		return blessName(es.IdentifierOpt.Token.S())
	}
	return
}

func (t *Translator) enumSpec(base *CTypeSpec, typ cc.Type, isRef bool) *CEnumSpec {
	tag := t.getEnumTag(typ)
	spec := &CEnumSpec{
		Tag:       tag,
		Arrays:    base.Arrays,
		VarArrays: base.VarArrays,
		Pointers:  base.Pointers,
	}

	// this would work with patched cznic/cc, replace when workaround is ready
	enumSpecifier := typ.TypeSpecifier().EnumSpecifier

	var idx int
	switch enumSpecifier.Case {
	case 0: // "enum" IdentifierOpt '{' EnumeratorList CommaOpt '}'
		list := enumSpecifier.EnumeratorList
		for list != nil {
			m := t.walkEnumerator(list.Enumerator)
			list = list.EnumeratorList

			switch {
			case m.Value == nil:
				if idx > 0 {
					// get previous enumerator
					prevM := spec.Members[idx-1]
					m.Value = incVal(prevM.Value)
					switch t.constRules[ConstEnum] {
					case ConstExpandFull:
						m.Expression = prevM.Expression + "+1"
					case ConstExpand:
						m.Expression = prevM.Name + "+1"
					case ConstCGOAlias:
						m.Expression = fmt.Sprintf("C.%s", blessName([]byte(m.Name)))
					}
					m.Spec = prevM.Spec.Copy()
				} else {
					m.Value = int32(0)
					m.Expression = "0"
				}
			case t.constRules[ConstEnum] == ConstCGOAlias:
				m.Expression = fmt.Sprintf("C.%s", blessName([]byte(m.Name)))
			default:
				m.Expression = fmt.Sprintf("%d", m.Value)
			}

			idx++
			m.Spec = spec.PromoteType(m.Value)
			spec.Members = append(spec.Members, m)
			t.valueMap[m.Name] = m.Value
			t.exprMap[m.Name] = m.Expression
		}
	case 2: // "enum" IDENTIFIER
		spec.PromoteType(int32(0))
		return spec
	}
	return spec
}

func (t *Translator) walkEnumerator(e *cc.Enumerator) *CDecl {
	decl := &CDecl{
		Name: blessName(e.EnumerationConstant.Token.S()),
	}
	switch e.Case {
	case 0: // EnumerationConstant
	case 1: // EnumerationConstant '=' ConstantExpression
		decl.Value = e.ConstantExpression.Value
		decl.Expression = string(e.ConstantExpression.Expression.Token.S())
	}
	return decl
}

func (t *Translator) structSpec(base *CTypeSpec, typ cc.Type, isRef bool) *CStructSpec {
	tag := t.getStructTag(typ)
	spec := &CStructSpec{
		Tag:       tag,
		Arrays:    base.Arrays,
		VarArrays: base.VarArrays,
		Pointers:  base.Pointers,
		IsUnion:   typ.Kind() == cc.Union,
	}
	if isRef {
		return spec
	}
	members, _ := typ.Members()
	for _, m := range members {
		spec.Members = append(spec.Members, &CDecl{
			Name: memberName(m),
			Spec: t.typeSpec(m.Type, isRef, false),
			Pos:  m.Declarator.Pos(),
		})
	}
	return spec
}

func (t *Translator) functionSpec(base *CTypeSpec, typ cc.Type, isRef bool) *CFunctionSpec {
	spec := &CFunctionSpec{
		Arrays:    base.Arrays,
		VarArrays: base.VarArrays,
		Pointers:  base.Pointers,
	}
	if isRef {
		return spec
	}
	if ret := typ.Result(); ret != nil && ret.Kind() != cc.Void {
		spec.Return = t.typeSpec(ret, isRef, true)
	}
	params, _ := typ.Parameters()
	for _, p := range params {
		spec.Params = append(spec.Params, &CDecl{
			Name: paramName(p),
			Spec: t.typeSpec(p.Type, isRef, false),
			Pos:  p.Declarator.Pos(),
		})
	}
	return spec
}

func (t *Translator) typeSpec(typ cc.Type, isRef, isRet bool) CType {
	spec := &CTypeSpec{
		Const: typ.Specifier().IsConst(),
	}
	if !isRet {
		spec.Raw = typedefNameOf(typ)
	}
	for typ.Kind() == cc.Array {
		size := typ.Elements()
		typ = typ.Element()
		if size >= 0 {
			spec.AddArray(uint64(size))
		}
	}
	for typ.Kind() == cc.Ptr {
		typ = typ.Element()
		spec.Pointers++
	}

	switch typ.Kind() {
	case cc.Void:
		spec.Base = "void"
	case cc.Ptr:
		spec.Base = "void"
		spec.Pointers = 1
	case cc.Char:
		spec.Base = "char"
	case cc.SChar:
		spec.Base = "char"
		spec.Signed = true
	case cc.UChar:
		spec.Base = "char"
		spec.Unsigned = true
	case cc.Short:
		spec.Base = "int"
		spec.Short = true
	case cc.UShort:
		spec.Base = "int"
		spec.Short = true
		spec.Unsigned = true
	case cc.Int:
		spec.Base = "int"
	case cc.UInt:
		spec.Base = "int"
		spec.Unsigned = true
	case cc.Long:
		spec.Base = "int"
		spec.Long = true
	case cc.ULong:
		spec.Base = "int"
		spec.Long = true
		spec.Unsigned = true
	case cc.LongLong:
		spec.Base = "long"
		spec.Long = true
	case cc.ULongLong:
		spec.Base = "long"
		spec.Long = true
		spec.Unsigned = true
	case cc.Float:
		spec.Base = "float"
	case cc.Double:
		spec.Base = "double"
	case cc.LongDouble:
		spec.Base = "double"
		spec.Long = true
	case cc.Bool:
		spec.Base = "_Bool"
	case cc.FloatComplex:
		spec.Base = "complexfloat"
		spec.Complex = true
	case cc.DoubleComplex:
		spec.Base = "complexdouble"
		spec.Complex = true
	case cc.LongDoubleComplex:
		spec.Base = "complexdouble"
		spec.Long = true
		spec.Complex = true
	case cc.Enum:
		s := t.enumSpec(spec, typ, isRef)
		if !isRet {
			s.Typedef = typedefNameOf(typ)
		}
		return s
	case cc.Struct, cc.Union:
		isRef := false
		tag := t.getStructTag(typ)
		if len(tag) > 0 {
			// avoid recursive structs
			if t.typeCache.Get(tag) {
				isRef = true
			} else {
				t.typeCache.Set(tag)
			}
		}
		s := t.structSpec(spec, typ, isRef)
		if !isRet {
			s.Typedef = typedefNameOf(typ)
		}
		t.typeCache.Delete(tag)
		return s
	case cc.Function:
		s := t.functionSpec(spec, typ, isRef)
		if !isRet && !typ.Specifier().IsTypedef() {
			s.Typedef = typedefNameOf(typ)
			if s.Return != nil {
				s.Return.SetRaw(s.Typedef)
			}
		}
		return s
	default:
		panic("unknown type " + typ.String())
	}

	return spec
}

func paramName(p cc.Parameter) string {
	if p.Name == 0 {
		return ""
	}
	return blessName(xc.Dict.S(p.Name))
}

func memberName(m cc.Member) string {
	if m.Name == 0 {
		return ""
	}
	return blessName(xc.Dict.S(m.Name))
}

func typedefNameOf(typ cc.Type) string {
	rawSpec := typ.Declarator().RawSpecifier()
	if name := rawSpec.TypedefName(); name > 0 {
		return blessName(xc.Dict.S(name))
	} else if rawSpec.IsTypedef() {
		return identifierOf(typ.Declarator().DirectDeclarator)
	}
	return ""
}

func identifierOf(dd *cc.DirectDeclarator) string {
	switch dd.Case {
	case 0: // IDENTIFIER
		if dd.Token.Val == 0 {
			return ""
		}
		return blessName(dd.Token.S())
	case 1: // '(' Declarator ')'
		return identifierOf(dd.Declarator.DirectDeclarator)
	default:
		//	DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
		//	DirectDeclarator '[' "static" TypeQualifierListOpt Expression ']'
		//	DirectDeclarator '[' TypeQualifierList "static" Expression ']'
		//	DirectDeclarator '[' TypeQualifierListOpt '*' ']'
		//	DirectDeclarator '(' ParameterTypeList ')'
		//	DirectDeclarator '(' IdentifierListOpt ')'
		return identifierOf(dd.DirectDeclarator)
	}
}
