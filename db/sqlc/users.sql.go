// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  uid,
  username,
  country_code
) VALUES (
  $1, $2, $3
) RETURNING uid, username, self_intro, profile_image_url, country_code, created_at, updated_at
`

type CreateUserParams struct {
	Uid         string  `json:"uid"`
	Username    string  `json:"username"`
	CountryCode *string `json:"country_code"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Uid, arg.Username, arg.CountryCode)
	var i User
	err := row.Scan(
		&i.Uid,
		&i.Username,
		&i.SelfIntro,
		&i.ProfileImageUrl,
		&i.CountryCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFollowUsers = `-- name: GetFollowUsers :many
SELECT
  users.uid, 
  users.username
FROM users
JOIN follows
ON users.uid = follows.follow_user_id
WHERE follows.uid = $1
`

type GetFollowUsersRow struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
}

func (q *Queries) GetFollowUsers(ctx context.Context, uid string) ([]GetFollowUsersRow, error) {
	rows, err := q.db.Query(ctx, getFollowUsers, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFollowUsersRow{}
	for rows.Next() {
		var i GetFollowUsersRow
		if err := rows.Scan(&i.Uid, &i.Username); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT uid, username, self_intro, profile_image_url, country_code, created_at, updated_at FROM users
WHERE uid = $1
`

func (q *Queries) GetUser(ctx context.Context, uid string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, uid)
	var i User
	err := row.Scan(
		&i.Uid,
		&i.Username,
		&i.SelfIntro,
		&i.ProfileImageUrl,
		&i.CountryCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET
  username = COALESCE($1, username),
  self_intro = COALESCE($2, self_intro),
  profile_image_url = COALESCE($3, profile_image_url)
WHERE uid = $4
RETURNING uid, username, self_intro, profile_image_url, country_code, created_at, updated_at
`

type UpdateUserParams struct {
	Username        *string `json:"username"`
	SelfIntro       *string `json:"self_intro"`
	ProfileImageUrl *string `json:"profile_image_url"`
	Uid             string  `json:"uid"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Username,
		arg.SelfIntro,
		arg.ProfileImageUrl,
		arg.Uid,
	)
	var i User
	err := row.Scan(
		&i.Uid,
		&i.Username,
		&i.SelfIntro,
		&i.ProfileImageUrl,
		&i.CountryCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
