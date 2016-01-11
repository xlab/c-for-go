package generator

import (
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func (gen *Generator) writeTypeTypedef(wr io.Writer, decl *tl.CDecl) {
	goSpec := gen.tr.TranslateSpec(decl.Spec)
	goTypeName := gen.tr.TransformName(tl.TargetType, decl.Name)
	fmt.Fprintf(wr, "// %s type as declared in %s\n", goTypeName,
		gen.tr.SrcLocation(tl.TargetType, decl.Name, decl.Pos))
	fmt.Fprintf(wr, "type %s %s", goTypeName, goSpec.UnderlyingString())
	writeSpace(wr, 1)
}

func (gen *Generator) writeEnumTypedef(wr io.Writer, decl *tl.CDecl) {
	cName, ok := getName(decl)
	if !ok {
		return
	}
	goName := gen.tr.TransformName(tl.TargetType, cName)
	typeRef := gen.tr.TranslateSpec(decl.Spec).UnderlyingString()
	if typeName := string(goName); typeName != typeRef {
		fmt.Fprintf(wr, "// %s as declared in %s\n", goName,
			gen.tr.SrcLocation(tl.TargetConst, cName, decl.Pos))
		fmt.Fprintf(wr, "type %s %s", goName, typeRef)
		writeSpace(wr, 1)
	}
}

func (gen *Generator) writeFunctionTypedef(wr io.Writer, decl *tl.CDecl) {
	var returnRef string
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		// defaults to ref for the return values
		ptrTip := tl.TipPtrRef
		if ptrTipRx, ok := gen.tr.PtrTipRx(tl.TipScopeFunction, decl.Name); ok {
			if tip := ptrTipRx.Self(); tip.IsValid() {
				ptrTip = tip
			}
		}
		returnRef = gen.tr.TranslateSpec(spec.Return, ptrTip).UnderlyingString()
	}

	ptrTipRx, _ := gen.tr.PtrTipRx(tl.TipScopeFunction, decl.Name)
	goFuncName := gen.tr.TransformName(tl.TargetType, decl.Name)
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrTipRx.Self())
	fmt.Fprintf(wr, "// %s type as declared in %s\n", goFuncName,
		gen.tr.SrcLocation(tl.TargetFunction, decl.Name, decl.Pos))
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

func getName(decl *tl.CDecl) (string, bool) {
	if base := decl.Spec.GetBase(); len(base) > 0 {
		return base, true
	}
	if len(decl.Name) > 0 {
		return decl.Name, true
	}
	return "", false
}

func (gen *Generator) writeStructTypedef(wr io.Writer, decl *tl.CDecl, raw bool) {
	cName, ok := getName(decl)
	if !ok {
		return
	}
	goName := gen.tr.TransformName(tl.TargetType, cName)

	if !decl.Spec.IsComplete() {
		// opaque struct
		fmt.Fprintf(wr, "// %s as declared in %s\n", goName,
			gen.tr.SrcLocation(tl.TargetType, cName, decl.Pos))
		fmt.Fprintf(wr, "type %s C.%s", goName, decl.Spec.CGoName())
		writeSpace(wr, 1)
		return
	}

	fmt.Fprintf(wr, "// %s as declared in %s\n", goName,
		gen.tr.SrcLocation(tl.TargetType, cName, decl.Pos))
	fmt.Fprintf(wr, "type %s struct {", goName)
	writeSpace(wr, 1)
	gen.submitHelper(cgoAllocMap)
	gen.writeStructMembers(wr, cName, decl.Spec)
	writeEndStruct(wr)
	writeSpace(wr, 1)
	if raw {
		for _, helper := range gen.getRawStructHelpers(goName, decl.Spec) {
			gen.submitHelper(helper)
		}
		return
	}
	for _, helper := range gen.getStructHelpers(goName, cName, decl.Spec) {
		gen.submitHelper(helper)
	}
}

func (gen *Generator) writeUnionTypedef(wr io.Writer, decl *tl.CDecl) {
	cName, ok := getName(decl)
	if !ok {
		return
	}
	goName := gen.tr.TransformName(tl.TargetType, cName)
	typeRef := gen.tr.TranslateSpec(decl.Spec).UnderlyingString()

	if typeName := string(goName); typeName != typeRef {
		fmt.Fprintf(wr, "// %s as declared in %s\n", goName,
			gen.tr.SrcLocation(tl.TargetType, cName, decl.Pos))
		fmt.Fprintf(wr, "type %s %s", goName, typeRef)
		writeSpace(wr, 1)
		return
	}
}
