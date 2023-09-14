package opt

type branch[T comparable] struct {
	flag bool
	val  T
}

func NotNil[T comparable](v *T) branch[T] {
	if v == nil {
		return branch[T]{flag: false}
	}

	return branch[T]{
		flag: true,
		val:  *v,
	}
}

func (br branch[T]) Otherwise(falsy T) T {
	if br.flag {
		return br.val
	}

	return falsy
}

type conditional[T comparable] struct {
	br branch[T]
}

func If[T comparable](flag bool) *conditional[T] {
	return &conditional[T]{
		branch[T]{
			flag: flag,
		},
	}
}

func (cond *conditional[T]) Then(truthy T) branch[T] {
	cond.br.val = truthy

	return cond.br
}
