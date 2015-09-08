package translator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCGoSpecAtLevel(t *testing.T) {
	assert := assert.New(t)
	tbl := []struct {
		Spec     CGoSpec
		Expected []string
	}{
		{CGoSpec{},
			[]string{"", "", ""}},
		{CGoSpec{Base: "char"},
			[]string{"char", "char", "char"}},
		{CGoSpec{Base: "char", Pointers: 1},
			[]string{"*char", "char", "char"}},
		{CGoSpec{Base: "char", Pointers: 2},
			[]string{"**char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: []uint64{1}},
			[]string{"*char", "char", "char"}},
		{CGoSpec{Base: "char", Arrays: []uint64{1, 2}},
			[]string{"*[2]char", "[2]char", "char"}},
		{CGoSpec{Base: "char", Arrays: []uint64{1, 2, 3}},
			[]string{"*[2][3]char", "[2][3]char", "[3]char", "char"}},
		{CGoSpec{Base: "char", Arrays: []uint64{1, 2, 3}, Pointers: 1},
			[]string{"*[2][3]*char", "[2][3]*char", "[3]*char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: []uint64{1, 2, 3}, Pointers: 2},
			[]string{"*[2][3]**char", "[2][3]**char", "[3]**char", "**char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: []uint64{1, 2}, Pointers: 3},
			[]string{"*[2]***char", "[2]***char", "***char", "**char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: []uint64{1, 2}, Pointers: 1},
			[]string{"*[2]*char", "[2]*char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: []uint64{1}, Pointers: 2},
			[]string{"***char", "**char", "*char", "char"}},
		{CGoSpec{Base: "char", Arrays: []uint64{1}, Pointers: 1},
			[]string{"**char", "*char", "char"}},
	}
	for _, test := range tbl {
		for level, exp := range test.Expected {
			assert.Equal(exp, test.Spec.AtLevel(uint8(level)), fmt.Sprintf("at level %d", level))
		}
	}
}
