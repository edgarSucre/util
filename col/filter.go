package col

type FilterFunc[T comparable] func(T) bool

func Filter[T comparable](data []T, fn FilterFunc[T]) []T {
	res := make([]T, 0, len(data))
	for _, v := range data {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}
