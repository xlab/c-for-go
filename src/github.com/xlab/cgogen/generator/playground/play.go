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
	passData(d)
	log.Println(d)
}
