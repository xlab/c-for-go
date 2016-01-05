package generator

import (
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

func (gen *Generator) writeTypeDeclaration(wr io.Writer, decl *tl.CDecl, ptrTip tl.Tip, public bool) {
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip)
	fmt.Fprintf(wr, "%s %s", goName, goSpec)
}

func (gen *Generator) writeArgType(wr io.Writer, decl *tl.CDecl, ptrTip tl.Tip, public bool) {
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip)
	if len(goSpec.Arrays) > 0 {
		fmt.Fprintf(wr, "%s *%s", goName, goSpec)
	} else {
		fmt.Fprintf(wr, "%s %s", goName, goSpec)
	}
}

func (gen *Generator) writeEnumDeclaration(wr io.Writer, decl *tl.CDecl, ptrTip tl.Tip, public bool) {
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	typeRef := gen.tr.TranslateSpec(decl.Spec, ptrTip).String()
	fmt.Fprintf(wr, "%s %s", goName, typeRef)
	writeSpace(wr, 1)
}

func (gen *Generator) writeFunctionAsArg(wr io.Writer, decl *tl.CDecl, ptrTip tl.Tip, public bool) {
	var returnRef string
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetFunction, cName, public))
	spec := decl.Spec.(*tl.CFunctionSpec)
	if len(spec.Typedef) > 0 {
		funcName := checkName(gen.tr.TransformName(tl.TargetType, spec.Typedef, true))
		fmt.Fprintf(wr, "%s %s", goName, funcName)
		return
	}
	if spec.Return != nil {
		returnRef = gen.tr.TranslateSpec(spec.Return, ptrTip).String()
	}
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip)
	fmt.Fprintf(wr, "%s %s", goName, goSpec)
	gen.writeFunctionParams(wr, cName, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
}

func (gen *Generator) writeFunctionDeclaration(wr io.Writer, decl *tl.CDecl, ptrTip tl.Tip, public bool) {
	var returnRef string
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		returnRef = gen.tr.TranslateSpec(spec.Return, ptrTip).String()
	}
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetFunction, cName, public))
	if returnRef == string(goName) {
		goName = gen.tr.TransformName(tl.TargetFunction, "new_"+cName, public)
	}
	fmt.Fprintf(wr, "// %s function as declared in %s\n", goName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "func %s", goName)
	gen.writeFunctionParams(wr, cName, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
	gen.writeFunctionBody(wr, decl)
	writeSpace(wr, 1)
}

func (gen *Generator) writeArgStruct(wr io.Writer, decl *tl.CDecl, ptrTip tl.Tip, public bool) {
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	if tag := decl.Spec.GetBase(); len(tag) > 0 {
		goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip)
		fmt.Fprintf(wr, "%s %s", goName, goSpec)
		return
	}
	if !decl.Spec.IsComplete() {
		return
	}

	fmt.Fprintf(wr, "%s struct {", goName)
	writeSpace(wr, 1)
	gen.submitHelper(cgoAllocMap)
	gen.writeStructMembers(wr, cName, decl.Spec)
	writeEndStruct(wr)
}

func (gen *Generator) writeArgUnion(wr io.Writer, decl *tl.CDecl, ptrTip tl.Tip, public bool) {
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	cgoSpec := gen.tr.CGoSpec(decl.Spec)
	fmt.Fprintf(wr, "%s [unsafe.Sizeof(%s)]byte", goName, cgoSpec.Base)
	return
}

func (gen *Generator) writeStructDeclaration(wr io.Writer, decl *tl.CDecl, ptrTip tl.Tip, public bool) {
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	if tag := decl.Spec.GetBase(); len(tag) > 0 {
		goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip)
		fmt.Fprintf(wr, "var %s %s", goName, goSpec)
		return
	}
	if !decl.Spec.IsComplete() {
		return
	}

	fmt.Fprintf(wr, "var %s struct {", goName)
	writeSpace(wr, 1)
	gen.submitHelper(cgoAllocMap)
	gen.writeStructMembers(wr, cName, decl.Spec)
	writeEndStruct(wr)
	writeSpace(wr, 1)
}

func (gen *Generator) writeUnionDeclaration(wr io.Writer, decl *tl.CDecl, ptrTip tl.Tip, public bool) {
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	cgoSpec := gen.tr.CGoSpec(decl.Spec)
	fmt.Fprintf(wr, "var %s [unsafe.Sizeof(%s)]byte", goName, cgoSpec.Base)
	return
}
