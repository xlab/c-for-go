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
	if declr == nil {
		return nil, nil
	}

	// read type spec
	spec := &CTypeSpec{}
	nextSpec := walkDeclarationSpec(declr.DeclarationSpecifiers, spec)
	for nextSpec != nil {
		nextSpec = walkDeclarationSpec(nextSpec, spec)
	}

	// prepare to collect declarations
	var declarations []CTypeDecl

	cdecl := &CTypeDecl{Spec: spec.Copy()}
	if declr.InitDeclaratorListOpt != nil {
		nextList := walkInitDeclaratorList(declr.InitDeclaratorListOpt.InitDeclaratorList, cdecl)
		declarations = append(declarations, *cdecl)

		for nextList != nil {
			cdecl = &CTypeDecl{Spec: spec.Copy()}
			nextList = walkInitDeclaratorList(nextList, cdecl)
			declarations = append(declarations, *cdecl)
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
		spec := &CTypeSpec{}
		nextDeclr := walkDeclarationSpec(declr.DeclarationSpecifiers, spec)
		for nextDeclr != nil {
			nextDeclr = walkDeclarationSpec(nextDeclr, spec)
		}

		decl = &CTypeDecl{Spec: spec}
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
	case 6: // DirectDeclarator '(' DirectDeclarator2
		next = declr.DirectDeclarator
		decl.Spec = &CFunctionSpec{
			Returns: CTypeDecl{
				Spec: decl.Spec,
				Pos:  decl.Pos,
			},
		}
		walkDirectDeclarator2(declr.DirectDeclarator2, decl)
	default:
		unmanagedCaseWarn(declr.Case, declr.Token.Pos())
		//	|       '(' Declarator ')'                                                           // Case 1
		//	|       DirectDeclarator '[' TypeQualifierListOpt AssignmentExpressionOpt ']'        // Case 2
		//	|       DirectDeclarator '[' "static" TypeQualifierListOpt AssignmentExpression ']'  // Case 3
		//	|       DirectDeclarator '[' TypeQualifierList "static" AssignmentExpression ']'     // Case 4
		//	|       DirectDeclarator '[' TypeQualifierListOpt '*' ']'                            // Case 5
		//	|       DirectDeclarator '(' DirectDeclarator2                                       // Case 6
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

func walkDeclarationSpec(declr *cc.DeclarationSpecifiers, spec *CTypeSpec) (next *cc.DeclarationSpecifiers) {
	switch declr.Case {
	case 0: // StorageClassSpecifier DeclarationSpecifiersOpt
		next = declr.DeclarationSpecifiersOpt.DeclarationSpecifiers
	case 2: // TypeQualifier DeclarationSpecifiersOpt
		spec.Const = (declr.TypeQualifier.Case == 0)
		next = declr.DeclarationSpecifiersOpt.DeclarationSpecifiers
	case 1: // TypeSpecifier DeclarationSpecifiersOpt
		walkTypeSpec(declr.TypeSpecifier, spec)
		if declr.DeclarationSpecifiersOpt != nil {
			next = declr.DeclarationSpecifiersOpt.DeclarationSpecifiers
		}
	case 3:
		unmanagedCaseWarn(declr.Case, declr.FunctionSpecifier.Token.Pos())
	}
	return
}

func walkTypeSpec(typeSpec *cc.TypeSpecifier, spec *CTypeSpec) {
	if typeSpec == nil {
		return
	}
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
	// case 7: signed
	case 8:
		spec.Unsigned = true
	case 9:
		spec.Base = "_Bool"
	case 10:
		spec.Base = "_Complex"
	case 11:
		spec.Struct = true
	// case 12: TODO enums
	case 13:
		spec.Base = string(typeSpec.Token.S())
	}
}
