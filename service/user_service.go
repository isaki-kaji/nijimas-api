package service

import (
	"context"
	"database/sql"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/domain"
)

type UserService struct {
	repository db.Repository
}

func NewUserService(repository db.Repository) domain.UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(ctx context.Context, arg domain.CreateUserRequest) (db.User, error) {
	param := db.CreateUserParams{
		Uid:      arg.Uid,
		Username: arg.Username,
		CountryCode: sql.NullString{
			String: arg.CountryCode,
			Valid:  arg.CountryCode != "",
		},
	}
	return s.repository.CreateUser(ctx, param)
}
