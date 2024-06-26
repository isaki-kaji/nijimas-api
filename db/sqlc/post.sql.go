// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: post.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO "post" (
  "post_id",
  "uid",
  "main_category",
  "post_text",
  "photo_url",
  "expense",
  "location",
  "public_type_no"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING post_id, uid, main_category, post_text, photo_url, expense, location, public_type_no, created_at
`

type CreatePostParams struct {
	PostID       uuid.UUID `json:"post_id"`
	Uid          string    `json:"uid"`
	MainCategory string    `json:"main_category"`
	PostText     *string   `json:"post_text"`
	PhotoUrl     *string   `json:"photo_url"`
	Expense      *int64    `json:"expense"`
	Location     *string   `json:"location"`
	PublicTypeNo string    `json:"public_type_no"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost,
		arg.PostID,
		arg.Uid,
		arg.MainCategory,
		arg.PostText,
		arg.PhotoUrl,
		arg.Expense,
		arg.Location,
		arg.PublicTypeNo,
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
	)
	return i, err
}

const getPostById = `-- name: GetPostById :one
SELECT
  p."post_id",
  u."uid",
  u."username",
  u."profile_image_url",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
WHERE p."post_id" = $1
`

type GetPostByIdRow struct {
	PostID          uuid.UUID `json:"post_id"`
	Uid             string    `json:"uid"`
	Username        string    `json:"username"`
	ProfileImageUrl *string   `json:"profile_image_url"`
	MainCategory    string    `json:"main_category"`
	SubCategory     *string   `json:"sub_category"`
	SubCategory_2   *string   `json:"sub_category_2"`
	PostText        *string   `json:"post_text"`
	PhotoUrl        *string   `json:"photo_url"`
	Expense         *int64    `json:"expense"`
	Location        *string   `json:"location"`
	PublicTypeNo    string    `json:"public_type_no"`
	CreatedAt       time.Time `json:"created_at"`
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
		&i.SubCategory,
		&i.SubCategory_2,
		&i.PostText,
		&i.PhotoUrl,
		&i.Expense,
		&i.Location,
		&i.PublicTypeNo,
		&i.CreatedAt,
	)
	return i, err
}

const getPostsByCategory = `-- name: GetPostsByCategory :many
SELECT
  p."post_id",
  u."uid",
  u."username",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
WHERE 
  (p."main_category" = $1 OR $1 IS NULL) AND
  (ps1."sub_category" = $2 OR $2 IS NULL) AND
  (ps2."sub_category" = $3 OR $3 IS NULL)
ORDER BY p."created_at" DESC
LIMIT 50
`

type GetPostsByCategoryParams struct {
	MainCategory  string `json:"main_category"`
	SubCategory   string `json:"sub_category"`
	SubCategory_2 string `json:"sub_category_2"`
}

type GetPostsByCategoryRow struct {
	PostID        uuid.UUID `json:"post_id"`
	Uid           string    `json:"uid"`
	Username      string    `json:"username"`
	MainCategory  string    `json:"main_category"`
	SubCategory   *string   `json:"sub_category"`
	SubCategory_2 *string   `json:"sub_category_2"`
	PostText      *string   `json:"post_text"`
	PhotoUrl      *string   `json:"photo_url"`
	Expense       *int64    `json:"expense"`
	Location      *string   `json:"location"`
	PublicTypeNo  string    `json:"public_type_no"`
	CreatedAt     time.Time `json:"created_at"`
}

func (q *Queries) GetPostsByCategory(ctx context.Context, arg GetPostsByCategoryParams) ([]GetPostsByCategoryRow, error) {
	rows, err := q.db.Query(ctx, getPostsByCategory, arg.MainCategory, arg.SubCategory, arg.SubCategory_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsByCategoryRow{}
	for rows.Next() {
		var i GetPostsByCategoryRow
		if err := rows.Scan(
			&i.PostID,
			&i.Uid,
			&i.Username,
			&i.MainCategory,
			&i.SubCategory,
			&i.SubCategory_2,
			&i.PostText,
			&i.PhotoUrl,
			&i.Expense,
			&i.Location,
			&i.PublicTypeNo,
			&i.CreatedAt,
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

const getPostsByFollowing = `-- name: GetPostsByFollowing :many
SELECT 
  p."post_id",
  u."uid",
  u."username",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
JOIN "follow_user" AS f ON f."follow_uid" = p."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
WHERE f."uid" = $1
ORDER BY p."created_at" DESC
LIMIT 50
`

type GetPostsByFollowingRow struct {
	PostID        uuid.UUID `json:"post_id"`
	Uid           string    `json:"uid"`
	Username      string    `json:"username"`
	MainCategory  string    `json:"main_category"`
	SubCategory   *string   `json:"sub_category"`
	SubCategory_2 *string   `json:"sub_category_2"`
	PostText      *string   `json:"post_text"`
	PhotoUrl      *string   `json:"photo_url"`
	Expense       *int64    `json:"expense"`
	Location      *string   `json:"location"`
	PublicTypeNo  string    `json:"public_type_no"`
	CreatedAt     time.Time `json:"created_at"`
}

func (q *Queries) GetPostsByFollowing(ctx context.Context, uid string) ([]GetPostsByFollowingRow, error) {
	rows, err := q.db.Query(ctx, getPostsByFollowing, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsByFollowingRow{}
	for rows.Next() {
		var i GetPostsByFollowingRow
		if err := rows.Scan(
			&i.PostID,
			&i.Uid,
			&i.Username,
			&i.MainCategory,
			&i.SubCategory,
			&i.SubCategory_2,
			&i.PostText,
			&i.PhotoUrl,
			&i.Expense,
			&i.Location,
			&i.PublicTypeNo,
			&i.CreatedAt,
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

const getPostsByMainCategory = `-- name: GetPostsByMainCategory :many
SELECT
  p."post_id",
  u."uid",
  u."username",
  u."profile_image_url",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at",
  f."uid" IS NOT NULL AS "is_favorite"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
LEFT JOIN "favorite" AS f
ON p."post_id" = f."post_id" AND f."uid" = $1
WHERE p."main_category" = $2
ORDER BY p."created_at" DESC
LIMIT 50
`

type GetPostsByMainCategoryParams struct {
	Uid          string `json:"uid"`
	MainCategory string `json:"main_category"`
}

type GetPostsByMainCategoryRow struct {
	PostID          uuid.UUID   `json:"post_id"`
	Uid             string      `json:"uid"`
	Username        string      `json:"username"`
	ProfileImageUrl *string     `json:"profile_image_url"`
	MainCategory    string      `json:"main_category"`
	SubCategory     *string     `json:"sub_category"`
	SubCategory_2   *string     `json:"sub_category_2"`
	PostText        *string     `json:"post_text"`
	PhotoUrl        *string     `json:"photo_url"`
	Expense         *int64      `json:"expense"`
	Location        *string     `json:"location"`
	PublicTypeNo    string      `json:"public_type_no"`
	CreatedAt       time.Time   `json:"created_at"`
	IsFavorite      interface{} `json:"is_favorite"`
}

func (q *Queries) GetPostsByMainCategory(ctx context.Context, arg GetPostsByMainCategoryParams) ([]GetPostsByMainCategoryRow, error) {
	rows, err := q.db.Query(ctx, getPostsByMainCategory, arg.Uid, arg.MainCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsByMainCategoryRow{}
	for rows.Next() {
		var i GetPostsByMainCategoryRow
		if err := rows.Scan(
			&i.PostID,
			&i.Uid,
			&i.Username,
			&i.ProfileImageUrl,
			&i.MainCategory,
			&i.SubCategory,
			&i.SubCategory_2,
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

const getPostsBySubCategory = `-- name: GetPostsBySubCategory :many
SELECT
  p."post_id",
  u."uid",
  u."username",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
WHERE 
  (ps1."sub_category" = $1 OR $1 IS NULL) AND
  (ps2."sub_category" = $2 OR $2 IS NULL)
ORDER BY p."created_at" DESC
LIMIT 50
`

type GetPostsBySubCategoryParams struct {
	SubCategory   string `json:"sub_category"`
	SubCategory_2 string `json:"sub_category_2"`
}

type GetPostsBySubCategoryRow struct {
	PostID        uuid.UUID `json:"post_id"`
	Uid           string    `json:"uid"`
	Username      string    `json:"username"`
	MainCategory  string    `json:"main_category"`
	SubCategory   *string   `json:"sub_category"`
	SubCategory_2 *string   `json:"sub_category_2"`
	PostText      *string   `json:"post_text"`
	PhotoUrl      *string   `json:"photo_url"`
	Expense       *int64    `json:"expense"`
	Location      *string   `json:"location"`
	PublicTypeNo  string    `json:"public_type_no"`
	CreatedAt     time.Time `json:"created_at"`
}

func (q *Queries) GetPostsBySubCategory(ctx context.Context, arg GetPostsBySubCategoryParams) ([]GetPostsBySubCategoryRow, error) {
	rows, err := q.db.Query(ctx, getPostsBySubCategory, arg.SubCategory, arg.SubCategory_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsBySubCategoryRow{}
	for rows.Next() {
		var i GetPostsBySubCategoryRow
		if err := rows.Scan(
			&i.PostID,
			&i.Uid,
			&i.Username,
			&i.MainCategory,
			&i.SubCategory,
			&i.SubCategory_2,
			&i.PostText,
			&i.PhotoUrl,
			&i.Expense,
			&i.Location,
			&i.PublicTypeNo,
			&i.CreatedAt,
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

const getPostsByUid = `-- name: GetPostsByUid :many
SELECT
  p."post_id",
  u."uid",
  u."username",
  u."profile_image_url",
  p."main_category",
  ps1."sub_category",
  ps2."sub_category",
  p."post_text",
  p."photo_url",
  p."expense",
  p."location",
  p."public_type_no",
  p."created_at",
  f."uid" IS NOT NULL AS "is_favorite"
FROM "post" AS p
JOIN "user" AS u ON p."uid" = u."uid"
LEFT JOIN "post_subcategory" AS ps1
ON p."post_id" = ps1."post_id" AND ps1."subcategory_no" = '1'
LEFT JOIN "post_subcategory" AS ps2
ON p."post_id" = ps2."post_id" AND ps2."subcategory_no" = '2'
LEFT JOIN "favorite" AS f
ON p."post_id" = f."post_id" AND f."uid" = $1
WHERE p."uid" = $2
ORDER BY p."created_at" DESC
LIMIT 50
`

type GetPostsByUidParams struct {
	Uid   string `json:"uid"`
	Uid_2 string `json:"uid_2"`
}

type GetPostsByUidRow struct {
	PostID          uuid.UUID   `json:"post_id"`
	Uid             string      `json:"uid"`
	Username        string      `json:"username"`
	ProfileImageUrl *string     `json:"profile_image_url"`
	MainCategory    string      `json:"main_category"`
	SubCategory     *string     `json:"sub_category"`
	SubCategory_2   *string     `json:"sub_category_2"`
	PostText        *string     `json:"post_text"`
	PhotoUrl        *string     `json:"photo_url"`
	Expense         *int64      `json:"expense"`
	Location        *string     `json:"location"`
	PublicTypeNo    string      `json:"public_type_no"`
	CreatedAt       time.Time   `json:"created_at"`
	IsFavorite      interface{} `json:"is_favorite"`
}

func (q *Queries) GetPostsByUid(ctx context.Context, arg GetPostsByUidParams) ([]GetPostsByUidRow, error) {
	rows, err := q.db.Query(ctx, getPostsByUid, arg.Uid, arg.Uid_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsByUidRow{}
	for rows.Next() {
		var i GetPostsByUidRow
		if err := rows.Scan(
			&i.PostID,
			&i.Uid,
			&i.Username,
			&i.ProfileImageUrl,
			&i.MainCategory,
			&i.SubCategory,
			&i.SubCategory_2,
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

const updatePost = `-- name: UpdatePost :one
UPDATE "post" SET
  "main_category" = COALESCE($1, "main_category"),
  "post_text" = COALESCE($2, "post_text"),
  "photo_url" = COALESCE($3, "photo_url"),
  "expense" = COALESCE($4, "expense"),
  "public_type_no" = COALESCE($5, "public_type_no")
WHERE "post_id" = $6
RETURNING post_id, uid, main_category, post_text, photo_url, expense, location, public_type_no, created_at
`

type UpdatePostParams struct {
	MainCategory *string   `json:"main_category"`
	PostText     *string   `json:"post_text"`
	PhotoUrl     *string   `json:"photo_url"`
	Expense      *int64    `json:"expense"`
	PublicTypeNo *string   `json:"public_type_no"`
	PostID       uuid.UUID `json:"post_id"`
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
	)
	return i, err
}
