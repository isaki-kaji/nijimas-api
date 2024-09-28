package db

import (
	"context"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

// func TestGetPostsByUid(t *testing.T) {
// 	postsUid, numPosts := createRandomPostsByUid(t)
// 	t.Logf("postsUid: %s, numPosts: %d\n", postsUid, numPosts)
// 	resultPosts, err := testRepository.GetPostsByUid(context.Background(), postsUid)
// 	require.NoError(t, err)
// 	require.Equal(t, len(resultPosts), numPosts)

// 	for i := 1; i < len(resultPosts); i++ {
// 		require.True(t, resultPosts[i-1].CreatedAt.Before(resultPosts[i].CreatedAt) || resultPosts[i-1].CreatedAt.Equal(resultPosts[i].CreatedAt))
// 	}
// }

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

func createRandomFullPost(t *testing.T) GetPostsByUidRow {
	user1 := createRandomUser(t)
	arg := CreatePostTxParam{
		PostID:       uuid.New(),
		Uid:          user1.Uid,
		MainCategory: util.RandomMainCategory(),
		SubCategory1: util.RandomString(5),
		SubCategory2: util.RandomString(5),
		PostText:     util.ToPointerOrNil(util.RandomString(100)),
		PhotoUrl:     util.ToPointerOrNil(util.RandomString(100)),
		Expense:      util.ToPointerOrNil(rand.Int63n(10000)),
		Location:     util.ToPointerOrNil(util.RandomString(20)),
		PublicTypeNo: util.RandomPublicTypeNo(),
	}

	post, err := testRepository.CreatePostTx(context.Background(), arg)
	require.NoError(t, err)

	user2, err := testRepository.GetUser(context.Background(), user1.Uid)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	subCategory1, err := testRepository.GetPostSubCategory1ByPostId(context.Background(), post.PostID)
	require.NoError(t, err)
	require.NotEmpty(t, subCategory1)

	subCategory2, err := testRepository.GetPostSubCategory2ByPostId(context.Background(), post.PostID)
	require.NoError(t, err)
	require.NotEmpty(t, subCategory2)

	postByUid := GetPostsByUidRow{
		PostID:        post.PostID,
		Uid:           post.Uid,
		Username:      user2.Username,
		MainCategory:  post.MainCategory,
		SubCategory:   &subCategory1.SubCategory,
		SubCategory_2: &subCategory2.SubCategory,
		PostText:      post.PostText,
		PhotoUrl:      post.PhotoUrl,
		Expense:       post.Expense,
		Location:      post.Location,
		PublicTypeNo:  post.PublicTypeNo,
		CreatedAt:     post.CreatedAt,
	}

	require.NotEmpty(t, post)
	require.Equal(t, arg.PostID, postByUid.PostID)
	require.Equal(t, arg.Uid, postByUid.Uid)
	require.Equal(t, user1.Username, postByUid.Username)
	require.Equal(t, arg.MainCategory, postByUid.MainCategory)
	require.Equal(t, arg.SubCategory1, *postByUid.SubCategory)
	require.Equal(t, arg.SubCategory2, *postByUid.SubCategory_2)
	require.Equal(t, arg.PostText, postByUid.PostText)
	require.Equal(t, arg.PhotoUrl, postByUid.PhotoUrl)
	require.Equal(t, arg.Expense, postByUid.Expense)
	require.Equal(t, arg.Location, postByUid.Location)
	require.Equal(t, arg.PublicTypeNo, postByUid.PublicTypeNo)
	require.NotZero(t, postByUid.CreatedAt)

	return postByUid
}

func createRandomPostsByUid(t *testing.T) (uid string, numPosts int) {
	user := createRandomUser(t)
	numPosts = rand.Intn(6)
	for i := 0; i < numPosts; i++ {
		arg := CreatePostTxParam{
			PostID:       uuid.New(),
			Uid:          user.Uid,
			MainCategory: util.RandomMainCategory(),
			SubCategory1: util.RandomString(5),
			SubCategory2: util.RandomString(5),
			PostText:     util.ToPointerOrNil(util.RandomString(100)),
			PhotoUrl:     util.ToPointerOrNil(util.RandomString(100)),
			Expense:      util.ToPointerOrNil(rand.Int63n(10000)),
			Location:     util.ToPointerOrNil(util.RandomString(20)),
			PublicTypeNo: util.RandomPublicTypeNo(),
		}
		_, err := testRepository.CreatePostTx(context.Background(), arg)
		require.NoError(t, err)
	}
	return user.Uid, numPosts
}
