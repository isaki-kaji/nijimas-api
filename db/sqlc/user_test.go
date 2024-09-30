package db

// import (
// 	"context"
// 	"testing"

// 	"github.com/isaki-kaji/nijimas-api/util"
// 	"github.com/stretchr/testify/require"
// )

// func createRandomUser(t *testing.T) User {
// 	arg := CreateUserParams{
// 		Uid:         util.RandomUid(),
// 		Username:    util.RandomString(5),
// 		CountryCode: util.RandomCountryCode(),
// 	}

// 	user, err := testRepository.CreateUser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user)

// 	require.Equal(t, arg.Uid, user.Uid)
// 	require.Equal(t, arg.Username, user.Username)
// 	require.Equal(t, arg.CountryCode, user.CountryCode)
// 	require.Nil(t, user.SelfIntro)
// 	require.Nil(t, user.ProfileImageUrl)
// 	require.Nil(t, user.BannerImageUrl)
// 	require.NotZero(t, user.CreatedAt)

// 	return user
// }

// func TestCreateUser(t *testing.T) {
// 	createRandomUser(t)
// }

// func TestGetUser(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	user2, err := testRepository.GetUser(context.Background(), user1.Uid)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.Uid, user2.Uid)
// 	require.Equal(t, user1.Username, user2.Username)
// 	require.Equal(t, user1.CountryCode, user2.CountryCode)
// 	require.Equal(t, user1.SelfIntro, user2.SelfIntro)
// 	require.Equal(t, user1.ProfileImageUrl, user2.ProfileImageUrl)
// 	require.Equal(t, user1.BannerImageUrl, user2.BannerImageUrl)
// 	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
// }
