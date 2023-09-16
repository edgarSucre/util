package col

import "github.com/edgarSucre/util/opt"

type FilterFunc[T comparable] func(T) bool

func Filter[T comparable](data []T, fn FilterFunc[T]) []T {
	res := make([]T, 0, len(data))
	for _, v := range data {
		res = opt.If[T](fn(v)).Append(res, v)
	}

	return res
}
