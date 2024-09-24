package controller

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewPostController),
	fx.Provide(NewFavoriteController),
)
