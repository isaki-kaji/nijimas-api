package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
)

func NewPostRouter(router *gin.Engine, authRouter gin.IRoutes, controller *controller.PostController) {
	authRouter.POST("/posts", controller.Create)
}
