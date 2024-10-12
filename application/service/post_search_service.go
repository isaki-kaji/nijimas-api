package service

import (
	"context"

	"github.com/isaki-kaji/nijimas-api/apperror"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

type PostSearchService interface {
	GetPostsByUid(ctx context.Context, ownUid string, uid string) ([]PostResponse, error)
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
