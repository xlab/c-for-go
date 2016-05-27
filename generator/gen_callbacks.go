package generator

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	tl "github.com/xlab/cgogen/translator"
)

func unexportName(name string) string {
	if len(name) == 0 {
		return name
	}
	return strings.ToLower(name[:1]) + name[1:]
}

func (gen *Generator) getCallbackHelpers(goFuncName, cFuncName string, spec tl.CType) (helpers []*Helper) {
	crc := getRefCRC(spec)
	cbCName := fmt.Sprintf("%s_%2x", cFuncName, crc)
	cbGoName := fmt.Sprintf("%s%2X", unexportName(goFuncName), crc)
	funcSpec := spec.(*tl.CFunctionSpec)

	var params []string
	var paramNames []string
	var paramNamesGo []string
	for i, param := range funcSpec.Params {
		if len(param.Name) == 0 {
			param.Name = fmt.Sprintf("arg%d", i)
		}
		paramSpec := gen.tr.NormalizeSpecPointers(param.Spec)
		params = append(params, fmt.Sprintf("%s %s", paramSpec.AtLevel(0), param.Name))
		paramNames = append(paramNames, param.Name)
		goName := checkName(gen.tr.TransformName(tl.TargetType, param.Name, false))
		paramNamesGo = append(paramNamesGo, fmt.Sprintf("%s%2x", goName, crc))
	}
	paramList := strings.Join(params, ", ")
	paramNamesList := strings.Join(paramNames, ", ")
	paramNamesGoList := strings.Join(paramNamesGo, ", ")

	buf := new(bytes.Buffer)
	retSpec := "void"
	if funcSpec.Return != nil {
		retSpec = funcSpec.Return.String()
	}
	fmt.Fprintf(buf, "%s %s(%s);", retSpec, cbCName, paramList)
	helpers = append(helpers, &Helper{
		Name:        cbCName,
		Description: fmt.Sprintf("%s is a proxy for callback %s.", cbCName, cFuncName),
		Source:      buf.String(),
		Side:        CHSide,
	})

	var ret string
	if funcSpec.Return != nil {
		ret = "return "
	}
	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "%s %s(%s) {\n", retSpec, cbCName, paramList)
	fmt.Fprintf(buf, "\t%s%s(%s);\n", ret, cbGoName, paramNamesList)
	buf.WriteRune('}')
	helpers = append(helpers, &Helper{
		Name:   cbCName,
		Source: buf.String(),
		Side:   CCSide,
	})

	cgoSpec := gen.tr.CGoSpec(&tl.CTypeSpec{
		Base: cFuncName,
	}, true)
	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x %s) PassRef() (ref *%s, allocs *cgoAllocMap)", goFuncName, cgoSpec)
	fmt.Fprintf(buf, `{
		if x == nil {
			return nil, nil
		}
		if %sFunc == nil {
 			%sFunc = x
 		}
		return (*%s)(C.%s), nil
	}`, cbGoName, cbGoName, cgoSpec, cbCName)
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.PassRef", goFuncName),
		Description: "PassRef returns a reference.",
		Source:      buf.String(),
	})

	if spec.GetPointers() > 0 {
		buf = new(bytes.Buffer)
		fmt.Fprintf(buf, "func (x %s) PassValue() (ref %s, allocs *cgoAllocMap)", goFuncName, cgoSpec)
		fmt.Fprintf(buf, `{
		if x == nil {
			return nil, nil
		}
		if %sFunc == nil {
 			%sFunc = x
 		}
		return (%s)(C.%s), nil
	}`, cbGoName, cbGoName, cgoSpec, cbCName)
		helpers = append(helpers, &Helper{
			Name:        fmt.Sprintf("%s.PassValue", goFuncName),
			Description: "PassValue returns a value.",
			Source:      buf.String(),
		})
	}

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "//export %s\n", cbGoName)
	cbGoDecl := &tl.CDecl{
		Name: cbGoName,
		Spec: spec,
	}

	proxyLines := gen.createCallbackProxies(cFuncName, funcSpec)
	proxySrc := new(bytes.Buffer)
	for i := range proxyLines {
		proxySrc.WriteString(proxyLines[i].Decl)
	}

	gen.writeCallbackProxyFunc(buf, cbGoDecl)
	fmt.Fprintln(buf, "{")
	fmt.Fprintf(buf, "if %sFunc != nil {\n", cbGoName)
	buf.WriteString(proxySrc.String())
	if funcSpec.Return != nil {
		ret := fmt.Sprintf("ret%2x", crc)
		fmt.Fprintf(buf, "%s := %sFunc(%s)\n", ret, cbGoName, paramNamesGoList)
		ptrTipRx, typeTipRx, memTipRx := gen.tr.TipRxsForSpec(tl.TipScopeFunction, cFuncName, funcSpec)
		retGoSpec := gen.tr.TranslateSpec(funcSpec.Return, ptrTipRx.Self(), typeTipRx.Self())
		retCGoSpec := gen.tr.CGoSpec(funcSpec.Return, true) // asArg?
		retProxy, _ := gen.proxyArgFromGo(memTipRx.Self(), ret, retGoSpec, retCGoSpec)
		fmt.Fprintf(buf, "ret, _ := %s\n", retProxy)
		fmt.Fprintf(buf, "return ret\n")
	} else {
		fmt.Fprintf(buf, "%sFunc(%s)\n", cbGoName, paramNamesGoList)
	}
	fmt.Fprintln(buf, "}")
	fmt.Fprintln(buf, `panic("callback func has not been set (race?)")`)
	fmt.Fprintln(buf, "}")

	fmt.Fprintf(buf, "\n\nvar %sFunc %s", cbGoName, goFuncName)
	helpers = append(helpers, &Helper{
		Name:   cbGoName,
		Source: buf.String(),
	})
	return
}

func (gen *Generator) writeCallbackProxyFunc(wr io.Writer, decl *tl.CDecl) {
	var returnRef string
	funcSpec := decl.Spec.(*tl.CFunctionSpec)
	if funcSpec.Return != nil {
		cgoSpec := gen.tr.CGoSpec(funcSpec.Return, true) // asArg?
		returnRef = cgoSpec.String()
	}
	fmt.Fprintf(wr, "func %s", decl.Name)
	gen.writeCallbackProxyFuncParams(wr, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
}

func (gen *Generator) writeCallbackProxyFuncParams(wr io.Writer, spec tl.CType) {
	funcSpec := spec.(*tl.CFunctionSpec)
	const public = false

	writeStartParams(wr)
	for i, param := range funcSpec.Params {
		declName := checkName(gen.tr.TransformName(tl.TargetType, param.Name, public))
		if len(param.Name) == 0 {
			declName = []byte(fmt.Sprintf("arg%d", i))
		}
		cgoSpec := gen.tr.CGoSpec(param.Spec, true)
		fmt.Fprintf(wr, "c%s %s", declName, cgoSpec.AtLevel(0))

		if i < len(funcSpec.Params)-1 {
			fmt.Fprintf(wr, ", ")
		}
	}
	writeEndParams(wr)
}

func (gen *Generator) createCallbackProxies(funcName string, funcSpec tl.CType) (to []proxyDecl) {
	spec := funcSpec.(*tl.CFunctionSpec)
	to = make([]proxyDecl, 0, len(spec.Params))

	crc := getRefCRC(funcSpec)
	ptrTipRx, typeTipRx, memTipRx := gen.tr.TipRxsForSpec(tl.TipScopeFunction, funcName, funcSpec)
	for i, param := range spec.Params {
		var goSpec tl.GoTypeSpec
		ptrTip := ptrTipRx.TipAt(i)
		typeTip := typeTipRx.TipAt(i)
		goSpec = gen.tr.TranslateSpec(param.Spec, ptrTip, typeTip)
		cgoSpec := gen.tr.CGoSpec(param.Spec, true)
		const public = false
		refName := string(gen.tr.TransformName(tl.TargetType, param.Name, public))
		if len(param.Name) == 0 {
			refName = fmt.Sprintf("arg%d", i)
		}
		toBuf := new(bytes.Buffer)
		cname := "c" + refName
		refName = fmt.Sprintf("%s%2x", refName, crc)
		toProxy, _ := gen.proxyCallbackArgToGo(memTipRx.TipAt(i), refName, cname, goSpec, cgoSpec)
		if len(toProxy) > 0 {
			fmt.Fprintln(toBuf, toProxy)
			to = append(to, proxyDecl{Name: cname, Decl: toBuf.String()})
		}
	}
	return
}

func (gen *Generator) proxyCallbackArgToGo(memTip tl.Tip, varName, ptrName string,
	goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec) (proxy string, nillable bool) {
	nillable = true

	if getHelper, ok := toGoHelperMap[goSpec]; ok {
		helper := getHelper(gen, cgoSpec)
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s := %s(%s)", varName, helper.Name, ptrName)
		return proxy, helper.Nillable
	}
	gen.submitHelper(cgoAllocMap)

	isPlain := (memTip == tl.TipMemRaw) || goSpec.IsPlain() || goSpec.IsPlainKind()
	switch {
	case !isPlain && (goSpec.Slices > 0 || len(goSpec.OuterArr) > 0), // ex: []string
		isPlain && goSpec.Slices > 0 && len(goSpec.OuterArr) > 0, // ex: [4][]byte
		isPlain && goSpec.Slices > 1:                             // ex: [][]byte
		helper := gen.getPackHelper(memTip, goSpec, cgoSpec)
		gen.submitHelper(helper)
		if len(goSpec.OuterArr) > 0 {
			ptrName = fmt.Sprintf("%s := (*%s)(unsafe.Pointer(&%s))", varName, cgoSpec, ptrName)
		}
		gen.submitHelper(sliceHeader)
		proxy = fmt.Sprintf("var %s %s\n%s(%s, %s)", varName, goSpec, helper.Name, varName, ptrName)
		return proxy, helper.Nillable
	case isPlain && goSpec.Slices != 0: // ex: []byte
		gen.submitHelper(sliceHeader)
		buf := new(bytes.Buffer)
		postfix := gen.randPostfix()
		fmt.Fprintf(buf, "var %s %s\n", varName, goSpec)
		fmt.Fprintf(buf, "hx%2x := (*sliceHeader)(unsafe.Pointer(&%s))\n", postfix, varName)
		fmt.Fprintf(buf, "hx%2x.Data = uintptr(unsafe.Pointer(%s))\n", postfix, ptrName)
		fmt.Fprintf(buf, "hx%2x.Cap = 0x7fffffff\n", postfix)
		fmt.Fprintf(buf, "// hx%2x.Len = ?\n", postfix)
		proxy = buf.String()
		return
	case isPlain: // ex: byte, [4]byte
		var ref, ptr string
		if (goSpec.Kind == tl.PlainTypeKind || goSpec.Kind == tl.EnumKind) &&
			len(goSpec.OuterArr) == 0 && goSpec.Pointers == 0 {
			proxy = fmt.Sprintf("%s := (%s)(%s)", varName, goSpec, ptrName)
			return
		} else if goSpec.Pointers == 0 {
			ref = "&"
			ptr = "*"
		}
		proxy = fmt.Sprintf("%s := %s(%s%s)(unsafe.Pointer(%s%s))", varName, ptr, ptr, goSpec, ref, ptrName)
		return
	default: // ex: *SomeType
		var ref, deref string
		if cgoSpec.Pointers == 0 {
			deref = "*"
			ref = "&"
		}
		proxy = fmt.Sprintf("%s := %sNew%sRef(%s%s)", varName, deref, goSpec.Raw, ref, ptrName)
		return
	}
}
