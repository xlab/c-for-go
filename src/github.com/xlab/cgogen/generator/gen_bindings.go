package generator

import (
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

type HelperSide int

const (
	GoSide HelperSide = 1
	CSide  HelperSide = 2
)

type Helper struct {
	Name        string
	Side        HelperSide
	Description string
	Source      string
}

var proxyHelperMap = map[string]*Helper{
	tl.StringSpec.String(): stringHelper,
}

type proxyDecl struct {
	Name string
	Decl string
}

func (gen *Generator) createProxies(decl tl.CDecl) []proxyDecl {
	spec := decl.Spec.(*tl.CFunctionSpec)
	proxies := make([]proxyDecl, len(spec.ParamList))
	for i, param := range spec.ParamList {
		refType := gen.tr.TranslateSpec(param.Spec)
		refName := gen.tr.TransformName(tl.TargetPrivate, param.Name)
		proxies[i].Name = "c" + string(refName)
		if helper, ok := proxyHelperMap[refType.String()]; ok {
			gen.helpersChan <- helper
			proxies[i].Decl = fmt.Sprintf("%s(%s)", helper.Name, refName)
			continue
		}
		proxies[i].Decl = fmt.Sprintf("*(*C.%s)(%s)", refType.InnerCGO, refName)
	}
	return proxies
}

func (gen *Generator) writeFunctionBody(wr io.Writer, decl tl.CDecl) {
	writeStartFuncBody(wr)
	proxies := gen.createProxies(decl)
	for _, pDecl := range proxies {
		fmt.Fprintf(wr, "%s := %s\n", pDecl.Name, pDecl.Decl)
	}
	writeSpace(wr, 1)

	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		fmt.Fprint(wr, "ret := ")
	}
	fmt.Fprintf(wr, "C.%s", decl.Name)
	writeStartParams(wr)
	for i := range spec.ParamList {
		fmt.Fprint(wr, proxies[i].Name)
		if i < len(spec.ParamList)-1 {
			fmt.Fprint(wr, ", ")
		}
	}
	writeEndParams(wr)
	writeSpace(wr, 1)
	writeEndFuncBody(wr)
}

var (
	stringHelper = &Helper{
		Name: "cStr",
		Side: GoSide,
		Description: `cStr helps to avoid copying Go strings around by referencing the data
string has directly into C land as a *C.char.`,
		Source: `func cStr(str string) *C.char {
			hdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
			return (*C.char)(unsafe.Pointer(hdr.Data))
		}`,
	}
)
