package api

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) SetupRouter() *gin.Engine {
	router := gin.Default()

	publicRouter := router.Group("")
	//privateRouter := router.Group("")

	publicRouter.POST("/signup", server.Signup)

	return router
}
