package service

import (
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/shopspring/decimal"
)

type PostResponse struct {
	PostID          uuid.UUID       `json:"post_id"`
	Uid             string          `json:"uid"`
	Username        string          `json:"username"`
	ProfileImageUrl *string         `json:"profile_image_url"`
	MainCategory    string          `json:"main_category"`
	SubCategory1    *string         `json:"sub_category1"`
	SubCategory2    *string         `json:"sub_category2"`
	PostText        *string         `json:"post_text"`
	PhotoUrl        []string        `json:"photo_url"`
	Expense         decimal.Decimal `json:"expense"`
	Location        *string         `json:"location"`
	CreatedAt       time.Time       `json:"created_at"`
	IsFavorite      bool            `json:"is_favorite"`
}

func transformPosts[T any](postsRow []T) ([]PostResponse, error) {
	response := make([]PostResponse, 0, len(postsRow))

	for _, post := range postsRow {
		p := transformPost(post)
		response = append(response, p)
	}
	return response, nil
}

// 汎用的なポスト変換関数
func transformPost(post any) PostResponse {
	postVal := reflect.ValueOf(post)

	p := PostResponse{}

	// 各フィールドをコピー
	p.PostID = postVal.FieldByName("PostID").Interface().(uuid.UUID)
	p.Uid = postVal.FieldByName("Uid").Interface().(string)
	p.Username = postVal.FieldByName("Username").Interface().(string)
	p.ProfileImageUrl = postVal.FieldByName("ProfileImageUrl").Interface().(*string)
	p.MainCategory = postVal.FieldByName("MainCategory").Interface().(string)

	// Subcategory1 と Subcategory2 の処理: stringをポインタに変換
	subCategory1 := postVal.FieldByName("Subcategory1").Interface().(string)
	p.SubCategory1 = util.ToPointerOrNil(subCategory1)

	subCategory2 := postVal.FieldByName("Subcategory2").Interface().(string)
	p.SubCategory2 = util.ToPointerOrNil(subCategory2)

	p.PostText = postVal.FieldByName("PostText").Interface().(*string)
	p.PhotoUrl = splitPhotoUrl(postVal.FieldByName("PhotoUrl").Interface().(*string))
	p.Expense = postVal.FieldByName("Expense").Interface().(decimal.Decimal)
	p.Location = postVal.FieldByName("Location").Interface().(*string)
	p.CreatedAt = postVal.FieldByName("CreatedAt").Interface().(time.Time)
	p.IsFavorite = postVal.FieldByName("IsFavorite").Interface().(bool)

	return p
}

// PhotoUrlの分割
func splitPhotoUrl(photoUrl *string) []string {
	if photoUrl == nil {
		return []string{}
	}
	return strings.Split(*photoUrl, ",")
}
