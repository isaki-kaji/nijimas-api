package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
)

func NewSubCategoryRouter(router *gin.Engine, authRouter gin.IRoutes, controller *controller.SubCategoryController) {
	authRouter.GET("/me/sub-categories", controller.GetUserUsedSubCategories)
}
