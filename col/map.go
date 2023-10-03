package col

type MapFunc[E, K any] func(E) K

// Map returns a slice of K, by calling MapFunc() on every element
func Map[S ~[]E, E, K any](data S, fn MapFunc[E, K]) []K {
	res := make([]K, len(data))
	for i, v := range data {
		res[i] = fn(v)
	}

	return res
}
