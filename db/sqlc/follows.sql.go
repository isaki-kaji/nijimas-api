// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: follows.sql

package db

import (
	"context"
)

const createFollow = `-- name: CreateFollow :one
INSERT INTO follows (
  uid,
  following_uid
) VALUES (
  $1, $2
) RETURNING uid, following_uid, created_at
`

type CreateFollowParams struct {
	Uid          string `json:"uid"`
	FollowingUid string `json:"following_uid"`
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) (Follow, error) {
	row := q.db.QueryRow(ctx, createFollow, arg.Uid, arg.FollowingUid)
	var i Follow
	err := row.Scan(&i.Uid, &i.FollowingUid, &i.CreatedAt)
	return i, err
}

const deleteFollow = `-- name: DeleteFollow :one
DELETE FROM follows
WHERE uid = $1 AND following_uid = $2
RETURNING uid, following_uid, created_at
`

type DeleteFollowParams struct {
	Uid          string `json:"uid"`
	FollowingUid string `json:"following_uid"`
}

func (q *Queries) DeleteFollow(ctx context.Context, arg DeleteFollowParams) (Follow, error) {
	row := q.db.QueryRow(ctx, deleteFollow, arg.Uid, arg.FollowingUid)
	var i Follow
	err := row.Scan(&i.Uid, &i.FollowingUid, &i.CreatedAt)
	return i, err
}

const getFollow = `-- name: GetFollow :one
SELECT uid, following_uid, created_at FROM follows
WHERE uid = $1 AND following_uid = $2
`

type GetFollowParams struct {
	Uid          string `json:"uid"`
	FollowingUid string `json:"following_uid"`
}

func (q *Queries) GetFollow(ctx context.Context, arg GetFollowParams) (Follow, error) {
	row := q.db.QueryRow(ctx, getFollow, arg.Uid, arg.FollowingUid)
	var i Follow
	err := row.Scan(&i.Uid, &i.FollowingUid, &i.CreatedAt)
	return i, err
}

const getFollowInfo = `-- name: GetFollowInfo :one
SELECT 
EXISTS (
        SELECT 1 
        FROM follows f2
        WHERE f2.uid = $2 AND f2.following_uid = $1
    ) AS is_following,
       COUNT(CASE WHEN f.uid = $1 THEN 1 ELSE NULL END) AS following_count,
       COUNT(CASE WHEN f.following_uid = $1 THEN 1 ELSE NULL END) AS followers_count
FROM follows f
WHERE f.uid = $1 or f.following_uid = $1
`

type GetFollowInfoParams struct {
	FollowingUid string `json:"following_uid"`
	OwnUid       string `json:"own_uid"`
}

type GetFollowInfoRow struct {
	IsFollowing    bool  `json:"is_following"`
	FollowingCount int64 `json:"following_count"`
	FollowersCount int64 `json:"followers_count"`
}

func (q *Queries) GetFollowInfo(ctx context.Context, arg GetFollowInfoParams) (GetFollowInfoRow, error) {
	row := q.db.QueryRow(ctx, getFollowInfo, arg.FollowingUid, arg.OwnUid)
	var i GetFollowInfoRow
	err := row.Scan(&i.IsFollowing, &i.FollowingCount, &i.FollowersCount)
	return i, err
}
