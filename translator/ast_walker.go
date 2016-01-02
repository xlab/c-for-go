package translator

import "github.com/xlab/c/cc"

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
			t.registerTagsOf(decl.Spec)
			declared = append(declared, decl)
			if init != nil {
				t.valueMap[decl.Name] = decl.Value
				t.exprMap[decl.Name] = decl.Expression
			}
			list = list.InitDeclaratorList
		}
	} else if declr := d.Declarator(); declr != nil {
		decl := t.declarator(declr)
		t.registerTagsOf(decl.Spec)
		declared = append(declared, decl)
	}
	return
}

func (t *Translator) declarator(d *cc.Declarator) *CDecl {
	s := d.RawSpecifier()
	dd := d.DirectDeclarator
	if dd.DirectDeclarator != nil {
		dd = dd.DirectDeclarator
	}
	decl := &CDecl{
		Spec:      t.typeSpec(d.Type),
		Name:      blessName(dd.Token.S()),
		IsTypedef: s.IsTypedef(),
		IsStatic:  s.IsStatic(),
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

func memberName(m cc.Member) string {
	dd := m.Declarator.DirectDeclarator
	if dd.DirectDeclarator != nil {
		dd = dd.DirectDeclarator
	}
	return blessName(dd.Token.S())
}

func (t *Translator) enumSpec(base *CTypeSpec, typ cc.Type) *CEnumSpec {
	tag := t.getEnumTag(typ)
	spec := &CEnumSpec{
		// TODO: Type, // (check out the augmented enums)
		Tag:       tag,
		Arrays:    base.Arrays,
		VarArrays: base.VarArrays,
		Pointers:  base.Pointers,
	}
	// TODO: fix enums lol
	members, _ := typ.Members()
	for _, m := range members {
		spec.Members = append(spec.Members, CEnumMember{
			Name: memberName(m),
			Type: t.typeSpec(m.Type),
		})
	}
	return spec
}

func (t *Translator) structSpec(base *CTypeSpec, typ cc.Type) *CStructSpec {
	tag := t.getStructTag(typ)
	spec := &CStructSpec{
		Tag:       tag,
		Arrays:    base.Arrays,
		VarArrays: base.VarArrays,
		Pointers:  base.Pointers,
	}
	members, _ := typ.Members()
	for _, m := range members {
		spec.Members = append(spec.Members, CStructMember{
			Name: memberName(m),
			Type: t.typeSpec(m.Type),
		})
	}
	return spec
}

func (t *Translator) unionSpec(base *CTypeSpec, typ cc.Type) *CStructSpec {
	tag := t.getStructTag(typ)
	spec := &CStructSpec{
		Tag:       tag,
		IsUnion:   true,
		Arrays:    base.Arrays,
		VarArrays: base.VarArrays,
		Pointers:  base.Pointers,
	}
	members, _ := typ.Members()
	for _, m := range members {
		spec.Members = append(spec.Members, CStructMember{
			Name: memberName(m),
			Type: t.typeSpec(m.Type),
		})
	}
	return spec
}

func paramName(p cc.Parameter) string {
	dd := p.Declarator.DirectDeclarator
	if dd.DirectDeclarator != nil {
		dd = dd.DirectDeclarator
	}
	return blessName(dd.Token.S())
}

func (t *Translator) functionSpec(base *CTypeSpec, typ cc.Type) *CFunctionSpec {
	spec := &CFunctionSpec{
		Arrays:    base.Arrays,
		VarArrays: base.VarArrays,
		Pointers:  base.Pointers,
	}
	if ret := typ.Result(); ret != nil && ret.Kind() != cc.Void {
		spec.Return = t.typeSpec(ret)
	}
	params, _ := typ.Parameters()
	for _, p := range params {
		spec.Params = append(spec.Params, CFunctionParam{
			Name: paramName(p),
			Type: t.typeSpec(p.Type),
		})
	}
	return spec
}

func (t *Translator) typeSpec(typ cc.Type) CType {
	id, _ := typ.Declarator().Identifier()
	cached, ok := t.typeCache.Get(id)
	if ok {
		return cached
	}

	spec := &CTypeSpec{}
	for typ.Kind() == cc.Ptr {
		typ = typ.Element()
		spec.Pointers++
	}
	for typ.Kind() == cc.Array {
		size := typ.Elements()
		typ = typ.Element()
		if size >= 0 {
			spec.AddArray(uint64(size))
			size = typ.Elements()
		}
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
		spec.Base = "bool"
	case cc.FloatComplex:
		spec.Base = "float"
		spec.Complex = true
	case cc.DoubleComplex:
		spec.Base = "double"
		spec.Complex = true
	case cc.LongDoubleComplex:
		spec.Base = "double"
		spec.Long = true
		spec.Complex = true
	case cc.Enum:
		s := t.enumSpec(spec, typ)
		t.typeCache.Set(id, s)
		return s
	case cc.Struct:
		// avoid recursive structs
		t.typeCache.Set(id, &CStructSpec{})
		s := t.structSpec(spec, typ)
		t.typeCache.Set(id, s)
		return s
	case cc.Union:
		// avoid recursive unions
		t.typeCache.Set(id, &CStructSpec{IsUnion: true})
		s := t.unionSpec(spec, typ)
		t.typeCache.Set(id, s)
		return s
	case cc.Function:
		s := t.functionSpec(spec, typ)
		t.typeCache.Set(id, s)
		return s
	case cc.TypedefName:
		// id := typ.Specifier().TypedefName()
		// spec.Base = blessName(xc.Dict.S(id))
		panic("typedef name? wtf?")
	default:
		panic("unknown type " + typ.String())
	}

	t.typeCache.Set(id, spec)
	return spec
}

// AUGMENTED ENUM SUPPORT BELOW

// func (t *Translator) walkEnumSpecifier(enSpec *cc.EnumSpecifier, decl *CDecl) {
// 	decl.Pos = enSpec.Token.Pos()
// 	switch enSpec.Case {
// 	case 0, // EnumSpecifier0 '{' EnumeratorList '}'
// 		1: // EnumSpecifier0 '{' EnumeratorList ',' '}'
// 		decl.Spec = &CEnumSpec{}
// 		enumSpec := decl.Spec.(*CEnumSpec)
// 		if enSpec.EnumeratorList.IdentifierOpt != nil {
// 			enumSpec.Tag = blessName(enSpec.EnumSpecifier0.IdentifierOpt.Token.S())
// 		}
// 		enumIdx := len(enumSpec.Enumerators)
// 		nextList := enSpec.EnumeratorList
// 		for nextList != nil {
// 			enumDecl := t.walkEnumerator(nextList.Enumerator)
// 			if enumDecl.Value == nil {
// 				if enumIdx > 0 {
// 					// get previous enumerator from the list
// 					prevEnum := enumSpec.Enumerators[enumIdx-1]
// 					enumDecl.Value = incVal(prevEnum.Value)
// 					switch t.constRules[ConstEnum] {
// 					case ConstExpandFull:
// 						enumDecl.Expression = append(prevEnum.Expression, '+', '1')
// 					case ConstExpand:
// 						enumDecl.Expression = []byte(prevEnum.Name + "+1")
// 					case ConstCGOAlias:
// 						enumDecl.Expression = []byte(fmt.Sprintf("C.%s", blessName([]byte(enumDecl.Name))))
// 					default:
// 						enumDecl.Expression = []byte(fmt.Sprintf("%d", enumDecl.Value))
// 					}
// 					enumDecl.Spec = prevEnum.Spec.Copy()
// 				} else {
// 					enumDecl.Value = int32(0)
// 					enumDecl.Expression = []byte("0")
// 				}
// 			} else if t.constRules[ConstEnum] == ConstCGOAlias {
// 				enumDecl.Expression = []byte(fmt.Sprintf("C.%s", blessName([]byte(enumDecl.Name))))
// 			}
// 			enumIdx++
// 			enumDecl.Spec = enumSpec.PromoteType(enumDecl.Value)
// 			enumSpec.Enumerators = append(enumSpec.Enumerators, enumDecl)
// 			t.valueMap[enumDecl.Name] = enumDecl.Value
// 			t.exprMap[enumDecl.Name] = enumDecl.Expression
// 			nextList = nextList.EnumeratorList
// 		}
// 	case 2: // "enum" IDENTIFIER
// 		decl.Spec = &CEnumSpec{
// 			Tag: blessName(enSpec.Token.S()),
// 		}
// 	}
// }
