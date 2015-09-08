package generator

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func (gen *Generator) writeTypeTypedef(wr io.Writer, decl tl.CDecl) {
	goSpec := gen.tr.TranslateSpec(decl.Spec)
	goTypeName := gen.tr.TransformName(tl.TargetType, decl.Name)
	fmt.Fprintf(wr, "// %s type as declared in %s\n", goTypeName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s %s", goTypeName, goSpec)
	writeSpace(wr, 1)
}

func (gen *Generator) writeEnumTypedef(wr io.Writer, decl tl.CDecl) {
	var goEnumName []byte
	if len(decl.Name) > 0 {
		goEnumName = gen.tr.TransformName(tl.TargetType, decl.Name)
	} else {
		return
	}
	typeRef := gen.tr.TranslateSpec(decl.Spec).String()
	if typeName := string(goEnumName); typeName != typeRef {
		fmt.Fprintf(wr, "// %s as declared in %s\n", goEnumName, tl.SrcLocation(decl.Pos))
		fmt.Fprintf(wr, "type %s %s", goEnumName, typeRef)
		writeSpace(wr, 1)
	}
}

func (gen *Generator) writeFunctionTypedef(wr io.Writer, decl tl.CDecl) {
	var returnRef string
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		// defaults to ref for the returns
		ptrTip := tl.TipPtrRef
		if ptrTipRx, ok := gen.tr.PtrTipRx(tl.TipScopeFunction, decl.Name); ok {
			if tip := ptrTipRx.Self(); tip.IsValid() {
				ptrTip = tip
			}
		}
		returnRef = gen.tr.TranslateSpec(spec.Return.Spec, ptrTip).String()
	}

	goFuncName := gen.tr.TransformName(tl.TargetType, decl.Name)
	fmt.Fprintf(wr, "// %s type as declared in %s\n", goFuncName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s %s", goFuncName, decl.Spec)
	gen.writeFunctionParams(wr, decl.Name, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
	writeSpace(wr, 1)
}

func (gen *Generator) writeStructTypedef(wr io.Writer, decl tl.CDecl, raw bool) {
	var goStructName []byte
	if len(decl.Name) > 0 {
		goStructName = gen.tr.TransformName(tl.TargetType, decl.Name)
	} else {
		return
	}

	if !decl.IsTemplate() {
		// not a template, so a struct referenced by a tag declares a new type
		typeRef := gen.tr.TranslateSpec(decl.Spec).String()

		if typeName := string(goStructName); typeName != typeRef {
			fmt.Fprintf(wr, "// %s as declared in %s\n", goStructName, tl.SrcLocation(decl.Pos))
			fmt.Fprintf(wr, "type %s %s", goStructName, typeRef)
			writeSpace(wr, 1)
			return
		}
	}

	fmt.Fprintf(wr, "// %s as declared in %s\n", goStructName, tl.SrcLocation(decl.Pos))
	fmt.Fprintf(wr, "type %s struct {", goStructName)
	writeSpace(wr, 1)
	gen.writeStructMembers(wr, decl.Name, decl.Spec)
	writeEndStruct(wr)
	writeSpace(wr, 1)
	if raw {
		for _, helper := range gen.getRawStructHelpers(goStructName, decl.Spec) {
			gen.submitHelper(helper)
		}
		return
	}
	for _, helper := range gen.getStructHelpers(goStructName, decl.Name, decl.Spec) {
		gen.submitHelper(helper)
	}
}

func (gen *Generator) getStructHelpers(goStructName []byte, cStructName string, spec tl.CType) (helpers []*Helper) {
	crc := getRefCRC(spec)
	cgoSpec := gen.tr.CGoSpec(spec)

	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x *%s) Ref() *%s", goStructName, cgoSpec)
	fmt.Fprintf(buf, `{
		if x == nil {
			return nil
		}
		return x.ref%2x
	}`, crc)
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.Ref", goStructName),
		Description: "Ref returns a reference.",
		Source:      buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x *%s) Free()", goStructName)
	fmt.Fprintf(buf, `{
		if x != nil && x.ref%2x != nil {
			runtime.SetFinalizer(x, nil)
			C.free(unsafe.Pointer(x.ref%2x))
			x.ref%2x = nil
		}
	}`, crc, crc, crc)
	helpers = append(helpers, &Helper{
		Name: fmt.Sprintf("%s.Free", goStructName),
		Description: "Free cleanups the memory using the free stdlib function on C side.\n" +
			"Does nothing if object has no pointer.",
		Source: buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func New%sRef(ref *%s) *%s", goStructName, cgoSpec, goStructName)
	fmt.Fprintf(buf, `{
		obj := &%s{}
		if ref == nil {
			obj.ref%2x = new(%s)
			return obj
		} 
		obj.ref%2x = ref
		runtime.SetFinalizer(obj, func(x *%s) {
			x.Free()
		})
		return obj
	}`, goStructName, crc, cgoSpec, crc, goStructName)
	name := fmt.Sprintf("New%sRef", goStructName)
	helpers = append(helpers, &Helper{
		Name:        name,
		Description: name + " initialises a new struct holding the reference to the originaitng C struct.",
		Source:      buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x *%s) PassRef() *%s {\n", goStructName, cgoSpec)
	buf.Write(gen.getPassRefSource(goStructName, cStructName, spec))
	buf.WriteRune('}')
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.PassRef", goStructName),
		Description: "PassRef returns a reference and creates new C object if no refernce yet.",
		Source:      buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x *%s) Deref() {\n", goStructName)
	buf.Write(gen.getDerefSource(cStructName, spec))
	buf.WriteRune('}')
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.Deref", goStructName),
		Description: "Deref reads the internal fields of struct from its C pointer.",
		Source:      buf.String(),
	})
	return
}

func (gen *Generator) getRawStructHelpers(goStructName []byte, spec tl.CType) (helpers []*Helper) {
	cgoSpec := gen.tr.CGoSpec(spec)

	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x *%s) Ref() *%s", goStructName, cgoSpec)
	fmt.Fprintf(buf, `{
		if x == nil {
			return nil
		}
		return (*%s)(unsafe.Pointer(x))
	}`, cgoSpec)
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.Ref", goStructName),
		Description: "Ref returns a reference.",
		Source:      buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x *%s) Free()", goStructName)
	fmt.Fprint(buf, `{
		if x != nil  {
			C.free(unsafe.Pointer(x))
		}
	}`)
	helpers = append(helpers, &Helper{
		Name: fmt.Sprintf("%s.Free", goStructName),
		Description: "Free cleanups the memory using the free stdlib function on C side.\n" +
			"Does nothing if object has no pointer.",
		Source: buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func New%sRef(ref *%s) *%s", goStructName, cgoSpec, goStructName)
	fmt.Fprintf(buf, `{
		return (*%s)(unsafe.Pointer(ref))
	}`, goStructName)
	name := fmt.Sprintf("New%sRef", goStructName)
	helpers = append(helpers, &Helper{
		Name:        name,
		Description: name + " initialises a new struct.",
		Source:      buf.String(),
	})

	buf = new(bytes.Buffer)
	fmt.Fprintf(buf, "func (x *%s) PassRef() *%s", goStructName, cgoSpec)
	fmt.Fprintf(buf, `{
		if x == nil {
			x = new(%s)
		}
		return (*%s)(unsafe.Pointer(x))
	}`, goStructName, cgoSpec)
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.PassRef", goStructName),
		Description: "PassRef returns a reference and creates new object if no refernce yet.",
		Source:      buf.String(),
	})
	return
}

func (gen *Generator) getPassRefSource(goStructName []byte, cStructName string, spec tl.CType) []byte {
	cgoSpec := gen.tr.CGoSpec(spec)
	structSpec := spec.(*tl.CStructSpec)
	buf := new(bytes.Buffer)
	crc := getRefCRC(spec)
	fmt.Fprintf(buf, `if x == nil {
		return nil
	} else if x.ref%2x != nil {
		return x.ref%2x
	}`, crc, crc)
	writeSpace(buf, 1)
	fmt.Fprintf(buf, "ref%2x := new(%s)\n", crc, cgoSpec.Base)
	writeSpace(buf, 1)

	ptrTipRx, memTipRx := gen.tr.PtrMemTipRxForSpec(tl.TipScopeStruct, cStructName, spec)
	for i, mem := range structSpec.Members {
		if len(mem.Name) == 0 {
			continue
		}
		memTip := memTipRx.TipAt(i)
		if tag := mem.Spec.GetBase(); len(tag) > 0 && !memTip.IsValid() {
			if memTipRx, ok := gen.tr.MemTipRx(tag); ok {
				memTip = memTipRx.TipAt(i)
			}
		}
		var goSpec tl.GoTypeSpec
		if ptrTip := ptrTipRx.TipAt(i); ptrTip.IsValid() {
			goSpec = gen.tr.TranslateSpec(mem.Spec, ptrTip)
		} else {
			goSpec = gen.tr.TranslateSpec(mem.Spec)
		}
		cgoSpec := gen.tr.CGoSpec(mem.Spec)
		goName := "x." + string(gen.tr.TransformName(tl.TargetPublic, mem.Name))
		fromProxy, nillable := gen.proxyValueFromGo(memTip, goName, goSpec, cgoSpec)
		if nillable {
			fmt.Fprintf(buf, "if %s != nil {\nref%2x.%s = %s\n}\n", goName, crc, mem.Name, fromProxy)
		} else {
			fmt.Fprintf(buf, "ref%2x.%s = %s\n", crc, mem.Name, fromProxy)
		}
	}
	fmt.Fprintf(buf, `x.ref%2x = ref%2x
		return ref%2x`, crc, crc, crc)
	writeSpace(buf, 1)
	return buf.Bytes()
}

func getRefCRC(spec tl.CType) uint32 {
	return crc32.ChecksumIEEE([]byte(spec.String()))
}

func (gen *Generator) getDerefSource(cStructName string, spec tl.CType) []byte {
	structSpec := spec.(*tl.CStructSpec)
	buf := new(bytes.Buffer)
	crc := getRefCRC(spec)
	fmt.Fprintf(buf, `if x.ref%2x == nil {
		return
	}`, crc)
	writeSpace(buf, 1)

	ptrTipRx, memTipRx := gen.tr.PtrMemTipRxForSpec(tl.TipScopeStruct, cStructName, spec)
	for i, mem := range structSpec.Members {
		if len(mem.Name) == 0 {
			continue
		}
		memTip := memTipRx.TipAt(i)
		if tag := mem.Spec.GetBase(); len(tag) > 0 && !memTip.IsValid() {
			if memTipRx, ok := gen.tr.MemTipRx(tag); ok {
				memTip = memTipRx.TipAt(i)
			}
		}
		var goSpec tl.GoTypeSpec
		if ptrTip := ptrTipRx.TipAt(i); ptrTip.IsValid() {
			goSpec = gen.tr.TranslateSpec(mem.Spec, ptrTip)
		} else {
			goSpec = gen.tr.TranslateSpec(mem.Spec)
		}
		goName := "x." + string(gen.tr.TransformName(tl.TargetPublic, mem.Name))
		cgoName := fmt.Sprintf("x.ref%2x.%s", crc, mem.Name)
		cgoSpec := gen.tr.CGoSpec(mem.Spec)
		toProxy, _ := gen.proxyValueToGo(memTip, goName, cgoName, goSpec, cgoSpec)
		fmt.Fprintln(buf, toProxy)
	}
	return buf.Bytes()
}
