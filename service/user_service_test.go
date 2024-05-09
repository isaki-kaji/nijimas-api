package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/isaki-kaji/nijimas-api/db/mock"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/domain"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	testUser := randomNewUser()

	testCases := []struct {
		name       string
		body       domain.CreateUserRequest
		buildStubs func(repository *mockdb.MockRepository)
		check      func(t *testing.T, user db.User, err error)
	}{
		{
			name: "OK",
			body: domain.CreateUserRequest{
				Uid:         testUser.Uid,
				Username:    testUser.Username,
				CountryCode: util.StringPointerToString(testUser.CountryCode),
			},
			buildStubs: func(repository *mockdb.MockRepository) {
				repository.EXPECT().
					GetUser(gomock.Any(), testUser.Uid).
					Times(1).
					Return(db.User{}, pgx.ErrNoRows)
				repository.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(testUser, nil)
			},
			check: func(t *testing.T, user db.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.Uid, testUser.Uid)
				require.Equal(t, user.Username, testUser.Username)
				require.Equal(t, user.CountryCode, testUser.CountryCode)
				require.Nil(t, user.SelfIntro)
				require.Nil(t, user.ProfileImageUrl)
				require.Nil(t, user.BannerImageUrl)
			},
		},
		{
			name: "User already exists",
			body: domain.CreateUserRequest{
				Uid:         testUser.Uid,
				Username:    testUser.Username,
				CountryCode: util.StringPointerToString(testUser.CountryCode),
			},
			buildStubs: func(repository *mockdb.MockRepository) {
				repository.EXPECT().
					GetUser(gomock.Any(), testUser.Uid).
					Times(1).
					Return(testUser, nil)
				repository.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			check: func(t *testing.T, user db.User, err error) {
				require.Error(t, err)
				require.EqualError(t, err, util.UserAlreadyExists)
			},
		}, {
			name: "Get user error",
			body: domain.CreateUserRequest{
				Uid:         testUser.Uid,
				Username:    testUser.Username,
				CountryCode: util.StringPointerToString(testUser.CountryCode),
			},
			buildStubs: func(repository *mockdb.MockRepository) {
				repository.EXPECT().
					GetUser(gomock.Any(), testUser.Uid).
					Times(1).
					Return(db.User{}, errors.New("unexpected error"))
				repository.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			check: func(t *testing.T, user db.User, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "Create user error",
			body: domain.CreateUserRequest{
				Uid:         testUser.Uid,
				Username:    testUser.Username,
				CountryCode: util.StringPointerToString(testUser.CountryCode),
			},
			buildStubs: func(repository *mockdb.MockRepository) {
				repository.EXPECT().
					GetUser(gomock.Any(), testUser.Uid).
					Times(1).
					Return(db.User{}, pgx.ErrNoRows)
				repository.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, errors.New("unexpected error"))
			},
			check: func(t *testing.T, user db.User, err error) {
				require.Error(t, err)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mockdb.NewMockRepository(ctrl)
			tc.buildStubs(repository)

			service := NewUserService(repository)
			user, err := service.CreateUser(context.Background(), tc.body)
			tc.check(t, user, err)
		})
	}
}

func TestGetUser(t *testing.T) {
	testUser := randomNewUser()

	testCases := []struct {
		name       string
		uid        string
		buildStubs func(repository *mockdb.MockRepository)
		check      func(t *testing.T, user db.User, err error)
	}{
		{
			name: "OK",
			uid:  testUser.Uid,
			buildStubs: func(repository *mockdb.MockRepository) {
				repository.EXPECT().
					GetUser(gomock.Any(), testUser.Uid).
					Times(1).
					Return(testUser, nil)
			},
			check: func(t *testing.T, user db.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.Uid, testUser.Uid)
				require.Equal(t, user.Username, testUser.Username)
				require.Equal(t, user.CountryCode, testUser.CountryCode)
			},
		},
		{
			name: "User not found",
			uid:  testUser.Uid,
			buildStubs: func(repository *mockdb.MockRepository) {
				repository.EXPECT().
					GetUser(gomock.Any(), testUser.Uid).
					Times(1).
					Return(db.User{}, pgx.ErrNoRows)
			},
			check: func(t *testing.T, user db.User, err error) {
				require.Error(t, err)
				require.EqualError(t, err, pgx.ErrNoRows.Error())
			},
		},
		{
			name: "Other error",
			uid:  testUser.Uid,
			buildStubs: func(repository *mockdb.MockRepository) {
				repository.EXPECT().
					GetUser(gomock.Any(), testUser.Uid).
					Times(1).
					Return(db.User{}, errors.New("unexpected error"))
			},
			check: func(t *testing.T, user db.User, err error) {
				require.Error(t, err)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mockdb.NewMockRepository(ctrl)
			tc.buildStubs(repository)

			service := NewUserService(repository)
			user, err := service.GetUser(context.Background(), tc.uid)
			tc.check(t, user, err)
		})
	}
}

func randomNewUser() (user db.User) {
	user = db.User{
		Uid:         util.RandomUid(),
		Username:    util.RandomString(5),
		CountryCode: util.RandomCountryCode(),
	}
	return
}
