package col_test

import (
	"testing"

	"github.com/edgarSucre/util/col"
	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}

	even := col.Filter(data, func(v int) bool { return v%2 == 0 })
	assert.Equal(t, []int{2, 4, 6, 8}, even)

	odd := col.Filter(data, func(v int) bool { return v%2 != 0 })
	assert.Equal(t, []int{1, 3, 5, 7}, odd)
}
