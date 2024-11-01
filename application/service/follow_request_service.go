package service

import (
	"context"
	"errors"

	"github.com/isaki-kaji/nijimas-api/apperror"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jackc/pgx/v5"
)

type FollowRequestService interface {
	DoFollowRequest(ctx context.Context, arg DoFollowRequestParams) (db.FollowRequest, error)
	CancelFollowRequest(ctx context.Context, arg CancelFollowRequestParams) (db.FollowRequest, error)
}

func NewFollowRequestService(repository db.Repository) FollowRequestService {
	return &FollowRequestServiceImpl{repository: repository}
}

type FollowRequestServiceImpl struct {
	repository db.Repository
}

type DoFollowRequestParams struct {
	Uid          string `json:"-"`
	FollowingUid string `json:"following_uid" binding:"required"`
}

func (s *FollowRequestServiceImpl) DoFollowRequest(ctx context.Context, arg DoFollowRequestParams) (db.FollowRequest, error) {
	fArg := db.GetFollowParams(arg)
	frArg := db.GetFollowRequestParams(arg)

	_, err := s.repository.GetFollow(ctx, fArg)
	if err == nil {
		return db.FollowRequest{}, apperror.DataConflict.Wrap(ErrFollowAlreadyExists, "follow already exists")
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return db.FollowRequest{}, apperror.GetDataFailed.Wrap(err, "failed to get follow")
	}

	_, err = s.repository.GetFollowRequest(ctx, frArg)
	if err == nil {
		return db.FollowRequest{}, apperror.DataConflict.Wrap(ErrFollowRequestAlreadyExists, "follow request already exists")
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return db.FollowRequest{}, apperror.GetDataFailed.Wrap(err, "failed to get follow request")
	}

	frId, err := util.GenerateUUID()
	if err != nil {
		return db.FollowRequest{}, apperror.OtherInternalErr.Wrap(err, "failed to generate follow request id")
	}
	createParams := db.CreateFollowRequestParams{
		RequestID:    frId,
		Uid:          arg.Uid,
		FollowingUid: arg.FollowingUid,
	}

	followRequest, err := s.repository.CreateFollowRequest(ctx, createParams)
	if err != nil {
		return db.FollowRequest{}, apperror.InsertDataFailed.Wrap(err, "failed to create follow request")
	}

	return followRequest, nil
}

type CancelFollowRequestParams struct {
	Uid          string `json:"-"`
	FollowingUid string `json:"following_uid" binding:"required"`
}

func (s *FollowRequestServiceImpl) CancelFollowRequest(ctx context.Context, arg CancelFollowRequestParams) (db.FollowRequest, error) {

	gArg := db.GetFollowRequestParams(arg)
	dArg := db.DeleteFollowRequestParams(arg)

	_, err := s.repository.GetFollowRequest(ctx, gArg)
	if err != nil {
		return db.FollowRequest{}, apperror.GetDataFailed.Wrap(err, "failed to get follow request")
	}

	followRequest, err := s.repository.DeleteFollowRequest(ctx, dArg)
	if err != nil {
		return db.FollowRequest{}, apperror.DeleteDataFailed.Wrap(err, "failed to delete follow request")
	}

	return followRequest, nil
}
