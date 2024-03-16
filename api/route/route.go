package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	publicRouter := router.Group("")
	NewSignupRouter(publicRouter)

	return router
}
