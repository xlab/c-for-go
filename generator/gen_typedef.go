package generator

import (
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func (gen *Generator) writeTypeTypedef(wr io.Writer, decl *tl.CDecl) {
	goSpec := gen.tr.TranslateSpec(decl.Spec)
	goTypeName := gen.tr.TransformName(tl.TargetType, decl.Name)
	fmt.Fprintf(wr, "// %s type as declared in %s\n", goTypeName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s %s", goTypeName, goSpec)
	writeSpace(wr, 1)
}

func (gen *Generator) writeEnumTypedef(wr io.Writer, decl *tl.CDecl) {
	var goEnumName []byte
	if len(decl.Name) > 0 {
		goEnumName = gen.tr.TransformName(tl.TargetType, decl.Name)
	} else {
		return
	}
	typeRef := gen.tr.TranslateSpec(decl.Spec).String()
	if typeName := string(goEnumName); typeName != typeRef {
		fmt.Fprintf(wr, "// %s as declared in %s\n", goEnumName, tl.SrcLocation(decl.Pos))
		fmt.Fprintf(wr, "type %s %s", goEnumName, typeRef)
		writeSpace(wr, 1)
	}
}

func (gen *Generator) writeFunctionTypedef(wr io.Writer, decl *tl.CDecl) {
	var returnRef string
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		// defaults to ref for the returns
		ptrTip := tl.TipPtrRef
		if ptrTipRx, ok := gen.tr.PtrTipRx(tl.TipScopeFunction, decl.Name); ok {
			if tip := ptrTipRx.Self(); tip.IsValid() {
				ptrTip = tip
			}
		}
		returnRef = gen.tr.TranslateSpec(spec.Return, ptrTip).String()
	}

	ptrTipRx, _ := gen.tr.PtrTipRx(tl.TipScopeFunction, decl.Name)
	goFuncName := gen.tr.TransformName(tl.TargetType, decl.Name)
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTipRx.Self())
	fmt.Fprintf(wr, "// %s type as declared in %s\n", goFuncName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s %s", goFuncName, goSpec)
	gen.writeFunctionParams(wr, decl.Name, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
	for _, helper := range gen.getCallbackHelpers(string(goFuncName), decl.Name, decl.Spec) {
		gen.submitHelper(helper)
	}
	writeSpace(wr, 1)
}

func (gen *Generator) writeStructTypedef(wr io.Writer, decl *tl.CDecl, raw bool) {
	var goStructName []byte
	if len(decl.Name) > 0 {
		goStructName = gen.tr.TransformName(tl.TargetType, decl.Name)
	} else {
		return
	}

	if !decl.Spec.IsComplete() {
		// not a template, so a struct referenced by a tag declares a new type
		typeRef := gen.tr.TranslateSpec(decl.Spec).String()

		if typeName := string(goStructName); typeName != typeRef {
			fmt.Fprintf(wr, "// %s as declared in %s\n", goStructName, tl.SrcLocation(decl.Pos))
			fmt.Fprintf(wr, "type %s %s", goStructName, typeRef)
			writeSpace(wr, 1)
			return
		}
		return
	}

	fmt.Fprintf(wr, "// %s as declared in %s\n", goStructName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s struct {", goStructName)
	writeSpace(wr, 1)
	gen.writeStructMembers(wr, decl.Name, decl.Spec)
	writeEndStruct(wr)
	writeSpace(wr, 1)
	if raw {
		for _, helper := range gen.getRawStructHelpers(goStructName, decl.Spec) {
			gen.submitHelper(helper)
		}
		return
	}
	for _, helper := range gen.getStructHelpers(goStructName, decl.Name, decl.Spec) {
		gen.submitHelper(helper)
	}
}

func (gen *Generator) writeUnionTypedef(wr io.Writer, decl *tl.CDecl) {
	var goUnionName []byte
	if len(decl.Name) > 0 {
		goUnionName = gen.tr.TransformName(tl.TargetType, decl.Name)
	} else {
		return
	}
	typeRef := gen.tr.TranslateSpec(decl.Spec).String()

	if typeName := string(goUnionName); typeName != typeRef {
		fmt.Fprintf(wr, "// %s as declared in %s\n", goUnionName, tl.SrcLocation(decl.Pos))
		fmt.Fprintf(wr, "type %s %s", goUnionName, typeRef)
		writeSpace(wr, 1)
		return
	}
}
