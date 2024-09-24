package api

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/util"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewServer),
)

type Server struct {
	config util.Config
	router *gin.Engine
}

func NewServer(config *util.Config, router *gin.Engine) (*Server, error) {
	server := &Server{
		config: *config,
		router: router,
	}
	return server, nil
}

func (server *Server) Start() error {
	return server.router.Run(server.config.ServerAddress)
}
