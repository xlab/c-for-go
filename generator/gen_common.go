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
	const public = true
	for i, member := range structSpec.Members {
		ptrTip := ptrTipRx.TipAt(i)
		if !ptrTip.IsValid() {
			ptrTip = tl.TipPtrArr
		}
		declName := gen.transformDeclName(member.Name, public)
		switch member.Spec.Kind() {
		case tl.TypeKind:
			goSpec := gen.tr.TranslateSpec(member.Spec, ptrTip)
			fmt.Fprintf(wr, "// %s member as declared in %s\n", declName, tl.SrcLocation(member.Pos))
			fmt.Fprintf(wr, "%s %s", declName, goSpec)
		case tl.StructKind, tl.OpaqueStructKind:
			if tag := member.Spec.GetBase(); len(tag) > 0 {
				goSpec := gen.tr.TranslateSpec(member.Spec, ptrTip)
				fmt.Fprintf(wr, "%s %s", declName, goSpec)
			}
		case tl.UnionKind:
			cgoSpec := gen.tr.CGoSpec(member.Spec)
			fmt.Fprintf(wr, "%s [unsafe.Sizeof(%s)]byte", declName, cgoSpec.Base)
		case tl.EnumKind:
			typeRef := gen.tr.TranslateSpec(member.Spec, ptrTip).String()
			if declName != typeRef {
				fmt.Fprintf(wr, "%s %s", declName, typeRef)
			}
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
}

func (gen *Generator) writeFunctionParams(wr io.Writer, funcName string, funcSpec tl.CType) {
	spec := funcSpec.(*tl.CFunctionSpec)
	ptrTipSpecRx, _ := gen.tr.PtrTipRx(tl.TipScopeFunction, funcName)
	const public = false

	writeStartParams(wr)
	for i, param := range spec.ParamList {
		ptrTip := ptrTipSpecRx.TipAt(i)
		if !ptrTip.IsValid() {
			ptrTip = tl.TipPtrArr
		}
		declName := gen.transformDeclName(param.Name, public)
		switch param.Spec.Kind() {
		case tl.TypeKind:
			goSpec := gen.tr.TranslateSpec(param.Spec, ptrTip)
			if len(goSpec.Arrays) > 0 {
				fmt.Fprintf(wr, "%s *%s", declName, goSpec)
			} else {
				fmt.Fprintf(wr, "%s %s", declName, goSpec)
			}
		case tl.StructKind, tl.OpaqueStructKind:
			if tag := param.Spec.GetBase(); len(tag) > 0 {
				goSpec := gen.tr.TranslateSpec(param.Spec, ptrTip)
				if len(goSpec.Arrays) > 0 {
					fmt.Fprintf(wr, "%s *%s", declName, goSpec)
				} else {
					fmt.Fprintf(wr, "%s %s", declName, goSpec)
				}
			}
		case tl.UnionKind:
			cgoSpec := gen.tr.CGoSpec(param.Spec)
			fmt.Fprintf(wr, "%s [unsafe.Sizeof(%s)]byte", declName, cgoSpec.Base)
		case tl.EnumKind:
			typeRef := gen.tr.TranslateSpec(param.Spec, ptrTip).String()
			if declName != typeRef {
				fmt.Fprintf(wr, "%s %s", declName, typeRef)
			}
		case tl.FunctionKind:
			gen.writeFunctionAsArg(wr, param, ptrTip, public)
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
