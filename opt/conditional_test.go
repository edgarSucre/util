package opt_test

import (
	"testing"

	"github.com/edgarSucre/util/opt"
	"github.com/stretchr/testify/assert"
)

func Test_Conditional_If(t *testing.T) {
	a := opt.If[int](true).Then(10).Otherwise(20)
	assert.Equal(t, 10, a)

	b := opt.If[int](false).Then(10).Otherwise(20)
	assert.Equal(t, 20, b)
}

func Test_Conditional_IfNil(t *testing.T) {
	var nilPointer *int
	valuePointer := new(int)

	a := opt.NotNil(nilPointer).Otherwise(10)
	assert.Equal(t, 10, a)

	b := opt.NotNil(valuePointer).Otherwise(10)
	assert.Equal(t, 0, b)
}
