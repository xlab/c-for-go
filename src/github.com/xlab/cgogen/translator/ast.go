package translator

import "github.com/cznic/c/internal/cc"

func (t *Translator) walkAST(unit *cc.TranslationUnit) ([]CTypeDecl, error) {
	var declarations []CTypeDecl

	next, cdecls, err := walkUnit(unit)
	if err != nil {
		return nil, err
	}
	declarations = append(declarations, cdecls...)

	for next != nil {
		next, cdecls, err = walkUnit(next)
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, cdecls...)
	}
	return declarations, nil
}

func walkUnit(unit *cc.TranslationUnit) (next *cc.TranslationUnit, declarations []CTypeDecl, err error) {
	if unit == nil {
		return
	}
	if unit.ExternalDeclaration != nil {
		declarations, err = walkDeclaration(unit.ExternalDeclaration.Declaration)
	}
	next = unit.TranslationUnit
	return
}

func walkDeclaration(declr *cc.Declaration) ([]CTypeDecl, error) {
	// read type spec into a reference type declaration
	refDecl := &CTypeDecl{Spec: &CTypeSpec{}}
	nextSpec := walkDeclarationSpec(declr.DeclarationSpecifiers, refDecl)
	for nextSpec != nil {
		nextSpec = walkDeclarationSpec(nextSpec, refDecl)
	}

	// prepare to collect declarations
	var declarations []CTypeDecl

	decl := &CTypeDecl{Spec: refDecl.Spec.Copy()}
	if declr.InitDeclaratorListOpt != nil {
		nextList := walkInitDeclaratorList(declr.InitDeclaratorListOpt.InitDeclaratorList, decl)
		declarations = append(declarations, *decl)

		for nextList != nil {
			decl = &CTypeDecl{Spec: refDecl.Spec.Copy()}
			nextList = walkInitDeclaratorList(nextList, decl)
			declarations = append(declarations, *decl)
		}
	}

	return declarations, nil
}

func walkInitDeclaratorList(declr *cc.InitDeclaratorList, decl *CTypeDecl) (next *cc.InitDeclaratorList) {
	next = declr.InitDeclaratorList
	walkPointers(declr.InitDeclarator.Declarator.PointerOpt, decl)

	nextDeclarator := walkDirectDeclarator(declr.InitDeclarator.Declarator.DirectDeclarator, decl)
	for nextDeclarator != nil {
		nextDeclarator = walkDirectDeclarator(nextDeclarator, decl)
	}
	return
}

func walkParameterList(list *cc.ParameterList) (next *cc.ParameterList, decl *CTypeDecl) {
	next = list.ParameterList
	declr := list.ParameterDeclaration
	switch declr.Case {
	case 0: // DeclarationSpecifiers Declarator
		decl = &CTypeDecl{Spec: &CTypeSpec{}}
		nextDeclr := walkDeclarationSpec(declr.DeclarationSpecifiers, decl)
		for nextDeclr != nil {
			nextDeclr = walkDeclarationSpec(nextDeclr, decl)
		}

		walkPointers(declr.Declarator.PointerOpt, decl)
		nextDeclarator := walkDirectDeclarator(declr.Declarator.DirectDeclarator, decl)
		for nextDeclarator != nil {
			nextDeclarator = walkDirectDeclarator(nextDeclarator, decl)
		}
	case 1: // DeclarationSpecifiers AbstractDeclaratorOpt
		unmanagedCaseWarn(declr.Case, list.Token.Pos())
	}
	return
}

func walkDirectDeclarator2(declr *cc.DirectDeclarator2, decl *CTypeDecl) {
	switch declr.Case {
	case 0: // ParameterTypeList ')'
		spec := decl.Spec.(*CFunctionSpec)
		nextList, paramDecl := walkParameterList(declr.ParameterTypeList.ParameterList)
		if paramDecl != nil {
			spec.ParamList = append(spec.ParamList, *paramDecl)
		}
		for nextList != nil {
			nextList, paramDecl = walkParameterList(nextList)
			if paramDecl != nil {
				spec.ParamList = append(spec.ParamList, *paramDecl)
			}
		}
	case 1: // IdentifierListOpt ')'
		unmanagedCaseWarn(declr.Case, declr.Token.Pos())
	}
}

func walkDirectDeclarator(declr *cc.DirectDeclarator, decl *CTypeDecl) (next *cc.DirectDeclarator) {
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
		var n int64
		if declr.AssignmentExpressionOpt != nil {
			primary := walkAssigmentExperessionToPrimary(declr.AssignmentExpressionOpt.AssignmentExpression)
			n = walkPrimaryExpressionToInt64(primary)
			if n < 0 {
				n = 0
			}
		}
		decl.AddArray(uint64(n))
		next = declr.DirectDeclarator
	case 6: // DirectDeclarator '(' DirectDeclarator2
		next = declr.DirectDeclarator
		decl.Spec = &CFunctionSpec{
			Returns: CTypeDecl{
				Spec: decl.Spec,
				Pos:  decl.Pos,
			},
		}
		walkDirectDeclarator2(declr.DirectDeclarator2, decl)
	}
	return
}

func walkPointers(popt *cc.PointerOpt, decl *CTypeDecl) {
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

func walkDeclarationSpec(declr *cc.DeclarationSpecifiers, decl *CTypeDecl) (next *cc.DeclarationSpecifiers) {
	switch declr.Case {
	case 0: // StorageClassSpecifier DeclarationSpecifiersOpt
		next = declr.DeclarationSpecifiersOpt.DeclarationSpecifiers
	case 1: // TypeSpecifier DeclarationSpecifiersOpt
		walkTypeSpec(declr.TypeSpecifier, decl)
		if declr.DeclarationSpecifiersOpt != nil {
			next = declr.DeclarationSpecifiersOpt.DeclarationSpecifiers
		}
	case 2: // TypeQualifier DeclarationSpecifiersOpt
		if spec, ok := decl.Spec.(*CTypeSpec); ok {
			spec.Const = (declr.TypeQualifier.Case == 0)
		}
		next = declr.DeclarationSpecifiersOpt.DeclarationSpecifiers
	case 3: // FunctionSpecifier DeclarationSpecifiersOpt
		unmanagedCaseWarn(declr.Case, declr.FunctionSpecifier.Token.Pos())
	}
	return
}

func walkSpecifierQualifierList(declr *cc.SpecifierQualifierList, decl *CTypeDecl) (next *cc.SpecifierQualifierList) {
	if declr.SpecifierQualifierListOpt != nil {
		next = declr.SpecifierQualifierListOpt.SpecifierQualifierList
	}
	switch declr.Case {
	case 0:
		walkTypeSpec(declr.TypeSpecifier, decl)
	case 1:
		if spec, ok := decl.Spec.(*CTypeSpec); ok {
			spec.Const = (declr.TypeQualifier.Case == 0)
		}
	}
	return
}

func walkStructDeclarator(declr *cc.StructDeclarator, decl *CTypeDecl) {
	switch declr.Case {
	case 0: // Declarator
		walkPointers(declr.Declarator.PointerOpt, decl)
		nextDeclr := declr.Declarator.DirectDeclarator
		for nextDeclr != nil {
			nextDeclr = walkDirectDeclarator(nextDeclr, decl)
		}
	case 1: // DeclaratorOpt ':' ConstantExpression
		unmanagedCaseWarn(declr.Case, declr.Token.Pos())
	}
}

func walkStructDeclaration(declr *cc.StructDeclaration) []CTypeDecl {
	refDecl := &CTypeDecl{Spec: &CTypeSpec{}}
	nextList := declr.SpecifierQualifierList
	for nextList != nil {
		nextList = walkSpecifierQualifierList(nextList, refDecl)
	}

	// prepare to collect declarations
	var declarations []CTypeDecl

	nextDeclaratorList := declr.StructDeclaratorListOpt.StructDeclaratorList
	for nextDeclaratorList != nil {
		decl := &CTypeDecl{Spec: refDecl.Spec.Copy()}
		walkStructDeclarator(nextDeclaratorList.StructDeclarator, decl)
		nextDeclaratorList = nextDeclaratorList.StructDeclaratorList
		declarations = append(declarations, *decl)
	}

	return declarations
}

func walkSUSpecifier(suSpec *cc.StructOrUnionSpecifier, decl *CTypeDecl) {
	switch suSpec.Case {
	case 0: // StructOrUnionSpecifier0 '{' StructDeclarationList '}'
		walkSUSpecifier0(suSpec.StructOrUnionSpecifier0, decl)
		structSpec := decl.Spec.(*CStructSpec)
		nextList := suSpec.StructDeclarationList
		for nextList != nil {
			declarations := walkStructDeclaration(nextList.StructDeclaration)
			structSpec.Members = append(structSpec.Members, declarations...)
			nextList = nextList.StructDeclarationList
		}
	case 1: // StructOrUnion IDENTIFIER
		switch suSpec.StructOrUnion.Case {
		case 0: // struct
			decl.Spec = &CStructSpec{}
		case 1: // union
			decl.Spec = &CStructSpec{
				Union: true,
			}
		}
	}
}

func walkSUSpecifier0(suSpec *cc.StructOrUnionSpecifier0, decl *CTypeDecl) {
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

func walkTypeSpec(typeSpec *cc.TypeSpecifier, decl *CTypeDecl) {
	if typeSpec == nil {
		return
	}

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
		walkSUSpecifier(typeSpec.StructOrUnionSpecifier, decl)
	case 12: // TODO: enums
	case 13:
		spec.Base = string(typeSpec.Token.S())
	}
}
