package generator

import (
	"bytes"
	"fmt"
	"io"
	"strings"

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

var fromGoHelperMap = map[string]*Helper{
	"string": cStringFunc,
}

var toGoHelperMap = map[string]*Helper{
	"string": goStringFunc,
}

type proxyDecl struct {
	Name string
	Decl string
}

// getHelperName transforms signatures like [4][4][]*string into A4A4SPString
// suitable for the generated proxy helper methods.
func getHelperName(gospec tl.GoTypeSpec) string {
	var name string
	for _, size := range gospec.Arrays {
		name += fmt.Sprintf("%sA%d", name, size)
	}
	name += strings.Repeat("S", int(gospec.Slices))
	name += strings.Repeat("P", int(gospec.Pointers))
	if gospec.Inner != nil {
		name += getHelperName(*gospec.Inner)
	} else {
		name += strings.Title(gospec.Base)
	}
	return name
}

func getCGoSpecHelperName(spec *tl.CGoSpec) string {
	base := strings.Replace(spec.Base, ".", " ", -1)
	base = strings.Title(strings.Replace(base, "_", " ", -1))
	base = strings.Replace(base, " ", "", -1)
	return strings.Repeat("P", int(spec.Pointers)) + base
}

func isPlainType(gospec tl.GoTypeSpec) bool {
	if gospec.Slices > 0 {
		return false
	}
	if gospec.Inner != nil {
		return isPlainType(gospec)
	}
	switch gospec.Base {
	case "int", "byte", "rune", "float32", "float64":
		return true
	case "string":
		return false
	}
	return false
}

type level struct {
	Array bool
	Slice bool
	Last  bool
}

type reverseBuffer []string

func (r *reverseBuffer) Linef(f string, a ...interface{}) {
	*r = append(*r, fmt.Sprintf(f, a...))
}

func (r reverseBuffer) String() string {
	buf := new(bytes.Buffer)
	for i := len(r) - 1; i >= 0; i-- {
		buf.WriteString(r[i])
	}
	return buf.String()
}
func (r reverseBuffer) WriteTo(w io.Writer) {
	buf := new(bytes.Buffer)
	for i := len(r) - 1; i >= 0; i-- {
		buf.WriteString(r[i])
	}
	buf.WriteTo(w)
}

func ptrs(n uint8) string {
	return strings.Repeat("*", int(n))
}

func (gen *Generator) getUnpackHelper(gospec tl.GoTypeSpec, spec *tl.CTypeSpec) *Helper {
	cgospec := gen.tr.CGoSpec(spec)
	name := "unpack" + getHelperName(gospec)
	code := new(bytes.Buffer)
	fmt.Fprintf(code, "func %s(x %s) %s {\n", name, gospec, cgospec)
	h := &Helper{
		Name:        name,
		Description: "TODO",
	}

	var level uint8
	buf1 := new(bytes.Buffer)
	buf2 := new(reverseBuffer)
	var walkSpec func(gospec tl.GoTypeSpec)
	walkSpec = func(gospec tl.GoTypeSpec) {
		for _, size := range gospec.Arrays {
			if level == 0 {
				fmt.Fprintf(buf1, "mem0 := make([]%s%s, %d)\n", ptrs(cgospec.Pointers-level-1), cgospec.Base, size)
				fmt.Fprintf(buf1, "for i0, x0 := range x {\n")
				buf2.Linef("return (%s)(unsafe.Pointer(&mem0[0]))", cgospec)
				buf2.Linef("}\n")
			} else {
				fmt.Fprintf(buf1, "mem%d := make([]%s%s, %d)\n", level,
					ptrs(cgospec.Pointers-level-1), cgospec.Base, size)
				fmt.Fprintf(buf1, "for i%d, x%d := range x%d {\n", level, level, level-1)
				buf2.Linef("mem%d[i%d] = (%s%s)(unsafe.Pointer(&mem%d[0]))\n",
					level-1, level-1, ptrs(cgospec.Pointers-level), cgospec.Base, level)
				buf2.Linef("}\n")
			}
			level++
		}
		for i := uint8(0); i < gospec.Slices; i++ {
			if level == 0 {
				fmt.Fprintf(buf1, "mem0 := make([]%s%s, len(x))\n", ptrs(cgospec.Pointers-level-1), cgospec.Base)
				fmt.Fprintf(buf1, "for i0, x0 := range x {\n")
				buf2.Linef("return (%s)(unsafe.Pointer(&mem0[0]))\n", cgospec)
				buf2.Linef("}\n")
			} else {
				fmt.Fprintf(buf1, "mem%d := make([]%s%s, len(x%d))\n", level,
					ptrs(cgospec.Pointers-level-1), cgospec.Base, level-1)
				fmt.Fprintf(buf1, "for i%d, x%d := range x%d {\n", level, level, level-1)
				buf2.Linef("mem%d[i%d] = (%s%s)(unsafe.Pointer(&mem%d[0]))\n",
					level-1, level-1, ptrs(cgospec.Pointers-level), cgospec.Base, level)
				buf2.Linef("}\n")
			}
			level++
		}
		if gospec.Inner == nil {
			return
		}
		if isPlainType(*gospec.Inner) {
			level = level - 1
			ptr := strings.Repeat("*", int(cgospec.Pointers-level))
			ref := "x"
			if gospec.Pointers == 0 {
				ref = "&x"
			}
			fmt.Fprintf(buf1, "mem%d[i%d] = (%s%s)(unsafe.Pointer(%s%d))\n", level, level, ptr, cgospec.Base, ref, level)
		} else if gospec.Inner.Slices > 0 || len(gospec.Inner.Arrays) > 0 {
			walkSpec(*gospec.Inner)
		} else if helper, ok := fromGoHelperMap[gospec.Inner.Base]; ok {
			level = level - 1
			h.Requires = append(h.Requires, helper)
			deref := strings.Repeat("*", int(gospec.Inner.Pointers))
			fmt.Fprintf(buf1, "mem%d[i%d] = %s(%sx%d)\n", level, level, helper.Name, deref, level)
		} else {
			level = level - 1
			fmt.Fprintf(buf1, "if x%d == nil {\ncontinue\n}\n", level)
			fmt.Fprintf(buf1, "if ref := x%d.Ref(); ref != nil{\nmem%d[i%d] = ref\n}", level, level, level)
			fmt.Fprintf(buf1, " else {\nmem%d[i%d] = x%d.NewRef()\n}\n", level, level, level)
		}
	}
	walkSpec(gospec)
	buf1.WriteTo(code)
	buf2.WriteTo(code)
	fmt.Fprintln(code, "}")
	h.Source = code.String()
	return h
}

func (gen *Generator) proxyFromGo(name string, decl tl.CDecl) (proxy string, free bool) {
	goSpec := gen.tr.TranslateSpec(decl.Spec)
	if helper, ok := fromGoHelperMap[goSpec.String()]; ok {
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s(%s)", helper.Name, name)
		return proxy, helper.ShouldFree
	}
	cgoSpec := gen.tr.CGoSpec(decl.Spec)
	if isPlainType(goSpec) {
		ref := name
		if goSpec.Pointers == 0 {
			ref = "&" + name
		}
		proxy = fmt.Sprintf("(%s)(unsafe.Pointer(%s))", cgoSpec, ref)
		return
	}
	switch spec := decl.Spec.(type) {
	case *tl.CTypeSpec:
		if goSpec.Slices > 1 || len(goSpec.Arrays) > 0 {
			helper := gen.getUnpackHelper(goSpec, spec)
			gen.submitHelper(helper)
			proxy = fmt.Sprintf("%s(%s)", helper.Name, name)
			return proxy, helper.ShouldFree
		}
		ref := name
		if goSpec.Slices > 0 {
			ref = fmt.Sprintf("&%s[0]", name)
		} else if goSpec.Pointers == 0 {
			ref = fmt.Sprintf("&%s", name)
		}
		proxy = fmt.Sprintf("(%s)(unsafe.Pointer(%s))", cgoSpec, ref)
		return
	default:
		proxy = fmt.Sprintf("%s.ToC())", cgoSpec, name)
		return
	}
}

func (gen *Generator) createGoProxy(pName []byte, decl tl.CDecl) (pDecl string, free bool) {
	refType := gen.tr.TranslateSpec(decl.Spec)
	if helper, ok := toGoHelperMap[refType.String()]; ok {
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
		refName := string(gen.tr.TransformName(tl.TargetPrivate, param.Name))
		proxies[i].Name = "c" + refName
		proxies[i].Decl, _ = gen.proxyFromGo(refName, param)
	}
	return proxies
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
	cStringSliceFunc = &Helper{
		Name:        "cStrSlice",
		Description: "cStrSlice represents the data from a lice of Go strings as **C.Char.",
		Source: `func cStrSlice(s []string) **C.char {
			mem := make([]*C.char, len(s))
			for i, str := range s {
				mem[i] = cStr(str)
			}
			return (**C.char)(unsafe.Pointer(&mem[0]))
		}`,
		Requires: Helpers{cStringFunc},
	}
	cStringFunc = &Helper{
		Name:        "cStr",
		Description: "cStr represents the data from Go string as *C.char and avoids copying.",
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
