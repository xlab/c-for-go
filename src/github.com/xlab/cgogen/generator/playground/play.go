package main

// #include "play.h"
import "C"
import (
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

const sizeofData = unsafe.Sizeof(C.data_t{})

type Data struct {
	Names []string
	Size  int
	Flag  bool
}

type Lol struct {
	Flag [4]bool
}

// goStr creates a string backed by *C.char and avoids copying.
func goStr(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*reflect.StringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

func packSlicedStr(mem0 [][][]string, ptr0 ****C.char, ns ...int) {
	const m = 1 << 30
	for i0, mem1 := range mem0 {
		ptr1 := (*(*[m]***C.char)(unsafe.Pointer(ptr0)))[i0]
		for i1, mem2 := range mem1 {
			ptr2 := (*(*[m]**C.char)(unsafe.Pointer(ptr1)))[i1]
			for i2 := range mem2 {
				ptr3 := (*(*[m]*C.char)(unsafe.Pointer(ptr2)))[i2]
				if len(ns) > 0 {
					mem2[i2] = goStrN(ptr3, ns[0])
					continue
				}
				mem2[i2] = goStr(ptr3)
			}
		}
	}
}

func unpackSlicedStr(s [][][]string) ****C.char {
	mem0 := make([]***C.char, len(s))
	for i0, s0 := range s {
		mem1 := make([]**C.char, len(s0))
		for i1, s1 := range s0 {
			mem2 := make([]*C.char, len(s1))
			for i2, o2 := range s1 {
				mem2[i2] = cStr(o2)
			}
			mem1[i1] = (**C.char)(unsafe.Pointer(&mem2[0]))
		}
		mem0[i0] = (***C.char)(unsafe.Pointer(&mem1[0]))
	}
	return (****C.char)(unsafe.Pointer(&mem0[0]))
}

func (d *Data) toC() *C.data_t {
	var mem [sizeofData]byte
	cobj := (*C.data_t)(unsafe.Pointer(&mem))
	cobj.names = cStrSlice(d.Names)
	cobj.size = C.size_t(d.Size)
	return cobj
}

func cStrSlice(s []string) **C.char {
	mem := make([]*C.char, len(s))
	for i, str := range s {
		mem[i] = cStr(str)
	}
	return (**C.char)(unsafe.Pointer(&mem[0]))
}

// goStrN creates a string backed by *C.char and avoids copying.
func goStrN(p *C.char, n int) (raw string) {
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
}

func cStr(str string) *C.char {
	h := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data))
}

func passData(d Data) {
	C.print_names(d.toC())
}

func printA() {
	fmt.Printf("type of a1 is %T\n", C.a1)
	fmt.Printf("type of a2 is %T\n", C.a2)
	fmt.Printf("type of a3 is %T\n", C.a3)
	fmt.Printf("type of a4 is %T\n", C.a4)
	fmt.Printf("type of a5 is %T\n", C.a5)
	fmt.Printf("type of a6 is %T\n", C.a6)
	fmt.Printf("type of a7 is %T\n", C.a7)
	fmt.Printf("type of a8 is %T\n", C.a8)
	fmt.Printf("type of a9 is %T\n", C.a9)
	fmt.Printf("type of a10 is %T\n", C.a10)
	fmt.Printf("type of a11 is %T\n", C.a11)
}

func main() {
	printA()
	names := []string{"Maxim", "Xlab", "Yo", "Lol"}
	d := Data{Names: names, Size: len(names)}
	//C.lol((*C.lol_t)(unsafe.Pointer(&Lol{Flag: [4]bool{true, true, false, true}})))

	cube := [][][]string{
		[][]string{
			[]string{"a1", "b1"},
			[]string{"c1", "d1"},
			[]string{"i1", "j1"},
			[]string{"k1", "l1"},
		},
		[][]string{
			[]string{"a2", "b2"},
			[]string{"c2", "d2"},
			[]string{"i2", "j2"},
			[]string{"k2", "l2"},
		},
		[][]string{
			[]string{"a3", "b3"},
			[]string{"c3", "d3"},
			[]string{"i3", "j3"},
			[]string{"k3", "l3"},
		},
	}
	passCube(cube, 3, 4, 2)

	passData(d)
	log.Println(d)
}

func passCube(cube [][][]string, x, y, z int) {
	ptr0 := unpackSlicedStr(cube)
	C.stringCube(ptr0, C.int(x), C.int(y), C.int(z))
	packSlicedStr(cube, ptr0, x, y, z)
	//log.Println(cube)
}
