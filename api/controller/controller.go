package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewPostController),
	fx.Provide(NewFavoriteController),
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
