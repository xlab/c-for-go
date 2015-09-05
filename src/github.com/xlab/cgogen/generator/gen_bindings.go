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

var toGoHelperMap = map[tl.GoTypeSpec]getHelperFunc{
	tl.StringSpec:  (*Generator).getPackStringHelper,
	tl.UStringSpec: (*Generator).getPackStringHelper,
}

type proxyDecl struct {
	Name string
	Decl string
}

// getHelperName transforms signatures like [4][4][]*string into A4A4SPString
// suitable for the generated proxy helper methods.
func getHelperName(goSpec tl.GoTypeSpec) string {
	buf := new(bytes.Buffer)
	for _, size := range tl.GetArraySizes(goSpec.Arrays) {
		fmt.Fprintf(buf, "A%d", size)
	}
	buf.WriteString(strings.Repeat("S", int(goSpec.Slices)))
	buf.WriteString(strings.Repeat("P", int(goSpec.Pointers)))
	if goSpec.Unsigned {
		buf.WriteRune('U')
	}
	buf.WriteString(strings.Title(goSpec.Base))
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

type level struct {
	Array bool
	Slice bool
	Last  bool
}

type reverseBuffer []string

func (r *reverseBuffer) Linef(f string, a ...interface{}) {
	if r == nil {
		*r = make(reverseBuffer, 0, 64)
	}
	*r = append(*r, fmt.Sprintf(f, a...))
}

func (r *reverseBuffer) Line(a ...interface{}) {
	if r == nil {
		*r = make(reverseBuffer, 0, 64)
	}
	*r = append(*r, fmt.Sprint(a...))
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

func genIndices(name string, level uint8) []byte {
	buf := new(bytes.Buffer)
	for i := uint8(0); i < level; i++ {
		fmt.Fprintf(buf, "[i%d]", i)
	}
	return buf.Bytes()
}

func unpackPlain(buf io.Writer, goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec, level uint8) {
	uplevel := level - 1
	if goSpec.Pointers > 0 {
		fmt.Fprintf(buf, "mem%d[i%d] = (%s)(unsafe.Pointer(x%d))\n",
			uplevel, uplevel, cgoSpec.AtLevel(level), level)
		return
	} else if len(goSpec.Arrays) > 0 {
		fmt.Fprintf(buf, "mem%d[i%d] = *(*%s)(unsafe.Pointer(x%d))\n",
			uplevel, uplevel, cgoSpec.AtLevel(level), level)
		return
	}
	fmt.Fprintf(buf, "mem%d[i%d] = *(*%s)(unsafe.Pointer(&x%d))\n",
		uplevel, uplevel, cgoSpec.Base, level)
}

func unpackPlainSlice(buf io.Writer, cgoSpec tl.CGoSpec, level uint8) {
	uplevel := level - 1
	fmt.Fprintf(buf, "mem%d[i%d] = (%s)(unsafe.Pointer(&x%s[0]))\n",
		uplevel, uplevel, cgoSpec.AtLevel(level), genIndices("i", level))
}

func (gen *Generator) unpackObj(buf io.Writer, goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec, level uint8) *Helper {
	uplevel := level - 1
	indices := genIndices("i", level)
	if getHelper, ok := fromGoHelperMap[goSpec]; ok {
		helper := getHelper(gen, cgoSpec)
		fmt.Fprintf(buf, "mem%d[i%d] = %s(%sx%s)\n",
			uplevel, uplevel, helper.Name, ptrs(goSpec.Pointers), indices)
		return helper
	}
	var deref string
	if cgoSpec.PointersAtLevel(level) == 0 {
		deref = "*"
	}
	fmt.Fprintf(buf, "mem%d[i%d] = %sx%s.PassRef()\n", uplevel, uplevel, deref, indices)
	return nil
}

func unpackArray(buf1 io.Writer, buf2 *reverseBuffer, size uint64, cgoSpec tl.CGoSpec, level uint8) {
	uplevel := level - 1
	if level == 0 {
		fmt.Fprintf(buf1, "var mem0 %s\n", cgoSpec)
		fmt.Fprintf(buf1, "for i0 := range x {\n")
		buf2.Linef("return (%s)(unsafe.Pointer(&mem0))", cgoSpec.AtLevel(level))
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "var mem%d %s\n", level, cgoSpec.AtLevel(level))
	fmt.Fprintf(buf1, "for i%d := range x%s {\n", level, genIndices("i", level))
	buf2.Linef("mem%d[i%d] = *(*%s)(unsafe.Pointer(&mem%d))\n",
		uplevel, uplevel, cgoSpec.AtLevel(level), level)
	buf2.Linef("}\n")
}

func notNilBarrier(buf io.Writer, name string) {
	fmt.Fprintf(buf, `if %s == nil {
		return nil
	}`, name)
}

func unpackSlice(buf1 io.Writer, buf2 *reverseBuffer, cgoSpec tl.CGoSpec, level uint8) {
	uplevel := level - 1
	if level == 0 {
		notNilBarrier(buf1, "x")
		writeSpace(buf1, 1)
		fmt.Fprintf(buf1, "mem0 := make([]%s, len(x))\n", cgoSpec.AtLevel(1))
		fmt.Fprintf(buf1, "for i0 := range x {\n")
		buf2.Linef("return (%s)(unsafe.Pointer(&mem0[0]))\n", cgoSpec.AtLevel(0))
		buf2.Linef("}\n")
		return
	}
	indices := genIndices("i", level)
	fmt.Fprintf(buf1, "mem%d := make([]%s, len(x%s))\n", level, cgoSpec.AtLevel(level+1), indices)
	fmt.Fprintf(buf1, "for i%d := range x%s {\n", level, indices)
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
			h := (*stringHeader)(unsafe.Pointer(&str))
			return (%s)(unsafe.Pointer(h.Data))
		}`, name, cgoSpec, cgoSpec),
		Requires: []*Helper{stringHeader},
	}
}

func (gen *Generator) getUnpackHelper(goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec) *Helper {
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
	isPlain := goSpec.IsPlain()

	switch {
	case isPlain && isSlice:
		unpackPlainSlice(buf1, cgoSpec, level)
	case isPlain:
		unpackPlain(buf1, goSpec, cgoSpec, level)
	case isSlice:
		unpackSlice(buf1, buf2, cgoSpec, level)
		goSpec.Slices = 0
		if helper := gen.unpackObj(buf1, goSpec, cgoSpec, level+1); helper != nil {
			h.Requires = append(h.Requires, helper)
		}
	default:
		if helper := gen.unpackObj(buf1, goSpec, cgoSpec, level); helper != nil {
			h.Requires = append(h.Requires, helper)
		}
	}

	buf1.WriteTo(code)
	buf2.WriteTo(code)
	fmt.Fprintln(code, "}")
	h.Source = code.String()
	return h
}

func (gen *Generator) proxyValueFromGo(name string,
	goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec) (proxy string, nillable bool) {

	if getHelper, ok := fromGoHelperMap[goSpec]; ok {
		helper := getHelper(gen, cgoSpec)
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s(%s)", helper.Name, name)
		return proxy, helper.Nillable
	}

	isPlain := goSpec.IsPlain()
	switch {
	case !isPlain && (goSpec.Slices > 0 || len(goSpec.Arrays) > 0), // ex: []string
		isPlain && goSpec.Slices > 0 && len(goSpec.Arrays) > 0, // ex: [4][]byte
		isPlain && goSpec.Slices > 1:                           // ex: [][]byte
		helper := gen.getUnpackHelper(goSpec, cgoSpec)
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
		proxy = fmt.Sprintf("*(*%s)(unsafe.Pointer(%s))", cgoSpec, name)
		return
	default: // ex: *SomeType
		var deref string
		if cgoSpec.Pointers == 0 {
			deref = "*"
		}
		proxy = fmt.Sprintf("%s%s.PassRef()", deref, name)
		return
	}
}

func (gen *Generator) proxyArgFromGo(name string,
	goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec) (proxy string, nillable bool) {

	if getHelper, ok := fromGoHelperMap[goSpec]; ok {
		helper := getHelper(gen, cgoSpec)
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s(%s)", helper.Name, name)
		return proxy, helper.Nillable
	}

	isPlain := goSpec.IsPlain()
	switch {
	case !isPlain && (goSpec.Slices > 0 || len(goSpec.Arrays) > 0), // ex: []string
		isPlain && goSpec.Slices > 0 && len(goSpec.Arrays) > 0, // ex: [4][]byte
		isPlain && goSpec.Slices > 1:                           // ex: [][]byte
		helper := gen.getUnpackHelper(goSpec, cgoSpec)
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
		proxy = fmt.Sprintf("(%s)(unsafe.Pointer(%s))", cgoSpec.AtLevel(0), name)
		return
	default: // ex: *SomeType
		var deref string
		if cgoSpec.Pointers == 0 {
			deref = "*"
		}
		proxy = fmt.Sprintf("%s%s.PassRef()", deref, name)
		return
	}
}

func packPlain(buf io.Writer, cgoSpec tl.CGoSpec, base string, pointers uint8, level uint8) {
	uplevel := level - 1
	fmt.Fprintf(buf, "mem%s = (%s%s)(unsafe.Pointer(%sptr%d))\n",
		genIndices("i", level), ptrs(pointers-uplevel), base, ptrs(cgoSpec.Pointers-level), level)
}

func packPlainSlice(buf io.Writer, base string, pointers uint8, level uint8) {
	fmt.Fprintf(buf, "h := (*sliceHeader)(unsafe.Pointer(&mem%s))\n", genIndices("i", level))
	fmt.Fprintf(buf, "h.Data = uintptr(unsafe.Pointer(ptr%d))\n", level)
	fmt.Fprintln(buf, "h.Cap = 0x7fffffff")
	fmt.Fprintln(buf, "// h.Len = ?")
}

func (gen *Generator) packObj(buf io.Writer, goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec, level uint8) *Helper {
	if getHelper, ok := toGoHelperMap[goSpec]; ok {
		helper := getHelper(gen, cgoSpec)
		fmt.Fprintf(buf, "mem%s = %s(%sptr%d)\n",
			genIndices("i", level), helper.Name, ptrs(goSpec.Pointers), level)
		return helper
	}
	var deref string
	var ref string
	if cgoSpec.PointersAtLevel(level) == 0 {
		deref = "*"
		ref = "&"
	}
	fmt.Fprintf(buf, "mem%s = %sNew%sRef(%sptr%d)\n", genIndices("i", level), deref, goSpec.Base, ref, level)
	return nil
}

func packArray(buf1 io.Writer, buf2 *reverseBuffer, cgoSpec tl.CGoSpec, level uint8) {
	if level == 0 {
		fmt.Fprintln(buf1, "const m = 0x7fffffff")
		fmt.Fprintln(buf1, "for i0 := range mem {")
		fmt.Fprintf(buf1, "ptr1 := ptr0[i0]\n")
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "for i%d := range mem%s {\n", level, genIndices("i", level))
	fmt.Fprintf(buf1, "ptr%d := ptr%d[i%d]\n", level+1, level, level)
	buf2.Linef("}\n")
}

func packSlice(buf1 io.Writer, buf2 *reverseBuffer, cgoSpec tl.CGoSpec, level uint8) {
	if level == 0 {
		fmt.Fprintln(buf1, "const m = 0x7fffffff")
		fmt.Fprintln(buf1, "for i0 := range mem {")
		fmt.Fprintf(buf1, "ptr1 := (*(*[m]%s)(unsafe.Pointer(ptr0)))[i0]\n", cgoSpec.AtLevel(level+1))
		buf2.Linef("}\n")
		return
	}
	fmt.Fprintf(buf1, "for i%d := range mem%s {\n", level, genIndices("i", level))
	fmt.Fprintf(buf1, "ptr%d := (*(*[m]%s)(unsafe.Pointer(ptr%d)))[i%d]\n",
		level+1, cgoSpec.AtLevel(level+1), level, level)
	buf2.Linef("}\n")
}

func (gen *Generator) getPackStringHelper(cgoSpec tl.CGoSpec) *Helper {
	cgoSpec = tl.CGoSpec{
		Pointers: 1,
		Base:     cgoSpec.Base,
	}
	name := "pack" + gen.getTypedHelperName("string", cgoSpec)
	return &Helper{
		Name:        name,
		Description: fmt.Sprintf("%s creates a Go string backed by %s and avoids copying.", name, cgoSpec),
		Source: fmt.Sprintf(`func %s(p %s) (raw string) {
			if p != nil && *p != 0 {
				h := (*stringHeader)(unsafe.Pointer(&raw))
				h.Data = uintptr(unsafe.Pointer(p))
				for *p != 0 {
					p = (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
				}
				h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
			}
			return
		}`, name, cgoSpec, cgoSpec),
		Requires: Helpers{stringHeader, rawString},
	}
}

func (gen *Generator) getPackHelper(goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec) *Helper {
	name := "pack" + getHelperName(goSpec)
	code := new(bytes.Buffer)
	var ref string
	if len(goSpec.Arrays) > 0 {
		ref = "*"
	}
	fmt.Fprintf(code, "func %s(mem %s, ptr0 %s%s) {\n", name, goSpec, ref, cgoSpec)
	h := &Helper{
		Name:        name,
		Description: fmt.Sprintf("%s reads sliced Go data structure out from plain C format.", name),
	}
	var level uint8
	buf1 := new(bytes.Buffer)
	buf2 := new(reverseBuffer)

	for range tl.GetArraySizes(goSpec.Arrays) {
		packArray(buf1, buf2, cgoSpec, level)
		level++
	}
	goSpec.Arrays = ""

	for goSpec.Slices > 1 {
		goSpec.Slices--
		packSlice(buf1, buf2, cgoSpec, level)
		level++
	}
	isSlice := goSpec.Slices > 0
	isPlain := goSpec.IsPlain()

	switch {
	case isPlain && isSlice:
		gen.submitHelper(sliceHeader)
		packPlainSlice(buf1, goSpec.Base, goSpec.Pointers, level)
	case isPlain:
		packPlain(buf1, cgoSpec, goSpec.Base, goSpec.Pointers, level)
	case isSlice:
		packSlice(buf1, buf2, cgoSpec, level)
		goSpec.Slices = 0
		if helper := gen.packObj(buf1, goSpec, cgoSpec, level+1); helper != nil {
			h.Requires = append(h.Requires, helper)
		}
	default:
		if helper := gen.packObj(buf1, goSpec, cgoSpec, level); helper != nil {
			h.Requires = append(h.Requires, helper)
		}
	}

	buf1.WriteTo(code)
	buf2.WriteTo(code)
	fmt.Fprintln(code, "}")
	h.Source = code.String()
	return h
}

func (gen *Generator) proxyArgToGo(memName, ptrName string,
	goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec) (proxy string, nillable bool) {
	nillable = true

	if getHelper, ok := toGoHelperMap[goSpec]; ok {
		helper := getHelper(gen, cgoSpec)
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s = %s(%s)", memName, helper.Name, ptrName)
		return proxy, helper.Nillable
	}

	isPlain := goSpec.IsPlain()
	switch {
	case !isPlain && (goSpec.Slices > 0 || len(goSpec.Arrays) > 0), // ex: []string
		isPlain && goSpec.Slices > 0 && len(goSpec.Arrays) > 0, // ex: [4][]byte
		isPlain && goSpec.Slices > 1:                           // ex: [][]byte
		helper := gen.getPackHelper(goSpec, cgoSpec)
		gen.submitHelper(helper)
		if len(goSpec.Arrays) > 0 {
			ptrName = fmt.Sprintf("(*%s)(unsafe.Pointer(%s))", cgoSpec, ptrName)
		}
		proxy = fmt.Sprintf("%s(%s, %s)", helper.Name, memName, ptrName)
		return proxy, helper.Nillable
	case isPlain && goSpec.Slices > 0: // ex: []byte
		// skip because slice data can be edited either way
		return
	case isPlain: // ex: byte, [4]byte
		if len(goSpec.Arrays) == 0 {
			proxy = fmt.Sprintf("*%s = *(%s)(%s)", memName, goSpec, ptrName)
			return
		}
		proxy = fmt.Sprintf("*%s = *(%s)(unsafe.Pointer(&%s))", memName, goSpec, ptrName)
		return
	default: // ex: *SomeType
		proxy = fmt.Sprintf("*%s = *(New%sRef(%s))", memName, goSpec.Base, ptrName)
		return
	}
}

func (gen *Generator) proxyValueToGo(memName, ptrName string,
	goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec) (proxy string, nillable bool) {
	nillable = true

	if getHelper, ok := toGoHelperMap[goSpec]; ok {
		helper := getHelper(gen, cgoSpec)
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s = %s(%s)", memName, helper.Name, ptrName)
		return proxy, helper.Nillable
	}

	isPlain := goSpec.IsPlain()
	switch {
	case !isPlain && (goSpec.Slices > 0 || len(goSpec.Arrays) > 0), // ex: []string
		isPlain && goSpec.Slices > 0 && len(goSpec.Arrays) > 0, // ex: [4][]byte
		isPlain && goSpec.Slices > 1:                           // ex: [][]byte
		helper := gen.getPackHelper(goSpec, cgoSpec)
		gen.submitHelper(helper)
		if len(goSpec.Arrays) > 0 {
			ptrName = fmt.Sprintf("(*%s)(unsafe.Pointer(%s))", cgoSpec, ptrName)
		}
		gen.submitHelper(sliceHeader)
		proxy = fmt.Sprintf("%s(%s, %s)", helper.Name, memName, ptrName)
		return proxy, helper.Nillable
	case isPlain && goSpec.Slices > 0: // ex: []byte
		gen.submitHelper(sliceHeader)
		buf := new(bytes.Buffer)
		fmt.Fprintf(buf, "h := (*sliceHeader)(unsafe.Pointer(&%s))\n", memName)
		fmt.Fprintf(buf, "h.Data = uintptr(unsafe.Pointer(%s))\n", ptrName)
		fmt.Fprintln(buf, "h.Cap = 0x7fffffff")
		fmt.Fprintln(buf, "// h.Len = ?")
		proxy = buf.String()
		return
	case isPlain: // ex: byte, [4]byte
		if len(goSpec.Arrays) == 0 {
			proxy = fmt.Sprintf("%s = (%s)(%s)", memName, goSpec, ptrName)
			return
		}
		proxy = fmt.Sprintf("%s = (%s)(unsafe.Pointer(&%s))", memName, goSpec, ptrName)
		return
	default: // ex: *SomeType
		var deref string
		var ref string
		if cgoSpec.Pointers == 0 {
			deref = "*"
			ref = "&"
		}
		proxy = fmt.Sprintf("%s = %sNew%sRef(%s%s)", memName, deref, goSpec.Base, ref, ptrName)
		return
	}
}

func (gen *Generator) proxyRetToGo(memName, ptrName string,
	goSpec tl.GoTypeSpec, cgoSpec tl.CGoSpec) (proxy string, nillable bool) {

	if getHelper, ok := toGoHelperMap[goSpec]; ok {
		helper := getHelper(gen, cgoSpec)
		gen.submitHelper(helper)
		proxy = fmt.Sprintf("%s := %s(%s)", memName, helper.Name, ptrName)
		return proxy, helper.Nillable
	}

	isPlain := goSpec.IsPlain()
	switch {
	case !isPlain && (goSpec.Slices > 0 || len(goSpec.Arrays) > 0), // ex: []string
		isPlain && goSpec.Slices > 0 && len(goSpec.Arrays) > 0, // ex: [4][]byte
		isPlain && goSpec.Slices > 1:                           // ex: [][]byte
		helper := gen.getPackHelper(goSpec, cgoSpec)
		gen.submitHelper(helper)
		var ref string
		if len(goSpec.Arrays) > 0 {
			ref = "&"
		}
		proxy = fmt.Sprintf("%s := %s(%s%s, %s)", memName, helper.Name, ref, memName, ptrName)
		return proxy, helper.Nillable
	case isPlain && goSpec.Slices > 0: // ex: []byte
		proxy = fmt.Sprintf("%s := (*(*[0x7fffffff]%s%s)(unsafe.Pointer(%s)))[:0]",
			memName, ptrs(goSpec.Pointers), goSpec.Base, ptrName)
		return
	case isPlain: // ex: byte, [4]byte
		if len(goSpec.Arrays) == 0 {
			proxy = fmt.Sprintf("%s := (%s)(%s)", memName, goSpec, ptrName)
			return
		}
		if goSpec.Pointers == 0 {
			ptrName = "&" + ptrName
		}
		var deref string
		if cgoSpec.Pointers == 0 {
			deref = "*"
		}
		proxy = fmt.Sprintf("%s := %s(%s%s)(unsafe.Pointer(%s))", memName, deref, deref, goSpec, ptrName)
		return
	default: // ex: *SomeType
		var deref string
		var ref string
		if cgoSpec.Pointers == 0 {
			deref = "*"
			ref = "&"
		}
		proxy = fmt.Sprintf("%s := %sNew%sRef(%s%s)", memName, deref, goSpec.Base, ref, ptrName)
		return
	}
}

func (gen *Generator) createProxies(funcName string, funcSpec tl.CType) (from, to []proxyDecl) {
	spec := funcSpec.(*tl.CFunctionSpec)
	from = make([]proxyDecl, len(spec.ParamList))
	to = make([]proxyDecl, 0, len(spec.ParamList))

	var nextPtrSpec tl.PointerSpec
	ptrLayout, fallback := gen.tr.PointerLayout(tl.PointerScopeFunction, funcName)

	for i, param := range spec.ParamList {
		nextPtrSpec, ptrLayout = tl.NextPointerSpec(ptrLayout, fallback)
		goSpec := gen.tr.TranslateSpec(param.Spec, nextPtrSpec)
		cgoSpec := gen.tr.CGoSpec(param.Spec)
		refName := string(gen.tr.TransformName(tl.TargetPrivate, param.Name))
		fromBuf := new(bytes.Buffer)
		toBuf := new(bytes.Buffer)
		name := "c" + refName
		fromProxy, nillable := gen.proxyArgFromGo(refName, goSpec, cgoSpec)
		if nillable {
			fmt.Fprintf(fromBuf, "var %s %s\n", name, cgoSpec)
			fmt.Fprintf(fromBuf, "if %s != nil {\n%s = %s\n}", refName, name, fromProxy)
		} else {
			fmt.Fprintf(fromBuf, "%s := %s", name, fromProxy)
		}
		from[i] = proxyDecl{Name: name, Decl: fromBuf.String()}

		isPlain := goSpec.IsPlain()
		switch {
		case !isPlain && (goSpec.Slices > 0 || len(goSpec.Arrays) > 0), // ex: []string
			isPlain && goSpec.Slices > 0 && len(goSpec.Arrays) > 0, // ex: [4][]byte
			isPlain && goSpec.Slices > 1:                           // ex: [][]byte
			// need to re-pack values into Go memory layout
			if toProxy, nillable := gen.proxyArgToGo(refName, name, goSpec, cgoSpec); len(toProxy) > 0 {
				if nillable {
					fmt.Fprintf(toBuf, "if %s != nil {\n%s\n}", refName, toProxy)
				} else {
					fmt.Fprintln(toBuf, toProxy)
				}
				to = append(to, proxyDecl{Name: name, Decl: toBuf.String()})
			}
		}
	}
	return
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
	wr2 := new(reverseBuffer)
	from, to := gen.createProxies(decl.Name, decl.Spec)
	for _, proxy := range from {
		fmt.Fprintln(wr, proxy.Decl)
	}
	for _, proxy := range to {
		// wr2 will be handled below
		wr2.Line(proxy.Decl)
	}
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		fmt.Fprint(wr, "ret := ")
	}
	fmt.Fprintf(wr, "C.%s", decl.Name)
	writeStartParams(wr)
	for i := range spec.ParamList {
		fmt.Fprint(wr, from[i].Name)
		if i < len(spec.ParamList)-1 {
			fmt.Fprint(wr, ", ")
		}
	}
	writeEndParams(wr)
	writeSpace(wr, 1)
	// wr2 being populated above
	wr2.WriteTo(wr)
	if spec.Return != nil {
		goSpec := gen.tr.TranslateSpec((*spec).Return.Spec, tl.PointerRef)
		cgoSpec := gen.tr.CGoSpec((*spec).Return.Spec)
		retProxy, nillable := gen.proxyRetToGo("mem", "ret", goSpec, cgoSpec)
		if nillable {
			fmt.Fprintln(wr, "if ret == nil {\nreturn nil\n}")
		}
		fmt.Fprintln(wr, retProxy)
		fmt.Fprintln(wr, "return mem")
	}
	writeEndFuncBody(wr)
}

var (
	sliceHeader = &Helper{
		Name: "sliceHeader",
		Source: `type sliceHeader struct {
	        Data uintptr
	        Len  int
	        Cap  int
		}`,
	}
	stringHeader = &Helper{
		Name: "stringHeader",
		Source: `type stringHeader struct {
        	Data uintptr
        	Len  int
		}`,
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
			h := (*stringHeader)(unsafe.Pointer(&raw))
			return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
		}`,
		Requires: []*Helper{stringHeader},
	}
)
