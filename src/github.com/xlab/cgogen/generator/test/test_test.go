package test

import (
	"bytes"
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

func TestTestSendMessage(t *testing.T) {
	var from, to [50]byte
	var message [140]byte
	copy(from[:], []byte("Gopher Bob"))
	copy(to[:], []byte("Gopher Anna"))
	copy(message[:], []byte("Hello!"))

	buf := make([]byte, 4096)
	msg := &TestMessage{
		From:      &from,
		To:        &to,
		Message:   &message,
		Signature: " --xxx",
	}
	size := TestSendMessage(msg, buf)
	assert.EqualValues(t, 249, size)
	packed := []byte(`msgGopher BobGopher AnnaHello! --xxx`)
	assert.Equal(t, packed, cleanBuf(buf))
}

func cleanBuf(buf []byte) []byte {
	tmp := new(bytes.Buffer)
	for i := range buf {
		if buf[i] != 0 {
			tmp.WriteByte(buf[i])
			continue
		}
	}
	return tmp.Bytes()
}

func TestTestA4Byte(t *testing.T) {
	a := [4]byte{'a', 'a', 'a', 'a'}
	b := [4]byte{'b', 'b', 'b', 'b'}
	TestA4Byte(&a)
	assert.Equal(t, b, a)
}

func TestTestA4String(t *testing.T) {
	a := [4]string{"g", "g", "g", "g"}
	b := [4]string{"go", "go", "go", "go"}
	TestA4String(&a)
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
	TestA4SByte(&a, 2)
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
	TestA4SString(&a, 2)
	assert.Equal(t, b, a)
}

func TestTestA2A2Byte(t *testing.T) {
	a := [2][2]byte{
		{'a', 'a'},
		{'a', 'a'},
	}
	b := [2][2]byte{
		{'b', 'b'},
		{'b', 'b'},
	}
	TestA2A2Byte(&a)
	assert.Equal(t, b, a)
}

func TestTestA2A2String(t *testing.T) {
	a := [2][2]string{
		{"g", "g"},
		{"g", "g"},
	}
	b := [2][2]string{
		{"go", "go"},
		{"go", "go"},
	}
	TestA2A2String(&a)
	assert.Equal(t, b, a)
}

func TestTestA2A2SByte(t *testing.T) {
	a := [2][2][]byte{
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
	}
	b := [2][2][]byte{
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
	}
	TestA2A2SByte(&a, 3)
	assert.Equal(t, b, a)
}

func TestTestA2A2SString(t *testing.T) {
	a := [2][2][]string{
		{{"g", "g", "g"}, {"g", "g", "g"}},
		{{"g", "g", "g"}, {"g", "g", "g"}},
	}
	b := [2][2][]string{
		{{"go", "go", "go"}, {"go", "go", "go"}},
		{{"go", "go", "go"}, {"go", "go", "go"}},
	}
	TestA2A2SString(&a, 3)
	assert.Equal(t, b, a)
}

func TestTestSSByte(t *testing.T) {
	a := [][]byte{
		{'a', 'a'},
		{'a', 'a'},
		{'a', 'a'},
		{'a', 'a'},
	}
	b := [][]byte{
		{'b', 'b'},
		{'b', 'b'},
		{'b', 'b'},
		{'b', 'b'},
	}
	TestSSByte(a, 4, 2)
	assert.Equal(t, b, a)
}

func TestTestSSString(t *testing.T) {
	a := [][]string{
		{"g", "g"},
		{"g", "g"},
		{"g", "g"},
		{"g", "g"},
	}
	b := [][]string{
		{"go", "go"},
		{"go", "go"},
		{"go", "go"},
		{"go", "go"},
	}
	TestSSString(a, 4, 2)
	assert.Equal(t, b, a)
}

func TestTestA4SSByte(t *testing.T) {
	a := [4][][]byte{
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
	}
	b := [4][][]byte{
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
	}
	TestA4SSByte(&a, 2, 3)
	assert.Equal(t, b, a)
}

func TestTestA4SSString(t *testing.T) {
	a := [4][][]string{
		{{"g", "g", "g"}, {"g", "g", "g"}},
		{{"g", "g", "g"}, {"g", "g", "g"}},
		{{"g", "g", "g"}, {"g", "g", "g"}},
		{{"g", "g", "g"}, {"g", "g", "g"}},
	}
	b := [4][][]string{
		{{"go", "go", "go"}, {"go", "go", "go"}},
		{{"go", "go", "go"}, {"go", "go", "go"}},
		{{"go", "go", "go"}, {"go", "go", "go"}},
		{{"go", "go", "go"}, {"go", "go", "go"}},
	}
	TestA4SSString(&a, 2, 3)
	assert.Equal(t, b, a)
}

func TestTestA2A2SSByte(t *testing.T) {
	a := [2][2][][]byte{
		{
			{{'a', 'a'}, {'a', 'a'}, {'a', 'a'}},
			{{'a', 'a'}, {'a', 'a'}, {'a', 'a'}},
		},
		{
			{{'a', 'a'}, {'a', 'a'}, {'a', 'a'}},
			{{'a', 'a'}, {'a', 'a'}, {'a', 'a'}},
		},
	}
	b := [2][2][][]byte{
		{
			{{'b', 'b'}, {'b', 'b'}, {'b', 'b'}},
			{{'b', 'b'}, {'b', 'b'}, {'b', 'b'}},
		},
		{
			{{'b', 'b'}, {'b', 'b'}, {'b', 'b'}},
			{{'b', 'b'}, {'b', 'b'}, {'b', 'b'}},
		},
	}
	TestA2A2SSByte(&a, 3, 2)
	assert.Equal(t, b, a)
}

func TestTestA2A2SSString(t *testing.T) {
	a := [2][2][][]string{
		{
			{{"g", "g"}, {"g", "g"}, {"g", "g"}},
			{{"g", "g"}, {"g", "g"}, {"g", "g"}},
		},
		{
			{{"g", "g"}, {"g", "g"}, {"g", "g"}},
			{{"g", "g"}, {"g", "g"}, {"g", "g"}},
		},
	}
	b := [2][2][][]string{
		{
			{{"go", "go"}, {"go", "go"}, {"go", "go"}},
			{{"go", "go"}, {"go", "go"}, {"go", "go"}},
		},
		{
			{{"go", "go"}, {"go", "go"}, {"go", "go"}},
			{{"go", "go"}, {"go", "go"}, {"go", "go"}},
		},
	}
	TestA2A2SSString(&a, 3, 2)
	assert.Equal(t, b, a)
}
