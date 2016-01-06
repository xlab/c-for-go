// WARNING: This file has automatically been generated on Wed, 06 Jan 2016 06:00:04 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package foo

/*
#include "foo.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

const (
	// IDLen as defined in test/foo.h:28
	IDLen = 4 + 1
)

// Fcb type as declared in test/foo.h:53
type Fcb func(_ int) int

// HasAnonTag as declared in test/foo.h:39
type HasAnonTag struct {
	A              int
	B              InnerAnonTag
	ref49edb189    *C.struct_foo_has_anon_tag
	allocs49edb189 interface{}
}

// AnonTag as declared in test/foo.h:33
type AnonTag struct {
	N              int
	ref3b932d23    *C.struct_foo_anon_tag
	allocs3b932d23 interface{}
}

// InnerAnonTag as declared in test/foo.h:43
type InnerAnonTag struct {
	N              int
	ref9a2478a3    *C.struct_foo_inner_anon_tag
	allocs9a2478a3 interface{}
}

// PassInt function as declared in test/foo.h:8
func PassInt(i1 int, i2 int) int {
	ci1, _ := (C.int)(i1), cgoAllocsUnknown
	ci2, _ := (C.int)(i2), cgoAllocsUnknown
	ret := C.foo_pass_int(ci1, ci2)
	v := (int)(ret)
	return v
}

// PassString function as declared in test/foo.h:9
func PassString(s1 string, s2 string) string {
	cs1, _ := unpackPUcharString(s1)
	cs2, _ := unpackPCharString(s2)
	ret := C.foo_pass_string(cs1, cs2)
	v := packPCharString(ret)
	return v
}

// PassBytes function as declared in test/foo.h:10
func PassBytes(b1 []byte, n1 uint, b2 []byte, n2 uint) []byte {
	cb1, _ := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&b1)).Data)), cgoAllocsUnknown
	cn1, _ := (C.size_t)(n1), cgoAllocsUnknown
	cb2, _ := (*C.uint8_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&b2)).Data)), cgoAllocsUnknown
	cn2, _ := (C.size_t)(n2), cgoAllocsUnknown
	ret := C.foo_pass_bytes(cb1, cn1, cb2, cn2)
	v := (*(*[0x7fffffff]byte)(unsafe.Pointer(ret)))[:0]
	return v
}

// FindChar function as declared in test/foo.h:11
func FindChar(s []byte, c byte) int {
	cs, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&s)).Data)), cgoAllocsUnknown
	cc, _ := (C.char)(c), cgoAllocsUnknown
	ret := C.foo_find_char(cs, cc)
	v := (int)(ret)
	return v
}

// A4Byte function as declared in test/foo.h:13
func A4Byte(b *[4]byte) {
	cb, _ := *(**C.uint8_t)(unsafe.Pointer(&b)), cgoAllocsUnknown
	C.foo_a4_byte(cb)
}

// A4String function as declared in test/foo.h:14
func A4String(s *[4]string) {
	cs, _ := unpackArgA4String(s)
	C.foo_a4_string(cs)
	packA4String(s, (*[4]*C.char)(unsafe.Pointer(cs)))
}

// A4SByte function as declared in test/foo.h:15
func A4SByte(b *[4][]byte, n uint) {
	cb, _ := unpackArgA4SByte(b)
	cn, _ := (C.size_t)(n), cgoAllocsUnknown
	C.foo_a4_s_byte(cb, cn)
	packA4SByte(b, (*[4]*C.uint8_t)(unsafe.Pointer(cb)))
}

// A4SString function as declared in test/foo.h:16
func A4SString(s *[4][]string, n uint) {
	cs, _ := unpackArgA4SString(s)
	cn, _ := (C.size_t)(n), cgoAllocsUnknown
	C.foo_a4_s_string(cs, cn)
	packA4SString(s, (*[4]**C.char)(unsafe.Pointer(cs)))
}

// A2A2Byte function as declared in test/foo.h:17
func A2A2Byte(b *[2][2]byte) {
	cb, _ := *(**[2]C.uint8_t)(unsafe.Pointer(&b)), cgoAllocsUnknown
	C.foo_a2_a2_byte(cb)
}

// A2A2String function as declared in test/foo.h:18
func A2A2String(s *[2][2]string) {
	cs, _ := unpackArgA2A2String(s)
	C.foo_a2_a2_string(cs)
	packA2A2String(s, (*[2][2]*C.char)(unsafe.Pointer(cs)))
}

// A2A2SByte function as declared in test/foo.h:19
func A2A2SByte(b *[2][2][]byte, n uint) {
	cb, _ := unpackArgA2A2SByte(b)
	cn, _ := (C.size_t)(n), cgoAllocsUnknown
	C.foo_a2_a2_s_byte(cb, cn)
	packA2A2SByte(b, (*[2][2]*C.uint8_t)(unsafe.Pointer(cb)))
}

// A2A2SString function as declared in test/foo.h:20
func A2A2SString(s *[2][2][]string, n uint) {
	cs, _ := unpackArgA2A2SString(s)
	cn, _ := (C.size_t)(n), cgoAllocsUnknown
	C.foo_a2_a2_s_string(cs, cn)
	packA2A2SString(s, (*[2][2]**C.char)(unsafe.Pointer(cs)))
}

// SSByte function as declared in test/foo.h:21
func SSByte(b [][]byte, n1 uint, n2 uint) {
	cb, _ := unpackArgSSByte(b)
	cn1, _ := (C.size_t)(n1), cgoAllocsUnknown
	cn2, _ := (C.size_t)(n2), cgoAllocsUnknown
	C.foo_s_s_byte(cb, cn1, cn2)
	packSSByte(b, cb)
}

// SSString function as declared in test/foo.h:22
func SSString(s [][]string, n1 uint, n2 uint) {
	cs, _ := unpackArgSSString(s)
	cn1, _ := (C.size_t)(n1), cgoAllocsUnknown
	cn2, _ := (C.size_t)(n2), cgoAllocsUnknown
	C.foo_s_s_string(cs, cn1, cn2)
	packSSString(s, cs)
}

// A4SSByte function as declared in test/foo.h:23
func A4SSByte(b *[4][][]byte, n1 uint, n2 uint) {
	cb, _ := unpackArgA4SSByte(b)
	cn1, _ := (C.size_t)(n1), cgoAllocsUnknown
	cn2, _ := (C.size_t)(n2), cgoAllocsUnknown
	C.foo_a4_s_s_byte(cb, cn1, cn2)
	packA4SSByte(b, (*[4]**C.uint8_t)(unsafe.Pointer(cb)))
}

// A4SSString function as declared in test/foo.h:24
func A4SSString(s *[4][][]string, n1 uint, n2 uint) {
	cs, _ := unpackArgA4SSString(s)
	cn1, _ := (C.size_t)(n1), cgoAllocsUnknown
	cn2, _ := (C.size_t)(n2), cgoAllocsUnknown
	C.foo_a4_s_s_string(cs, cn1, cn2)
	packA4SSString(s, (*[4]***C.char)(unsafe.Pointer(cs)))
}

// A2A2SSByte function as declared in test/foo.h:25
func A2A2SSByte(b *[2][2][][]byte, n1 uint, n2 uint) {
	cb, _ := unpackArgA2A2SSByte(b)
	cn1, _ := (C.size_t)(n1), cgoAllocsUnknown
	cn2, _ := (C.size_t)(n2), cgoAllocsUnknown
	C.foo_a2_a2_s_s_byte(cb, cn1, cn2)
	packA2A2SSByte(b, (*[2][2]**C.uint8_t)(unsafe.Pointer(cb)))
}

// A2A2SSString function as declared in test/foo.h:26
func A2A2SSString(s *[2][2][][]string, n1 uint, n2 uint) {
	cs, _ := unpackArgA2A2SSUString(s)
	cn1, _ := (C.size_t)(n1), cgoAllocsUnknown
	cn2, _ := (C.size_t)(n2), cgoAllocsUnknown
	C.foo_a2_a2_s_s_string(cs, cn1, cn2)
	packA2A2SSUString(s, (*[2][2]***C.uint8_t)(unsafe.Pointer(cs)))
}

// PassAnonTag function as declared in test/foo.h:37
func PassAnonTag(a AnonTag, b AnonTag) int {
	ca, _ := a.PassValue()
	cb, _ := b.PassValue()
	ret := C.foo_pass_anon_tag(ca, cb)
	v := (int)(ret)
	return v
}
