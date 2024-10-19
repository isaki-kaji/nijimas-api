package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
)

func NewFollowRouter(router *gin.Engine, authRouter gin.IRoutes, controller *controller.FollowController) {
	authRouter.POST("/follows", controller.ToggleFollow)
}
