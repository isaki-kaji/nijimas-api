package service

import (
	"context"
	"errors"

	"github.com/isaki-kaji/nijimas-api/apperror"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/jackc/pgx/v5"
)

type FollowService interface {
	ToggleFollow(ctx context.Context, arg ToggleFollowParams) (db.Follow, string, error)
}

func NewFollowService(repository db.Repository) FollowService {
	return &FollowServiceImpl{repository: repository}
}

type FollowServiceImpl struct {
	repository db.Repository
}

type ToggleFollowParams struct {
	Uid          string `json:"-"`
	FollowingUid string `json:"following_uid" binding:"required"`
}

func (s *FollowServiceImpl) ToggleFollow(ctx context.Context, arg ToggleFollowParams) (db.Follow, string, error) {
	getArg := db.GetFollowParams(arg)
	_, err := s.repository.GetFollow(ctx, getArg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			createArg := db.CreateFollowParams(arg)
			createdFollow, _ := s.repository.CreateFollow(ctx, createArg)
			return createdFollow, FlagCreated, nil
		}
		err = apperror.GetDataFailed.Wrap(err, "failed to get follow")
		return db.Follow{}, "", err
	}

	deleteArg := db.DeleteFollowParams(arg)
	deletedFollow, err := s.repository.DeleteFollow(ctx, deleteArg)
	if err != nil {
		err = apperror.DeleteDataFailed.Wrap(err, "failed to delete follow")
		return db.Follow{}, "", err
	}
	return deletedFollow, FlagDeleted, err
}
