package generator

import (
	"bytes"
	"fmt"
	"hash/crc32"

	tl "github.com/xlab/cgogen/translator"
)

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

	buf.Reset()
	fmt.Fprintf(buf, "func (x *%s) Free()", goStructName)
	fmt.Fprintf(buf, `{
		if x != nil && x.allocs%2x != nil {
			x.allocs%2x.(*cgoAllocMap).Free()
			x.ref%2x = nil
		}
	}`, crc, crc, crc)
	helpers = append(helpers, &Helper{
		Name: fmt.Sprintf("%s.Free", goStructName),
		Description: "Free cleanups the memory using the free stdlib function on C side.\n" +
			"Does nothing if object has no pointer.",
		Source: buf.String(),
	})

	buf.Reset()
	fmt.Fprintf(buf, "func New%sRef(ref interface{}) *%s", goStructName, goStructName)
	fmt.Fprintf(buf, `{
		if ref == nil {
			return nil
		}
		type ifaceHeader struct {
			_   uintptr
			Ref uintptr
		}
		hdr := (*(*ifaceHeader)(unsafe.Pointer(&ref)))
		obj := new(%s)
		obj.ref%2x = (*%s)(unsafe.Pointer(hdr.Ref))
		return obj
	}`, goStructName, crc, cgoSpec)

	name := fmt.Sprintf("New%sRef", goStructName)
	helpers = append(helpers, &Helper{
		Name:        name,
		Description: name + " initialises a new struct holding the reference to the originaitng C struct.",
		Source:      buf.String(),
	})

	buf.Reset()
	fmt.Fprintf(buf, "func (x *%s) PassRef() (*%s, *cgoAllocMap) {\n", goStructName, cgoSpec)
	buf.Write(gen.getPassRefSource(goStructName, cStructName, spec))
	buf.WriteRune('}')
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.PassRef", goStructName),
		Description: "PassRef returns a reference and creates new C object if no refernce yet.",
		Source:      buf.String(),
	})

	buf.Reset()
	fmt.Fprintf(buf, "func (x %s) PassValue() (%s, *cgoAllocMap) {\n", goStructName, cgoSpec)
	buf.Write(gen.getPassValueSource(goStructName, spec))
	buf.WriteRune('}')
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.PassValue", goStructName),
		Description: "PassValue creates a new C object if no refernce yet and returns the dereferenced value.",
		Source:      buf.String(),
	})

	buf.Reset()
	fmt.Fprintf(buf, "func (x *%s) Deref() {\n", goStructName)
	buf.Write(gen.getDerefSource(goStructName, cStructName, spec))
	buf.WriteRune('}')
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.Deref", goStructName),
		Description: "Deref reads the internal fields of struct from its C pointer.",
		Source:      buf.String(),
	})
	return
}

func (gen *Generator) getRawStructHelpers(goStructName []byte, spec tl.CType) (helpers []*Helper) {
	if spec.GetPointers() > 0 {
		return nil // can't addess a pointer receiver
	}
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

	buf.Reset()
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

	buf.Reset()
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

	buf.Reset()
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
		return nil, nil
	} else if x.ref%2x != nil {
		return x.ref%2x, nil
	}`, crc, crc)
	writeSpace(buf, 1)

	h := gen.getAllocMemoryHelper(tl.CGoSpec{Base: cgoSpec.Base})
	gen.submitHelper(h)

	fmt.Fprintf(buf, "mem%2x := %s(1)\n", crc, h.Name)
	fmt.Fprintf(buf, "ref%2x := (*%s)(mem%2x)\n", crc, cgoSpec.Base, crc)
	fmt.Fprintf(buf, "allocs%2x := new(cgoAllocMap)", crc)

	writeSpace(buf, 1)

	ptrTipRx, memTipRx := gen.tr.PtrMemTipRxForSpec(tl.TipScopeStruct, cStructName, spec)
	for i, m := range structSpec.Members {
		if len(m.Name) == 0 {
			continue
			// TODO: generate setters
		}

		typeName := m.Spec.GetBase()
		switch m.Spec.Kind() {
		case tl.StructKind, tl.OpaqueStructKind, tl.EnumKind:
			if !gen.tr.IsAcceptableName(tl.TargetType, typeName) {
				continue
			}
		}

		memTip := memTipRx.TipAt(i)
		if !memTip.IsValid() {
			memTip = gen.MemTipOf(m)
		}
		var goSpec tl.GoTypeSpec
		if ptrTip := ptrTipRx.TipAt(i); ptrTip.IsValid() {
			goSpec = gen.tr.TranslateSpec(m.Spec, ptrTip)
		} else if memTip == tl.TipMemRaw {
			goSpec = gen.tr.TranslateSpec(m.Spec, tl.TipPtrSRef)
		} else {
			goSpec = gen.tr.TranslateSpec(m.Spec)
		}

		cgoSpec := gen.tr.CGoSpec(m.Spec)
		const public = true
		goName := "x." + string(gen.tr.TransformName(tl.TargetType, m.Name, public))
		fromProxy, nillable := gen.proxyValueFromGo(memTip, goName, goSpec, cgoSpec)
		if nillable {
			fmt.Fprintf(buf, "if %s != nil {\n", goName)
		}
		fmt.Fprintf(buf, "var c%s_allocs *cgoAllocMap\n", m.Name)
		fmt.Fprintf(buf, "ref%2x.%s, c%s_allocs  = %s\n", crc, m.Name, m.Name, fromProxy)
		fmt.Fprintf(buf, "allocs%2x.Borrow(c%s_allocs)\n", crc, m.Name)
		if nillable {
			fmt.Fprintf(buf, "}\n\n")
		} else {
			fmt.Fprint(buf, "\n")
		}
	}
	fmt.Fprintf(buf, "x.ref%2x = ref%2x\n", crc, crc)
	fmt.Fprintf(buf, "x.allocs%2x = allocs%2x\n", crc, crc)
	fmt.Fprintf(buf, "return ref%2x, allocs%2x\n", crc, crc)
	writeSpace(buf, 1)
	return buf.Bytes()
}

func (gen *Generator) getPassValueSource(goStructName []byte, spec tl.CType) []byte {
	buf := new(bytes.Buffer)
	crc := getRefCRC(spec)
	fmt.Fprintf(buf, `if x.ref%2x != nil {
		return *x.ref%2x, nil
	}`, crc, crc)
	writeSpace(buf, 1)
	fmt.Fprintf(buf, "ref, allocs := x.PassRef()\n")
	fmt.Fprintf(buf, "return *ref, allocs\n")
	return buf.Bytes()
}

func getRefCRC(spec tl.CType) uint32 {
	return crc32.ChecksumIEEE([]byte(spec.String()))
}

func (gen *Generator) getDerefSource(goStructName []byte, cStructName string, spec tl.CType) []byte {
	structSpec := spec.(*tl.CStructSpec)
	buf := new(bytes.Buffer)
	crc := getRefCRC(spec)
	fmt.Fprintf(buf, `if x.ref%2x == nil {
		return
	}`, crc)
	writeSpace(buf, 1)

	ptrTipRx, memTipRx := gen.tr.PtrMemTipRxForSpec(tl.TipScopeStruct, cStructName, spec)
	for i, m := range structSpec.Members {
		if len(m.Name) == 0 {
			continue
			// TODO: generate getters
		}

		typeName := m.Spec.GetBase()
		switch m.Spec.Kind() {
		case tl.StructKind, tl.OpaqueStructKind, tl.EnumKind:
			if !gen.tr.IsAcceptableName(tl.TargetType, typeName) {
				continue
			}
		}

		memTip := memTipRx.TipAt(i)
		if !memTip.IsValid() {
			memTip = gen.MemTipOf(m)
		}
		var goSpec tl.GoTypeSpec
		if ptrTip := ptrTipRx.TipAt(i); ptrTip.IsValid() {
			goSpec = gen.tr.TranslateSpec(m.Spec, ptrTip)
		} else if memTip == tl.TipMemRaw {
			goSpec = gen.tr.TranslateSpec(m.Spec, tl.TipPtrSRef)
		} else {
			goSpec = gen.tr.TranslateSpec(m.Spec)
		}
		const public = true
		goName := "x." + string(gen.tr.TransformName(tl.TargetType, m.Name, public))
		cgoName := fmt.Sprintf("x.ref%2x.%s", crc, m.Name)
		cgoSpec := gen.tr.CGoSpec(m.Spec)
		toProxy, _ := gen.proxyValueToGo(memTip, goName, cgoName, goSpec, cgoSpec)
		fmt.Fprintln(buf, toProxy)
	}
	return buf.Bytes()
}
