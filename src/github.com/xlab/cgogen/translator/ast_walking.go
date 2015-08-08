package translator

import (
	"fmt"

	"github.com/cznic/c/internal/cc"
)

func (t *Translator) walkTranslationUnit(unit *cc.TranslationUnit) (next *cc.TranslationUnit) {
	t.walkExternalDeclaration(unit.ExternalDeclaration)
	return unit.TranslationUnit
}

func (t *Translator) walkExternalDeclaration(declr *cc.ExternalDeclaration) {
	switch declr.Case {
	case 0: // FunctionDefinition
		// (not an API definition)
	case 1: // Declaration
		declares := t.walkDeclaration(declr.Declaration)
		for _, decl := range declares {
			if decl.IsTypedef {
				t.typedefs = append(t.typedefs, decl)
				continue
			}
			t.declares = append(t.declares, decl)
		}
	}
}

func (t *Translator) walkDeclaration(declr *cc.Declaration) []CDecl {
	// read type spec into a reference type declaration
	refDecl := CDecl{Spec: &CTypeSpec{}}
	nextSpec := declr.DeclarationSpecifiers
	for nextSpec != nil {
		nextSpec = t.walkDeclarationSpecifiers(nextSpec, &refDecl)
	}

	switch refDecl.Spec.Kind() {
	case EnumKind:
		spec := refDecl.Spec.(*CEnumSpec)
		if len(spec.Tag) != 0 {
			t.tagMap[spec.Tag] = refDecl
		}
	case StructKind:
		spec := refDecl.Spec.(*CStructSpec)
		if len(spec.Tag) != 0 {
			t.tagMap[spec.Tag] = refDecl
		}
	}

	// prepare to collect declares
	var declares []CDecl

	if declr.InitDeclaratorListOpt != nil {
		nextList := declr.InitDeclaratorListOpt.InitDeclaratorList
		for nextList != nil {
			decl := CDecl{Spec: refDecl.Spec.Copy(), IsTypedef: refDecl.IsTypedef}
			nextList = t.walkInitDeclaratorList(nextList, &decl)
			declares = append(declares, decl)
			t.valueMap[decl.Name] = decl.Value
			t.exprMap[decl.Name] = decl.Expression
		}
	}
	return declares
}

func (t *Translator) walkInitializer(init *cc.Initializer, decl *CDecl) {
	switch init.Case {
	case 0:
		decl.Value = t.EvalAssignmentExpression(init.AssignmentExpression)
		decl.Expression = t.ExpandAssignmentExpression(init.AssignmentExpression)
	case 1, // '{' InitializerList '}'
		2: // '{' InitializerList ',' '}'
		unmanagedCaseWarn(init.Case, init.Token.Pos())
	}
}

func (t *Translator) walkInitDeclaratorList(declr *cc.InitDeclaratorList, decl *CDecl) (next *cc.InitDeclaratorList) {
	next = declr.InitDeclaratorList
	walkPointers(declr.InitDeclarator.Declarator.PointerOpt, decl)

	switch declr.InitDeclarator.Case {
	case 1: // Declarator '=' Initializer
		t.walkInitializer(declr.InitDeclarator.Initializer, decl)
	}

	nextDeclarator := declr.InitDeclarator.Declarator.DirectDeclarator
	for nextDeclarator != nil {
		nextDeclarator = t.walkDirectDeclarator(nextDeclarator, decl)
	}
	return
}

func (t *Translator) walkParameterList(list *cc.ParameterList) (next *cc.ParameterList, decl *CDecl) {
	next = list.ParameterList
	declr := list.ParameterDeclaration
	switch declr.Case {
	case 0: // DeclarationSpecifiers Declarator
		decl = &CDecl{Spec: &CTypeSpec{}}
		nextDeclr := declr.DeclarationSpecifiers
		for nextDeclr != nil {
			nextDeclr = t.walkDeclarationSpecifiers(nextDeclr, decl)
		}

		walkPointers(declr.Declarator.PointerOpt, decl)
		nextDeclarator := declr.Declarator.DirectDeclarator
		for nextDeclarator != nil {
			nextDeclarator = t.walkDirectDeclarator(nextDeclarator, decl)
		}
	case 1: // DeclarationSpecifiers AbstractDeclaratorOpt
		decl = &CDecl{Spec: &CTypeSpec{}}
		nextDeclr := declr.DeclarationSpecifiers
		for nextDeclr != nil {
			nextDeclr = t.walkDeclarationSpecifiers(nextDeclr, decl)
		}
		if declr.AbstractDeclaratorOpt != nil {
			walkPointers(declr.AbstractDeclaratorOpt.AbstractDeclarator.PointerOpt, decl)
		}
	}
	return
}

func (t *Translator) walkDirectDeclarator2(declr *cc.DirectDeclarator2, decl *CDecl) {
	switch declr.Case {
	case 0: // ParameterTypeList ')'
		spec := decl.Spec.(*CFunctionSpec)
		nextList := declr.ParameterTypeList.ParameterList
		var paramDecl *CDecl
		for nextList != nil {
			if nextList, paramDecl = t.walkParameterList(nextList); paramDecl != nil {
				spec.ParamList = append(spec.ParamList, *paramDecl)
			}
		}
	case 1: // IdentifierListOpt ')'
		if declr.IdentifierListOpt != nil {
			unmanagedCaseWarn(declr.Case, declr.Token.Pos())
		}
	}
}

func (t *Translator) walkDirectDeclarator(declr *cc.DirectDeclarator, decl *CDecl) (next *cc.DirectDeclarator) {
	decl.Pos = declr.Token.Pos()
	switch declr.Case {
	case 0: // IDENTIFIER
		decl.Name = string(declr.Token.S())
	case 1: // '(' Declarator ')'
		walkPointers(declr.Declarator.PointerOpt, decl)
		next = declr.Declarator.DirectDeclarator
	case 2, // DirectDeclarator '[' TypeQualifierListOpt AssignmentExpressionOpt ']'
		3, // DirectDeclarator '[' "static" TypeQualifierListOpt AssignmentExpression ']'
		4, // DirectDeclarator '[' TypeQualifierList "static" AssignmentExpression ']'
		5: // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
		assignmentExpr := declr.AssignmentExpression
		if declr.AssignmentExpressionOpt != nil {
			assignmentExpr = declr.AssignmentExpressionOpt.AssignmentExpression
		}
		val := t.ExpandAssignmentExpression(assignmentExpr)
		decl.AddArray(val)
		next = declr.DirectDeclarator
	case 6: // DirectDeclarator '(' DirectDeclarator2
		next = declr.DirectDeclarator
		decl.Spec = &CFunctionSpec{
			Returns: CDecl{
				Spec: decl.Spec,
				Pos:  decl.Pos,
			},
		}
		t.walkDirectDeclarator2(declr.DirectDeclarator2, decl)
	}
	return
}

func walkPointers(popt *cc.PointerOpt, decl *CDecl) {
	if popt == nil {
		return
	}
	nextPointer := popt.Pointer.Pointer
	pointers := uint8(1)
	for nextPointer != nil {
		nextPointer = nextPointer.Pointer
		pointers++
	}
	decl.SetPointers(pointers)
}

func (t *Translator) walkDeclarationSpecifiers(declr *cc.DeclarationSpecifiers, decl *CDecl) (next *cc.DeclarationSpecifiers) {
	switch declr.Case {
	case 0: // StorageClassSpecifier DeclarationSpecifiersOpt
		switch declr.StorageClassSpecifier.Case {
		case 0: // "typedef"
			decl.IsTypedef = true
		}
		next = declr.DeclarationSpecifiersOpt.DeclarationSpecifiers
	case 1: // TypeSpecifier DeclarationSpecifiersOpt
		t.walkTypeSpec(declr.TypeSpecifier, decl)
		if declr.DeclarationSpecifiersOpt != nil {
			next = declr.DeclarationSpecifiersOpt.DeclarationSpecifiers
		}
	case 2: // TypeQualifier DeclarationSpecifiersOpt
		if spec, ok := decl.Spec.(*CTypeSpec); ok {
			spec.Const = (declr.TypeQualifier.Case == 0)
		}
		if declr.DeclarationSpecifiersOpt != nil {
			next = declr.DeclarationSpecifiersOpt.DeclarationSpecifiers
		}
	case 3: // FunctionSpecifier DeclarationSpecifiersOpt
		unmanagedCaseWarn(declr.Case, declr.FunctionSpecifier.Token.Pos())
	}
	return
}

func (t *Translator) walkSpecifierQualifierList(declr *cc.SpecifierQualifierList, decl *CDecl) (next *cc.SpecifierQualifierList) {
	if declr.SpecifierQualifierListOpt != nil {
		next = declr.SpecifierQualifierListOpt.SpecifierQualifierList
	}
	switch declr.Case {
	case 0:
		t.walkTypeSpec(declr.TypeSpecifier, decl)
	case 1:
		if spec, ok := decl.Spec.(*CTypeSpec); ok {
			spec.Const = (declr.TypeQualifier.Case == 0)
		}
	}
	return
}

func (t *Translator) walkStructDeclarator(declr *cc.StructDeclarator, decl *CDecl) {
	switch declr.Case {
	case 0: // Declarator
		walkPointers(declr.Declarator.PointerOpt, decl)
		nextDeclr := declr.Declarator.DirectDeclarator
		for nextDeclr != nil {
			nextDeclr = t.walkDirectDeclarator(nextDeclr, decl)
		}
	case 1: // DeclaratorOpt ':' ConstantExpression
		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		// C should not have bitfields. Don't use them. They are poorly
		// specified, non-portable, inefficient, unsafe, and have poor semantics.
		// They have no good properties to compensate.
		//
		// -Rob Pike
		// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
		if declr.DeclaratorOpt != nil {
			walkPointers(declr.DeclaratorOpt.Declarator.PointerOpt, decl)
			nextDeclr := declr.DeclaratorOpt.Declarator.DirectDeclarator
			for nextDeclr != nil {
				nextDeclr = t.walkDirectDeclarator(nextDeclr, decl)
			}
		}
	}
}

func (t *Translator) walkStructDeclaration(declr *cc.StructDeclaration) []CDecl {
	refDecl := &CDecl{Spec: &CTypeSpec{}}
	nextList := declr.SpecifierQualifierList
	for nextList != nil {
		nextList = t.walkSpecifierQualifierList(nextList, refDecl)
	}

	// prepare to collect declares
	var declares []CDecl

	if declr.StructDeclaratorListOpt != nil {
		nextDeclaratorList := declr.StructDeclaratorListOpt.StructDeclaratorList
		for nextDeclaratorList != nil {
			decl := CDecl{Spec: refDecl.Spec.Copy()}
			t.walkStructDeclarator(nextDeclaratorList.StructDeclarator, &decl)
			nextDeclaratorList = nextDeclaratorList.StructDeclaratorList
			declares = append(declares, decl)
		}
	}

	return declares
}

func (t *Translator) walkSUSpecifier(suSpec *cc.StructOrUnionSpecifier, decl *CDecl) {
	switch suSpec.Case {
	case 0: // StructOrUnionSpecifier0 '{' StructDeclarationList '}'
		walkSUSpecifier0(suSpec.StructOrUnionSpecifier0, decl)
		structSpec := decl.Spec.(*CStructSpec)
		nextList := suSpec.StructDeclarationList
		for nextList != nil {
			declares := t.walkStructDeclaration(nextList.StructDeclaration)
			structSpec.Members = append(structSpec.Members, declares...)
			nextList = nextList.StructDeclarationList
		}
	case 1: // StructOrUnion IDENTIFIER
		switch suSpec.StructOrUnion.Case {
		case 0: // struct
			decl.Spec = &CStructSpec{
				Tag: string(suSpec.Token.S()),
			}
		case 1: // union
			decl.Spec = &CStructSpec{
				Union: true,
				Tag:   string(suSpec.Token.S()),
			}
		}
	}
}

func walkSUSpecifier0(suSpec *cc.StructOrUnionSpecifier0, decl *CDecl) {
	switch suSpec.StructOrUnion.Case {
	case 0: // struct
		decl.Spec = &CStructSpec{}
	case 1: // union
		decl.Spec = &CStructSpec{
			Union: true,
		}
	}
	if suSpec.IdentifierOpt != nil {
		decl.Spec.(*CStructSpec).Tag = string(suSpec.IdentifierOpt.Token.S())
	}
}

func (t *Translator) walkEnumSpecifier(enSpec *cc.EnumSpecifier, decl *CDecl) {
	switch enSpec.Case {
	case 0, // EnumSpecifier0 '{' EnumeratorList '}'
		1: // EnumSpecifier0 '{' EnumeratorList ',' '}'
		decl.Spec = &CEnumSpec{}
		enumSpec := decl.Spec.(*CEnumSpec)
		if enSpec.EnumSpecifier0.IdentifierOpt != nil {
			enumSpec.Tag = string(enSpec.EnumSpecifier0.IdentifierOpt.Token.S())
		}

		enumIdx := len(enumSpec.Enumerators)
		nextList := enSpec.EnumeratorList
		for nextList != nil {
			enumDecl := t.walkEnumerator(nextList.Enumerator)
			if enumDecl.Value == nil {
				if enumIdx > 0 {
					// get previous enumerator from the list
					prevEnum := enumSpec.Enumerators[enumIdx-1]
					enumDecl.Value = incVal(prevEnum.Value)
					switch t.constRules[ConstEnum] {
					case ConstExpandFull:
						enumDecl.Expression = append(prevEnum.Expression, '+', '1')
					case ConstExpand:
						enumDecl.Expression = []byte(prevEnum.Name + "+1")
					default:
						enumDecl.Expression = []byte(fmt.Sprintf("%d", enumDecl.Value))
					}
					enumDecl.Spec = prevEnum.Spec.Copy()
				} else {
					enumDecl.Value = int32(0)
					enumDecl.Expression = []byte("0")
				}
			}
			enumIdx++
			enumDecl.Spec = enumSpec.PromoteType(enumDecl.Value)
			enumSpec.Enumerators = append(enumSpec.Enumerators, enumDecl)
			t.valueMap[enumDecl.Name] = enumDecl.Value
			t.exprMap[enumDecl.Name] = enumDecl.Expression
			nextList = nextList.EnumeratorList
		}
	case 2: // "enum" IDENTIFIER
		decl.Spec = &CEnumSpec{
			Tag: string(enSpec.Token.S()),
		}
	}
}

func (t *Translator) walkEnumerator(enSpec *cc.Enumerator) (decl CDecl) {
	switch enSpec.Case {
	case 0: // EnumerationConstant
		decl.Name = string(enSpec.EnumerationConstant.Token.S())
	case 1: // EnumerationConstant '=' ConstantExpression
		decl.Name = string(enSpec.EnumerationConstant.Token.S())
		decl.Value = t.EvalConditionalExpression(enSpec.ConstantExpression.ConditionalExpression)
		decl.Expression = t.ExpandConditionalExpression(enSpec.ConstantExpression.ConditionalExpression)
	}
	return
}

func (t *Translator) walkTypeSpec(typeSpec *cc.TypeSpecifier, decl *CDecl) {
	spec := decl.Spec.(*CTypeSpec)

	switch typeSpec.Case {
	case 0:
		spec.Base = "void"
	case 1:
		spec.Base = "char"
	case 2:
		spec.Short = true
	case 3:
		spec.Base = "int"
	case 4:
		if spec.Long {
			spec.Base = "long"
		} else {
			spec.Long = true
		}
	case 5:
		spec.Base = "float"
	case 6:
		spec.Base = "double"
	case 7: // IGNORE: signed
	case 8:
		spec.Unsigned = true
	case 9:
		spec.Base = "_Bool"
	case 10:
		spec.Base = "_Complex"
	case 11:
		t.walkSUSpecifier(typeSpec.StructOrUnionSpecifier, decl)
	case 12:
		t.walkEnumSpecifier(typeSpec.EnumSpecifier, decl)
	case 13:
		spec.Base = string(typeSpec.Token.S())
	}
}
