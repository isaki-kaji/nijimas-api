// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: post_search.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const getPostsByUid = `-- name: GetPostsByUid :many
SELECT
  p.post_id,
  u.uid,
  u.username,
  u.profile_image_url,
  p.main_category,
  COALESCE(sc.sub_category1, '')::text AS subCategory1,
  COALESCE(sc.sub_category2, '')::text AS subCategory2,
  p.post_text,
  p.photo_url,
  p.expense,
  p.location,
  CASE WHEN f.uid IS NOT NULL THEN TRUE ELSE FALSE END AS is_favorite,
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
LEFT JOIN favorites f
  ON p.post_id = f.post_id AND f.uid = $1
WHERE p.uid = $2 AND p.public_type_no = '0'
ORDER BY p.post_id DESC
LIMIT 50
`

type GetPostsByUidParams struct {
	Uid   string `json:"uid"`
	Uid_2 string `json:"uid_2"`
}

type GetPostsByUidRow struct {
	PostID          uuid.UUID       `json:"post_id"`
	Uid             string          `json:"uid"`
	Username        string          `json:"username"`
	ProfileImageUrl *string         `json:"profile_image_url"`
	MainCategory    string          `json:"main_category"`
	Subcategory1    string          `json:"subcategory1"`
	Subcategory2    string          `json:"subcategory2"`
	PostText        *string         `json:"post_text"`
	PhotoUrl        *string         `json:"photo_url"`
	Expense         decimal.Decimal `json:"expense"`
	Location        *string         `json:"location"`
	IsFavorite      bool            `json:"is_favorite"`
	CreatedAt       time.Time       `json:"created_at"`
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
			&i.Subcategory1,
			&i.Subcategory2,
			&i.PostText,
			&i.PhotoUrl,
			&i.Expense,
			&i.Location,
			&i.IsFavorite,
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
