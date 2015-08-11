package generator

import (
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func (gen *Generator) writeTypeDeclaration(wr io.Writer, decl tl.CDecl, global bool) {
	var declName string
	if global {
		if name := gen.tr.TransformName(tl.TargetDeclare, decl.Name); len(name) > 0 {
			declName = string(name)
		} else {
			declName = "_"
		}
	} else if len(decl.Name) > 0 {
		declName = decl.Name
	} else {
		declName = "_"
	}

	fmt.Fprintf(wr, "%s %s", declName, gen.tr.TranslateSpec(tl.TargetTypedef, decl.Spec))
}

func (gen *Generator) writeFunctionDeclaration(wr io.Writer, decl tl.CDecl, global bool) {
	var returnRef string
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		returnRef = gen.tr.TranslateSpec(tl.TargetDeclare, spec.Return.Spec).String()
	}

	var declName string
	if global {
		declName = string(gen.tr.TransformName(tl.TargetDeclare, decl.Name))
		if returnRef == declName {
			declName = string(gen.tr.TransformName(tl.TargetDeclare, "new_"+decl.Name))
		}
		if len(declName) == 0 {
			declName = "_"
		}
	} else if len(decl.Name) > 0 {
		declName = decl.Name
	} else {
		declName = "_"
	}

	if global {
		fmt.Fprintf(wr, "func %s", declName)
	} else {
		fmt.Fprintf(wr, "%s %s", declName, gen.tr.TranslateSpec(tl.TargetDeclare, decl.Spec))
	}
	gen.writeFunctionParams(wr, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
}

func (gen *Generator) writeStructDeclaration(wr io.Writer, decl tl.CDecl, global bool) {
	var declName string
	if global {
		if name := gen.tr.TransformName(tl.TargetDeclare, decl.Name); len(name) > 0 {
			declName = string(name)
		} else {
			declName = "_"
		}
	} else if len(decl.Name) > 0 {
		declName = decl.Name
	} else {
		declName = "_"
	}

	if tag := decl.Spec.GetBase(); len(tag) > 0 {
		// refSpec := &tl.CTypeSpec{
		// 	Base:      spec.Tag,
		// 	Arrays:    spec.GetArrays(),
		// 	VarArrays: spec.GetVarArrays(),
		// 	Pointers:  spec.GetPointers(),
		// }
		fmt.Fprintf(wr, "%s %s", declName, gen.tr.TranslateSpec(tl.TargetTag, decl.Spec))
		return
	}
	if !decl.IsTemplate() {
		return
	}

	fmt.Fprintf(wr, "%s struct {", declName)
	writeSpace(wr, 1)
	gen.writeStructMembers(wr, decl.Spec)
	writeEndStruct(wr)
}

func (gen *Generator) writeEnumDeclaration(wr io.Writer, decl tl.CDecl, global bool) {
	var declName string
	if global {
		if name := gen.tr.TransformName(tl.TargetDeclare, decl.Name); len(name) > 0 {
			declName = string(name)
		} else {
			declName = "_"
		}
	} else if len(decl.Name) > 0 {
		declName = decl.Name
	} else {
		declName = "_"
	}

	typeRef := gen.tr.TranslateSpec(tl.TargetTag, decl.Spec).String()
	if declName != typeRef {
		fmt.Fprintf(wr, "%s %s", declName, typeRef)
		writeSpace(wr, 1)
	}
}
