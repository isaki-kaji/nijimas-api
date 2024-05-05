package db

import (
	"context"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) Post {
	user := createRandomUser(t)
	arg := CreatePostParams{
		PostID:       uuid.New(),
		Uid:          user.Uid,
		MainCategory: util.RandomMainCategory(),
		PostText:     util.ToPointerOrNil(util.RandomString(100)),
		PhotoUrl:     util.ToPointerOrNil(util.RandomString(100)),
		Expense:      util.ToPointerOrNil(rand.Int63n(10000)),
		Location:     util.ToPointerOrNil(util.RandomString(20)),
		PublicTypeNo: util.RandomPublicTypeNo(),
	}

	post, err := testRepository.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.PostID, post.PostID)
	require.Equal(t, arg.Uid, post.Uid)
	require.Equal(t, arg.MainCategory, post.MainCategory)
	require.Equal(t, arg.PostText, post.PostText)
	require.Equal(t, arg.PhotoUrl, post.PhotoUrl)
	require.Equal(t, arg.Expense, post.Expense)
	require.Equal(t, arg.Location, post.Location)
	require.Equal(t, arg.PublicTypeNo, post.PublicTypeNo)
	require.NotZero(t, post.CreatedAt)

	return post
}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}
