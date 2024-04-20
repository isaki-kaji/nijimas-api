package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/domain"
	"github.com/jackc/pgx/v5"
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

	param := db.CreatePostParams{
		PostID:       uuid,
		Uid:          arg.Uid,
		MainCategory: arg.MainCategory,
		PostText:     PointerOrNil(arg.PostText),
		PhotoUrl:     PointerOrNil(arg.PhotoUrl),
		Location:     arg.Location,
		MealFlag:     arg.MealFlag,
		PublicTypeNo: arg.PublicTypeNo,
	}
	newPost, err := s.repository.CreatePost(ctx, param)
	if err != nil {
		return db.Post{}, err
	}

	err = s.handleSubCategory(ctx, uuid, arg.SubCategory1, "1")
	if err != nil {
		return db.Post{}, err
	}
	err = s.handleSubCategory(ctx, uuid, arg.SubCategory2, "2")
	if err != nil {
		return db.Post{}, err
	}
	return newPost, nil
}

func (s *PostService) handleSubCategory(ctx context.Context, uuid uuid.UUID, subCategory string, subCategoryNo string) error {
	if subCategory == "" {
		return nil
	}
	if _, err := s.repository.GetSubCategory(ctx, subCategory); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
		_, err := s.repository.CreateSubCategory(ctx, subCategory)
		if err != nil {
			return err
		}
	}
	createPostSubCategoryParam := db.CreatePostSubCategoryParams{
		PostID:        uuid,
		SubcategoryNo: subCategoryNo,
		SubCategory:   subCategory,
	}
	_, err := s.repository.CreatePostSubCategory(ctx, createPostSubCategoryParam)
	if err != nil {
		return err
	}
	return nil
}
