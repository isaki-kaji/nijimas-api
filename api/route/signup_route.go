package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
)

func NewSignupRouter(router *gin.RouterGroup, controller *controller.AuthController) {
	router.POST("/signup", controller.Signup)
}
