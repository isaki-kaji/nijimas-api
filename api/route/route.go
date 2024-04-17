package route

import (
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
	"github.com/isaki-kaji/nijimas-api/api/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRouter),
)

func NewRouter(
	AuthController *controller.UserController,
	AuthClient *auth.Client,
) *gin.Engine {
	router := gin.Default()
	authRouter := router.Group("/").Use(middleware.AuthMiddleware(AuthClient))

	NewUserRouter(router, authRouter, AuthController)

	return router
}
