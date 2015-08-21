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

func TestTestProxyPtr3String(t *testing.T) {
	cube := [4][][][]string{
		[][][]string{
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
		},
		[][][]string{
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
		},
	}

	result := TestProxyPtr3String(cube)
	assert.Equal(t, [][][]string{}, result)
}
