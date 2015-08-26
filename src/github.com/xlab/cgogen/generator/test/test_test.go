package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestPassInt(t *testing.T) {
	expected := 5
	result := TestPassInt(2, 3)
	assert.Equal(t, expected, result)
}

func TestTestPassString(t *testing.T) {
	expected := "ab"
	result := TestPassString("a", "b")
	assert.Equal(t, expected, result)
}

func TestTestPassBytes(t *testing.T) {
	expected := []byte("abcd")
	result := TestPassBytes([]byte("ab"), 2, []byte("cd"), 2)
	assert.Equal(t, expected, result[:4])
}

func TestTestFindChar(t *testing.T) {
	tbl := []struct {
		Input  string
		Search byte
		Index  int64
	}{
		{"golang", 'c', -1},
		{"gopher", 'g', 0},
		{"fun", 'n', 2},
	}
	for _, test := range tbl {
		result := TestFindChar(test.Input, test.Search)
		assert.Equal(t, test.Index, result)
	}
}

func TestTestA4Byte(t *testing.T) {
	a := [4]byte{'a', 'a', 'a', 'a'}
	b := [4]byte{'b', 'b', 'b', 'b'}
	TestA4Byte(a)
	assert.Equal(t, b, a)
}

func TestTestA4String(t *testing.T) {
	a := [4]string{"g", "g", "g", "g"}
	b := [4]string{"go", "go", "go", "go"}
	TestA4String(a)
	assert.Equal(t, b, a)
}

func TestTestA4SByte(t *testing.T) {
	a := [4][]byte{
		{'a', 'a'},
		{'a', 'a'},
		{'a', 'a'},
		{'a', 'a'},
	}
	b := [4][]byte{
		{'b', 'b'},
		{'b', 'b'},
		{'b', 'b'},
		{'b', 'b'},
	}
	TestA4SByte(a, 2)
	assert.Equal(t, b, a)
}

func TestTestA4SString(t *testing.T) {
	a := [4][]string{
		{"g", "g"},
		{"g", "g"},
		{"g", "g"},
		{"g", "g"},
	}
	b := [4][]string{
		{"go", "go"},
		{"go", "go"},
		{"go", "go"},
		{"go", "go"},
	}
	TestA4SString(a, 2)
	assert.Equal(t, b, a)
}
