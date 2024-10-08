package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/api/controller"
)

func NewFavoriteRouter(router *gin.Engine, authRouter gin.IRoutes, controller *controller.FavoriteController) {
	authRouter.POST("/favorites", controller.ToggleFavorite)
}
