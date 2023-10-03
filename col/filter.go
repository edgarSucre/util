package col

type FilterFunc[E any] func(E) bool

// Filter returns a slice of elements when fn(element) is true
func Filter[S ~[]E, E any](data S, fn FilterFunc[E]) S {
	res := make(S, 0, len(data))
	for _, v := range data {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}
