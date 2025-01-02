package translator

import (
	"fmt"
	"strings"

	"modernc.org/cc/v4"
	"modernc.org/token"
)

func (t *Translator) walkTranslationUnit(unit *cc.TranslationUnit) {
	for unit != nil {
		t.walkExternalDeclaration(unit.ExternalDeclaration)
		unit = unit.TranslationUnit
	}
}

func (t *Translator) walkExternalDeclaration(d *cc.ExternalDeclaration) {
	if t.IsTokenIgnored(d.Position()) {
		return
	}
	switch d.Case {
	case cc.ExternalDeclarationFuncDef:
		var decl *CDecl
		if declr := d.FunctionDefinition.Declarator; declr != nil {
			decl = t.declarator(declr.Type(), identifierOf(declr.DirectDeclarator), declr.IsTypename(), declr.IsStatic(), declr.IsConst(), declr.Position())
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
	case cc.ExternalDeclarationDecl:
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
	if t.IsTokenIgnored(d.Position()) {
		return
	}
	list := d.InitDeclaratorList
	if list != nil {
		for list != nil {
			declr := list.InitDeclarator.Declarator
			decl := t.declarator(declr.Type(), identifierOf(declr.DirectDeclarator), declr.IsTypename(), declr.IsStatic(), declr.IsConst(), declr.Position())
			init := list.InitDeclarator.Initializer
			if init != nil && init.AssignmentExpression != nil {
				decl.Value = init.AssignmentExpression.Value
				decl.Expression = blessName(init.Token.SrcStr())
			}
			t.registerTagsOf(decl)
			declared = append(declared, decl)
			if init != nil {
				t.valueMap[decl.Name] = decl.Value
				t.exprMap[decl.Name] = decl.Expression
			}
			list = list.InitDeclaratorList
		}
	} else if declr := d.Declarator; declr != nil {
		decl := t.declarator(declr.Type(), identifierOf(declr.DirectDeclarator), declr.IsTypename(), declr.IsStatic(), declr.IsConst(), declr.Position())
		t.registerTagsOf(decl)
		declared = append(declared, decl)
	} else if dspec := d.DeclarationSpecifiers; dspec != nil {
		typeSpec := dspec.TypeSpecifier
		if typeSpec != nil {
			if structSpec := typeSpec.StructOrUnionSpecifier; structSpec != nil {
				decl := t.declarator(structSpec.Type(), "", structSpec.Type().Typedef() != nil, false, false, structSpec.Position())
				t.registerTagsOf(decl)
				declared = append(declared, decl)
			} else if enumSpec := typeSpec.EnumSpecifier; enumSpec != nil {
				decl := t.declarator(enumSpec.Type(), "", enumSpec.Type().Typedef() != nil, false, false, enumSpec.Position())
				t.registerTagsOf(decl)
				declared = append(declared, decl)
			}
		}
	}
	return
}

func (t *Translator) declarator(typ cc.Type, name string, isTypedef bool, isStatic bool, isConst bool, position token.Position) *CDecl {
	decl := &CDecl{
		Spec:      t.typeSpec(typ, name, 0, isConst, false),
		Name:      name,
		IsTypedef: isTypedef,
		IsStatic:  isStatic,
		Position:  position,
	}
	return decl
}

func (t *Translator) getStructTag(typ cc.Type) (tag string) {
	var tagToken cc.Token
	if structType, ok := typ.(*cc.StructType); ok {
		tagToken = structType.Tag()
	} else if unionType, ok := typ.(*cc.UnionType); ok {
		tagToken = unionType.Tag()
	}
	return blessName(tagToken.SrcStr())
}

func (t *Translator) enumSpec(base *CTypeSpec, typ cc.Type) *CEnumSpec {
	enumType := typ.(*cc.EnumType)
	enumTag := enumType.Tag()
	tag := blessName(enumTag.SrcStr())
	spec := &CEnumSpec{
		Tag:      tag,
		Pointers: base.Pointers,
		OuterArr: base.OuterArr,
		InnerArr: base.InnerArr,
	}
	if enumType, ok := typ.(*cc.EnumType); ok {
		list := enumType.Enumerators()
		for _, en := range list {
			name := blessName(en.Token.SrcStr())
			m := &CDecl{
				Name:     name,
				Position: en.Token.Position(),
			}
			switch {
			case t.constRules[ConstEnum] == ConstCGOAlias:
				m.Expression = fmt.Sprintf("C.%s", name)
			case t.constRules[ConstEnum] == ConstExpand:
				enTokens := cc.NodeTokens(en.ConstantExpression)
				srcParts := make([]string, 0, len(enTokens))
				exprParts := make([]string, 0, len(enTokens))
				valid := true

				// TODO: needsTypecast is never true here
				// TODO: some state machine
				needsTypecast := false
				typecastValue := false
				typecastValueParens := 0

				for _, token := range enTokens {
					src := token.SrcStr()
					srcParts = append(srcParts, src)
					switch token.Ch {
					case rune(cc.IDENTIFIER):
						exprParts = append(exprParts, string(t.TransformName(TargetConst, src, true)))
					default:
						// TODO: state machine
						const (
							lparen = rune(40)
							rparen = rune(41)
						)
						switch {
						case needsTypecast && token.Ch == rparen:
							typecastValue = true
							needsTypecast = false
							exprParts = append(exprParts, src+"(")
						case typecastValue && token.Ch == lparen:
							typecastValueParens++
						case typecastValue && token.Ch == rparen:
							if typecastValueParens == 0 {
								typecastValue = false
								exprParts = append(exprParts, ")"+src)
							} else {
								typecastValueParens--
							}
						default:
							// somewhere in the world a helpless kitten died because of this
							if token.Ch == '~' {
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
					m.Value = en.Value()
				}
				m.Src = strings.Join(srcParts, " ")
			default:
				m.Value = en.Value()
			}
			m.Spec = spec.PromoteType(en.Value())
			spec.Members = append(spec.Members, m)
			t.valueMap[m.Name] = en.Value()
			t.exprMap[m.Name] = m.Expression
		}
	}

	return spec
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
	if structType, ok := typ.(*cc.StructType); ok {
		for i := 0; i < structType.NumFields(); i++ {
			m := structType.FieldByIndex(i)
			var position token.Position
			var name string
			var isConst bool
			if m.Declarator() != nil {
				position = m.Declarator().Position()
				name = identifierOf(m.Declarator().DirectDeclarator)
				isConst = m.Declarator().IsConst()
			}
			spec.Members = append(spec.Members, &CDecl{
				Name:     memberName(i, m),
				Spec:     t.typeSpec(m.Type(), name, deep+1, isConst, false),
				Position: position,
			})
		}
	}
	return spec
}

func (t *Translator) functionSpec(base *CTypeSpec, typ cc.Type, name string, isConst bool, deep int) *CFunctionSpec {
	spec := &CFunctionSpec{
		Pointers: base.Pointers,
	}
	if deep > 2 { // a function inside params of another function
		// Replicate cc v1 behaviour.
		spec.Raw = base.Raw
		spec.Typedef = base.Raw
	} else {
		spec.Raw = name
	}
	if deep > maxDeepLevel {
		return spec
	}
	if funcType, ok := typ.(*cc.FunctionType); ok {
		if ret := funcType.Result(); ret != nil && ret.Kind() != cc.Void {
			// function result type cannot be declarator of a function definition
			// so we use typedef here
			var name string
			var isResultConst = isConst
			if funcType.Result().Typedef() != nil {
				name = identifierOf(funcType.Result().Typedef().DirectDeclarator)
				isResultConst = isResultConst || funcType.Result().Typedef().IsConst()
			}
			spec.Return = t.typeSpec(ret, name, deep+1, isResultConst, true)
		}
		params := funcType.Parameters()
		for i, p := range params {
			if p.Type().Kind() != cc.Void { // cc v4 may return a "void" parameter, which means no parameters.
				var name string
				var isParamConst bool
				if p.Declarator != nil {
					name = identifierOf(p.Declarator.DirectDeclarator)
					isParamConst = p.Declarator.IsConst()
				}
				spec.Params = append(spec.Params, &CDecl{
					Name:     paramName(i, p),
					Spec:     t.typeSpec(p.Type(), name, deep+1, isParamConst, false),
					Position: p.Declarator.Position(),
				})
			}
		}
	}
	return spec
}

func (t *Translator) typeSpec(typ cc.Type, name string, deep int, isConst bool, isRet bool) CType {
	spec := &CTypeSpec{
		Const: isConst,
	}
	if !isRet {
		spec.Raw = typedefNameOf(typ)
	}

	for outerArrayType, ok := typ.(*cc.ArrayType); ok; outerArrayType, ok = typ.(*cc.ArrayType) {
		size := outerArrayType.Len()
		typ = outerArrayType.Elem()
		if size >= 0 {
			spec.AddOuterArr(uint64(size))
		}
	}
	var isVoidPtr bool
	for pointerType, ok := typ.(*cc.PointerType); ok; pointerType, ok = typ.(*cc.PointerType) {
		if next := pointerType.Elem(); next.Kind() == cc.Void {
			isVoidPtr = true
			spec.Base = "void*"
			break
		}
		typ = pointerType.Elem()
		spec.Pointers++
	}
	for innerArrayType, ok := typ.(*cc.ArrayType); ok; innerArrayType, ok = typ.(*cc.ArrayType) {
		size := innerArrayType.Len()
		typ = innerArrayType.Elem()
		if size >= 0 {
			spec.AddInnerArr(uint64(size))
		}
	}
	if isVoidPtr {
		return spec
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
	case cc.Int128:
		spec.Base = "int"
		spec.Long = true
	case cc.Long:
		spec.Base = "int"
		spec.Long = true
	case cc.ULong:
		spec.Base = "int"
		spec.Long = true
		spec.Unsigned = true
	case cc.UInt128:
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
	case cc.Float, cc.Float32, cc.Float32x:
		spec.Base = "float"
	case cc.Double, cc.Float64, cc.Float64x:
		spec.Base = "double"
	case cc.LongDouble, cc.Float128, cc.Float128x:
		spec.Base = "double"
		spec.Long = true
	case cc.Bool:
		spec.Base = "_Bool"
		// according to C99, _Bool is unsigned, but the flag is not used here.
	case cc.ComplexFloat:
		spec.Base = "complexfloat"
		spec.Complex = true
	case cc.ComplexDouble:
		spec.Base = "complexdouble"
		spec.Complex = true
	case cc.ComplexLongDouble:
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
		s := t.functionSpec(spec, typ, name, isConst, deep+1)
		if !isRet && typ.Typedef() == nil {
			if funcType, ok := typ.(*cc.FunctionType); ok {
				// According to comparison with cc v1 behavior, typedef is set to return value typedef
				// if this is not a function typedef.
				// This is somehow strange, and probably causes some bugs, but I'm trying to replicate previous behaviour here.
				retTypedef := typedefNameOf(funcType.Result())
				if s.Typedef == "" {
					s.Typedef = retTypedef
				}
				if s.Return != nil {
					s.Return.SetRaw(retTypedef)
				}
			}
		}
		return s
	default:
		panic("unknown type " + typ.String())
	}

	return spec
}

func paramName(n int, p *cc.Parameter) string {
	if len(p.Name()) == 0 {
		return fmt.Sprintf("arg%d", n)
	}
	return blessName(p.Name())
}

func memberName(n int, f *cc.Field) string {
	if len(f.Name()) == 0 {
		return fmt.Sprintf("field%d", n)
	}
	return blessName(f.Name())
}

func typedefNameOf(typ cc.Type) string {
	var name string
	typeDef := typ.Typedef()
	if typeDef != nil {
		name = blessName(typeDef.Name())

		if len(name) == 0 {
			name = identifierOf(typeDef.DirectDeclarator)
		}
	}

	if len(name) == 0 {
		if pointerType, ok := typ.(*cc.PointerType); ok {
			name = typedefNameOf(pointerType.Elem())
		} else if arrayType, ok := typ.(*cc.ArrayType); ok {
			name = typedefNameOf(arrayType.Elem())
		}
	}

	return name
}

func identifierOf(dd *cc.DirectDeclarator) string {
	if dd == nil {
		return ""
	}
	switch dd.Case {
	case cc.DirectDeclaratorIdent: // IDENTIFIER
		return blessName(dd.Token.SrcStr())
	case cc.DirectDeclaratorDecl: // '(' Declarator ')'
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
