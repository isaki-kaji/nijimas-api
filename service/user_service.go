package service

import (
	"context"
	"errors"
	"fmt"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/domain"
	"github.com/isaki-kaji/nijimas-api/util"
)

type UserService struct {
	repository db.Repository
}

func NewUserService(repository db.Repository) domain.UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(ctx context.Context, arg domain.CreateUserRequest) (db.User, error) {
	_, err := s.repository.GetUser(ctx, arg.Uid)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			param := db.CreateUserParams{
				Uid:         arg.Uid,
				Username:    arg.Username,
				CountryCode: &arg.CountryCode,
			}
			newUser, err := s.repository.CreateUser(ctx, param)
			if err != nil {
				return db.User{}, err
			}
			return newUser, nil
		}
		return db.User{}, err
	}
	return db.User{}, fmt.Errorf(util.UserAlreadyExists)
}

func (s *UserService) GetUser(ctx context.Context, uid string) (db.User, error) {
	user, err := s.repository.GetUser(ctx, uid)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
