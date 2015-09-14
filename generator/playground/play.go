package main

// #include "play.h"
import "C"
import "log"

func plus(a int) int {
	return a + 5
}

func main() {
	Lol(plus, 5)
	Lol(plus, 5)
}

// ------ generated below

func Lol(cb Fcb, a int) {
	log.Println("old cb:", fcbcStorage, "new cb:", cb)
	setFcbx(cb)
	ccb := cb.PassRef()
	C.lol(ccb, C.int(a))
}

type Fcb func(_ int) int

func (cb Fcb) PassRef() *C.fcb {
	return (*C.fcb)(C.fcbx)
}

//export fcbc
func fcbc(a int) int {
	if fcbcStorage != nil {
		return fcbcStorage(a)
	}
	panic("callback is nil")
}

func setFcbx(cb Fcb) {
	fcbcStorage = cb
}

var fcbcStorage Fcb
