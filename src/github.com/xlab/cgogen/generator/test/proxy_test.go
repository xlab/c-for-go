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
