package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRouter),
)

func NewRouter(
	AuthController *controller.UserController,
) *gin.Engine {
	router := gin.Default()
	authRouter := router.Group("/").Use(authMiddleware())

	NewUserRouter(router, authRouter, AuthController)

	return router
}
