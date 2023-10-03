/*
Package col provides util methods to manage slices such as: Filter, Map, Pop, Cut etc...
*/
package col

// Cut returns a new slice, removing elements from, to...
// marking removed elements to be garbage collected
func Cut[S ~[]E, E any](s S, from, to int) S {
	var x E
	rightIndex := len(s) - to + from

	copy(s[from:], s[to:])
	for k, n := rightIndex, len(s); k < n; k++ {
		s[k] = x
	}

	s = s[:rightIndex]
	return s
}

// Pop removes the last element from the slice and return that element.
func Pop[S ~[]E, E any](s S) (E, S) {
	i := len(s) - 1

	return s[i], s[:i]
}

// Shift removes the first lement from the slice and return that element
func Shift[S ~[]E, E any](s S) (E, S) {
	return s[0], s[1:]
}

// Unshift appends an element at the start of the slice
func Unshift[S ~[]E, E any](s S, x E) S {
	return append([]E{x}, s...)
}

// Insert returns a new slice with y inserted in x at i
func Insert[S ~[]E, E any](x, y S, i int) S {
	return append(x[:i], append(y, x[i:]...)...)
}
