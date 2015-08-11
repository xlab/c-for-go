package generator

import (
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func (gen *Generator) writeTypeTypedef(wr io.Writer, decl tl.CDecl) {
	goSpec := gen.tr.TranslateSpec(tl.TargetTypedef, decl.Spec)
	declName := gen.tr.TransformName(tl.TargetTypedef, decl.Name)
	fmt.Fprintf(wr, "// %s type as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s %s", declName, goSpec)
	writeSpace(wr, 1)
}

func (gen *Generator) writeFunctionTypedef(wr io.Writer, decl tl.CDecl) {
	var returnRef string
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		returnRef = gen.tr.TranslateSpec(tl.TargetDeclare, spec.Return.Spec).String()
	}

	declName := gen.tr.TransformName(tl.TargetTypedef, decl.Name)
	fmt.Fprintf(wr, "// %s type as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s %s", declName, decl.Spec)
	gen.writeFunctionParams(wr, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
	writeSpace(wr, 1)
}

func (gen *Generator) writeStructTypedef(wr io.Writer, decl tl.CDecl) {
	var declName []byte
	if len(decl.Name) > 0 {
		declName = gen.tr.TransformName(tl.TargetTypedef, decl.Name)
	} else {
		return
	}

	if !decl.IsTemplate() {
		// not a template, so a struct referenced by a tag declares a new type
		typeRef := gen.tr.TranslateSpec(tl.TargetTag, decl.Spec).String()

		if typeName := string(declName); typeName != typeRef {
			fmt.Fprintf(wr, "// %s as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
			fmt.Fprintf(wr, "type %s %s", declName, typeRef)
			writeSpace(wr, 1)
			return
		}
	}

	fmt.Fprintf(wr, "// %s as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s struct {", declName)
	writeSpace(wr, 1)
	gen.writeStructMembers(wr, decl.Spec)
	writeEndStruct(wr)
	writeSpace(wr, 1)
}

func (gen *Generator) writeEnumTypedef(wr io.Writer, decl tl.CDecl) {
	var declName []byte
	if len(decl.Name) > 0 {
		declName = gen.tr.TransformName(tl.TargetTypedef, decl.Name)
	} else {
		return
	}
	typeRef := gen.tr.TranslateSpec(tl.TargetTag, decl.Spec).String()
	if typeName := string(declName); typeName != typeRef {
		fmt.Fprintf(wr, "// %s as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
		fmt.Fprintf(wr, "type %s %s", declName, typeRef)
		writeSpace(wr, 1)
	}
}
