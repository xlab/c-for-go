package generator

import (
	"fmt"
	"io"
	"strings"

	tl "github.com/xlab/cgogen/translator"
)

var (
	skipName    = []byte("_")
	skipNameStr = "_"
)

func (gen *Generator) writeStructMembers(wr io.Writer, structName string, spec tl.CType) {
	structSpec := spec.(*tl.CStructSpec)
	ptrTipRx, memTipRx := gen.tr.PtrMemTipRxForSpec(tl.TipScopeStruct, structName, structSpec)
	const public = true
	for i, member := range structSpec.Members {
		ptrTip := ptrTipRx.TipAt(i)
		if !ptrTip.IsValid() {
			ptrTip = tl.TipPtrArr
		}
		declName := checkName(gen.tr.TransformName(tl.TargetType, member.Name, public))
		switch member.Spec.Kind() {
		case tl.TypeKind:
			goSpec := gen.tr.TranslateSpec(member.Spec, ptrTip)
			fmt.Fprintf(wr, "%s %s", declName, goSpec)
		case tl.StructKind, tl.OpaqueStructKind:
			if !gen.tr.IsAcceptableName(tl.TargetType, member.Spec.GetBase()) {
				continue
			}
			goSpec := gen.tr.TranslateSpec(member.Spec, ptrTip)
			fmt.Fprintf(wr, "%s %s", declName, goSpec)
		case tl.EnumKind:
			if !gen.tr.IsAcceptableName(tl.TargetType, member.Spec.GetBase()) {
				continue
			}
			typeRef := gen.tr.TranslateSpec(member.Spec, ptrTip).String()
			fmt.Fprintf(wr, "%s %s", declName, typeRef)
		case tl.FunctionKind:
			gen.writeFunctionAsArg(wr, member, ptrTip, public)
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
	fmt.Fprintf(wr, "allocs%2x interface{}\n", crc)
}

func (gen *Generator) writeFunctionParams(wr io.Writer, funcName string, funcSpec tl.CType) {
	spec := funcSpec.(*tl.CFunctionSpec)
	ptrTipSpecRx, _ := gen.tr.PtrTipRx(tl.TipScopeFunction, funcName)
	const public = false

	writeStartParams(wr)
	for i, param := range spec.Params {
		ptrTip := ptrTipSpecRx.TipAt(i)
		if !ptrTip.IsValid() {
			ptrTip = tl.TipPtrArr
		}
		declName := checkName(gen.tr.TransformName(tl.TargetType, param.Name, public))
		switch param.Spec.Kind() {
		case tl.TypeKind:
			goSpec := gen.tr.TranslateSpec(param.Spec, ptrTip)
			if len(goSpec.Arrays) > 0 {
				fmt.Fprintf(wr, "%s *%s", declName, goSpec)
			} else {
				fmt.Fprintf(wr, "%s %s", declName, goSpec)
			}
		case tl.StructKind, tl.OpaqueStructKind:
			goSpec := gen.tr.TranslateSpec(param.Spec, ptrTip)
			if len(goSpec.Arrays) > 0 {
				fmt.Fprintf(wr, "%s *%s", declName, goSpec)
			} else {
				fmt.Fprintf(wr, "%s %s", declName, goSpec)
			}
		case tl.EnumKind:
			typeRef := gen.tr.TranslateSpec(param.Spec, ptrTip).String()
			fmt.Fprintf(wr, "%s %s", declName, typeRef)
		case tl.FunctionKind:
			gen.writeFunctionAsArg(wr, param, ptrTip, public)
		}
		if i < len(spec.Params)-1 {
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
