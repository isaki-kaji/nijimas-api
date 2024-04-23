package util

func PointerOrNil[T comparable](value T) *T {
	var zero T
	if value == zero {
		return nil
	}
	return &value
}
