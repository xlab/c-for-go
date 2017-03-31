package parser

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/cznic/cc"
)

func TestParse(t *testing.T) {
	unit, err := ParseWith(NewConfig("testdata/parser_test.h"))
	if err != nil {
		t.Fatal(err)
	} else if len(defines["LOL"]) == 0 {
		t.Fatal("LOL is not defined")
	}
	testUnit(t, unit)
}

func testUnit(t *testing.T, u *cc.TranslationUnit) {
	buf, err := ioutil.ReadFile("testdata/parser_test.out")
	if err != nil {
		t.Fatal(err)
	} else if len(buf) == 0 {
		t.Fatal("no reference output provided")
	}
	if u == nil {
		t.Fatal("no translation unit returned")
	}
	if u.String() != string(buf) {
		log.Println(u)
		t.Error("output doesn't match reference")
	}
}
