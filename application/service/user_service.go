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
	FollowingStatus string  `json:"following_status"`
	FollowingCount  int     `json:"following_count"`
	FollowersCount  int     `json:"followers_count"`
	PostCount       int     `json:"post_count"`
}

func (s *UserServiceImpl) GetUserDetailByUid(ctx context.Context, uid string, ownUid string) (UserDetailResponse, error) {

	var followingCount int
	var followersCount int
	var followingStatus string
	var postCount int64

	type followResult struct {
		followingCount  int64
		followersCount  int64
		followingStatus string
		err             error
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
		getFollowInfoParam := db.GetFollowInfoParams{
			FollowingUid: uid,
			OwnUid:       ownUid,
		}

		getFollowRequestParams := db.GetFollowRequestParams{
			Uid:          ownUid,
			FollowingUid: uid,
		}

		follow, err := s.repository.GetFollowInfo(ctx, getFollowInfoParam)
		if err != nil {
			err = apperror.GetDataFailed.Wrap(err, "failed to get follow count")
			followChan <- followResult{followingCount: 0, followersCount: 0, followingStatus: "", err: err}
			return
		}

		var followingStatus string
		if follow.IsFollowing {
			followingStatus = StatusFollowing
		} else {
			_, err := s.repository.GetFollowRequest(ctx, getFollowRequestParams)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					followingStatus = StatusNotFollowing
				} else {
					err = apperror.GetDataFailed.Wrap(err, "failed to get follow request")
					followChan <- followResult{followingCount: 0, followersCount: 0, followingStatus: "", err: err}
					return
				}
			} else {
				followingStatus = StatusRequested
			}
		}

		followChan <- followResult{followingCount: follow.FollowingCount, followersCount: follow.FollowersCount, followingStatus: followingStatus, err: nil}
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
			followingCount = int(fr.followingCount)
			followersCount = int(fr.followersCount)
			followingStatus = fr.followingStatus
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
		FollowingCount:  followingCount,
		FollowersCount:  followersCount,
		FollowingStatus: followingStatus,
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
