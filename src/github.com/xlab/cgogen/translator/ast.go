package translator

import "github.com/cznic/c/internal/cc"

func (t *Translator) walkAST(unit *cc.TranslationUnit) ([]CTypeDecl, error) {
	var declarations []CTypeDecl

	next, decl, err := walkUnit(unit)
	if err != nil {
		return nil, err
	} else if decl != nil {
		declarations = append(declarations, *decl)
	}

	for next != nil {
		next, decl, err = walkUnit(next)
		if err != nil {
			return nil, err
		} else if decl != nil {
			declarations = append(declarations, *decl)
		}
	}
	return declarations, nil
}

func walkUnit(unit *cc.TranslationUnit) (next *cc.TranslationUnit, decl *CTypeDecl, err error) {
	if unit == nil {
		return
	}
	if unit.ExternalDeclaration != nil {
		decl, err = walkDeclaration(unit.ExternalDeclaration.Declaration)
	}
	next = unit.TranslationUnit
	return
}

func walkDeclaration(decl *cc.Declaration) (*CTypeDecl, error) {
	if decl == nil {
		return nil, nil
	}

	spec := &CTypeSpec{}
	next := collectDeclarationSpec(decl.DeclarationSpecifiers, spec)
	for next != nil {
		next = collectDeclarationSpec(next, spec)
	}

	ctdecl := &CTypeDecl{
		Pos:  decl.Token.Pos(),
		Spec: spec,
	}

	// TODO
	return ctdecl, nil
}

func collectDeclarationSpec(declSpec *cc.DeclarationSpecifiers, spec *CTypeSpec) (next *cc.DeclarationSpecifiers) {
	if declSpec == nil {
		return nil
	}
	switch declSpec.Case {
	case 0: // StorageClassSpecifier DeclarationSpecifiersOpt
		next = declSpec.DeclarationSpecifiersOpt.DeclarationSpecifiers
	case 2: // TypeQualifier DeclarationSpecifiersOpt
		spec.Const = (declSpec.TypeQualifier.Case == 0)
		next = declSpec.DeclarationSpecifiersOpt.DeclarationSpecifiers
	case 1: // TypeSpecifier DeclarationSpecifiersOpt
		collectTypeSpec(declSpec.TypeSpecifier, spec)
		if declSpec.DeclarationSpecifiersOpt != nil {
			next = declSpec.DeclarationSpecifiersOpt.DeclarationSpecifiers
		}
	}
	return
}

func collectTypeSpec(typeSpec *cc.TypeSpecifier, spec *CTypeSpec) {
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
	// case 7:
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
