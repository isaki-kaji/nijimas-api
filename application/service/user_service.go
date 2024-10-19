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
	GetUserDetailByUid(ctx context.Context, uid string, ownUid string) (UserDetailResponse, error)
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

type UserDetailResponse struct {
	Uid             string  `json:"uid"`
	Username        string  `json:"username"`
	SelfIntro       *string `json:"self_intro"`
	ProfileImageUrl *string `json:"profile_image_url"`
	IsFollowing     bool    `json:"is_following"`
	FollowingCount  int     `json:"following_count"`
	FollowersCount  int     `json:"followers_count"`
	PostCount       int     `json:"post_count"`
}

func (s *UserServiceImpl) GetUserDetailByUid(ctx context.Context, uid string, ownUid string) (UserDetailResponse, error) {

	var follow db.GetFollowCountRow
	var postCount int64

	type followResult struct {
		follow db.GetFollowCountRow
		err    error
	}

	type postCountResult struct {
		count int64
		err   error
	}

	followChan := make(chan followResult)
	postCountChan := make(chan postCountResult)

	defer close(followChan)
	defer close(postCountChan)

	user, err := s.repository.GetUser(ctx, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = apperror.DataNotFound.Wrap(err, "user not found")
		} else {
			err = apperror.GetDataFailed.Wrap(err, "failed to get user")
		}
		return UserDetailResponse{}, err
	}

	go func() {
		getFollowCountParam := db.GetFollowCountParams{
			FollowingUid: uid,
			OwnUid:       ownUid,
		}
		follow, err := s.repository.GetFollowCount(ctx, getFollowCountParam)
		if err != nil {
			err = apperror.GetDataFailed.Wrap(err, "failed to get follow count")
			followChan <- followResult{follow: db.GetFollowCountRow{}, err: err}
			return
		}
		followChan <- followResult{follow: follow, err: nil}
	}()

	go func() {
		count, err := s.repository.GetPostsCount(ctx, uid)
		if err != nil {
			err = apperror.GetDataFailed.Wrap(err, "failed to get post count")
			postCountChan <- postCountResult{count: 0, err: err}
			return
		}
		postCountChan <- postCountResult{count: count, err: nil}
	}()

	for i := 0; i < 2; i++ {
		select {
		case fr := <-followChan:
			if fr.err != nil {
				return UserDetailResponse{}, fr.err
			}
			follow = fr.follow
		case pr := <-postCountChan:
			if pr.err != nil {
				return UserDetailResponse{}, pr.err
			}
			postCount = pr.count
		}
	}

	userDetailResponse := UserDetailResponse{
		Uid:             user.Uid,
		Username:        user.Username,
		SelfIntro:       user.SelfIntro,
		ProfileImageUrl: user.ProfileImageUrl,
		IsFollowing:     follow.IsFollowing,
		FollowingCount:  int(follow.FollowingCount),
		FollowersCount:  int(follow.FollowersCount),
		PostCount:       int(postCount),
	}

	return userDetailResponse, nil
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
