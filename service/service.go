package service

import (
	"github.com/gin-gonic/gin"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"go.uber.org/fx"
)

type Service struct {
	repository db.Repository
}

func NewService(repository db.Repository) *Service {
	return &Service{repository: repository}
}

var Module = fx.Options(
	fx.Provide(NewService),
)

func (s *Service) CreateUser(ctx *gin.Context, arg db.CreateUserParams) (db.User, error) {
	return s.repository.CreateUser(ctx, arg)
}
