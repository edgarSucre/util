/*
Package cnd provides conditional operators such as NilCoalesing ??
in the form of functions
*/
package cnd

// NotNil returns x if not nil, otherwise returns y
func NotNil[T comparable](x *T, y T) T {
	if x != nil {
		return *x
	}

	return y
}

// Or is the ternary operator ??, it returns truthy if flag is true
// otherwiese returns falsy
func Or[T comparable](flag bool, truthy, falsy T) T {
	if flag {
		return truthy
	}

	return falsy
}
