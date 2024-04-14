package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
)

func NewUserRouter(router *gin.Engine, authRouter gin.IRoutes, controller *controller.UserController) {
	router.POST("/users", controller.Create)
	authRouter.GET("/users/:id", controller.Get)
}
