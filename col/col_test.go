package col_test

import (
	"testing"

	"github.com/edgarSucre/util/col"
	"github.com/stretchr/testify/assert"
)

func Test_Cut(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a = col.Cut(a, 5, 9)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 10}, a)

	b := []string{"a", "b", "c", "d", "e"}
	b = col.Cut(b, 1, 4)
	assert.Equal(t, []string{"a", "e"}, b)

	p1 := new(int)
	p2 := new(int)
	p3 := new(int)
	p4 := new(int)

	c := []*int{p1, p2, p3, p4}
	c = col.Cut(c, 1, 3)

	assert.Equal(t, []*int{p1, p4}, c)
}

func Test_Pop(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x, a := col.Pop(a)

	assert.Equal(t, 10, x)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, a)
}

func Test_Shift(t *testing.T) {
	a := []int{1, 2, 3}
	x, a := col.Shift(a)

	assert.Equal(t, 1, x)
	assert.Equal(t, []int{2, 3}, a)
}

func Test_UnShift(t *testing.T) {
	a := []int{1, 2, 3}
	a = col.Unshift(a, 0)

	assert.Equal(t, []int{0, 1, 2, 3}, a)
}

func Test_Insert(t *testing.T) {
	a := []int{1, 2, 3, 6, 7}
	b := []int{4, 5}

	a = col.Insert(a, b, 3)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, a)
}
