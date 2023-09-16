package col

type MapFunc[T, K comparable] func(T) K

func Map[T, K comparable](data []T, fn MapFunc[T, K]) []K {
	res := make([]K, len(data))
	for i, v := range data {
		res[i] = fn(v)
	}

	return res
}
