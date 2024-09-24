package service

import (
	"context"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jinzhu/copier"
)

type PostService interface {
	CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error)
	GetPostsByUid(ctx context.Context, param db.GetPostsByUidParams) ([]PostResponse, error)
	GetPostsByMainCategory(ctx context.Context, param db.GetPostsByMainCategoryParams) ([]PostResponse, error)
}

func NewPostService(repository db.Repository, store *firestore.Client) PostService {
	return &PostServiceImpl{repository: repository}
}

type PostServiceImpl struct {
	repository db.Repository
}

type CreatePostRequest struct {
	PostID       string `json:"post_id" binding:"required,uuid"`
	Uid          string `json:"uid" binding:"required"`
	MainCategory string `json:"main_category" binding:"required,max=255"`
	SubCategory1 string `json:"sub_category1" binding:"max=255"`
	SubCategory2 string `json:"sub_category2" binding:"max=255"`
	PostText     string `json:"post_text"`
	PhotoUrl     string `json:"photo_url" binding:"max=2000"`
	Expense      int64  `json:"expense" binding:"lt=100000000"`
	Location     string `json:"location"`
	PublicTypeNo string `json:"public_type_no" binding:"required,oneof=1 2 3"`
}

func (s *PostServiceImpl) CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error) {
	uuid, err := uuid.Parse(arg.PostID)
	if err != nil {
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
		Expense:      util.ToPointerOrNil(arg.Expense),
		PublicTypeNo: arg.PublicTypeNo,
	}

	post, err := s.repository.CreatePostTx(ctx, param)
	if err != nil {
		return db.Post{}, err
	}

	// go func() {
	// 	err := function.CalcUserPosts(arg.Uid)
	// 	if err != nil {
	// 		slog.Error("failed to calc user posts: %v", err)
	// 	}
	// }()
	return post, nil
}

type PostResponse struct {
	PostID          string    `json:"post_id"`
	Uid             string    `json:"uid"`
	Username        string    `json:"username"`
	ProfileImageUrl *string   `json:"profile_image_url"`
	MainCategory    string    `json:"main_category"`
	SubCategory1    *string   `json:"sub_category1"`
	SubCategory2    *string   `json:"sub_category2"`
	PostText        *string   `json:"post_text"`
	PhotoUrl        []string  `json:"photo_url"`
	Expense         *int64    `json:"expense"`
	Location        *string   `json:"location"`
	PublicTypeNo    string    `json:"public_type_no"`
	CreatedAt       time.Time `json:"created_at"`
	IsFavorite      bool      `json:"is_favorite"`
}

func (s *PostServiceImpl) GetPostsByUid(ctx context.Context, param db.GetPostsByUidParams) ([]PostResponse, error) {
	posts, err := s.repository.GetPostsByUid(ctx, param)
	if err != nil {
		return nil, err
	}
	return transformPosts(posts)
}

func (s *PostServiceImpl) GetPostsByMainCategory(ctx context.Context, param db.GetPostsByMainCategoryParams) ([]PostResponse, error) {
	posts, err := s.repository.GetPostsByMainCategory(ctx, param)
	if err != nil {
		return nil, err
	}
	return transformPosts(posts)
}

// IDで一つだけ取得する可能性があるから、PostResponseを返すようにするべきかも
func transformPosts[T any](postsRow []T) ([]PostResponse, error) {
	response := make([]PostResponse, 0, len(postsRow))

	for _, post := range postsRow {
		var commonRow CommonGetPostsRow
		if err := copier.Copy(&commonRow, post); err != nil {
			return nil, err
		}

		p := PostResponse{}
		err := copier.Copy(&p, commonRow)
		if err != nil {
			return nil, err
		}
		p.PhotoUrl = splitPhotoUrl(commonRow.PhotoUrl)
		p.SubCategory1 = commonRow.SubCategory
		p.SubCategory2 = commonRow.SubCategory_2
		response = append(response, p)
	}

	return response, nil
}

func splitPhotoUrl(photoUrl *string) []string {
	if photoUrl == nil {
		return []string{}
	}
	return strings.Split(*photoUrl, ",")
}

type CommonGetPostsRow struct {
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
	IsFavorite      any       `json:"is_favorite"`
}
