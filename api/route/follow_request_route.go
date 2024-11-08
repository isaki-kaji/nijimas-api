package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
)

func NewFollowRequestRouter(router *gin.Engine, authRouter gin.IRoutes, controller *controller.FollowRequestController) {
	authRouter.GET("/me/follow-requests", controller.GetFollowRequests)
	authRouter.POST("/follow-requests", controller.DoFollowRequest)
	authRouter.DELETE("/follow-requests", controller.CancelFollowRequest)
	authRouter.PUT("/follow-requests", controller.HandleFollowRequest)
}
