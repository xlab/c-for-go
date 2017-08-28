package translator

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/cznic/cc"
	"github.com/cznic/xc"
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
		var decl *CDecl
		if declr := d.FunctionDefinition.Declarator; declr != nil {
			decl = t.declarator(declr)
			t.registerTagsOf(decl)
		} else {
			return
		}
		if decl.IsTypedef {
			t.typedefs = append(t.typedefs, decl)
			t.typedefsSet[decl.Name] = struct{}{}
			return
		}
		t.declares = append(t.declares, decl)
	case 1: // Declaration
		declares := t.walkDeclaration(d.Declaration)
		for _, decl := range declares {
			if decl.IsTypedef {
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
			if init != nil && init.Expression != nil {
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
		Spec:      t.typeSpec(d.Type, 0, false),
		Name:      identifierOf(d.DirectDeclarator),
		IsTypedef: specifier.IsTypedef(),
		IsStatic:  specifier.IsStatic(),
		Pos:       d.Pos(),
	}
	return decl
}

func (t *Translator) getStructTag(typ cc.Type) (tag string) {
	b := t.fileScope.Lookup(cc.NSTags, typ.Tag())
	switch v := b.Node.(type) {
	case *cc.StructOrUnionSpecifier:
		if v.IdentifierOpt != nil {
			return blessName(v.IdentifierOpt.Token.S())
		}
	case xc.Token:
		return blessName(v.S())
	}
	return
}

func (t *Translator) enumSpec(base *CTypeSpec, typ cc.Type) *CEnumSpec {
	tag := blessName(xc.Dict.S(typ.Tag()))
	spec := &CEnumSpec{
		Tag:      tag,
		Pointers: base.Pointers,
		OuterArr: base.OuterArr,
		InnerArr: base.InnerArr,
	}
	list := typ.EnumeratorList()
	for _, en := range list {
		name := blessName(en.DefTok.S())
		m := &CDecl{
			Name: name,
			Pos:  en.DefTok.Pos(),
		}
		switch {
		case en.Value == nil:
			panic("value cannot be nil in enum")
		case t.constRules[ConstEnum] == ConstCGOAlias:
			m.Expression = fmt.Sprintf("C.%s", name)
		case t.constRules[ConstEnum] == ConstExpand:
			srcParts := make([]string, 0, len(en.Tokens))
			exprParts := make([]string, 0, len(en.Tokens))
			valid := true

			// TODO: some state machine
			needsTypecast := false
			typecastValue := false
			typecastValueParens := 0

			for _, token := range en.Tokens {
				src := cc.TokSrc(token)
				srcParts = append(srcParts, src)
				switch token.Rune {
				case cc.IDENTIFIER:
					exprParts = append(exprParts, string(t.TransformName(TargetConst, src, true)))
				default:
					// TODO: state machine
					const (
						lparen = rune(40)
						rparen = rune(41)
					)
					switch {
					case needsTypecast && token.Rune == rparen:
						typecastValue = true
						needsTypecast = false
						exprParts = append(exprParts, src+"(")
					case typecastValue && token.Rune == lparen:
						typecastValueParens++
					case typecastValue && token.Rune == rparen:
						if typecastValueParens == 0 {
							typecastValue = false
							exprParts = append(exprParts, ")"+src)
						} else {
							typecastValueParens--
						}
					default:
						// somewhere in the world a helpless kitten died because of this
						if token.Rune == '~' {
							src = "^"
						}
						if runes := []rune(src); len(runes) > 0 && isNumeric(runes) {
							// TODO(xlab): better const handling
							src = readNumeric(runes)
						}
						exprParts = append(exprParts, src)
					}
				}
				if !valid {
					break
				}
			}
			if typecastValue {
				// still in typecast value, need to close paren
				exprParts = append(exprParts, ")")
				typecastValue = false
			}

			if len(exprParts) > 0 {
				m.Expression = strings.Join(exprParts, " ")
			} else {
				m.Value = en.Value
			}
			m.Src = strings.Join(srcParts, " ")
		default:
			m.Value = en.Value
		}
		m.Spec = spec.PromoteType(en.Value)
		spec.Members = append(spec.Members, m)
		t.valueMap[m.Name] = en.Value
		t.exprMap[m.Name] = m.Expression
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

const maxDeepLevel = 3

func (t *Translator) structSpec(base *CTypeSpec, typ cc.Type, deep int) *CStructSpec {
	tag := t.getStructTag(typ)
	spec := &CStructSpec{
		Tag:      tag,
		IsUnion:  typ.Kind() == cc.Union,
		Pointers: base.Pointers,
		OuterArr: base.OuterArr,
		InnerArr: base.InnerArr,
	}
	if deep > maxDeepLevel {
		return spec
	}
	members, _ := typ.Members()
	for i, m := range members {
		var pos token.Pos
		if m.Declarator != nil {
			pos = m.Declarator.Pos()
		}
		spec.Members = append(spec.Members, &CDecl{
			Name: memberName(i, m),
			Spec: t.typeSpec(m.Type, deep+1, false),
			Pos:  pos,
		})
	}
	return spec
}

func (t *Translator) functionSpec(base *CTypeSpec, typ cc.Type, deep int) *CFunctionSpec {
	spec := &CFunctionSpec{
		Pointers: base.Pointers,
	}
	if deep > 2 { // a function inside params of another function
		spec.Raw = typedefNameOf(typ)
	} else {
		spec.Raw = identifierOf(typ.Declarator().DirectDeclarator)
	}
	if deep > maxDeepLevel {
		return spec
	}
	if ret := typ.Result(); ret != nil && ret.Kind() != cc.Void {
		spec.Return = t.typeSpec(ret, deep+1, true)
	}
	params, _ := typ.Parameters()
	for i, p := range params {
		spec.Params = append(spec.Params, &CDecl{
			Name: paramName(i, p),
			Spec: t.typeSpec(p.Type, deep+1, false),
			Pos:  p.Declarator.Pos(),
		})
	}
	return spec
}

func (t *Translator) typeSpec(typ cc.Type, deep int, isRet bool) CType {
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
			spec.AddOuterArr(uint64(size))
		}
	}
	for typ.Kind() == cc.Ptr {
		typ = typ.Element()
		spec.Pointers++
	}
	for typ.Kind() == cc.Array {
		size := typ.Elements()
		typ = typ.Element()
		if size >= 0 {
			spec.AddInnerArr(uint64(size))
		}
	}

	switch typ.Kind() {
	case cc.Void:
		spec.Base = "void"
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
		s := t.enumSpec(spec, typ)
		if !isRet {
			s.Typedef = typedefNameOf(typ)
		}
		return s
	case cc.Union:
		return &CStructSpec{
			Tag:      t.getStructTag(typ),
			IsUnion:  true,
			Pointers: spec.Pointers,
			OuterArr: spec.OuterArr,
			InnerArr: spec.InnerArr,
			Typedef:  typedefNameOf(typ),
		}
	case cc.Struct:
		s := t.structSpec(spec, typ, deep+1)
		if !isRet {
			s.Typedef = typedefNameOf(typ)
		}
		return s
	case cc.Function:
		s := t.functionSpec(spec, typ, deep+1)
		if !isRet && !typ.Specifier().IsTypedef() {
			s.Typedef = typedefNameOf(typ)
			retTyp := typ.Result().RawDeclarator().Type
			retTypedef := typedefNameOf(retTyp)
			if s.Return != nil {
				s.Return.SetRaw(retTypedef)
			}
		}
		return s
	default:
		panic("unknown type " + typ.String())
	}

	return spec
}

func paramName(n int, p cc.Parameter) string {
	if p.Name == 0 {
		return fmt.Sprintf("arg%d", n)
	}
	return blessName(xc.Dict.S(p.Name))
}

func memberName(n int, m cc.Member) string {
	if m.Name == 0 {
		return fmt.Sprintf("field%d", n)
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
