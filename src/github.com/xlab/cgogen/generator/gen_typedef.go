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
	for _, helper := range gen.getStructHelpers(declName, decl.Spec) {
		gen.submitHelper(helper)
	}
}

func (gen *Generator) getStructHelpers(structName []byte, spec tl.CType) (helpers []*Helper) {
	cgoSpec := gen.tr.CGoSpec(spec)
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x *%s) Ref() *%s", structName, cgoSpec)
	fmt.Fprintln(buf, `{
		return x.__ref
	}`)
	helpers = append(helpers, &Helper{
		Name:        "Ref",
		Description: "Ref returns a reference.",
		Source:      buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x *%s) Free()", structName)
	fmt.Fprintln(buf, `{
		if x.__ref != nil {
			C.free(unsafe.Pointer(x.__ref))
			x.__ref = nil
		}
	}`)
	helpers = append(helpers, &Helper{
		Name: "Free",
		Description: "Free cleanups the memory using the free stdlib function on C side.\n" +
			"Does nothing if object has no pointer.",
		Source: buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func New%s(ref *%s) *%s", structName, cgoSpec, structName)
	fmt.Fprintf(buf, `{
		return &%s{
			__ref: ref,
		}
	}`, structName)
	name := fmt.Sprintf("New%s", structName)
	helpers = append(helpers, &Helper{
		Name:        name,
		Description: name + " initialises a new struct holding the reference to the originaitng C struct.",
		Source:      buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x *%s) PassRef() *%s {\n", structName, cgoSpec)
	buf.Write(gen.getPassRefSource(spec))
	buf.WriteRune('}')
	helpers = append(helpers, &Helper{
		Name:        "PassRef",
		Description: "PassRef returns a reference and creates new C object if no refernce yet.",
		Source:      buf.String(),
	})
	return
}

func (gen *Generator) getPassRefSource(structSpec tl.CType) []byte {
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
