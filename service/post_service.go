package service

import db "github.com/isaki-kaji/nijimas-api/db/sqlc"

type PostService struct {
	repository db.Repository
}

func NewPostService(repository db.Repository) *PostService {
	return &PostService{repository: repository}
}

func (s *PostService) CreatePost() {

}
