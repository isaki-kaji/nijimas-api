package service

import (
	"context"
	"errors"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jackc/pgx/v5"
)

type UserService interface {
	CreateUser(ctx context.Context, arg CreateUserRequest) (db.User, error)
	GetUser(ctx context.Context, uid string) (db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
}

func NewUserService(repository db.Repository) UserService {
	return &UserServiceImpl{repository: repository}
}

type UserServiceImpl struct {
	repository db.Repository
}

type CreateUserRequest struct {
	Uid         string `json:"uid" binding:"required"`
	Username    string `json:"username" binding:"required,max=14,min=2"`
	CountryCode string `json:"country_code" binding:"omitempty,len=2"`
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, arg CreateUserRequest) (db.User, error) {
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
		CountryCode: util.ToPointerOrNil(arg.CountryCode),
	}
	newUser, err := s.repository.CreateUser(ctx, param)
	if err != nil {
		return db.User{}, err
	}
	return newUser, nil
}

func (s *UserServiceImpl) GetUser(ctx context.Context, uid string) (db.User, error) {
	user, err := s.repository.GetUser(ctx, uid)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	_, err := s.repository.GetUser(ctx, arg.Uid)
	if err != nil {
		return db.User{}, err
	}
	param := db.UpdateUserParams{
		Uid:             arg.Uid,
		Username:        arg.Username,
		SelfIntro:       arg.SelfIntro,
		ProfileImageUrl: arg.ProfileImageUrl,
		BannerImageUrl:  arg.BannerImageUrl,
	}
	updatedUser, err := s.repository.UpdateUser(ctx, param)
	if err != nil {
		return db.User{}, err
	}
	return updatedUser, nil
}
