package service

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/isaki-kaji/nijimas-api/apperror"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/shopspring/decimal"
)

type PostService interface {
	CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error)
	GetOwnPosts(ctx context.Context, uid string) ([]PostResponse, error)
	// GetPostsByMainCategory(ctx context.Context, param db.GetPostsByMainCategoryParams) ([]PostResponse, error)
}

func NewPostService(repository db.Repository) PostService {
	return &PostServiceImpl{repository: repository}
}

type PostServiceImpl struct {
	repository db.Repository
}

type CreatePostRequest struct {
	Uid          string `json:"-"`
	MainCategory string `json:"main_category" binding:"required,max=255"`
	SubCategory1 string `json:"sub_category1" binding:"max=255"`
	SubCategory2 string `json:"sub_category2" binding:"max=255"`
	PostText     string `json:"post_text"`
	PhotoUrl     string `json:"photo_url" binding:"max=2000"`
	Expense      string `json:"expense" binding:"lt=100000000"`
	Location     string `json:"location"`
	PublicTypeNo string `json:"public_type_no" binding:"required,oneof=0 1 2"`
}

func (s *PostServiceImpl) CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error) {
	uuid, err := util.GenerateUUID()
	if err != nil {
		err = apperror.OtherInternalErr.Wrap(err, "failed to generate uuid")
		return db.Post{}, err
	}

	param := db.CreatePostTxParam{
		PostID:       uuid,
		Uid:          arg.Uid,
		MainCategory: arg.MainCategory,
		PostText:     util.ToPointerOrNil(arg.PostText),
		PhotoUrl:     util.ToPointerOrNil(arg.PhotoUrl),
		SubCategory1: arg.SubCategory1,
		SubCategory2: arg.SubCategory2,
		Location:     util.ToPointerOrNil(arg.Location),
		Expense:      arg.Expense,
		PublicTypeNo: arg.PublicTypeNo,
		CreatedAt:    time.Now(),
	}

	post, err := s.repository.CreatePostTx(ctx, param)
	if err != nil {
		err = apperror.InsertDataFailed.Wrap(err, "failed to create post")
		return db.Post{}, err
	}
	return post, nil
}

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
	PublicTypeNo    string          `json:"public_type_no"`
	CreatedAt       time.Time       `json:"created_at"`
	IsFavorite      bool            `json:"is_favorite"`
}

func (s *PostServiceImpl) GetOwnPosts(ctx context.Context, uid string) ([]PostResponse, error) {
	posts, err := s.repository.GetOwnPosts(ctx, uid)
	if err != nil {
		err = apperror.GetDataFailed.Wrap(err, "failed to get posts")
		return nil, err
	}
	postsResponse, err := transformPosts(posts)
	if err != nil {
		err = apperror.GetDataFailed.Wrap(err, "failed to transform posts")
		return nil, err
	}
	return postsResponse, nil
}

// func (s *PostServiceImpl) GetPostsByMainCategory(ctx context.Context, param db.GetPostsByMainCategoryParams) ([]PostResponse, error) {
// 	posts, err := s.repository.GetPostsByMainCategory(ctx, param)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return transformPosts(posts)
// }

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
	p.PublicTypeNo = postVal.FieldByName("PublicTypeNo").Interface().(string)
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
