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
	for _, param := range funcSpec.ParamList {
		params = append(params, param.String())
		paramNames = append(paramNames, param.Name)
		goName := checkName(gen.tr.TransformName(tl.TargetType, param.Name, false))
		paramNamesGo = append(paramNamesGo, string(goName))
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
		Base:     cFuncName,
		Pointers: 1,
	})
	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x %s) PassRef() %s", goFuncName, cgoSpec)
	fmt.Fprintf(buf, `{
 		%sFunc = x
		return (%s)(C.%s)
	}`, cbGoName, cgoSpec, cbCName)
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.PassRef", goFuncName),
		Description: "PassRef returns a reference.",
		Source:      buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "//export %s\n", cbGoName)
	cbGoDecl := tl.CDecl{
		Name: cbGoName,
		Spec: spec,
	}
	gen.writeCallbackProxy(buf, cbGoDecl)
	fmt.Fprintf(buf, `{
		if %sFunc != nil {
			return %sFunc(%s)
		}
		panic("callback func has not been set (race?)")
	}`, cbGoName, cbGoName, paramNamesGoList)
	fmt.Fprintf(buf, "\n\nvar %sFunc %s", cbGoName, goFuncName)
	helpers = append(helpers, &Helper{
		Name:   cbGoName,
		Source: buf.String(),
	})

	return
}

func (gen *Generator) writeCallbackProxy(wr io.Writer, decl tl.CDecl) {
	var returnRef string
	funcSpec := decl.Spec.(*tl.CFunctionSpec)
	if funcSpec.Return != nil {
		cgoSpec := gen.tr.CGoSpec(funcSpec.Return.Spec)
		returnRef = cgoSpec.String()
	}
	declName := gen.tr.TransformName(tl.TargetFunction, decl.Name, false)
	fmt.Fprintf(wr, "func %s", declName)
	gen.writeCallbackProxyParams(wr, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
}

func (gen *Generator) writeCallbackProxyParams(wr io.Writer, spec tl.CType) {
	funcSpec := spec.(*tl.CFunctionSpec)
	const public = false

	writeStartParams(wr)
	for i, param := range funcSpec.ParamList {
		declName := checkName(gen.tr.TransformName(tl.TargetType, param.Name, public))
		cgoSpec := gen.tr.CGoSpec(param.Spec)
		fmt.Fprintf(wr, "%s %s", declName, cgoSpec)

		if i < len(funcSpec.ParamList)-1 {
			fmt.Fprintf(wr, ", ")
		}
	}
	writeEndParams(wr)
}
