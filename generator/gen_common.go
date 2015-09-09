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
	ptrTipRx, memTipRx := gen.tr.PtrMemTipRxForSpec(tl.TipScopeStruct, structName, structSpec)
	for i, member := range structSpec.Members {
		ptrTip := ptrTipRx.TipAt(i)
		if !ptrTip.IsValid() {
			ptrTip = tl.TipPtrArr
		}
		// declName := gen.tr.TransformName(tl.TargetPublic, member.Name)
		// fmt.Fprintf(wr, "// %s member as declared in %s\n", declName, tl.SrcLocation(member.Pos))
		switch member.Spec.Kind() {
		case tl.TypeKind:
			gen.writeTypeDeclaration(wr, member, ptrTip, true)
		case tl.StructKind, tl.OpaqueStructKind:
			gen.writeArgStruct(wr, member, ptrTip, true)
		case tl.EnumKind:
			gen.writeEnumDeclaration(wr, member, ptrTip, true)
		case tl.FunctionKind:
			gen.writeArgFunction(wr, member, ptrTip, true)
		}
		writeSpace(wr, 1)
	}

	if memTipRx.Self() == tl.TipMemRaw {
		return
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
	ptrTipSpecRx, _ := gen.tr.PtrTipRx(tl.TipScopeFunction, funcName)

	writeStartParams(wr)
	for i, param := range spec.ParamList {
		ptrTip := ptrTipSpecRx.TipAt(i)
		if !ptrTip.IsValid() {
			ptrTip = tl.TipPtrArr
		}
		switch param.Spec.Kind() {
		case tl.TypeKind:
			gen.writeTypeDeclaration(wr, param, ptrTip, false)
		case tl.StructKind, tl.OpaqueStructKind:
			gen.writeArgStruct(wr, param, ptrTip, false)
		case tl.EnumKind:
			gen.writeEnumDeclaration(wr, param, ptrTip, false)
		case tl.FunctionKind:
			gen.writeArgFunction(wr, param, ptrTip, false)
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
