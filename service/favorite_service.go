package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/isaki-kaji/nijimas-api/apperror"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/jackc/pgx/v5"
)

type FavoriteService interface {
	ToggleFavorite(ctx context.Context, arg ToggleFavoriteParams) (db.Favorite, string, error)
}

func NewFavoriteService(repository db.Repository) FavoriteService {
	return &FavoriteServiceImpl{repository: repository}
}

type FavoriteServiceImpl struct {
	repository db.Repository
}

type ToggleFavoriteParams struct {
	PostID uuid.UUID `json:"post_id" binding:"required"`
	Uid    string    `json:"-"`
}

func (s *FavoriteServiceImpl) ToggleFavorite(ctx context.Context, arg ToggleFavoriteParams) (db.Favorite, string, error) {
	getArg := db.GetFavoriteParams(arg)
	_, err := s.repository.GetFavorite(ctx, getArg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			createArg := db.CreateFavoriteParams(arg)
			createdFavorite, _ := s.repository.CreateFavorite(ctx, createArg)
			return createdFavorite, FlagCreated, nil
		}
		err = apperror.GetDataFailed.Wrap(err, "failed to get favorite")
		return db.Favorite{}, "", err
	}

	deleteArg := db.DeleteFavoriteParams(arg)
	deletedFavorite, err := s.repository.DeleteFavorite(ctx, deleteArg)
	if err != nil {
		err = apperror.DeleteDataFailed.Wrap(err, "failed to delete favorite")
		return db.Favorite{}, "", err
	}
	return deletedFavorite, FlagDeleted, err
}
