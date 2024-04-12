package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
)

func NewUserRouter(router *gin.RouterGroup, controller *controller.UserController) {
	router.POST("/users", controller.Create)
}
