package generator

import (
	"fmt"
	"io"
	"path/filepath"

	tl "github.com/xlab/c-for-go/translator"
)

func checkName(name []byte) []byte {
	if len(name) > 0 {
		return name
	}
	return skipName
}

func (gen *Generator) writeTypeDeclaration(wr io.Writer, decl *tl.CDecl,
	ptrTip, typeTip tl.Tip, public bool, seenNames map[string]bool) {

	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip, typeTip)
	fmt.Fprintf(wr, "%s %s", goName, goSpec)
}

func (gen *Generator) writeArgType(wr io.Writer, decl *tl.CDecl,
	ptrTip, typeTip tl.Tip, public bool) {

	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip, typeTip)
	if len(goSpec.OuterArr) > 0 {
		fmt.Fprintf(wr, "%s *%s", goName, goSpec)
	} else {
		fmt.Fprintf(wr, "%s %s", goName, goSpec)
	}
}

func (gen *Generator) writeEnumDeclaration(wr io.Writer, decl *tl.CDecl,
	ptrTip, typeTip tl.Tip, public bool) {
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	typeRef := gen.tr.TranslateSpec(decl.Spec, ptrTip, typeTip).String()
	fmt.Fprintf(wr, "%s %s", goName, typeRef)
	writeSpace(wr, 1)
}

func (gen *Generator) writeFunctionAsArg(wr io.Writer, decl *tl.CDecl,
	ptrTip, typeTip tl.Tip, public bool) {

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
		returnRef = gen.tr.TranslateSpec(spec.Return, ptrTip, typeTip).String()
	}
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip, typeTip)
	fmt.Fprintf(wr, "%s %s", goName, goSpec)
	gen.writeFunctionParams(wr, cName, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
}

func (gen *Generator) writeFunctionDeclaration(wr io.Writer, decl *tl.CDecl,
	ptrTip, typeTip tl.Tip, public bool) {

	var returnRef string
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		returnRef = gen.tr.TranslateSpec(spec.Return, ptrTip, typeTip).String()
	}
	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetFunction, cName, public))
	if returnRef == string(goName) {
		goName = gen.tr.TransformName(tl.TargetFunction, "new_"+cName, public)
	}
	fmt.Fprintf(wr, "// %s function as declared in %s\n", goName,
		filepath.ToSlash(gen.tr.SrcLocation(tl.TargetFunction, decl.Name, decl.Pos)))
	fmt.Fprintf(wr, "func")
	gen.writeInstanceObjectParam(wr, cName, decl.Spec)
	fmt.Fprintf(wr, " %s", goName)
	gen.writeFunctionParams(wr, cName, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
	gen.writeFunctionBody(wr, decl)
	writeSpace(wr, 1)
}

func (gen *Generator) writeArgStruct(wr io.Writer, decl *tl.CDecl,
	ptrTip, typeTip tl.Tip, public bool) {

	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	if tag := decl.Spec.GetBase(); len(tag) > 0 {
		goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip, typeTip)
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

func (gen *Generator) writeArgUnion(wr io.Writer, decl *tl.CDecl,
	ptrTip, typeTip tl.Tip, public bool) {

	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	if tag := decl.Spec.GetBase(); len(tag) > 0 {
		goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip, typeTip)
		fmt.Fprintf(wr, "%s %s", goName, goSpec)
	}
}

func (gen *Generator) writeStructDeclaration(wr io.Writer, decl *tl.CDecl,
	ptrTip, typeTip tl.Tip, public bool) {

	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	if tag := decl.Spec.GetBase(); len(tag) > 0 {
		goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip, typeTip)
		if string(goName) != goSpec.String() {
			fmt.Fprintf(wr, "var %s %s", goName, goSpec)
		}
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

func (gen *Generator) writeUnionDeclaration(wr io.Writer, decl *tl.CDecl,
	ptrTip, typeTip tl.Tip, public bool) {

	cName, _ := getName(decl)
	goName := checkName(gen.tr.TransformName(tl.TargetType, cName, public))
	if tag := decl.Spec.GetBase(); len(tag) > 0 {
		goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTip, typeTip)
		if string(goName) != goSpec.String() {
			fmt.Fprintf(wr, "var %s %s", goName, goSpec)
		}
	}
}
