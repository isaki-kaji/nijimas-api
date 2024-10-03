package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
)

func NewSummaryRouter(router *gin.Engine, authRouter gin.IRoutes, controller *controller.SummaryController) {
	authRouter.GET("/me/summaries/:year/:month", controller.GetMonthlySummary)
}
