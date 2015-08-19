package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestProxyString(t *testing.T) {
	expected := "HELLO"
	result := TestProxyString("a", "b", "c")
	assert.Equal(t, expected, result)
}

func TestTestProxyBytes(t *testing.T) {
	expected := []byte("abcd")
	result := TestProxyBytes([]byte("ab"), []byte("cd"))
	assert.Equal(t, expected, result)
}
