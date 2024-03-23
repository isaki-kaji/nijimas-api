package service

import (
	"context"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/domain"
)

type AuthService struct {
	repository db.Repository
}

func NewAuthService(repository db.Repository) domain.AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) SignupUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return s.repository.CreateUser(ctx, arg)
}
