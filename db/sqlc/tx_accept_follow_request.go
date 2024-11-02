package db

import (
	"context"

	"github.com/google/uuid"
)

type AcceptFollowRequestTxParams struct {
	RequestId    uuid.UUID
	Uid          string
	FollowingUid string
}

func (r *SQLRepository) AcceptFollowRequestTx(ctx context.Context, param AcceptFollowRequestTxParams) (Follow, error) {
	tx, err := r.connPool.Begin(ctx)
	if err != nil {
		return Follow{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()
	qtx := r.WithTx(tx)

	_, err = qtx.UpdateFollowRequestToAccepted(ctx, param.RequestId)
	if err != nil {
		return Follow{}, err
	}

	createFollowParams := CreateFollowParams{
		Uid:          param.Uid,
		FollowingUid: param.FollowingUid,
	}

	post, err := qtx.CreateFollow(ctx, createFollowParams)
	if err != nil {
		return Follow{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return Follow{}, err
	}

	return post, nil
}
