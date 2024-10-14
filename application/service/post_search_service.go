package service

import (
	"context"

	"github.com/isaki-kaji/nijimas-api/apperror"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

type PostSearchService interface {
	GetPostsByUid(ctx context.Context, ownUid string, uid string) ([]PostResponse, error)
	GetPostsByMainCategory(ctx context.Context, ownUid string, mainCategory string) ([]PostResponse, error)
	GetPostsBySubCategory(ctx context.Context, ownUid string, subCategory string) ([]PostResponse, error)
	GetPostsByMainCategoryAndSubCategory(ctx context.Context, ownUid string, mainCategory string, subCategory string) ([]PostResponse, error)
}

func NewPostSearchService(repository db.Repository) PostSearchService {
	return &PostSearchServiceImpl{repository: repository}
}

type PostSearchServiceImpl struct {
	repository db.Repository
}

func (s *PostSearchServiceImpl) GetPostsByUid(ctx context.Context, ownUid string, uid string) ([]PostResponse, error) {
	getPostsByUidParam := db.GetPostsByUidParams{
		Uid:   ownUid,
		Uid_2: uid,
	}

	posts, err := s.repository.GetPostsByUid(ctx, getPostsByUidParam)
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

func (s *PostSearchServiceImpl) GetPostsByMainCategory(ctx context.Context, ownUid string, mainCategory string) ([]PostResponse, error) {
	getPostsByMainCategoryParam := db.GetPostsByMainCategoryParams{
		Uid:          ownUid,
		MainCategory: mainCategory,
	}

	posts, err := s.repository.GetPostsByMainCategory(ctx, getPostsByMainCategoryParam)
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

func (s *PostSearchServiceImpl) GetPostsBySubCategory(ctx context.Context, ownUid string, subCategory string) ([]PostResponse, error) {
	getPostsBySubCategoryParam := db.GetPostsBySubCategoryParams{
		Uid:          ownUid,
		CategoryName: subCategory,
	}

	posts, err := s.repository.GetPostsBySubCategory(ctx, getPostsBySubCategoryParam)
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

func (s *PostSearchServiceImpl) GetPostsByMainCategoryAndSubCategory(ctx context.Context, ownUid string, mainCategory string, subCategory string) ([]PostResponse, error) {
	getPostsByMainCategoryAndSubCategoryParam := db.GetPostsByMainCategoryAndSubCategoryParams{
		Uid:          ownUid,
		MainCategory: mainCategory,
		SubCategory:  subCategory,
	}

	posts, err := s.repository.GetPostsByMainCategoryAndSubCategory(ctx, getPostsByMainCategoryAndSubCategoryParam)
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
