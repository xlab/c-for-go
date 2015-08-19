package generator

import (
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

type HelperSide string

const (
	NoSide HelperSide = ""
	GoSide HelperSide = "go"
	CSide  HelperSide = "c"
)

type Helpers []*Helper

type Helper struct {
	Name        string
	Side        HelperSide
	Description string
	Source      string
	ShouldFree  bool
	Requires    []*Helper
}

var cProxyHelperMap = map[string]*Helper{
	tl.StringSpec.String(): cStringFunc,
}

var goProxyHelperMap = map[string]*Helper{
	tl.StringSpec.String(): goStringFunc,
}

type proxyDecl struct {
	Name string
	Decl string
}

func (gen *Generator) submitHelper(h *Helper) {
	if h == nil {
		return
	}
	gen.helpersChan <- h
	reqs := h.Requires
	for len(reqs) > 0 {
		var newReqs Helpers
		for _, req := range reqs {
			gen.helpersChan <- req
			newReqs = append(newReqs, req.Requires...)
		}
		reqs = newReqs
	}
}

func (gen *Generator) createCProxy(pName []byte, decl tl.CDecl) (pDecl string, free bool) {
	refType := gen.tr.TranslateSpec(decl.Spec)
	if helper, ok := cProxyHelperMap[refType.String()]; ok {
		gen.submitHelper(helper)
		pDecl = fmt.Sprintf("%s(%s)", helper.Name, pName)
		return pDecl, helper.ShouldFree
	}
	cgoSpec := gen.tr.CGoSpec(decl.Spec)
	var ref = string(pName)
	if refType.Slices > 0 {
		ref = fmt.Sprintf("%s[0]", pName)
		pDecl = fmt.Sprintf("(%s)(unsafe.Pointer(&%s))", cgoSpec, ref)
		return
	}
	pDecl = fmt.Sprintf("(%s)(unsafe.Pointer(%s))", cgoSpec, ref)
	return
}

func (gen *Generator) createGoProxy(pName []byte, decl tl.CDecl) (pDecl string, free bool) {
	refType := gen.tr.TranslateSpec(decl.Spec)
	if helper, ok := goProxyHelperMap[refType.String()]; ok {
		gen.submitHelper(helper)
		pDecl = fmt.Sprintf("%s(%s)", helper.Name, pName)
		return pDecl, helper.ShouldFree
	}
	pDecl = fmt.Sprintf("*(*%s)(unsafe.Pointer(%s))", refType, pName)
	return
}

func (gen *Generator) createProxies(decl tl.CDecl) []proxyDecl {
	spec := decl.Spec.(*tl.CFunctionSpec)
	proxies := make([]proxyDecl, len(spec.ParamList))
	for i, param := range spec.ParamList {
		refName := gen.tr.TransformName(tl.TargetPrivate, param.Name)
		proxies[i].Name = "c" + string(refName)
		proxies[i].Decl, _ = gen.createCProxy(refName, param)
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
	if spec.Return != nil {
		retProxyDecl, free := gen.createGoProxy([]byte("ret"), *spec.Return)
		if free {
			fmt.Fprintln(wr, "C.free(ret)")
		}
		fmt.Fprintln(wr, "return "+retProxyDecl)
	}
	writeEndFuncBody(wr)
}

var (
	cStringFunc = &Helper{
		Name:        "cStr",
		Description: "cStr gets data from Go string as *C.char and avoids copying.",
		Source: `func cStr(str string) *C.char {
			h := (*reflect.StringHeader)(unsafe.Pointer(&str))
			return (*C.char)(unsafe.Pointer(h.Data))
		}`,
	}
	goStringFunc = &Helper{
		Name:        "goStr",
		Description: "goStr creates a string backed by *C.char and avoids copying.",
		Source: `func goStr(p *C.char) (raw string) {
			if p != nil && *p != 0 {
				h := (*reflect.StringHeader)(unsafe.Pointer(&raw))
				h.Data = uintptr(unsafe.Pointer(p))
				for *p != 0 {
					p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
				}
				h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
			}
			return
		}`,
		Requires: Helpers{rawString},
	}
	rawString = &Helper{
		Name:        "RawString",
		Description: "RawString reperesents a string backed by data on the C side.",
		Source:      `type RawString string`,
		Requires:    Helpers{rawStringCopy},
	}
	rawStringCopy = &Helper{
		Name:        "RawString.Copy",
		Description: "Copy returns a Go-managed copy of raw string.",
		Source: `func (raw RawString) Copy() string {
			if len(raw) == 0 {
				return ""
			}
			h := (*reflect.StringHeader)(unsafe.Pointer(&raw))
			return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
		}`,
	}
)
