package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CreatePostTxParam struct {
	PostID       uuid.UUID
	Uid          string
	MainCategory string
	SubCategory1 string
	SubCategory2 string
	PostText     *string
	PhotoUrl     *string
	Expense      *int64
	MealFlag     bool
	Location     *string
	PublicTypeNo string
}

func (r *SQLRepository) CreatePostTx(ctx context.Context, param CreatePostTxParam) (Post, error) {
	tx, err := r.connPool.Begin(ctx)
	if err != nil {
		return Post{}, err
	}
	defer tx.Rollback(ctx)
	qtx := r.WithTx(tx)

	dbParam := CreatePostParams{
		PostID:       param.PostID,
		Uid:          param.Uid,
		MainCategory: param.MainCategory,
		PostText:     param.PostText,
		PhotoUrl:     param.PhotoUrl,
		Expense:      param.Expense,
		Location:     param.Location,
		PublicTypeNo: param.PublicTypeNo,
	}

	post, err := qtx.CreatePost(ctx, dbParam)
	if err != nil {
		return Post{}, err
	}

	err = r.handleSubCategory(ctx, param.PostID, param.SubCategory1, "1", qtx)
	if err != nil {
		return Post{}, err
	}
	err = r.handleSubCategory(ctx, param.PostID, param.SubCategory2, "2", qtx)
	if err != nil {
		return Post{}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (r *SQLRepository) handleSubCategory(ctx context.Context, postID uuid.UUID, subCategory string, subCategoryNo string, qtx *Queries) error {
	if subCategory == "" {
		return nil
	}

	if _, err := qtx.GetSubCategory(ctx, subCategory); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
		_, err := qtx.CreateSubCategory(ctx, subCategory)
		if err != nil {
			return err
		}
	}

	createPostSubCategoryParam := CreatePostSubCategoryParams{
		PostID:        postID,
		SubcategoryNo: subCategoryNo,
		SubCategory:   subCategory,
	}
	_, err := qtx.CreatePostSubCategory(ctx, createPostSubCategoryParam)
	if err != nil {
		return err
	}
	return nil
}
