// WARNING: This file has automatically been generated on Wed, 06 Jan 2016 06:00:04 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package foo

/*
#include "foo.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// PassRef returns a reference.
func (x Fcb) PassRef() (ref *C.foo_fcb, allocs *cgoAllocMap) {
	if fcbE3897987Func == nil {
		fcbE3897987Func = x
	}
	return (*C.foo_fcb)(C.foo_fcb_e3897987), nil
}

//export fcbE3897987
func fcbE3897987(carg0 C.int) C.int {
	if fcbE3897987Func != nil {
		arg0 := (int)(carg0)
		rete3897987 := fcbE3897987Func(arg0)
		ret, _ := (C.int)(rete3897987), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var fcbE3897987Func Fcb

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocStruct_HasAnonTagMemory allocates memory for type C.struct_foo_has_anon_tag in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStruct_HasAnonTagMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStruct_HasAnonTagValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStruct_HasAnonTagValue = unsafe.Sizeof([1]C.struct_foo_has_anon_tag{})

// Ref returns a reference.
func (x *HasAnonTag) Ref() *C.struct_foo_has_anon_tag {
	if x == nil {
		return nil
	}
	return x.ref49edb189
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *HasAnonTag) Free() {
	if x != nil && x.allocs49edb189 != nil {
		runtime.SetFinalizer(x, nil)
		x.allocs49edb189.(*cgoAllocMap).Free()
		x.ref49edb189 = nil
	}
}

// NewHasAnonTagRef initialises a new struct holding the reference to the originaitng C struct.
func NewHasAnonTagRef(ref *C.struct_foo_has_anon_tag) *HasAnonTag {
	if ref == nil {
		return nil
	}
	obj := new(HasAnonTag)
	obj.ref49edb189 = ref
	// enable this if the reference is unmanaged:
	// runtime.SetFinalizer(obj, func(x *HasAnonTag) {
	// 	C.free(unsafe.Pointer(x.ref49edb189))
	// })
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *HasAnonTag) PassRef() (ref *C.struct_foo_has_anon_tag, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref49edb189 != nil {
		return x.ref49edb189, nil
	}
	mem49edb189 := allocStruct_HasAnonTagMemory(1)
	ref49edb189 := (*C.struct_foo_has_anon_tag)(mem49edb189)
	allocs49edb189 := new(cgoAllocMap)
	var ca_allocs *cgoAllocMap
	ref49edb189.a, ca_allocs = (C.int)(x.A), cgoAllocsUnknown
	allocs49edb189.Borrow(ca_allocs)

	var cb_allocs *cgoAllocMap
	ref49edb189.b, cb_allocs = x.B.PassValue()
	allocs49edb189.Borrow(cb_allocs)

	x.ref49edb189 = ref49edb189
	x.allocs49edb189 = allocs49edb189
	return ref49edb189, allocs49edb189

}

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x *HasAnonTag) PassValue() (value C.struct_foo_has_anon_tag, allocs *cgoAllocMap) {
	if x == nil {
		x = NewHasAnonTagRef(nil)
	} else if x.ref49edb189 != nil {
		return *x.ref49edb189, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *HasAnonTag) Deref() {
	if x.ref49edb189 == nil {
		return
	}
	x.A = (int)(x.ref49edb189.a)
	x.B = *NewInnerAnonTagRef(&x.ref49edb189.b)
}

// allocStruct_AnonTagMemory allocates memory for type C.struct_foo_anon_tag in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStruct_AnonTagMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStruct_AnonTagValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStruct_AnonTagValue = unsafe.Sizeof([1]C.struct_foo_anon_tag{})

// Ref returns a reference.
func (x *AnonTag) Ref() *C.struct_foo_anon_tag {
	if x == nil {
		return nil
	}
	return x.ref3b932d23
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *AnonTag) Free() {
	if x != nil && x.allocs3b932d23 != nil {
		runtime.SetFinalizer(x, nil)
		x.allocs3b932d23.(*cgoAllocMap).Free()
		x.ref3b932d23 = nil
	}
}

// NewAnonTagRef initialises a new struct holding the reference to the originaitng C struct.
func NewAnonTagRef(ref *C.struct_foo_anon_tag) *AnonTag {
	if ref == nil {
		return nil
	}
	obj := new(AnonTag)
	obj.ref3b932d23 = ref
	// enable this if the reference is unmanaged:
	// runtime.SetFinalizer(obj, func(x *AnonTag) {
	// 	C.free(unsafe.Pointer(x.ref3b932d23))
	// })
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *AnonTag) PassRef() (ref *C.struct_foo_anon_tag, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3b932d23 != nil {
		return x.ref3b932d23, nil
	}
	mem3b932d23 := allocStruct_AnonTagMemory(1)
	ref3b932d23 := (*C.struct_foo_anon_tag)(mem3b932d23)
	allocs3b932d23 := new(cgoAllocMap)
	var cn_allocs *cgoAllocMap
	ref3b932d23.n, cn_allocs = (C.int)(x.N), cgoAllocsUnknown
	allocs3b932d23.Borrow(cn_allocs)

	x.ref3b932d23 = ref3b932d23
	x.allocs3b932d23 = allocs3b932d23
	return ref3b932d23, allocs3b932d23

}

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x *AnonTag) PassValue() (value C.struct_foo_anon_tag, allocs *cgoAllocMap) {
	if x == nil {
		x = NewAnonTagRef(nil)
	} else if x.ref3b932d23 != nil {
		return *x.ref3b932d23, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *AnonTag) Deref() {
	if x.ref3b932d23 == nil {
		return
	}
	x.N = (int)(x.ref3b932d23.n)
}

// allocStruct_InnerAnonTagMemory allocates memory for type C.struct_foo_inner_anon_tag in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStruct_InnerAnonTagMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStruct_InnerAnonTagValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStruct_InnerAnonTagValue = unsafe.Sizeof([1]C.struct_foo_inner_anon_tag{})

// Ref returns a reference.
func (x *InnerAnonTag) Ref() *C.struct_foo_inner_anon_tag {
	if x == nil {
		return nil
	}
	return x.ref9a2478a3
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *InnerAnonTag) Free() {
	if x != nil && x.allocs9a2478a3 != nil {
		runtime.SetFinalizer(x, nil)
		x.allocs9a2478a3.(*cgoAllocMap).Free()
		x.ref9a2478a3 = nil
	}
}

// NewInnerAnonTagRef initialises a new struct holding the reference to the originaitng C struct.
func NewInnerAnonTagRef(ref *C.struct_foo_inner_anon_tag) *InnerAnonTag {
	if ref == nil {
		return nil
	}
	obj := new(InnerAnonTag)
	obj.ref9a2478a3 = ref
	// enable this if the reference is unmanaged:
	// runtime.SetFinalizer(obj, func(x *InnerAnonTag) {
	// 	C.free(unsafe.Pointer(x.ref9a2478a3))
	// })
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *InnerAnonTag) PassRef() (ref *C.struct_foo_inner_anon_tag, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9a2478a3 != nil {
		return x.ref9a2478a3, nil
	}
	mem9a2478a3 := allocStruct_InnerAnonTagMemory(1)
	ref9a2478a3 := (*C.struct_foo_inner_anon_tag)(mem9a2478a3)
	allocs9a2478a3 := new(cgoAllocMap)
	var cn_allocs *cgoAllocMap
	ref9a2478a3.n, cn_allocs = (C.int)(x.N), cgoAllocsUnknown
	allocs9a2478a3.Borrow(cn_allocs)

	x.ref9a2478a3 = ref9a2478a3
	x.allocs9a2478a3 = allocs9a2478a3
	return ref9a2478a3, allocs9a2478a3

}

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x *InnerAnonTag) PassValue() (value C.struct_foo_inner_anon_tag, allocs *cgoAllocMap) {
	if x == nil {
		x = NewInnerAnonTagRef(nil)
	} else if x.ref9a2478a3 != nil {
		return *x.ref9a2478a3, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *InnerAnonTag) Deref() {
	if x.ref9a2478a3 == nil {
		return
	}
	x.N = (int)(x.ref9a2478a3.n)
}

// unpackPUcharString represents the data from Go string as *C.uchar and avoids copying.
func unpackPUcharString(str string) (*C.uchar, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.uchar)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// allocA4PCharMemory allocates memory for type [4]*C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA4PCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA4PCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA4PCharValue = unsafe.Sizeof([1][4]*C.char{})

// unpackArgA4String transforms a sliced Go data structure into plain C format.
func unpackArgA4String(x *[4]string) (unpacked **C.char, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	mem0 := allocA4PCharMemory(1)
	allocs.Add(mem0)
	v0 := (*[4]*C.char)(mem0)
	for i0 := range x {
		v0[i0], _ = unpackPCharString(x[i0])
	}
	unpacked = (**C.char)(mem0)
	return
}

// packA4String reads sliced Go data structure out from plain C format.
func packA4String(v *[4]string, ptr0 *[4]*C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := ptr0[i0]
		v[i0] = packPCharString(ptr1)
	}
}

// allocA4PUint8Memory allocates memory for type [4]*C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA4PUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA4PUint8Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA4PUint8Value = unsafe.Sizeof([1][4]*C.uint8_t{})

// unpackArgA4SByte transforms a sliced Go data structure into plain C format.
func unpackArgA4SByte(x *[4][]byte) (unpacked **C.uint8_t, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.uint8_t) {
		go allocs.Free()
	})

	mem0 := allocA4PUint8Memory(1)
	allocs.Add(mem0)
	v0 := (*[4]*C.uint8_t)(mem0)
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.uint8_t)(unsafe.Pointer(h.Data))
	}
	unpacked = (**C.uint8_t)(mem0)
	return
}

// packA4SByte reads sliced Go data structure out from plain C format.
func packA4SByte(v *[4][]byte, ptr0 *[4]*C.uint8_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := ptr0[i0]
		hxfc4425b := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxfc4425b.Data = uintptr(unsafe.Pointer(ptr1))
		hxfc4425b.Cap = 0x7fffffff
		// hxfc4425b.Len = ?
	}
}

// allocA4PPCharMemory allocates memory for type [4]**C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA4PPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA4PPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA4PPCharValue = unsafe.Sizeof([1][4]**C.char{})

// allocPCharMemory allocates memory for type *C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPCharValue = unsafe.Sizeof([1]*C.char{})

// unpackArgA4SString transforms a sliced Go data structure into plain C format.
func unpackArgA4SString(x *[4][]string) (unpacked ***C.char, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(****C.char) {
		go allocs.Free()
	})

	mem0 := allocA4PPCharMemory(1)
	allocs.Add(mem0)
	v0 := (*[4]**C.char)(mem0)
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocPCharMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]*C.char)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			v1[i1], _ = unpackPCharString(x[i0][i1])
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (**C.char)(unsafe.Pointer(h.Data))
	}
	unpacked = (***C.char)(mem0)
	return
}

// packA4SString reads sliced Go data structure out from plain C format.
func packA4SString(v *[4][]string, ptr0 *[4]**C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := ptr0[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m]*C.char)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = packPCharString(ptr2)
		}
	}
}

// allocA2A2PCharMemory allocates memory for type [2][2]*C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA2A2PCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA2A2PCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA2A2PCharValue = unsafe.Sizeof([1][2][2]*C.char{})

// allocA2PCharMemory allocates memory for type [2]*C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA2PCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA2PCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA2PCharValue = unsafe.Sizeof([1][2]*C.char{})

// unpackArgA2A2String transforms a sliced Go data structure into plain C format.
func unpackArgA2A2String(x *[2][2]string) (unpacked *[2]*C.char, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**[2]*C.char) {
		go allocs.Free()
	})

	mem0 := allocA2A2PCharMemory(1)
	allocs.Add(mem0)
	v0 := (*[2][2]*C.char)(mem0)
	for i0 := range x {
		mem1 := allocA2PCharMemory(1)
		allocs.Add(mem1)
		v1 := (*[2]*C.char)(mem1)
		for i1 := range x[i0] {
			v1[i1], _ = unpackPCharString(x[i0][i1])
		}
		v0[i0] = *(*[2]*C.char)(mem1)
	}
	unpacked = (*[2]*C.char)(mem0)
	return
}

// packA2A2String reads sliced Go data structure out from plain C format.
func packA2A2String(v *[2][2]string, ptr0 *[2][2]*C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := ptr0[i0]
		for i1 := range v[i0] {
			ptr2 := ptr1[i1]
			v[i0][i1] = packPCharString(ptr2)
		}
	}
}

// allocA2A2PUint8Memory allocates memory for type [2][2]*C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA2A2PUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA2A2PUint8Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA2A2PUint8Value = unsafe.Sizeof([1][2][2]*C.uint8_t{})

// allocA2PUint8Memory allocates memory for type [2]*C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA2PUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA2PUint8Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA2PUint8Value = unsafe.Sizeof([1][2]*C.uint8_t{})

// unpackArgA2A2SByte transforms a sliced Go data structure into plain C format.
func unpackArgA2A2SByte(x *[2][2][]byte) (unpacked *[2]*C.uint8_t, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**[2]*C.uint8_t) {
		go allocs.Free()
	})

	mem0 := allocA2A2PUint8Memory(1)
	allocs.Add(mem0)
	v0 := (*[2][2]*C.uint8_t)(mem0)
	for i0 := range x {
		mem1 := allocA2PUint8Memory(1)
		allocs.Add(mem1)
		v1 := (*[2]*C.uint8_t)(mem1)
		for i1 := range x[i0] {
			h := (*sliceHeader)(unsafe.Pointer(&x[i0][i1]))
			v1[i1] = (*C.uint8_t)(unsafe.Pointer(h.Data))
		}
		v0[i0] = *(*[2]*C.uint8_t)(mem1)
	}
	unpacked = (*[2]*C.uint8_t)(mem0)
	return
}

// packA2A2SByte reads sliced Go data structure out from plain C format.
func packA2A2SByte(v *[2][2][]byte, ptr0 *[2][2]*C.uint8_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := ptr0[i0]
		for i1 := range v[i0] {
			ptr2 := ptr1[i1]
			hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&v[i0][i1]))
			hxf95e7c8.Data = uintptr(unsafe.Pointer(ptr2))
			hxf95e7c8.Cap = 0x7fffffff
			// hxf95e7c8.Len = ?
		}
	}
}

// allocA2A2PPCharMemory allocates memory for type [2][2]**C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA2A2PPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA2A2PPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA2A2PPCharValue = unsafe.Sizeof([1][2][2]**C.char{})

// allocA2PPCharMemory allocates memory for type [2]**C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA2PPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA2PPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA2PPCharValue = unsafe.Sizeof([1][2]**C.char{})

// unpackArgA2A2SString transforms a sliced Go data structure into plain C format.
func unpackArgA2A2SString(x *[2][2][]string) (unpacked *[2]**C.char, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**[2]**C.char) {
		go allocs.Free()
	})

	mem0 := allocA2A2PPCharMemory(1)
	allocs.Add(mem0)
	v0 := (*[2][2]**C.char)(mem0)
	for i0 := range x {
		mem1 := allocA2PPCharMemory(1)
		allocs.Add(mem1)
		v1 := (*[2]**C.char)(mem1)
		for i1 := range x[i0] {
			len2 := len(x[i0][i1])
			mem2 := allocPCharMemory(len2)
			allocs.Add(mem2)
			h2 := &sliceHeader{
				Data: uintptr(mem2),
				Cap:  len2,
				Len:  len2,
			}
			v2 := *(*[]*C.char)(unsafe.Pointer(h2))
			for i2 := range x[i0][i1] {
				v2[i2], _ = unpackPCharString(x[i0][i1][i2])
			}
			h := (*sliceHeader)(unsafe.Pointer(&v2))
			v1[i1] = (**C.char)(unsafe.Pointer(h.Data))
		}
		v0[i0] = *(*[2]**C.char)(mem1)
	}
	unpacked = (*[2]**C.char)(mem0)
	return
}

// packA2A2SString reads sliced Go data structure out from plain C format.
func packA2A2SString(v *[2][2][]string, ptr0 *[2][2]**C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := ptr0[i0]
		for i1 := range v[i0] {
			ptr2 := ptr1[i1]
			for i2 := range v[i0][i1] {
				ptr3 := (*(*[m]*C.char)(unsafe.Pointer(ptr2)))[i2]
				v[i0][i1][i2] = packPCharString(ptr3)
			}
		}
	}
}

// allocPUint8Memory allocates memory for type *C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPUint8Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPUint8Value = unsafe.Sizeof([1]*C.uint8_t{})

// unpackArgSSByte transforms a sliced Go data structure into plain C format.
func unpackArgSSByte(x [][]byte) (unpacked **C.uint8_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.uint8_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPUint8Memory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.uint8_t)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.uint8_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.uint8_t)(unsafe.Pointer(h.Data))
	return
}

// packSSByte reads sliced Go data structure out from plain C format.
func packSSByte(v [][]byte, ptr0 **C.uint8_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m]*C.uint8_t)(unsafe.Pointer(ptr0)))[i0]
		hxff2234b := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxff2234b.Data = uintptr(unsafe.Pointer(ptr1))
		hxff2234b.Cap = 0x7fffffff
		// hxff2234b.Len = ?
	}
}

// allocPPCharMemory allocates memory for type **C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPPCharValue = unsafe.Sizeof([1]**C.char{})

// unpackArgSSString transforms a sliced Go data structure into plain C format.
func unpackArgSSString(x [][]string) (unpacked ***C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(****C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]**C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocPCharMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]*C.char)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			v1[i1], _ = unpackPCharString(x[i0][i1])
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (**C.char)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (***C.char)(unsafe.Pointer(h.Data))
	return
}

// packSSString reads sliced Go data structure out from plain C format.
func packSSString(v [][]string, ptr0 ***C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m]**C.char)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m]*C.char)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = packPCharString(ptr2)
		}
	}
}

// allocA4PPUint8Memory allocates memory for type [4]**C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA4PPUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA4PPUint8Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA4PPUint8Value = unsafe.Sizeof([1][4]**C.uint8_t{})

// unpackArgA4SSByte transforms a sliced Go data structure into plain C format.
func unpackArgA4SSByte(x *[4][][]byte) (unpacked ***C.uint8_t, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(****C.uint8_t) {
		go allocs.Free()
	})

	mem0 := allocA4PPUint8Memory(1)
	allocs.Add(mem0)
	v0 := (*[4]**C.uint8_t)(mem0)
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocPUint8Memory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]*C.uint8_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			h := (*sliceHeader)(unsafe.Pointer(&x[i0][i1]))
			v1[i1] = (*C.uint8_t)(unsafe.Pointer(h.Data))
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (**C.uint8_t)(unsafe.Pointer(h.Data))
	}
	unpacked = (***C.uint8_t)(mem0)
	return
}

// packA4SSByte reads sliced Go data structure out from plain C format.
func packA4SSByte(v *[4][][]byte, ptr0 *[4]**C.uint8_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := ptr0[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m]*C.uint8_t)(unsafe.Pointer(ptr1)))[i1]
			hxff73280 := (*sliceHeader)(unsafe.Pointer(&v[i0][i1]))
			hxff73280.Data = uintptr(unsafe.Pointer(ptr2))
			hxff73280.Cap = 0x7fffffff
			// hxff73280.Len = ?
		}
	}
}

// allocA4PPPCharMemory allocates memory for type [4]***C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA4PPPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA4PPPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA4PPPCharValue = unsafe.Sizeof([1][4]***C.char{})

// unpackArgA4SSString transforms a sliced Go data structure into plain C format.
func unpackArgA4SSString(x *[4][][]string) (unpacked ****C.char, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(*****C.char) {
		go allocs.Free()
	})

	mem0 := allocA4PPPCharMemory(1)
	allocs.Add(mem0)
	v0 := (*[4]***C.char)(mem0)
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocPPCharMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]**C.char)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			len2 := len(x[i0][i1])
			mem2 := allocPCharMemory(len2)
			allocs.Add(mem2)
			h2 := &sliceHeader{
				Data: uintptr(mem2),
				Cap:  len2,
				Len:  len2,
			}
			v2 := *(*[]*C.char)(unsafe.Pointer(h2))
			for i2 := range x[i0][i1] {
				v2[i2], _ = unpackPCharString(x[i0][i1][i2])
			}
			h := (*sliceHeader)(unsafe.Pointer(&v2))
			v1[i1] = (**C.char)(unsafe.Pointer(h.Data))
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (***C.char)(unsafe.Pointer(h.Data))
	}
	unpacked = (****C.char)(mem0)
	return
}

// packA4SSString reads sliced Go data structure out from plain C format.
func packA4SSString(v *[4][][]string, ptr0 *[4]***C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := ptr0[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m]**C.char)(unsafe.Pointer(ptr1)))[i1]
			for i2 := range v[i0][i1] {
				ptr3 := (*(*[m]*C.char)(unsafe.Pointer(ptr2)))[i2]
				v[i0][i1][i2] = packPCharString(ptr3)
			}
		}
	}
}

// allocA2A2PPUint8Memory allocates memory for type [2][2]**C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA2A2PPUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA2A2PPUint8Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA2A2PPUint8Value = unsafe.Sizeof([1][2][2]**C.uint8_t{})

// allocA2PPUint8Memory allocates memory for type [2]**C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA2PPUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA2PPUint8Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA2PPUint8Value = unsafe.Sizeof([1][2]**C.uint8_t{})

// unpackArgA2A2SSByte transforms a sliced Go data structure into plain C format.
func unpackArgA2A2SSByte(x *[2][2][][]byte) (unpacked *[2]**C.uint8_t, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**[2]**C.uint8_t) {
		go allocs.Free()
	})

	mem0 := allocA2A2PPUint8Memory(1)
	allocs.Add(mem0)
	v0 := (*[2][2]**C.uint8_t)(mem0)
	for i0 := range x {
		mem1 := allocA2PPUint8Memory(1)
		allocs.Add(mem1)
		v1 := (*[2]**C.uint8_t)(mem1)
		for i1 := range x[i0] {
			len2 := len(x[i0][i1])
			mem2 := allocPUint8Memory(len2)
			allocs.Add(mem2)
			h2 := &sliceHeader{
				Data: uintptr(mem2),
				Cap:  len2,
				Len:  len2,
			}
			v2 := *(*[]*C.uint8_t)(unsafe.Pointer(h2))
			for i2 := range x[i0][i1] {
				h := (*sliceHeader)(unsafe.Pointer(&x[i0][i1][i2]))
				v2[i2] = (*C.uint8_t)(unsafe.Pointer(h.Data))
			}
			h := (*sliceHeader)(unsafe.Pointer(&v2))
			v1[i1] = (**C.uint8_t)(unsafe.Pointer(h.Data))
		}
		v0[i0] = *(*[2]**C.uint8_t)(mem1)
	}
	unpacked = (*[2]**C.uint8_t)(mem0)
	return
}

// packA2A2SSByte reads sliced Go data structure out from plain C format.
func packA2A2SSByte(v *[2][2][][]byte, ptr0 *[2][2]**C.uint8_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := ptr0[i0]
		for i1 := range v[i0] {
			ptr2 := ptr1[i1]
			for i2 := range v[i0][i1] {
				ptr3 := (*(*[m]*C.uint8_t)(unsafe.Pointer(ptr2)))[i2]
				hxfa9955c := (*sliceHeader)(unsafe.Pointer(&v[i0][i1][i2]))
				hxfa9955c.Data = uintptr(unsafe.Pointer(ptr3))
				hxfa9955c.Cap = 0x7fffffff
				// hxfa9955c.Len = ?
			}
		}
	}
}

// allocA2A2PPPUint8Memory allocates memory for type [2][2]***C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA2A2PPPUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA2A2PPPUint8Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA2A2PPPUint8Value = unsafe.Sizeof([1][2][2]***C.uint8_t{})

// allocA2PPPUint8Memory allocates memory for type [2]***C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA2PPPUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA2PPPUint8Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA2PPPUint8Value = unsafe.Sizeof([1][2]***C.uint8_t{})

// allocPPUint8Memory allocates memory for type **C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPPUint8Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPPUint8Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPPUint8Value = unsafe.Sizeof([1]**C.uint8_t{})

// unpackArgA2A2SSUString transforms a sliced Go data structure into plain C format.
func unpackArgA2A2SSUString(x *[2][2][][]string) (unpacked *[2]***C.uint8_t, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**[2]***C.uint8_t) {
		go allocs.Free()
	})

	mem0 := allocA2A2PPPUint8Memory(1)
	allocs.Add(mem0)
	v0 := (*[2][2]***C.uint8_t)(mem0)
	for i0 := range x {
		mem1 := allocA2PPPUint8Memory(1)
		allocs.Add(mem1)
		v1 := (*[2]***C.uint8_t)(mem1)
		for i1 := range x[i0] {
			len2 := len(x[i0][i1])
			mem2 := allocPPUint8Memory(len2)
			allocs.Add(mem2)
			h2 := &sliceHeader{
				Data: uintptr(mem2),
				Cap:  len2,
				Len:  len2,
			}
			v2 := *(*[]**C.uint8_t)(unsafe.Pointer(h2))
			for i2 := range x[i0][i1] {
				len3 := len(x[i0][i1][i2])
				mem3 := allocPUint8Memory(len3)
				allocs.Add(mem3)
				h3 := &sliceHeader{
					Data: uintptr(mem3),
					Cap:  len3,
					Len:  len3,
				}
				v3 := *(*[]*C.uint8_t)(unsafe.Pointer(h3))
				for i3 := range x[i0][i1][i2] {
					v3[i3], _ = unpackPUint8String(x[i0][i1][i2][i3])
				}
				h := (*sliceHeader)(unsafe.Pointer(&v3))
				v2[i2] = (**C.uint8_t)(unsafe.Pointer(h.Data))
			}
			h := (*sliceHeader)(unsafe.Pointer(&v2))
			v1[i1] = (***C.uint8_t)(unsafe.Pointer(h.Data))
		}
		v0[i0] = *(*[2]***C.uint8_t)(mem1)
	}
	unpacked = (*[2]***C.uint8_t)(mem0)
	return
}

// unpackPUint8String represents the data from Go string as *C.uint8_t and avoids copying.
func unpackPUint8String(str string) (*C.uint8_t, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.uint8_t)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

// packA2A2SSUString reads sliced Go data structure out from plain C format.
func packA2A2SSUString(v *[2][2][][]string, ptr0 *[2][2]***C.uint8_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := ptr0[i0]
		for i1 := range v[i0] {
			ptr2 := ptr1[i1]
			for i2 := range v[i0][i1] {
				ptr3 := (*(*[m]**C.uint8_t)(unsafe.Pointer(ptr2)))[i2]
				for i3 := range v[i0][i1][i2] {
					ptr4 := (*(*[m]*C.uint8_t)(unsafe.Pointer(ptr3)))[i3]
					v[i0][i1][i2][i3] = packPUint8String(ptr4)
				}
			}
		}
	}
}

// packPUint8String creates a Go string backed by *C.uint8_t and avoids copying.
func packPUint8String(p *C.uint8_t) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}
