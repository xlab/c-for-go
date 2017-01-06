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
	ptrTipRx, typeTipRx, memTipRx := gen.tr.TipRxsForSpec(tl.TipScopeType, structName, structSpec)
	const public = true
	for i, member := range structSpec.Members {
		ptrTip := ptrTipRx.TipAt(i)
		if !ptrTip.IsValid() {
			ptrTip = tl.TipPtrArr
		}
		typeTip := typeTipRx.TipAt(i)
		if !typeTip.IsValid() {
			typeTip = tl.TipTypeNamed
		}
		memTip := memTipRx.TipAt(i)
		if !memTip.IsValid() {
			memTip = gen.MemTipOf(member)
		}
		if memTip == tl.TipMemRaw {
			ptrTip = tl.TipPtrSRef
		}
		declName := checkName(gen.tr.TransformName(tl.TargetType, member.Name, public))
		switch member.Spec.Kind() {
		case tl.TypeKind:
			goSpec := gen.tr.TranslateSpec(member.Spec, ptrTip, typeTip)
			fmt.Fprintf(wr, "%s %s", declName, goSpec)
		case tl.StructKind, tl.OpaqueStructKind, tl.UnionKind:
			if !gen.tr.IsAcceptableName(tl.TargetType, member.Spec.GetBase()) {
				continue
			}
			goSpec := gen.tr.TranslateSpec(member.Spec, ptrTip, typeTip)
			fmt.Fprintf(wr, "%s %s", declName, goSpec)
		case tl.EnumKind:
			if !gen.tr.IsAcceptableName(tl.TargetType, member.Spec.GetBase()) {
				continue
			}
			typeRef := gen.tr.TranslateSpec(member.Spec, ptrTip, typeTip).String()
			fmt.Fprintf(wr, "%s %s", declName, typeRef)
		case tl.FunctionKind:
			gen.writeFunctionAsArg(wr, member, ptrTip, typeTip, public)
		}
		writeSpace(wr, 1)
	}

	if memTipRx.Self() == tl.TipMemRaw {
		return
	}

	crc := getRefCRC(structSpec)
	cgoSpec := gen.tr.CGoSpec(structSpec, false)
	if len(cgoSpec.Base) == 0 {
		return
	}
	fmt.Fprintf(wr, "ref%2x *%s\n", crc, cgoSpec)
	fmt.Fprintf(wr, "allocs%2x interface{}\n", crc)
}

func (gen *Generator) writeFunctionParams(wr io.Writer, funcName string, funcSpec tl.CType) {
	spec := funcSpec.(*tl.CFunctionSpec)
	ptrTipSpecRx, _ := gen.tr.PtrTipRx(tl.TipScopeFunction, funcName)
	typeTipSpecRx, _ := gen.tr.TypeTipRx(tl.TipScopeFunction, funcName)
	const public = false

	writeStartParams(wr)
	for i, param := range spec.Params {
		ptrTip := ptrTipSpecRx.TipAt(i)
		if !ptrTip.IsValid() {
			ptrTip = tl.TipPtrArr
		}
		typeTip := typeTipSpecRx.TipAt(i)
		if !typeTip.IsValid() {
			// try to use type tip for the type itself
			if tip, ok := gen.tr.TypeTipRx(tl.TipScopeType, param.Spec.CGoName()); ok {
				if tip := tip.Self(); tip.IsValid() {
					typeTip = tip
				}
			}
		}
		declName := checkName(gen.tr.TransformName(tl.TargetType, param.Name, public))
		switch param.Spec.Kind() {
		case tl.TypeKind:
			goSpec := gen.tr.TranslateSpec(param.Spec, ptrTip, typeTip)
			if len(goSpec.OuterArr) > 0 {
				fmt.Fprintf(wr, "%s *%s", declName, goSpec)
			} else {
				fmt.Fprintf(wr, "%s %s", declName, goSpec)
			}
		case tl.StructKind, tl.OpaqueStructKind, tl.UnionKind:
			goSpec := gen.tr.TranslateSpec(param.Spec, ptrTip, typeTip)
			if len(goSpec.OuterArr) > 0 {
				fmt.Fprintf(wr, "%s *%s", declName, goSpec)
			} else {
				fmt.Fprintf(wr, "%s %s", declName, goSpec)
			}
		case tl.EnumKind:
			typeRef := gen.tr.TranslateSpec(param.Spec, ptrTip, typeTip).String()
			fmt.Fprintf(wr, "%s %s", declName, typeRef)
		case tl.FunctionKind:
			gen.writeFunctionAsArg(wr, param, ptrTip, typeTip, public)
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
