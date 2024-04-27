package domain

import (
	"context"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

type PostService interface {
	CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error)
}

type CreatePostRequest struct {
	PostID       string `json:"post_id" binding:"required,uuid"`
	Uid          string `json:"uid" binding:"required"`
	MainCategory string `json:"main_category" binding:"required,max=255"`
	SubCategory1 string `json:"sub_category1" binding:"max=255"`
	SubCategory2 string `json:"sub_category2" binding:"max=255"`
	PostText     string `json:"post_text"`
	PhotoUrl     string `json:"photo_url" binding:"max=2000"`
	Expense      int64  `json:"expense"`
	Location     string `json:"location"`
	PublicTypeNo string `json:"public_type_no" binding:"required,oneof=1 2 3"`
}
