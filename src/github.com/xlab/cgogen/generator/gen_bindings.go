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
	Nillable    bool
	Requires    []*Helper
}

var fromGoHelperMap = map[string]*Helper{
	"string": cStringFunc,
}

var toGoHelperMap = map[string]*Helper{
	"string":  goStringFunc,
	"stringN": goStringNFunc,
}

type proxyDecl struct {
	Name string
	Decl string
}

// getHelperName transforms signatures like [4][4][]*string into A4A4SPString
// suitable for the generated proxy helper methods.
func getHelperName(gospec tl.GoTypeSpec) string {
	buf := new(bytes.Buffer)
	for _, size := range gospec.Arrays {
		fmt.Fprintf(buf, "A%d", size)
	}
	buf.WriteString(strings.Repeat("S", int(gospec.Slices)))
	buf.WriteString(strings.Repeat("P", int(gospec.Pointers)))
	buf.WriteString(strings.Title(gospec.Base))
	return buf.String()
}

func getCGoSpecHelperName(spec *tl.CGoSpec) string {
	base := strings.Replace(spec.Base, ".", " ", -1)
	base = strings.Title(strings.Replace(base, "_", " ", -1))
	base = strings.Replace(base, " ", "", -1)
	return strings.Repeat("P", int(spec.Pointers)) + base
}

func isPlainType(base string) bool {
	switch base {
	case "int", "byte", "rune", "float32", "float64", "void":
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

func unpackPlain(buf io.Writer, base string, cptr, gptr uint8, level uint8) {
	ref := "x"
	if gptr == 0 {
		ref = "&x"
	}
	if cptr > 0 {
		fmt.Fprintf(buf, "mem%d[i%d] = (%s%s)(unsafe.Pointer(%s%d))\n",
			level, level, ptrs(cptr-level), base, ref, level)
		return
	}
	fmt.Fprintf(buf, "mem%d[i%d] = *(*%s)(unsafe.Pointer(%s%d))\n",
		level, level, base, ref, level)
}

func unpackPlainSlice(buf io.Writer, base string, pointers uint8, level uint8) {
	fmt.Fprintf(buf, "mem%d[i%d] = (%s%s)(unsafe.Pointer(&x%d[0]))\n",
		level, level, ptrs(pointers-level), base, level)
}

func unpackObj(buf io.Writer, base string, pointers uint8, level uint8) *Helper {
	if helper, ok := fromGoHelperMap[base]; ok {
		fmt.Fprintf(buf, "mem%d[i%d] = %s(%sx%d)\n", level, level, helper.Name, ptrs(pointers), level)
		return helper
	}
	fmt.Fprintf(buf, "if x%d == nil {\ncontinue\n}\n", level)
	fmt.Fprintf(buf, "mem%d[i%d] = x%d.Ref()\n", level, level, level)
	return nil
}

func unpackArray(buf1 io.Writer, buf2 *reverseBuffer, size uint64, base string, pointers uint8, level uint8) {
	if level == 0 {
		fmt.Fprintf(buf1, "var mem0 [%d]%s%s\n", size, ptrs(pointers-level-1), base)
		fmt.Fprintf(buf1, "for i0, x0 := range x {\n")
		buf2.Linef("return (%s%s)(unsafe.Pointer(&mem0))", ptrs(pointers), base)
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "var mem%d [%d]%s%s\n", level, size, ptrs(pointers-level-1), base)
	fmt.Fprintf(buf1, "for i%d, x%d := range x%d {\n", level, level, level-1)
	buf2.Linef("mem%d[i%d] = (%s%s)(unsafe.Pointer(&mem%d[0]))\n",
		level-1, level-1, ptrs(pointers-level), base, level)
	buf2.Linef("}\n")
}

func unpackSlice(buf1 io.Writer, buf2 *reverseBuffer, base string, pointers uint8, level uint8) {
	if level == 0 {
		fmt.Fprintf(buf1, "mem0 := make([]%s%s, len(x))\n", ptrs(pointers-level-1), base)
		fmt.Fprintf(buf1, "for i0, x0 := range x {\n")
		buf2.Linef("return (%s%s)(unsafe.Pointer(&mem0[0]))\n", ptrs(pointers), base)
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "mem%d := make([]%s%s, len(x%d))\n", level, ptrs(pointers-level-1), base, level-1)
	fmt.Fprintf(buf1, "for i%d, x%d := range x%d {\n", level, level, level-1)
	buf2.Linef("mem%d[i%d] = (%s%s)(unsafe.Pointer(&mem%d[0]))\n",
		level-1, level-1, ptrs(pointers-level), base, level)
	buf2.Linef("}\n")
}

func (gen *Generator) getUnpackHelper(gospec tl.GoTypeSpec, spec tl.CType) *Helper {
	cgospec := gen.tr.CGoSpec(spec)
	name := "unpack" + getHelperName(gospec)
	code := new(bytes.Buffer)
	fmt.Fprintf(code, "func %s(x %s) %s {\n", name, gospec, cgospec)
	h := &Helper{
		Name:        name,
		Description: fmt.Sprintf("%s transforms a sliced Go data structure into plain C format.", name),
	}
	var level uint8
	buf1 := new(bytes.Buffer)
	buf2 := new(reverseBuffer)

	for _, size := range gospec.Arrays {
		unpackArray(buf1, buf2, size, cgospec.Base, cgospec.Pointers, level)
		level++
	}
	gospec.Arrays = nil

	for gospec.Slices > 1 {
		gospec.Slices--
		unpackSlice(buf1, buf2, cgospec.Base, cgospec.Pointers, level)
		level++
	}
	isSlice := gospec.Slices > 0
	isPlain := isPlainType(gospec.Base)

	switch {
	case isPlain && isSlice:
		unpackPlainSlice(buf1, cgospec.Base, cgospec.Pointers, level-1)
	case isPlain:
		unpackPlain(buf1, gospec.Base, cgospec.Pointers, gospec.Pointers, level-1)
	case isSlice:
		unpackSlice(buf1, buf2, cgospec.Base, cgospec.Pointers, level)
		if helper := unpackObj(buf1, gospec.Base, gospec.Pointers, level-1); helper != nil {
			h.Requires = append(h.Requires, helper)
		}
	default:
		if helper := unpackObj(buf1, gospec.Base, gospec.Pointers, level-1); helper != nil {
			h.Requires = append(h.Requires, helper)
		}
	}

	buf1.WriteTo(code)
	buf2.WriteTo(code)
	fmt.Fprintln(code, "}")
	h.Source = code.String()
	return h
}

func (gen *Generator) proxyFromGo(name string, decl tl.CDecl) (proxy string, nillable bool) {
	goSpec := gen.tr.TranslateSpec(decl.Spec)
	if helper, ok := fromGoHelperMap[goSpec.String()]; ok {
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s(%s)", helper.Name, name)
		return proxy, helper.Nillable
	}

	cgoSpec := gen.tr.CGoSpec(decl.Spec)
	isPlain := isPlainType(goSpec.Base)
	switch {
	case !isPlain && (goSpec.Slices > 0 || len(goSpec.Arrays) > 0), // ex: []string
		isPlain && goSpec.Slices > 0 && len(goSpec.Arrays) > 0, // ex: [4][]byte
		isPlain && goSpec.Slices > 1:                           // ex: [][]byte
		helper := gen.getUnpackHelper(goSpec, decl.Spec)
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s(%s)", helper.Name, name)
		return proxy, helper.Nillable
	case isPlain && goSpec.Slices > 0: // ex: []byte
		proxy = fmt.Sprintf("(%s)(unsafe.Pointer(&%s[0]))", cgoSpec, name)
		return
	case isPlain: // ex: byte, [4]byte
		if len(goSpec.Arrays) == 0 {
			proxy = fmt.Sprintf("(%s)(%s)", cgoSpec, name)
			return
		}
		ref := name
		if goSpec.Pointers == 0 {
			ref = "&" + name
		}
		proxy = fmt.Sprintf("(%s)(unsafe.Pointer(%s))", cgoSpec, ref)
		return
	default: // ex: *SomeType
		nillable = true
		proxy = fmt.Sprintf("%s.PassRef()", name)
		return
	}
}

func packPlain(buf io.Writer, base string, cptr, gptr uint8, level uint8) {
	fmt.Fprintf(buf, "mem%d[i%d] = (%s%s)(unsafe.Pointer(%sx%d))\n",
		level, level, ptrs(gptr-level), base, ptrs(cptr-level), level)
}

func packPlainSlice(buf io.Writer, base string, pointers uint8, level uint8) {
	fmt.Fprintf(buf, "mem%d[i%d] = (%s%s)(unsafe.Pointer(&x%d[0]))\n",
		level, level, ptrs(pointers-level), base, level)
}

func packObj(buf io.Writer, base string, pointers uint8, level uint8) *Helper {
	if helper, ok := toGoHelperMap[base]; ok {
		fmt.Fprintf(buf, "mem%d[i%d] = %s(%sptr%d)\n",
			level, level, helper.Name, ptrs(pointers), level)
		return helper
	}
	fmt.Fprintf(buf, "if x%d == nil {\ncontinue\n}\n", level)
	fmt.Fprintf(buf, "mem%d[i%d] = x%d.PassRef()\n", level, level, level)
	return nil
}

func packArray(buf1 io.Writer, buf2 *reverseBuffer, size uint64, base string, pointers uint8, level uint8, last bool) {
	if level == 0 {
		fmt.Fprintln(buf1, "for i0, mem1 := range mem0 {")
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "ptr%d := (*(*[%d]%s%s)(unsafe.Pointer(ptr%d)))[i%d]\n",
		size, level, ptrs(pointers-level-1), base, level-1, level-1)
	if last {
		fmt.Fprintf(buf1, "for i%d := range mem%d {\n", level, level+1)
	} else {
		fmt.Fprintf(buf1, "for i%d, mem%d := range mem%d {\n", level, level+1, level-1)
	}
	buf2.Linef("}\n")
}

func packSlice(buf1 io.Writer, buf2 *reverseBuffer, base string, pointers uint8, level uint8, last bool) {
	if level == 0 {
		fmt.Fprintln(buf1, "const m = 1 << 30")
		fmt.Fprintln(buf1, "for i0, mem1 := range mem0 {")
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "ptr%d := (*(*[m]%s%s)(unsafe.Pointer(ptr%d)))[i%d]\n",
		level, ptrs(pointers-level-1), base, level-1, level-1)
	if last {
		fmt.Fprintf(buf1, "for i%d := range mem%d {\n", level, level+1)
	} else {
		fmt.Fprintf(buf1, "for i%d, mem%d := range mem%d {\n", level, level+1, level)
	}
	buf2.Linef("}\n")
}

func (gen *Generator) getPackHelper(gospec tl.GoTypeSpec, spec tl.CType) *Helper {
	cgospec := gen.tr.CGoSpec(spec)
	name := "pack" + getHelperName(gospec)
	code := new(bytes.Buffer)
	fmt.Fprintf(code, "func %s(mem0 %s, ptr0 %s) {\n", name, gospec, cgospec)
	h := &Helper{
		Name:        name,
		Description: fmt.Sprintf("%s reads sliced Go data structure out from plain C format.", name),
	}
	var level uint8
	buf1 := new(bytes.Buffer)
	buf2 := new(reverseBuffer)

	for i, size := range gospec.Arrays {
		last := i >= len(gospec.Arrays)-1
		packArray(buf1, buf2, size, cgospec.Base, cgospec.Pointers, level, last)
		level++
	}
	gospec.Arrays = nil

	for gospec.Slices > 1 {
		gospec.Slices--
		last := gospec.Slices == 0
		packSlice(buf1, buf2, cgospec.Base, cgospec.Pointers, level, last)
		level++
	}
	isSlice := gospec.Slices > 0
	isPlain := isPlainType(gospec.Base)

	switch {
	case isPlain && isSlice:
		packPlainSlice(buf1, cgospec.Base, cgospec.Pointers, level-1)
	case isPlain:
		packPlain(buf1, gospec.Base, cgospec.Pointers, gospec.Pointers, level-1)
	case isSlice:
		packSlice(buf1, buf2, cgospec.Base, cgospec.Pointers, level, true)
		if helper := packObj(buf1, gospec.Base, gospec.Pointers, level); helper != nil {
			h.Requires = append(h.Requires, helper)
		}
	default:
		if helper := packObj(buf1, gospec.Base, gospec.Pointers, level); helper != nil {
			h.Requires = append(h.Requires, helper)
		}
	}

	buf1.WriteTo(code)
	buf2.WriteTo(code)
	fmt.Fprintln(code, "}")
	h.Source = code.String()
	return h
}

func (gen *Generator) proxyToGo(name string, decl tl.CDecl) (proxy string, nillable bool) {
	goSpec := gen.tr.TranslateSpec(decl.Spec)
	nillable = true

	if helper, ok := toGoHelperMap[goSpec.String()]; ok {
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s(%s)", helper.Name, name)
		return proxy, helper.Nillable
	}

	cgoSpec := gen.tr.CGoSpec(decl.Spec)
	isPlain := isPlainType(goSpec.Base)
	switch {
	case !isPlain && (goSpec.Slices > 0 || len(goSpec.Arrays) > 0), // ex: []string
		isPlain && goSpec.Slices > 0 && len(goSpec.Arrays) > 0, // ex: [4][]byte
		isPlain && goSpec.Slices > 1:                           // ex: [][]byte
		helper := gen.getPackHelper(goSpec, decl.Spec)
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s(%s)", helper.Name, name)
		return proxy, helper.Nillable
	case isPlain && goSpec.Slices > 0: // ex: []byte
		proxy = fmt.Sprintf("(*(*[1<<30]%s%s)(unsafe.Pointer(%s)))[:0]",
			ptrs(goSpec.Pointers), goSpec.Base, name)
		return
	case isPlain: // ex: byte, [4]byte
		nillable = false
		if len(goSpec.Arrays) == 0 {
			proxy = fmt.Sprintf("(%s)(%s)", goSpec, name)
			return
		}
		ref := name
		if goSpec.Pointers == 0 {
			ref = "&" + name
		}
		var deref string
		if cgoSpec.Pointers == 0 {
			deref = "*"
		}
		proxy = fmt.Sprintf("%s(%s%s)(unsafe.Pointer(%s))", deref, deref, goSpec, ref)
		return
	default: // ex: *SomeType
		proxy = fmt.Sprintf("New%s(%s)", goSpec.Base, name)
		return
	}
}

func (gen *Generator) createProxies(decl tl.CDecl) []proxyDecl {
	spec := decl.Spec.(*tl.CFunctionSpec)
	proxies := make([]proxyDecl, len(spec.ParamList))
	for i, param := range spec.ParamList {
		buf := new(bytes.Buffer)
		refName := string(gen.tr.TransformName(tl.TargetPrivate, param.Name))
		name := "c" + refName
		proxy, nillable := gen.proxyFromGo(refName, param)
		if nillable {
			fmt.Fprintf(buf, "var %s %s\n", name, gen.tr.CGoSpec(param.Spec))
			fmt.Fprintf(buf, "if %s != nil {\n%s = %s\n}", refName, name, proxy)
		} else {
			fmt.Fprintf(buf, "%s := %s", name, proxy)
		}
		proxies[i] = proxyDecl{Name: name, Decl: buf.String()}
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
	for _, proxy := range proxies {
		fmt.Fprintln(wr, proxy.Decl)
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
		retProxyDecl, nillable := gen.proxyToGo("ret", *spec.Return)
		if nillable {
			fmt.Fprintln(wr, "if ret == nil {\nreturn nil\n}")
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
	goStringNFunc = &Helper{
		Name:        "goStrN",
		Description: "goStrN creates a string backed by *C.char (of max length N) and avoids copying.",
		Source: `func goStr(p *C.char, n int) (raw string) {
			if p != nil && *p != 0 && n > 0 {
				h := (*reflect.StringHeader)(unsafe.Pointer(&raw))
				h.Data = uintptr(unsafe.Pointer(p))
				for *p != 0 && n > 0 {
					n--
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
