package service

import (
	"context"

	"github.com/isaki-kaji/nijimas-api/apperror"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

type SubCategoryService interface {
	GetUserUsedSubCategories(ctx context.Context, uid string) ([]db.GetUserUsedSubCategoriesRow, error)
}

func NewSubCategoryService(repository db.Repository) SubCategoryService {
	return &SubCategoryServiceImpl{repository: repository}
}

type SubCategoryServiceImpl struct {
	repository db.Repository
}

func (s *SubCategoryServiceImpl) GetUserUsedSubCategories(ctx context.Context, uid string) ([]db.GetUserUsedSubCategoriesRow, error) {
	subCategories, err := s.repository.GetUserUsedSubCategories(ctx, uid)
	if err != nil {
		err = apperror.GetDataFailed.Wrap(err, "failed to get user used sub categories")
		return []db.GetUserUsedSubCategoriesRow{}, err
	}

	return subCategories, nil
}
