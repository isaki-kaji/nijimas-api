package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/isaki-kaji/nijimas-api/util"
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
	Expense      string
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

	numericExpense, err := util.ToNumeric(param.Expense)
	if err != nil {
		return Post{}, err
	}

	dbParam := CreatePostParams{
		PostID:       param.PostID,
		Uid:          param.Uid,
		MainCategory: param.MainCategory,
		PostText:     param.PostText,
		PhotoUrl:     param.PhotoUrl,
		Expense:      numericExpense,
		Location:     param.Location,
		PublicTypeNo: param.PublicTypeNo,
	}

	post, err := qtx.CreatePost(ctx, dbParam)
	if err != nil {
		return Post{}, err
	}

	err = handleSubCategory(ctx, param.PostID, param.SubCategory1, "1", qtx)
	if err != nil {
		return Post{}, err
	}
	err = handleSubCategory(ctx, param.PostID, param.SubCategory2, "2", qtx)
	if err != nil {
		return Post{}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func handleSubCategory(ctx context.Context, postID uuid.UUID, categoryName string, categoryNo string, qtx *Queries) error {
	if categoryName == "" {
		return nil
	}

	categoryId, err := registerSubCategory(ctx, categoryName, qtx)
	if err != nil {
		return err
	}

	err = registerPostSubCategory(ctx, postID, categoryId, categoryNo, qtx)
	if err != nil {
		return err
	}
	return nil
}

func registerSubCategory(ctx context.Context, categoryName string, qtx *Queries) (uuid.UUID, error) {
	subCategory, err := qtx.GetSubCategoryByName(ctx, categoryName)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return uuid.Nil, err
		}

		categoryId, err := util.GenerateUUID()
		if err != nil {
			return uuid.Nil, err
		}

		createSubCategoryParam := CreateSubCategoryParams{
			CategoryID:   categoryId,
			CategoryName: categoryName,
		}

		subCategory, err = qtx.CreateSubCategory(ctx, createSubCategoryParam)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return subCategory.CategoryID, nil
}

func registerPostSubCategory(ctx context.Context, postId uuid.UUID, categoryId uuid.UUID, categoryNo string, qtx *Queries) error {
	createPostSubCategoryParam := CreatePostSubCategoryParams{
		PostID:     postId,
		CategoryID: categoryId,
		CategoryNo: categoryNo,
	}
	_, err := qtx.CreatePostSubCategory(ctx, createPostSubCategoryParam)
	if err != nil {
		return err
	}
	return nil
}
