package generator

import (
	"bytes"
	"fmt"
	"hash/crc32"

	tl "github.com/xlab/cgogen/translator"
)

func (gen *Generator) getStructHelpers(goStructName []byte, cStructName string, spec tl.CType) (helpers []*Helper) {
	crc := getRefCRC(spec)
	cgoSpec := gen.tr.CGoSpec(spec, true)

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
		Description: "Ref returns the underlying reference to C object or nil if struct is nil.",
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
		Description: "Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.\n" +
			"Does nothing if struct is nil or has no allocation map.",
		Source: buf.String(),
	})

	buf.Reset()
	fmt.Fprintf(buf, "func New%sRef(ref unsafe.Pointer) *%s", goStructName, goStructName)
	fmt.Fprintf(buf, `{
		if ref == nil {
			return nil
		}
		obj := new(%s)
		obj.ref%2x = (*%s)(unsafe.Pointer(ref))
		return obj
	}`, goStructName, crc, cgoSpec)

	name := fmt.Sprintf("New%sRef", goStructName)
	helpers = append(helpers, &Helper{
		Name: name,
		Description: name + " creates a new wrapper struct with underlying reference set to the original C object.\n" +
			"Returns nil if the provided pointer to C object is nil too.",
		Source: buf.String(),
	})

	buf.Reset()
	fmt.Fprintf(buf, "func (x *%s) PassRef() (*%s, *cgoAllocMap) {\n", goStructName, cgoSpec)
	buf.Write(gen.getPassRefSource(goStructName, cStructName, spec))
	buf.WriteRune('}')
	helpers = append(helpers, &Helper{
		Name: fmt.Sprintf("%s.PassRef", goStructName),
		Description: "PassRef returns the underlying C object, otherwise it will allocate one and set its values\n" +
			"from this wrapping struct, counting allocations into an allocation map.",
		Source: buf.String(),
	})

	buf.Reset()
	fmt.Fprintf(buf, "func (x %s) PassValue() (%s, *cgoAllocMap) {\n", goStructName, cgoSpec)
	buf.Write(gen.getPassValueSource(goStructName, spec))
	buf.WriteRune('}')
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.PassValue", goStructName),
		Description: "PassValue does the same as PassRef except that it will try to dereference the returned pointer.",
		Source:      buf.String(),
	})

	buf.Reset()
	fmt.Fprintf(buf, "func (x *%s) Deref() {\n", goStructName)
	buf.Write(gen.getDerefSource(goStructName, cStructName, spec))
	buf.WriteRune('}')
	helpers = append(helpers, &Helper{
		Name: fmt.Sprintf("%s.Deref", goStructName),
		Description: "Deref uses the underlying reference to C object and fills the wrapping struct with values.\n" +
			"Do not forget to call this method whether you get a struct for C object and want to read its values.",
		Source: buf.String(),
	})
	return
}

func (gen *Generator) getRawStructHelpers(goStructName []byte, spec tl.CType) (helpers []*Helper) {
	if spec.GetPointers() > 0 {
		return nil // can't addess a pointer receiver
	}
	cgoSpec := gen.tr.CGoSpec(spec, true)

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
		Description: "Ref returns a reference to C object as it is.",
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
		Name:        fmt.Sprintf("%s.Free", goStructName),
		Description: "Free cleanups the referenced memory using C free.",
		Source:      buf.String(),
	})

	buf.Reset()
	fmt.Fprintf(buf, "func New%sRef(ref *%s) *%s", goStructName, cgoSpec, goStructName)
	fmt.Fprintf(buf, `{
		return (*%s)(unsafe.Pointer(ref))
	}`, goStructName)
	name := fmt.Sprintf("New%sRef", goStructName)
	helpers = append(helpers, &Helper{
		Name:        name,
		Description: name + " converts the C object reference into a raw struct reference without wrapping.",
		Source:      buf.String(),
	})

	buf.Reset()
	allocHelper := gen.getAllocMemoryHelper(cgoSpec)
	fmt.Fprintf(buf, "func New%s() *%s", goStructName, goStructName)
	fmt.Fprintf(buf, `{
		return (*%s)(%s(1))
	}`, goStructName, allocHelper.Name)
	name = fmt.Sprintf("New%s", goStructName)
	helpers = append(helpers, &Helper{
		Name: name,
		Description: name + " allocates a new C object of this type and converts the reference into\n" +
			"a raw struct reference without wrapping.",
		Source:   buf.String(),
		Requires: []*Helper{allocHelper},
	})

	buf.Reset()
	fmt.Fprintf(buf, "func (x *%s) PassRef() *%s", goStructName, cgoSpec)
	fmt.Fprintf(buf, `{
		if x == nil {
			x = (*%s)(%s(1))
		}
		return (*%s)(unsafe.Pointer(x))
	}`, goStructName, allocHelper.Name, cgoSpec)
	helpers = append(helpers, &Helper{
		Name:        fmt.Sprintf("%s.PassRef", goStructName),
		Description: "PassRef returns a reference to C object as it is or allocates a new C object of this type.",
		Source:      buf.String(),
		Requires:    []*Helper{allocHelper},
	})
	return
}

func (gen *Generator) getPassRefSource(goStructName []byte, cStructName string, spec tl.CType) []byte {
	cgoSpec := gen.tr.CGoSpec(spec, false)
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

	ptrTipRx, typeTipRx, memTipRx := gen.tr.TipRxsForSpec(tl.TipScopeStruct, cStructName, spec)
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
		ptrTip := ptrTipRx.TipAt(i)
		if memTip == tl.TipMemRaw {
			ptrTip = tl.TipPtrSRef
		}
		typeTip := typeTipRx.TipAt(i)
		goSpec := gen.tr.TranslateSpec(m.Spec, ptrTip, typeTip)
		cgoSpec := gen.tr.CGoSpec(m.Spec, false)
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

	ptrTipRx, typeTipRx, memTipRx := gen.tr.TipRxsForSpec(tl.TipScopeStruct, cStructName, spec)
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
		ptrTip := ptrTipRx.TipAt(i)
		if memTip == tl.TipMemRaw {
			ptrTip = tl.TipPtrSRef
		}
		typeTip := typeTipRx.TipAt(i)
		goSpec := gen.tr.TranslateSpec(m.Spec, ptrTip, typeTip)
		const public = true
		goName := "x." + string(gen.tr.TransformName(tl.TargetType, m.Name, public))
		cgoName := fmt.Sprintf("x.ref%2x.%s", crc, m.Name)
		cgoSpec := gen.tr.CGoSpec(m.Spec, false)
		toProxy, _ := gen.proxyValueToGo(memTip, goName, cgoName, goSpec, cgoSpec)
		fmt.Fprintln(buf, toProxy)
	}
	return buf.Bytes()
}
