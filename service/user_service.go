package service

import (
	"context"
	"errors"

	"github.com/isaki-kaji/nijimas-api/apperror"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jackc/pgx/v5"
)

type UserService interface {
	CreateUser(ctx context.Context, arg CreateUserRequest) (db.User, error)
	GetUser(ctx context.Context, uid string) (db.User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (db.User, error)
}

func NewUserService(repository db.Repository) UserService {
	return &UserServiceImpl{repository: repository}
}

type UserServiceImpl struct {
	repository db.Repository
}

type CreateUserRequest struct {
	Uid         string `json:"-"`
	Username    string `json:"username" binding:"required,max=14,min=2"`
	CountryCode string `json:"country_code" binding:"omitempty,len=2"`
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, arg CreateUserRequest) (db.User, error) {
	_, err := s.repository.GetUser(ctx, arg.Uid)
	if err == nil {
		err = apperror.DataConflict.Wrap(ErrUserAlreadyExists, "user already exists")
		return db.User{}, err
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		err = apperror.GetDataFailed.Wrap(err, "failed to get user")
		return db.User{}, err
	}
	param := db.CreateUserParams{
		Uid:         arg.Uid,
		Username:    arg.Username,
		CountryCode: util.ToPointerOrNil(arg.CountryCode),
	}
	newUser, err := s.repository.CreateUser(ctx, param)
	if err != nil {
		err = apperror.InsertDataFailed.Wrap(err, "failed to create user")
		return db.User{}, err
	}
	return newUser, nil
}

func (s *UserServiceImpl) GetUser(ctx context.Context, uid string) (db.User, error) {
	user, err := s.repository.GetUser(ctx, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = apperror.DataNotFound.Wrap(err, "user not found")
			return db.User{}, err
		}
		err = apperror.GetDataFailed.Wrap(err, "failed to get user")
		return db.User{}, err
	}
	return user, nil
}

type UpdateUserParams struct {
	Uid             string `json:"-"`
	Username        string `json:"username" binding:"required,max=14,min=2"`
	SelfIntro       string `json:"self_intro" binding:"max=200"`
	ProfileImageUrl string `json:"profile_image_url" binding:"max=2000"`
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, arg UpdateUserParams) (db.User, error) {
	_, err := s.repository.GetUser(ctx, arg.Uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := apperror.DataNotFound.Wrap(err, "user not found")
			return db.User{}, err
		}
		err = apperror.GetDataFailed.Wrap(err, "failed to get user")
		return db.User{}, err
	}
	param := db.UpdateUserParams{
		Uid:             arg.Uid,
		Username:        util.ToPointerOrNil(arg.Username),
		SelfIntro:       util.ToPointerOrNil(arg.SelfIntro),
		ProfileImageUrl: util.ToPointerOrNil(arg.ProfileImageUrl),
	}
	updatedUser, err := s.repository.UpdateUser(ctx, param)
	if err != nil {
		err = apperror.UpdateDataFailed.Wrap(err, "failed to update user")
		return db.User{}, err
	}
	return updatedUser, nil
}
