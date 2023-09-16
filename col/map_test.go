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
}
