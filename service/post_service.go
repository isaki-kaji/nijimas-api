package service

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jinzhu/copier"
)

type PostService interface {
	CreatePost(ctx context.Context, arg CreatePostRequest) (db.Post, error)
	GetPostsByUid(ctx context.Context, uid string) ([]PostResponse, error)
}

func NewPostService(repository db.Repository) PostService {
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
		PublicTypeNo: arg.PublicTypeNo}

	return s.repository.CreatePostTx(ctx, param)
}

type PostResponse struct {
	PostID       string    `json:"post_id"`
	Uid          string    `json:"uid"`
	Username     string    `json:"username"`
	MainCategory string    `json:"main_category"`
	SubCategory1 *string   `json:"sub_category1"`
	SubCategory2 *string   `json:"sub_category2"`
	PostText     *string   `json:"post_text"`
	PhotoUrl     []string  `json:"photo_url"`
	Expense      *int64    `json:"expense"`
	Location     *string   `json:"location"`
	PublicTypeNo string    `json:"public_type_no"`
	CreatedAt    time.Time `json:"created_at"`
}

func (s *PostServiceImpl) GetPostsByUid(ctx context.Context, uid string) ([]PostResponse, error) {
	response := []PostResponse{}
	posts, err := s.repository.GetPostsByUid(ctx, uid)
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		p := PostResponse{}
		err := copier.Copy(&p, post)
		if err != nil {
			return nil, err
		}
		p.PhotoUrl = splitPhotoUrl(post.PhotoUrl)
		p.SubCategory1 = post.SubCategory
		p.SubCategory2 = post.SubCategory_2
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
