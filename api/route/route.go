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
	AuthController *controller.AuthController,
) *gin.Engine {
	router := gin.Default()

	publicRouter := router.Group("")
	//privateRouter := router.Group("")

	NewSignupRouter(publicRouter, AuthController)

	return router
}
