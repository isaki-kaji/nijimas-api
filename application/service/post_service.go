package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/isaki-kaji/nijimas-api/apperror"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
)

type PostService interface {
	CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error)
	GetOwnPosts(ctx context.Context, uid string) ([]PostResponse, error)
	GetTimelinePosts(ctx context.Context, uid string) ([]PostResponse, error)
}

func NewPostService(repository db.Repository) PostService {
	return &PostServiceImpl{repository: repository}
}

type PostServiceImpl struct {
	repository db.Repository
}

type CreatePostRequest struct {
	PostId       string `json:"post_id" binding:"required"`
	Uid          string `json:"-"`
	MainCategory string `json:"main_category" binding:"required,max=255"`
	SubCategory1 string `json:"sub_category1" binding:"max=255"`
	SubCategory2 string `json:"sub_category2" binding:"max=255"`
	PostText     string `json:"post_text"`
	PhotoUrl     string `json:"photo_url" binding:"max=2000"`
	Expense      string `json:"expense" binding:"required"`
	Location     string `json:"location"`
	PublicTypeNo string `json:"public_type_no" binding:"required,oneof=0 1 2"`
}

func (s *PostServiceImpl) CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error) {

	postId, err := uuid.Parse(arg.PostId)
	if err != nil {
		err = apperror.ValidationFailed.Wrap(err, "invalid post_id")
		return db.Post{}, err
	}

	param := db.CreatePostTxParam{
		PostID:       postId,
		Uid:          arg.Uid,
		MainCategory: arg.MainCategory,
		PostText:     util.ToPointerOrNil(arg.PostText),
		PhotoUrl:     util.ToPointerOrNil(arg.PhotoUrl),
		SubCategory1: arg.SubCategory1,
		SubCategory2: arg.SubCategory2,
		Location:     util.ToPointerOrNil(arg.Location),
		Expense:      arg.Expense,
		PublicTypeNo: arg.PublicTypeNo,
		CreatedAt:    time.Now(),
	}

	post, err := s.repository.CreatePostTx(ctx, param)
	if err != nil {
		err = apperror.InsertDataFailed.Wrap(err, "failed to create post")
		return db.Post{}, err
	}
	return post, nil
}

func (s *PostServiceImpl) GetOwnPosts(ctx context.Context, uid string) ([]PostResponse, error) {
	posts, err := s.repository.GetOwnPosts(ctx, uid)
	if err != nil {
		err = apperror.GetDataFailed.Wrap(err, "failed to get posts")
		return nil, err
	}
	postsResponse, err := transformPosts(posts)
	if err != nil {
		err = apperror.GetDataFailed.Wrap(err, "failed to transform posts")
		return nil, err
	}
	return postsResponse, nil
}

func (s *PostServiceImpl) GetTimelinePosts(ctx context.Context, uid string) ([]PostResponse, error) {
	posts, err := s.repository.GetTimelinePosts(ctx, uid)
	if err != nil {
		err = apperror.GetDataFailed.Wrap(err, "failed to get posts")
		return nil, err
	}
	postsResponse, err := transformPosts(posts)
	if err != nil {
		err = apperror.GetDataFailed.Wrap(err, "failed to transform posts")
		return nil, err
	}
	return postsResponse, nil
}
