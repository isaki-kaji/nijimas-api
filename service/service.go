package service

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewPostService),
)

func PointerOrNil[T comparable](value T) *T {
	var zero T
	if value == zero {
		return nil
	}
	return &value
}
