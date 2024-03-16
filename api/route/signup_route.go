package route

import "github.com/gin-gonic/gin"

func NewSignupRouter(router *gin.RouterGroup) {
	router.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(200, "signup")
	})
}
