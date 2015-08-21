package main

// #include "play.h"
import "C"
import (
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

func cStr(str string) *C.char {
	h := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data))
}

func passData(d Data) {
	C.print_names(d.toC())
}

func main() {
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
	c, orig := unpackSlicedStr(cube)
	C.stringCube(c, C.int(x), C.int(y), C.int(z))
	log.Println(orig)
}
