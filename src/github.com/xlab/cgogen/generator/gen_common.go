package generator

import (
	"fmt"
	"io"
	"strings"

	tl "github.com/xlab/cgogen/translator"
)

var skipName = []byte("_")

func (gen *Generator) writeStructMembers(wr io.Writer, structSpec tl.CType) {
	spec := structSpec.(*tl.CStructSpec)
	for _, member := range spec.Members {
		switch member.Spec.Kind() {
		case tl.TypeKind:
			gen.writeTypeDeclaration(wr, member, false)
		case tl.StructKind:
			gen.writeStructDeclaration(wr, member, false)
		case tl.EnumKind:
			gen.writeEnumDeclaration(wr, member, false)
		case tl.FunctionKind:
			gen.writeFunctionDeclaration(wr, member, false)
		}
		writeSpace(wr, 1)
	}
}

func (gen *Generator) writeFunctionParams(wr io.Writer, funcSpec tl.CType) {
	spec := funcSpec.(*tl.CFunctionSpec)
	writeStartParams(wr)
	for _, param := range spec.ParamList {
		switch param.Spec.Kind() {
		case tl.TypeKind:
			gen.writeTypeDeclaration(wr, param, false)
		case tl.StructKind:
			gen.writeStructDeclaration(wr, param, false)
		case tl.EnumKind:
			gen.writeEnumDeclaration(wr, param, false)
		case tl.FunctionKind:
			gen.writeFunctionDeclaration(wr, param, false)
		}
		fmt.Fprintf(wr, ",")
	}
	writeEndParams(wr)
}

func writeStartParams(wr io.Writer) {
	fmt.Fprint(wr, "(")
}

func writeEndParams(wr io.Writer) {
	fmt.Fprint(wr, ")")
}

func writeEndStruct(wr io.Writer) {
	fmt.Fprintln(wr, "}")
}

func writeSpace(wr io.Writer, n int) {
	fmt.Fprint(wr, strings.Repeat("\n", n))
}

func writeError(wr io.Writer, err error) {
	fmt.Fprintf(wr, "// error: %v\n", err)
}
