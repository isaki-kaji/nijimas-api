package service

import (
	"context"

	"github.com/google/uuid"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/domain"
	"github.com/isaki-kaji/nijimas-api/util"
)

type PostService struct {
	repository db.Repository
}

func NewPostService(repository db.Repository) domain.PostService {
	return &PostService{repository: repository}
}

func (s *PostService) CreatePost(ctx context.Context, arg domain.CreatePostRequest) (db.Post, error) {
	uuid, err := uuid.Parse(arg.PostID)
	if err != nil {
		return db.Post{}, err
	}

	param := db.CreatePostTxParam{
		PostID:       uuid,
		Uid:          arg.Uid,
		MainCategory: arg.MainCategory,
		PostText:     util.PointerOrNil(arg.PostText),
		PhotoUrl:     util.PointerOrNil(arg.PhotoUrl),
		SubCategory1: arg.SubCategory1,
		SubCategory2: arg.SubCategory2,
		Location:     arg.Location,
		Expense:      util.PointerOrNil(arg.Expense),
		PublicTypeNo: arg.PublicTypeNo}

	return s.repository.CreatePostTx(ctx, param)
}
