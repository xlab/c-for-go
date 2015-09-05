package generator

import (
	"fmt"
	"io"
	"strings"

	tl "github.com/xlab/cgogen/translator"
)

var skipName = []byte("_")

func (gen *Generator) writeStructMembers(wr io.Writer, structName string, spec tl.CType) {
	structSpec := spec.(*tl.CStructSpec)

	var nextPtrSpec tl.PointerSpec
	ptrLayout, fallback := gen.tr.PointerLayout(tl.PointerScopeStruct, structName)

	for _, member := range structSpec.Members {
		nextPtrSpec, ptrLayout = tl.NextPointerSpec(ptrLayout, fallback)
		// declName := gen.tr.TransformName(tl.TargetPublic, member.Name)
		// fmt.Fprintf(wr, "// %s member as declared in %s\n", declName, tl.SrcLocation(member.Pos))
		switch member.Spec.Kind() {
		case tl.TypeKind:
			gen.writeTypeDeclaration(wr, member, nextPtrSpec, true)
		case tl.StructKind:
			gen.writeArgStruct(wr, member, nextPtrSpec, true)
		case tl.EnumKind:
			gen.writeEnumDeclaration(wr, member, nextPtrSpec, true)
		case tl.FunctionKind:
			gen.writeArgFunction(wr, member, nextPtrSpec, true)
		}
		writeSpace(wr, 1)
	}
	crc := getRefCRC(structSpec)
	cgoSpec := gen.tr.CGoSpec(structSpec)
	if len(cgoSpec.Base) == 0 {
		return
	}
	fmt.Fprintf(wr, "ref%2x *%s\n", crc, cgoSpec)
}

func (gen *Generator) writeFunctionParams(wr io.Writer, funcName string, funcSpec tl.CType) {
	spec := funcSpec.(*tl.CFunctionSpec)

	var nextPtrSpec tl.PointerSpec
	ptrLayout, fallback := gen.tr.PointerLayout(tl.PointerScopeFunction, funcName)

	writeStartParams(wr)
	for i, param := range spec.ParamList {
		nextPtrSpec, ptrLayout = tl.NextPointerSpec(ptrLayout, fallback)
		switch param.Spec.Kind() {
		case tl.TypeKind:
			gen.writeTypeDeclaration(wr, param, nextPtrSpec, false)
		case tl.StructKind:
			gen.writeArgStruct(wr, param, nextPtrSpec, false)
		case tl.EnumKind:
			gen.writeEnumDeclaration(wr, param, nextPtrSpec, false)
		case tl.FunctionKind:
			gen.writeArgFunction(wr, param, nextPtrSpec, false)
		}
		if i < len(spec.ParamList)-1 {
			fmt.Fprintf(wr, ", ")
		}
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
	fmt.Fprint(wr, "}")
}

func writeStartFuncBody(wr io.Writer) {
	fmt.Fprintln(wr, "{")
}

func writeEndFuncBody(wr io.Writer) {
	fmt.Fprintln(wr, "}")
}

func writeSpace(wr io.Writer, n int) {
	fmt.Fprint(wr, strings.Repeat("\n", n))
}

func writeError(wr io.Writer, err error) {
	fmt.Fprintf(wr, "// error: %v\n", err)
}
