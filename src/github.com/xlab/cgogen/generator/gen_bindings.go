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

type getHelperFunc func(gen *Generator, spec tl.CGoSpec) *Helper

var fromGoHelperMap = map[tl.GoTypeSpec]getHelperFunc{
	tl.StringSpec:  (*Generator).getUnpackStringHelper,
	tl.UStringSpec: (*Generator).getUnpackStringHelper,
}

var toGoHelperMap = map[tl.GoTypeSpec]*Helper{
	tl.StringSpec:  goStringFunc,
	tl.UStringSpec: goUStringFunc,
}

type proxyDecl struct {
	Name string
	Decl string
}

// getHelperName transforms signatures like [4][4][]*string into A4A4SPString
// suitable for the generated proxy helper methods.
func getHelperName(gospec tl.GoTypeSpec) string {
	buf := new(bytes.Buffer)
	for _, size := range tl.GetArraySizes(gospec.Arrays) {
		fmt.Fprintf(buf, "A%d", size)
	}
	buf.WriteString(strings.Repeat("S", int(gospec.Slices)))
	buf.WriteString(strings.Repeat("P", int(gospec.Pointers)))
	if gospec.Unsigned {
		buf.WriteRune('U')
	}
	buf.WriteString(strings.Title(gospec.Base))
	return buf.String()
}

func (gen *Generator) getTypedHelperName(goBase string, spec tl.CGoSpec) string {
	if strings.HasPrefix(spec.Base, "C.") {
		spec.Base = spec.Base[2:]
	}
	specName := gen.tr.TransformName(tl.TargetType, spec.Base)
	buf := new(bytes.Buffer)
	for _, size := range spec.Arrays {
		fmt.Fprintf(buf, "A%d", size)
	}
	buf.WriteString(strings.Repeat("P", int(spec.Pointers)))
	buf.Write(bytes.Title(specName))
	buf.WriteString(strings.Title(goBase))
	return buf.String()
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

func unpackPlain(buf io.Writer, cgoSpec tl.CGoSpec, pointers uint8, level uint8) {
	uplevel := level - 1
	ref := "x"
	if pointers == 0 {
		ref = "&x"
	}
	if cgoSpec.Pointers > 0 {
		fmt.Fprintf(buf, "mem%d[i%d] = (%s)(unsafe.Pointer(%s%d))\n",
			uplevel, uplevel, cgoSpec.AtLevel(level), ref, level)
		return
	}
	fmt.Fprintf(buf, "mem%d[i%d] = *(*%s)(unsafe.Pointer(%s%d))\n",
		uplevel, uplevel, cgoSpec.Base, ref, level)
}

func unpackPlainSlice(buf io.Writer, cgoSpec tl.CGoSpec, level uint8) {
	uplevel := level - 1
	fmt.Fprintf(buf, "mem%d[i%d] = (%s)(unsafe.Pointer(&x%d[0]))\n",
		uplevel, uplevel, cgoSpec.AtLevel(level), uplevel)
}

func (gen *Generator) unpackObj(buf io.Writer, cgoSpec tl.CGoSpec, goSpec tl.GoTypeSpec, level uint8) *Helper {
	uplevel := level - 1
	if getHelperFunc, ok := fromGoHelperMap[goSpec]; ok {
		helper := getHelperFunc(gen, cgoSpec)
		fmt.Fprintf(buf, "mem%d[i%d] = %s(%sx%d)\n",
			uplevel, uplevel, helper.Name, ptrs(goSpec.Pointers), uplevel)
		return helper
	}
	fmt.Fprintf(buf, "if x%d == nil {\ncontinue\n}\n", uplevel)
	fmt.Fprintf(buf, "mem%d[i%d] = x%d.Ref()\n", uplevel, uplevel, uplevel)
	return nil
}

func unpackArray(buf1 io.Writer, buf2 *reverseBuffer, size uint64, cgoSpec tl.CGoSpec, level uint8) {
	uplevel := level - 1
	if level == 0 {
		fmt.Fprintf(buf1, "var mem0 %s\n", cgoSpec)
		fmt.Fprintf(buf1, "for i0, x0 := range x {\n")
		buf2.Linef("return (%s)(unsafe.Pointer(&mem0))", cgoSpec.AtLevel(level))
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "var mem%d %s\n", level, cgoSpec.AtLevel(level))
	fmt.Fprintf(buf1, "for i%d, x%d := range x%d {\n", level, level, uplevel)
	buf2.Linef("mem%d[i%d] = *(*%s)(unsafe.Pointer(&mem%d))\n",
		uplevel, uplevel, cgoSpec.AtLevel(level), level)
	buf2.Linef("}\n")
}

func unpackSlice(buf1 io.Writer, buf2 *reverseBuffer, cgoSpec tl.CGoSpec, level uint8) {
	uplevel := level - 1
	if level == 0 {
		fmt.Fprintf(buf1, "mem0 := make([]%s, len(x))\n", cgoSpec.AtLevel(1))
		fmt.Fprintf(buf1, "for i0, x0 := range x {\n")
		buf2.Linef("return (%s)(unsafe.Pointer(&mem0[0]))\n", cgoSpec.AtLevel(0))
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "mem%d := make([]%s, len(x%d))\n", level, cgoSpec.AtLevel(level+1), uplevel)
	fmt.Fprintf(buf1, "for i%d, x%d := range x%d {\n", level, level, uplevel)
	buf2.Linef("mem%d[i%d] = (%s)(unsafe.Pointer(&mem%d[0]))\n", uplevel, uplevel, cgoSpec.AtLevel(level), level)
	buf2.Linef("}\n")
}

func (gen *Generator) getUnpackStringHelper(cgoSpec tl.CGoSpec) *Helper {
	cgoSpec = tl.CGoSpec{
		Pointers: 1,
		Base:     cgoSpec.Base,
	}
	name := "unpack" + gen.getTypedHelperName("string", cgoSpec)
	return &Helper{
		Name:        name,
		Description: fmt.Sprintf("%s represents the data from Go string as %s and avoids copying.", name, cgoSpec),
		Source: fmt.Sprintf(`func %s(str string) %s {
			h := (*reflect.StringHeader)(unsafe.Pointer(&str))
			return (%s)(unsafe.Pointer(h.Data))
		}`, name, cgoSpec, cgoSpec),
	}
}

func (gen *Generator) getUnpackHelper(cgoSpec tl.CGoSpec, goSpec tl.GoTypeSpec) *Helper {
	name := "unpack" + getHelperName(goSpec)
	code := new(bytes.Buffer)
	fmt.Fprintf(code, "func %s(x %s) %s {\n", name, goSpec, cgoSpec.AtLevel(0))
	h := &Helper{
		Name:        name,
		Description: fmt.Sprintf("%s transforms a sliced Go data structure into plain C format.", name),
	}
	var level uint8
	buf1 := new(bytes.Buffer)
	buf2 := new(reverseBuffer)

	for _, size := range tl.GetArraySizes(goSpec.Arrays) {
		unpackArray(buf1, buf2, size, cgoSpec, level)
		level++
	}
	goSpec.Arrays = ""

	for goSpec.Slices > 1 {
		goSpec.Slices--
		unpackSlice(buf1, buf2, cgoSpec, level)
		level++
	}
	isSlice := goSpec.Slices > 0
	isPlain := isPlainType(goSpec.Base)

	switch {
	case isPlain && isSlice:
		unpackPlainSlice(buf1, cgoSpec, level)
	case isPlain:
		unpackPlain(buf1, cgoSpec, goSpec.Pointers, level)
	case isSlice:
		unpackSlice(buf1, buf2, cgoSpec, level)
		goSpec.Slices = 0
		if helper := gen.unpackObj(buf1, cgoSpec, goSpec, level+1); helper != nil {
			h.Requires = append(h.Requires, helper)
		}
	default:
		if helper := gen.unpackObj(buf1, cgoSpec, goSpec, level); helper != nil {
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
	cgoSpec := gen.tr.CGoSpec(decl.Spec)
	if getHelperFunc, ok := fromGoHelperMap[goSpec]; ok {
		helper := getHelperFunc(gen, cgoSpec)
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s(%s)", helper.Name, name)
		return proxy, helper.Nillable
	}

	isPlain := isPlainType(goSpec.Base)
	switch {
	case !isPlain && (goSpec.Slices > 0 || len(goSpec.Arrays) > 0), // ex: []string
		isPlain && goSpec.Slices > 0 && len(goSpec.Arrays) > 0, // ex: [4][]byte
		isPlain && goSpec.Slices > 1:                           // ex: [][]byte
		helper := gen.getUnpackHelper(cgoSpec, goSpec)
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s(%s)", helper.Name, name)
		return proxy, helper.Nillable
	case isPlain && goSpec.Slices > 0: // ex: []byte
		proxy = fmt.Sprintf("(%s)(unsafe.Pointer(&%s[0]))", cgoSpec.AtLevel(0), name)
		return
	case isPlain: // ex: byte, [4]byte
		if len(goSpec.Arrays) == 0 {
			proxy = fmt.Sprintf("(%s)(%s)", cgoSpec.AtLevel(0), name)
			return
		}
		ref := name
		if goSpec.Pointers == 0 {
			ref = "&" + name
		}
		proxy = fmt.Sprintf("(%s)(unsafe.Pointer(%s))", cgoSpec.AtLevel(0), ref)
		return
	default: // ex: *SomeType
		nillable = true
		proxy = fmt.Sprintf("%s.PassRef()", name)
		return
	}
}

func packPlain(buf io.Writer, cgoSpec tl.CGoSpec, base string, pointers uint8, level uint8) {
	uplevel := level - 1
	fmt.Fprintf(buf, "mem%d[i%d] = (%s%s)(unsafe.Pointer(%sx%d))\n",
		uplevel, uplevel, ptrs(pointers-uplevel), base, ptrs(cgoSpec.Pointers-level), level)
}

func packPlainSlice(buf io.Writer, base string, pointers uint8, level uint8) {
	uplevel := level - 1
	fmt.Fprintf(buf, "mem%d[i%d] = (%s%s)(unsafe.Pointer(&x%d[0]))\n",
		uplevel, uplevel, ptrs(pointers-uplevel), base, level)
}

func packObj(buf io.Writer, goSpec tl.GoTypeSpec, level uint8) *Helper {
	uplevel := level - 1
	if helper, ok := toGoHelperMap[goSpec]; ok {
		fmt.Fprintf(buf, "mem%d[i%d] = %s(%sptr%d)\n",
			uplevel, uplevel, helper.Name, ptrs(goSpec.Pointers), level)
		return helper
	}
	fmt.Fprintf(buf, "if x%d == nil {\ncontinue\n}\n", level)
	fmt.Fprintf(buf, "mem%d[i%d] = x%d.PassRef()\n", uplevel, uplevel, level)
	return nil
}

func packArray(buf1 io.Writer, buf2 *reverseBuffer, size uint64, cgoSpec tl.CGoSpec, level uint8, last bool) {
	uplevel := level - 1
	if level == 0 {
		fmt.Fprintln(buf1, "for i0, mem1 := range mem0 {")
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "ptr%d := (*(*[%d]%s%s)(unsafe.Pointer(ptr%d)))[i%d]\n",
		size, level, cgoSpec.AtLevel(uplevel), uplevel, uplevel)
	if last {
		fmt.Fprintf(buf1, "for i%d := range mem%d {\n", level, uplevel)
	} else {
		fmt.Fprintf(buf1, "for i%d, mem%d := range mem%d {\n", level, level+1, uplevel)
	}
	buf2.Linef("}\n")
}

func packSlice(buf1 io.Writer, buf2 *reverseBuffer, cgoSpec tl.CGoSpec, level uint8, last bool) {
	uplevel := level - 1
	if level == 0 {
		fmt.Fprintln(buf1, "const m = 1 << 30")
		fmt.Fprintln(buf1, "for i0, mem1 := range mem0 {")
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "ptr%d := (*(*[m]%s%s)(unsafe.Pointer(ptr%d)))[i%d]\n",
		level, cgoSpec.AtLevel(uplevel), uplevel, uplevel)
	if last {
		fmt.Fprintf(buf1, "for i%d := range mem%d {\n", level, uplevel)
	} else {
		fmt.Fprintf(buf1, "for i%d, mem%d := range mem%d {\n", level, level+1, uplevel)
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

	arrays := tl.GetArraySizes(gospec.Arrays)
	for i, size := range arrays {
		last := i >= len(arrays)-1
		packArray(buf1, buf2, size, cgospec, level, last)
		level++
	}
	gospec.Arrays = ""

	for gospec.Slices > 1 {
		gospec.Slices--
		last := gospec.Slices == 0
		packSlice(buf1, buf2, cgospec, level, last)
		level++
	}
	isSlice := gospec.Slices > 0
	isPlain := isPlainType(gospec.Base)

	switch {
	case isPlain && isSlice:
		packPlainSlice(buf1, gospec.Base, gospec.Pointers, level)
	case isPlain:
		packPlain(buf1, cgospec, gospec.Base, gospec.Pointers, level)
	case isSlice:
		packSlice(buf1, buf2, cgospec, level, true)
		if helper := packObj(buf1, gospec, level); helper != nil {
			h.Requires = append(h.Requires, helper)
		}
	default:
		if helper := packObj(buf1, gospec, level); helper != nil {
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

	if helper, ok := toGoHelperMap[goSpec]; ok {
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
	cStringFunc = &Helper{
		Name:        "cStr",
		Description: "cStr represents the data from Go string as *C.char and avoids copying.",
		Source: `func cStr(str string) *C.char {
			h := (*reflect.StringHeader)(unsafe.Pointer(&str))
			return (*C.char)(unsafe.Pointer(h.Data))
		}`,
	}
	cUStringFunc = &Helper{
		Name:        "cUStr",
		Description: "cUStr represents the data from Go string as *C.uchar and avoids copying.",
		Source: `func cUStr(str string) *C.uchar {
			h := (*reflect.StringHeader)(unsafe.Pointer(&str))
			return (*C.uchar)(unsafe.Pointer(h.Data))
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
	goUStringFunc = &Helper{
		Name:        "goUStr",
		Description: "goUStr creates a string backed by *C.uchar and avoids copying.",
		Source: `func goUStr(p *C.uchar) (raw string) {
			if p != nil && *p != 0 {
				h := (*reflect.StringHeader)(unsafe.Pointer(&raw))
				h.Data = uintptr(unsafe.Pointer(p))
				for *p != 0 {
					p = (*C.uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
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
