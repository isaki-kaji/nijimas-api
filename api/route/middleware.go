package route

import "github.com/gin-gonic/gin"

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
