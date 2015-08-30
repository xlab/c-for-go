package generator

import (
	"bytes"
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func (gen *Generator) writeTypeTypedef(wr io.Writer, decl tl.CDecl) {
	goSpec := gen.tr.TranslateSpec(decl.Spec)
	declName := gen.tr.TransformName(tl.TargetType, decl.Name)
	fmt.Fprintf(wr, "// %s type as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s %s", declName, goSpec)
	writeSpace(wr, 1)
}

func (gen *Generator) writeEnumTypedef(wr io.Writer, decl tl.CDecl) {
	var declName []byte
	if len(decl.Name) > 0 {
		declName = gen.tr.TransformName(tl.TargetType, decl.Name)
	} else {
		return
	}
	typeRef := gen.tr.TranslateSpec(decl.Spec).String()
	if typeName := string(declName); typeName != typeRef {
		fmt.Fprintf(wr, "// %s as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
		fmt.Fprintf(wr, "type %s %s", declName, typeRef)
		writeSpace(wr, 1)
	}
}

func (gen *Generator) writeFunctionTypedef(wr io.Writer, decl tl.CDecl) {
	var returnRef string
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		returnRef = gen.tr.TranslateSpec(spec.Return.Spec).String()
	}

	declName := gen.tr.TransformName(tl.TargetType, decl.Name)
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
		declName = gen.tr.TransformName(tl.TargetType, decl.Name)
	} else {
		return
	}

	if !decl.IsTemplate() {
		// not a template, so a struct referenced by a tag declares a new type
		typeRef := gen.tr.TranslateSpec(decl.Spec).String()

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
	gen.writeStructMethods(wr, string(declName), decl.Spec)
}

func (gen *Generator) writeStructMethods(wr io.Writer, structName string, spec tl.CType) {
	cgoSpec := gen.tr.CGoSpec(spec)
	fmt.Fprintf(wr, "func (x *%s) Ref() *%s", structName, cgoSpec)
	fmt.Fprintln(wr, `{
		return x.__ref
	}`)
	writeSpace(wr, 1)
	fmt.Fprintf(wr, "func (x *%s) Free()", structName)
	fmt.Fprintln(wr, `{
		if x.__ref != nil {
			C.free(unsafe.Pointer(x.__ref))
			x.__ref = nil
		}
	}`)
	fmt.Fprintf(wr, "func New%s(ref *%s) *%s", structName, cgoSpec, structName)
	fmt.Fprintf(wr, `{
		return &%s{
			__ref: ref,
		}
	}`, structName)
	writeSpace(wr, 2)
	fmt.Fprintf(wr, "func (x *%s) PassRef() *%s {\n", structName, cgoSpec)
	wr.Write(gen.getPassRefSource(structName, spec))
	fmt.Fprintln(wr, "}")
}

func (gen *Generator) getPassRefSource(structName string, structSpec tl.CType) []byte {
	cgoSpec := gen.tr.CGoSpec(structSpec)
	spec := structSpec.(*tl.CStructSpec)
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, `if x.__ref != nil {
		return x.__ref
	}`)
	writeSpace(buf, 1)
	fmt.Fprintf(buf, "__ref := new(%s)\n", cgoSpec.Base)
	for _, mem := range spec.Members {
		if len(mem.Name) == 0 {
			continue
		}
		goSpec := gen.tr.TranslateSpec(mem.Spec)
		cgoSpec := gen.tr.CGoSpec(mem.Spec)
		goName := "x." + string(gen.tr.TransformName(tl.TargetPublic, mem.Name))
		fromProxy, nillable := gen.proxyValueFromGo(goName, goSpec, cgoSpec)
		if nillable {
			fmt.Fprintf(buf, "if %s != nil {\n__ref.%s = %s\n}\n", goName, mem.Name, fromProxy)
		} else {
			fmt.Fprintf(buf, "__ref.%s = %s\n", mem.Name, fromProxy)
		}
	}
	fmt.Fprintln(buf, `x.__ref = __ref
		return __ref`)
	return buf.Bytes()
}
