package foo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooPassInt(t *testing.T) {
	expected := 5
	result := FooPassInt(2, 3)
	assert.Equal(t, expected, result)
}

func TestFooPassString(t *testing.T) {
	expected := "ab"
	result := FooPassString("a", "b")
	assert.Equal(t, expected, result)
}

func TestFooPassBytes(t *testing.T) {
	expected := []byte("abcd")
	result := FooPassBytes([]byte("ab"), 2, []byte("cd"), 2)
	assert.Equal(t, expected, result[:4])
}

func TestFooFindChar(t *testing.T) {
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
		result := FooFindChar(test.Input, test.Search)
		assert.Equal(t, test.Index, result)
	}
}

func TestFooSendMessage(t *testing.T) {
	var from, to [50]byte
	var message [140]byte
	copy(from[:], []byte("Gopher Bob"))
	copy(to[:], []byte("Gopher Anna"))
	copy(message[:], []byte("Hello!"))

	buf := make([]byte, 4096)
	msg := &FooMessage{
		From:      &from,
		To:        &to,
		Message:   &message,
		Signature: " --xxx",
	}
	size := FooSendMessage(msg, buf)
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

func TestFooA4Byte(t *testing.T) {
	a := [4]byte{'a', 'a', 'a', 'a'}
	b := [4]byte{'b', 'b', 'b', 'b'}
	FooA4Byte(&a)
	assert.Equal(t, b, a)
}

func TestFooA4String(t *testing.T) {
	a := [4]string{"g", "g", "g", "g"}
	b := [4]string{"go", "go", "go", "go"}
	FooA4String(&a)
	assert.Equal(t, b, a)
}

func TestFooA4SByte(t *testing.T) {
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
	FooA4SByte(&a, 2)
	assert.Equal(t, b, a)
}

func TestFooA4SString(t *testing.T) {
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
	FooA4SString(&a, 2)
	assert.Equal(t, b, a)
}

func TestFooA2A2Byte(t *testing.T) {
	a := [2][2]byte{
		{'a', 'a'},
		{'a', 'a'},
	}
	b := [2][2]byte{
		{'b', 'b'},
		{'b', 'b'},
	}
	FooA2A2Byte(&a)
	assert.Equal(t, b, a)
}

func TestFooA2A2String(t *testing.T) {
	a := [2][2]string{
		{"g", "g"},
		{"g", "g"},
	}
	b := [2][2]string{
		{"go", "go"},
		{"go", "go"},
	}
	FooA2A2String(&a)
	assert.Equal(t, b, a)
}

func TestFooA2A2SByte(t *testing.T) {
	a := [2][2][]byte{
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
		{{'a', 'a', 'a'}, {'a', 'a', 'a'}},
	}
	b := [2][2][]byte{
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
		{{'b', 'b', 'b'}, {'b', 'b', 'b'}},
	}
	FooA2A2SByte(&a, 3)
	assert.Equal(t, b, a)
}

func TestFooA2A2SString(t *testing.T) {
	a := [2][2][]string{
		{{"g", "g", "g"}, {"g", "g", "g"}},
		{{"g", "g", "g"}, {"g", "g", "g"}},
	}
	b := [2][2][]string{
		{{"go", "go", "go"}, {"go", "go", "go"}},
		{{"go", "go", "go"}, {"go", "go", "go"}},
	}
	FooA2A2SString(&a, 3)
	assert.Equal(t, b, a)
}

func TestFooSSByte(t *testing.T) {
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
	FooSSByte(a, 4, 2)
	assert.Equal(t, b, a)
}

func TestFooSSString(t *testing.T) {
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
	FooSSString(a, 4, 2)
	assert.Equal(t, b, a)
}

func TestFooA4SSByte(t *testing.T) {
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
	FooA4SSByte(&a, 2, 3)
	assert.Equal(t, b, a)
}

func TestFooA4SSString(t *testing.T) {
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
	FooA4SSString(&a, 2, 3)
	assert.Equal(t, b, a)
}

func TestFooA2A2SSByte(t *testing.T) {
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
	FooA2A2SSByte(&a, 3, 2)
	assert.Equal(t, b, a)
}

func TestFooA2A2SSString(t *testing.T) {
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
	FooA2A2SSString(&a, 3, 2)
	assert.Equal(t, b, a)
}
