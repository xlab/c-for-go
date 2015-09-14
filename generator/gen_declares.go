package generator

import (
	"bytes"
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func checkName(name []byte) []byte {
	if len(name) > 0 {
		return name
	}
	return skipName
}

func (gen *Generator) writeTypeDeclaration(wr io.Writer, decl tl.CDecl, ptrTip tl.Tip, public bool) {
	declName := checkName(gen.tr.TransformName(tl.TargetType, decl.Name, public))
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip)
	fmt.Fprintf(wr, "%s %s", declName, goSpec)
}

func (gen *Generator) writeArgType(wr io.Writer, decl tl.CDecl, ptrTip tl.Tip, public bool) {
	declName := checkName(gen.tr.TransformName(tl.TargetType, decl.Name, public))
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip)
	if len(goSpec.Arrays) > 0 {
		fmt.Fprintf(wr, "%s *%s", declName, goSpec)
	} else {
		fmt.Fprintf(wr, "%s %s", declName, goSpec)
	}
}

func (gen *Generator) writeEnumDeclaration(wr io.Writer, decl tl.CDecl, ptrTip tl.Tip, public bool) {
	declName := checkName(gen.tr.TransformName(tl.TargetType, decl.Name, public))
	typeRef := gen.tr.TranslateSpec(decl.Spec, ptrTip).Bytes()
	if !bytes.Equal(declName, typeRef) {
		fmt.Fprintf(wr, "%s %s", declName, typeRef)
		writeSpace(wr, 1)
	}
}

func (gen *Generator) writeFunctionAsArg(wr io.Writer, decl tl.CDecl, ptrTip tl.Tip, public bool) {
	var returnRef string
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		returnRef = gen.tr.TranslateSpec(spec.Return.Spec, ptrTip).String()
	}
	declName := checkName(gen.tr.TransformName(tl.TargetFunction, decl.Name, public))
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip)
	fmt.Fprintf(wr, "%s %s", declName, goSpec)
	gen.writeFunctionParams(wr, decl.Name, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
}

func (gen *Generator) writeFunctionDeclaration(wr io.Writer, decl tl.CDecl, ptrTip tl.Tip, public bool) {
	var returnRef []byte
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		returnRef = gen.tr.TranslateSpec(spec.Return.Spec, ptrTip).Bytes()
	}
	declName := checkName(gen.tr.TransformName(tl.TargetFunction, decl.Name, public))
	if bytes.Equal(returnRef, declName) {
		declName = gen.tr.TransformName(tl.TargetFunction, "new_"+decl.Name, public)
	}
	fmt.Fprintf(wr, "// %s function as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "func %s", declName)
	gen.writeFunctionParams(wr, decl.Name, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
	gen.writeFunctionBody(wr, decl)
	writeSpace(wr, 1)
}

func (gen *Generator) writeArgStruct(wr io.Writer, decl tl.CDecl, ptrTip tl.Tip, public bool) {
	declName := checkName(gen.tr.TransformName(tl.TargetType, decl.Name, public))
	if tag := decl.Spec.GetBase(); len(tag) > 0 {
		goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip)
		fmt.Fprintf(wr, "%s %s", declName, goSpec)
		return
	}
	if !decl.IsTemplate() {
		return
	}

	fmt.Fprintf(wr, "%s struct {", declName)
	writeSpace(wr, 1)
	gen.writeStructMembers(wr, decl.Name, decl.Spec)
	writeEndStruct(wr)
}

func (gen *Generator) writeArgUnion(wr io.Writer, decl tl.CDecl, ptrTip tl.Tip, public bool) {
	declName := checkName(gen.tr.TransformName(tl.TargetType, decl.Name, public))
	cgoSpec := gen.tr.CGoSpec(decl.Spec)
	fmt.Fprintf(wr, "%s [unsafe.Sizeof(%s)]byte", declName, cgoSpec.Base)
	return
}

func (gen *Generator) writeStructDeclaration(wr io.Writer, decl tl.CDecl, ptrTip tl.Tip, public bool) {
	declName := checkName(gen.tr.TransformName(tl.TargetType, decl.Name, public))
	if tag := decl.Spec.GetBase(); len(tag) > 0 {
		goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip)
		fmt.Fprintf(wr, "var %s %s", declName, goSpec)
		return
	}
	if !decl.IsTemplate() {
		return
	}

	fmt.Fprintf(wr, "var %s struct {", declName)
	writeSpace(wr, 1)
	gen.writeStructMembers(wr, decl.Name, decl.Spec)
	writeEndStruct(wr)
	writeSpace(wr, 1)
}

func (gen *Generator) writeUnionDeclaration(wr io.Writer, decl tl.CDecl, ptrTip tl.Tip, public bool) {
	declName := checkName(gen.tr.TransformName(tl.TargetType, decl.Name, public))
	cgoSpec := gen.tr.CGoSpec(decl.Spec)
	fmt.Fprintf(wr, "var %s [unsafe.Sizeof(%s)]byte", declName, cgoSpec.Base)
	return
}
