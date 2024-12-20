// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Favorite struct {
	PostID    uuid.UUID `json:"post_id"`
	Uid       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
}

type Follow struct {
	Uid          string    `json:"uid"`
	FollowingUid string    `json:"following_uid"`
	CreatedAt    time.Time `json:"created_at"`
}

type FollowRequest struct {
	RequestID    uuid.UUID `json:"request_id"`
	Uid          string    `json:"uid"`
	FollowingUid string    `json:"following_uid"`
	// 0:申請中, 1:承認済, 2:拒否済
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MainCategory struct {
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
}

type Post struct {
	PostID       uuid.UUID       `json:"post_id"`
	Uid          string          `json:"uid"`
	MainCategory string          `json:"main_category"`
	PostText     *string         `json:"post_text"`
	PhotoUrl     *string         `json:"photo_url"`
	Expense      decimal.Decimal `json:"expense"`
	Location     *string         `json:"location"`
	// 0:公開、1:フォロワーにのみ公開、2:非公開
	PublicTypeNo string    `json:"public_type_no"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PostSubcategory struct {
	PostID     uuid.UUID `json:"post_id"`
	CategoryNo string    `json:"category_no"`
	CategoryID uuid.UUID `json:"category_id"`
}

type SubCategory struct {
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	Uid             string    `json:"uid"`
	Username        string    `json:"username"`
	SelfIntro       *string   `json:"self_intro"`
	ProfileImageUrl *string   `json:"profile_image_url"`
	CountryCode     *string   `json:"country_code"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type UserTopSubcategory struct {
	Uid        string    `json:"uid"`
	CategoryNo string    `json:"category_no"`
	CategoryID uuid.UUID `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}
