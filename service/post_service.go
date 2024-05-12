package service

import (
	"context"

	"github.com/google/uuid"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
)

type PostService interface {
	CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error)
}

func NewPostService(repository db.Repository) PostService {
	return &PostServiceImpl{repository: repository}
}

type PostServiceImpl struct {
	repository db.Repository
}

type CreatePostRequest struct {
	PostID       string `json:"post_id" binding:"required,uuid"`
	Uid          string `json:"uid" binding:"required"`
	MainCategory string `json:"main_category" binding:"required,max=255"`
	SubCategory1 string `json:"sub_category1" binding:"max=255"`
	SubCategory2 string `json:"sub_category2" binding:"max=255"`
	PostText     string `json:"post_text"`
	PhotoUrl     string `json:"photo_url" binding:"max=2000"`
	Expense      int64  `json:"expense" binding:"lt=100000000"`
	Location     string `json:"location"`
	PublicTypeNo string `json:"public_type_no" binding:"required,oneof=1 2 3"`
}

func (s *PostServiceImpl) CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error) {
	uuid, err := uuid.Parse(arg.PostID)
	if err != nil {
		return db.Post{}, err
	}

	param := db.CreatePostTxParam{
		PostID:       uuid,
		Uid:          arg.Uid,
		MainCategory: arg.MainCategory,
		PostText:     util.ToPointerOrNil(arg.PostText),
		PhotoUrl:     util.ToPointerOrNil(arg.PhotoUrl),
		SubCategory1: arg.SubCategory1,
		SubCategory2: arg.SubCategory2,
		Location:     util.ToPointerOrNil(arg.Location),
		Expense:      util.ToPointerOrNil(arg.Expense),
		PublicTypeNo: arg.PublicTypeNo}

	return s.repository.CreatePostTx(ctx, param)
}
