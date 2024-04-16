package route

import (
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
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
	authRouter := router.Group("/").Use(authMiddleware(AuthClient))

	NewUserRouter(router, authRouter, AuthController)

	return router
}
