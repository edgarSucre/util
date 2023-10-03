package col_test

import (
	"fmt"
	"testing"

	"github.com/edgarSucre/util/col"
	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	data := []int{1, 2, 3}

	str := col.Map(data, func(v int) string { return fmt.Sprint(v) })
	assert.Equal(t, []string{"1", "2", "3"}, str)

	p1, p2, p3 := new(int), new(int), new(int)

	fn := func(p *int) int {
		return *p
	}

	pointers := []*int{p1, p2, p3}
	num := col.Map(pointers, fn)
	assert.Equal(t, []int{0, 0, 0}, num)
}
