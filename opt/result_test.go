package opt_test

import (
	"fmt"
	"testing"

	"github.com/edgarSucre/util/opt"
	"github.com/stretchr/testify/assert"
)

func currying[T comparable](fn func() (T, error)) (T, error) {
	return fn()
}

func Test_Result(t *testing.T) {
	r1 := opt.Try(privateFunc1())
	r2 := opt.Map(r1, privateFunc2())
	r3 := opt.Map(r2, privateFunc3())

	v, err := r3.Get()

	assert.ErrorContains(t, err, "context added to parse3")
	assert.Equal(t, int64(0), v)
}

func privateFunc1() opt.TryFunc[int8] {
	result := opt.Touple(parse1(false))
	result.Wrap("context added to parse1")

	return func() (int8, error) {
		return result.Get()
	}
}

func privateFunc2() opt.MapFunc[int8, int16] {
	return func(v int8) *opt.Result[int16] {
		result := opt.Touple(parse2(v, false))
		result.Wrap("context aded to parse 2")

		return result
	}
}

func privateFunc3() opt.MapFunc[int16, int64] {
	return func(i int16) *opt.Result[int64] {
		result := opt.Touple(parse3(i, true))
		result.Wrap("context added to parse3")

		return result
	}
}

func parse1(fail bool) (int8, error) {
	if fail {
		return 0, fmt.Errorf("parse1 failed")
	}

	return 10, nil
}

func parse2(v int8, fail bool) (int16, error) {
	if fail {
		return 0, fmt.Errorf("parse2 failed")
	}

	return int16(500), nil
}

func parse3(v int16, fail bool) (int64, error) {
	if fail {
		return 0, fmt.Errorf("parse3 failed")
	}

	return int64(50000), nil
}
