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
	result := TestPassBytes([]byte("ab"), []byte("cd"))
	assert.Equal(t, expected, result)
}
