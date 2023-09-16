package opt

import "fmt"

type Result[T comparable] struct {
	val     T
	err     error
	wrapped bool
}

func (r *Result[T]) IsError() bool {
	return r.err != nil
}

func (r *Result[T]) Get() (T, error) {
	return r.val, r.err
}

func (r *Result[T]) Wrap(msg string) *Result[T] {
	if !r.wrapped && r.IsError() {
		r.err = fmt.Errorf("%w - %s", r.err, msg)
		r.wrapped = true
	}

	return r
}

type TryFunc[T comparable] func() (T, error)

func Try[T comparable](fn TryFunc[T]) *Result[T] {
	v, err := fn()
	return &Result[T]{val: v, err: err, wrapped: false}
}

func Touple[T comparable](v T, e error) *Result[T] {
	return &Result[T]{val: v, err: e, wrapped: false}
}

type MapFunc[T, K comparable] func(T) *Result[K]

func Map[T, K comparable](r *Result[T], fn MapFunc[T, K]) *Result[K] {
	if r.IsError() {
		return &Result[K]{err: r.err}
	}

	return fn(r.val)
}
