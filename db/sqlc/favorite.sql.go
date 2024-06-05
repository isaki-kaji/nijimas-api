// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: favorite.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createFavorite = `-- name: CreateFavorite :one
INSERT INTO "favorite" (
 "post_id",
 "uid"
) VALUES (
 $1,$2
) RETURNING favorite_id, post_id, uid, created_at
`

type CreateFavoriteParams struct {
	PostID uuid.UUID `json:"post_id"`
	Uid    string    `json:"uid"`
}

func (q *Queries) CreateFavorite(ctx context.Context, arg CreateFavoriteParams) (Favorite, error) {
	row := q.db.QueryRow(ctx, createFavorite, arg.PostID, arg.Uid)
	var i Favorite
	err := row.Scan(
		&i.FavoriteID,
		&i.PostID,
		&i.Uid,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFavorite = `-- name: DeleteFavorite :one
DELETE FROM "favorite"
WHERE "post_id" = $1 AND "uid" = $2
RETURNING favorite_id, post_id, uid, created_at
`

type DeleteFavoriteParams struct {
	PostID uuid.UUID `json:"post_id"`
	Uid    string    `json:"uid"`
}

func (q *Queries) DeleteFavorite(ctx context.Context, arg DeleteFavoriteParams) (Favorite, error) {
	row := q.db.QueryRow(ctx, deleteFavorite, arg.PostID, arg.Uid)
	var i Favorite
	err := row.Scan(
		&i.FavoriteID,
		&i.PostID,
		&i.Uid,
		&i.CreatedAt,
	)
	return i, err
}

const getFavorite = `-- name: GetFavorite :one
SELECT favorite_id, post_id, uid, created_at FROM "favorite"
WHERE "post_id" = $1 AND "uid" = $2
`

type GetFavoriteParams struct {
	PostID uuid.UUID `json:"post_id"`
	Uid    string    `json:"uid"`
}

func (q *Queries) GetFavorite(ctx context.Context, arg GetFavoriteParams) (Favorite, error) {
	row := q.db.QueryRow(ctx, getFavorite, arg.PostID, arg.Uid)
	var i Favorite
	err := row.Scan(
		&i.FavoriteID,
		&i.PostID,
		&i.Uid,
		&i.CreatedAt,
	)
	return i, err
}
