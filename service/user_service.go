package service

import (
	"context"
	"errors"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/domain"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	repository db.Repository
}

func NewUserService(repository db.Repository) domain.UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(ctx context.Context, arg domain.CreateUserRequest) (db.User, error) {
	_, err := s.repository.GetUser(ctx, arg.Uid)
	if err == nil {
		return db.User{}, errors.New(util.UserAlreadyExists)
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return db.User{}, err
	}
	param := db.CreateUserParams{
		Uid:         arg.Uid,
		Username:    arg.Username,
		CountryCode: util.PointerOrNil(arg.CountryCode),
	}
	newUser, err := s.repository.CreateUser(ctx, param)
	if err != nil {
		return db.User{}, err
	}
	return newUser, nil
}

func (s *UserService) GetUser(ctx context.Context, uid string) (db.User, error) {
	user, err := s.repository.GetUser(ctx, uid)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
