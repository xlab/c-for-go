package generator

import (
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func writeType(wr io.Writer, name []byte, spec tl.GoTypeSpec) {
	fmt.Fprintf(wr, "type %s %s\n", name, spec)
}

func writeVar(wr io.Writer, name []byte, spec tl.GoTypeSpec) {
	fmt.Fprintf(wr, "var %s %s\n", name, spec)
}

func writeMember(wr io.Writer, name []byte, spec tl.GoTypeSpec) {
	fmt.Fprintf(wr, "%s %s\n", name, spec)
}

func (gen *Generator) writeTypeDeclaration(wr io.Writer, typeDecl tl.CDecl) {
	goSpec := gen.tr.TranslateSpec(typeDecl.Spec)
	declName := gen.tr.TransformName(tl.TargetTypedef, typeDecl.Name)
	if len(declName) == 0 {
		return
	}
	fmt.Fprintf(wr, "// %s type as declared in %s\n", declName, tl.SrcLocation(typeDecl.Pos))
	writeType(wr, declName, goSpec)
}

func (gen *Generator) writeMemberDeclaration(wr io.Writer, decl tl.CDecl) {
	switch decl.Spec.Kind() {
	case tl.StructKind:
		gen.writeStructDeclaration(wr, decl)
	case tl.TypeKind:
		goSpec := gen.tr.TranslateSpec(decl.Spec)
		declName := gen.tr.TransformName(tl.TargetDeclare, decl.Name)
		if len(declName) == 0 {
			declName = []byte("noname")
		}
		writeMember(wr, declName, goSpec)
	default:
		writeError(wr, fmt.Errorf("yet unsupported type: %s", decl.Spec))
	}
}

func (gen *Generator) writeStructDeclaration(wr io.Writer, structDecl tl.CDecl) {
	if !structDecl.IsTypedef {
		return
	}
	spec, ok := structDecl.Spec.(*tl.CStructSpec)
	if !ok {
		return
	} else if spec.IsUnion {
		return
	}

	var typeName []byte
	if len(structDecl.Name) > 0 {
		typeName = gen.tr.TransformName(tl.TargetTypedef, structDecl.Name)
	} else if len(spec.Tag) > 0 {
		typeName = gen.tr.TransformName(tl.TargetTag, spec.Tag)
	} else {
		return
	}

	if len(spec.Members) == 0 {
		typeRef := gen.tr.TranslateSpec(spec).String()
		if string(typeName) == typeRef {
			// looks that a type alias turned out to be redundant
			// after translation rules.
			return
		}
		fmt.Fprintf(wr, "// %s as declared in %s\n", typeName, tl.SrcLocation(structDecl.Pos))
		fmt.Fprintf(wr, "type %s %s\n", typeName, typeRef)
		return
	}

	fmt.Fprintf(wr, "// %s as declared in %s\n", typeName, tl.SrcLocation(structDecl.Pos))
	fmt.Fprintf(wr, "type %s struct {\n", typeName)

	for _, memDecl := range spec.Members {
		memName := gen.tr.TransformName(tl.TargetDeclare, memDecl.Name)
		if len(memName) == 0 {
			memName = []byte("noname")
		}
		fmt.Fprintf(wr, "// %s member as declared in %s\n", memName, tl.SrcLocation(memDecl.Pos))
		gen.writeMemberDeclaration(wr, memDecl)
	}
	writeEndStruct(wr)
}

func writeStartStruct(wr io.Writer) {
	fmt.Fprintln(wr, "struct {")
}

func writeEndStruct(wr io.Writer) {
	fmt.Fprintln(wr, "}")
}
