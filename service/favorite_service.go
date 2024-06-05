package service

import (
	"context"
	"errors"

	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/jackc/pgx/v5"
)

type FavoriteService interface {
	ToggleFavorite(ctx context.Context, arg db.GetFavoriteParams) (db.Favorite, string, error)
}

func NewFavoriteService(repository db.Repository) FavoriteService {
	return &FavoriteServiceImpl{repository: repository}
}

type FavoriteServiceImpl struct {
	repository db.Repository
}

func (s *FavoriteServiceImpl) ToggleFavorite(ctx context.Context, arg db.GetFavoriteParams) (db.Favorite, string, error) {
	_, err := s.repository.GetFavorite(ctx, arg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			createArg := db.CreateFavoriteParams(arg)
			createdFavorite, _ := s.repository.CreateFavorite(ctx, createArg)
			return createdFavorite, "created", nil
		}
		return db.Favorite{}, "", err
	}

	deleteArg := db.DeleteFavoriteParams(arg)
	deletedFavorite, err := s.repository.DeleteFavorite(ctx, deleteArg)
	return deletedFavorite, "deleted", err
}
