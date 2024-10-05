package api

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/configs"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewServer),
)

type Server struct {
	config configs.Config
	router *gin.Engine
}

func NewServer(config *configs.Config, router *gin.Engine) (*Server, error) {
	server := &Server{
		config: *config,
		router: router,
	}
	return server, nil
}

func (server *Server) Start() error {
	return server.router.Run(server.config.ServerAddress)
}
