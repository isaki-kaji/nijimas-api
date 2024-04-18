package domain

import (
	"context"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

type PostService interface {
	CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error)
}

type CreatePostRequest struct {
	PostID       string `json:"post_id" binding:"required"`
	Uid          string `json:"uid" binding:"required"`
	MainCategory string `json:"main_category" binding:"required"`
}
