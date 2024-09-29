// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: posts.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
  post_id,
  uid,
  main_category,
  post_text,
  photo_url,
  expense,
  location,
  public_type_no
) VALUES (
  $1,
  $2,
  $3,
  CASE WHEN $6::text = '' THEN NULL ELSE $6::text END,
  CASE WHEN $7::text = '' THEN NULL ELSE $7::text END,
  $4,
  CASE WHEN $8::text = '' THEN NULL ELSE $8::text END,
  $5
) RETURNING post_id, uid, main_category, post_text, photo_url, expense, location, public_type_no, created_at, updated_at
`

type CreatePostParams struct {
	PostID       uuid.UUID       `json:"post_id"`
	Uid          string          `json:"uid"`
	MainCategory string          `json:"main_category"`
	Expense      decimal.Decimal `json:"expense"`
	PublicTypeNo string          `json:"public_type_no"`
	PostText     *string         `json:"post_text"`
	PhotoUrl     *string         `json:"photo_url"`
	Location     *string         `json:"location"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost,
		arg.PostID,
		arg.Uid,
		arg.MainCategory,
		arg.Expense,
		arg.PublicTypeNo,
		arg.PostText,
		arg.PhotoUrl,
		arg.Location,
	)
	var i Post
	err := row.Scan(
		&i.PostID,
		&i.Uid,
		&i.MainCategory,
		&i.PostText,
		&i.PhotoUrl,
		&i.Expense,
		&i.Location,
		&i.PublicTypeNo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOwnPosts = `-- name: GetOwnPosts :many
SELECT
  p.post_id,
  u.uid,
  u.username,
  u.profile_image_url,
  p.main_category,
  sc.sub_category1 AS subcategory1,
  sc.sub_category2 AS subcategory2,
  p.post_text,
  p.photo_url,
  p.expense,
  p.location,
  p.public_type_no,
  p.created_at,
  CASE WHEN f.uid IS NOT NULL THEN TRUE ELSE FALSE END AS is_favorite
FROM posts p
JOIN users u ON p.uid = u.uid
LEFT JOIN (
  SELECT
    ps.post_id,
    MAX(CASE WHEN ps.category_no = '1' THEN s.category_name ELSE NULL END) AS sub_category1,
    MAX(CASE WHEN ps.category_no = '2' THEN s.category_name ELSE NULL END) AS sub_category2
  FROM post_subcategories ps
  JOIN sub_categories s ON ps.category_id = s.category_id
  GROUP BY ps.post_id
) sc ON p.post_id = sc.post_id
LEFT JOIN favorites f
  ON p.post_id = f.post_id AND f.uid = $1
WHERE p.uid = $1
ORDER BY p.post_id DESC
LIMIT 50
`

type GetOwnPostsRow struct {
	PostID          uuid.UUID       `json:"post_id"`
	Uid             string          `json:"uid"`
	Username        string          `json:"username"`
	ProfileImageUrl *string         `json:"profile_image_url"`
	MainCategory    string          `json:"main_category"`
	Subcategory1    interface{}     `json:"subcategory1"`
	Subcategory2    interface{}     `json:"subcategory2"`
	PostText        *string         `json:"post_text"`
	PhotoUrl        *string         `json:"photo_url"`
	Expense         decimal.Decimal `json:"expense"`
	Location        *string         `json:"location"`
	PublicTypeNo    string          `json:"public_type_no"`
	CreatedAt       time.Time       `json:"created_at"`
	IsFavorite      bool            `json:"is_favorite"`
}

func (q *Queries) GetOwnPosts(ctx context.Context, uid string) ([]GetOwnPostsRow, error) {
	rows, err := q.db.Query(ctx, getOwnPosts, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOwnPostsRow{}
	for rows.Next() {
		var i GetOwnPostsRow
		if err := rows.Scan(
			&i.PostID,
			&i.Uid,
			&i.Username,
			&i.ProfileImageUrl,
			&i.MainCategory,
			&i.Subcategory1,
			&i.Subcategory2,
			&i.PostText,
			&i.PhotoUrl,
			&i.Expense,
			&i.Location,
			&i.PublicTypeNo,
			&i.CreatedAt,
			&i.IsFavorite,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostById = `-- name: GetPostById :one
SELECT
  p.post_id,
  u.uid,
  u.username,
  u.profile_image_url,
  p.main_category,
  sc.sub_category1 AS subCategory1,
  sc.sub_category2 AS subCategory2,
  p.post_text,
  p.photo_url,
  p.expense,
  p.location,
  p.public_type_no,
  p.created_at
FROM posts p
JOIN users u ON p.uid = u.uid
LEFT JOIN (
  SELECT
    ps.post_id,
    MAX(CASE WHEN ps.category_no = '1' THEN s.category_name ELSE NULL END) AS sub_category1,
    MAX(CASE WHEN ps.category_no = '2' THEN s.category_name ELSE NULL END) AS sub_category2
  FROM post_subcategories ps
  JOIN sub_categories s ON ps.category_id = s.category_id
  GROUP BY ps.post_id
) sc ON p.post_id = sc.post_id
WHERE p.post_id = $1
`

type GetPostByIdRow struct {
	PostID          uuid.UUID       `json:"post_id"`
	Uid             string          `json:"uid"`
	Username        string          `json:"username"`
	ProfileImageUrl *string         `json:"profile_image_url"`
	MainCategory    string          `json:"main_category"`
	Subcategory1    interface{}     `json:"subcategory1"`
	Subcategory2    interface{}     `json:"subcategory2"`
	PostText        *string         `json:"post_text"`
	PhotoUrl        *string         `json:"photo_url"`
	Expense         decimal.Decimal `json:"expense"`
	Location        *string         `json:"location"`
	PublicTypeNo    string          `json:"public_type_no"`
	CreatedAt       time.Time       `json:"created_at"`
}

func (q *Queries) GetPostById(ctx context.Context, postID uuid.UUID) (GetPostByIdRow, error) {
	row := q.db.QueryRow(ctx, getPostById, postID)
	var i GetPostByIdRow
	err := row.Scan(
		&i.PostID,
		&i.Uid,
		&i.Username,
		&i.ProfileImageUrl,
		&i.MainCategory,
		&i.Subcategory1,
		&i.Subcategory2,
		&i.PostText,
		&i.PhotoUrl,
		&i.Expense,
		&i.Location,
		&i.PublicTypeNo,
		&i.CreatedAt,
	)
	return i, err
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts SET
  main_category = COALESCE($1, main_category),
  post_text = COALESCE($2, post_text),
  photo_url = COALESCE($3, photo_url),
  expense = COALESCE($4, expense),
  public_type_no = COALESCE($5, public_type_no)
WHERE post_id = $6
RETURNING post_id, uid, main_category, post_text, photo_url, expense, location, public_type_no, created_at, updated_at
`

type UpdatePostParams struct {
	MainCategory *string        `json:"main_category"`
	PostText     *string        `json:"post_text"`
	PhotoUrl     *string        `json:"photo_url"`
	Expense      pgtype.Numeric `json:"expense"`
	PublicTypeNo *string        `json:"public_type_no"`
	PostID       uuid.UUID      `json:"post_id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, updatePost,
		arg.MainCategory,
		arg.PostText,
		arg.PhotoUrl,
		arg.Expense,
		arg.PublicTypeNo,
		arg.PostID,
	)
	var i Post
	err := row.Scan(
		&i.PostID,
		&i.Uid,
		&i.MainCategory,
		&i.PostText,
		&i.PhotoUrl,
		&i.Expense,
		&i.Location,
		&i.PublicTypeNo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}