package api

import (
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/service"
	"github.com/isaki-kaji/nijimas-api/util"
	"go.uber.org/fx"
)

type Server struct {
	config  util.Config
	service service.Service
	router  *gin.Engine
}

func NewServer(config *util.Config, service *service.Service) (*Server, error) {

	server := &Server{
		config:  *config,
		service: *service,
	}

	server.router = server.SetupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

var Module = fx.Options(
	fx.Provide(NewServer),
)
