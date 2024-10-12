package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
)

func NewPostSearchRouter(router *gin.Engine, authRouter gin.IRoutes, controller *controller.PostSearchController) {
	authRouter.GET("/posts", controller.GetPostsByQuery)
}
