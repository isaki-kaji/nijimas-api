package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/apperror"
	"github.com/isaki-kaji/nijimas-api/application/service"
)

type PostController struct {
	service service.PostService
}

func NewPostController(service service.PostService) *PostController {
	return &PostController{service: service}
}

func (p *PostController) CreatePost(ctx *gin.Context) {
	var req service.CreatePostRequest
	ownUid, err := checkPostReq(ctx, &req)
	if err != nil {
		return
	}
	req.Uid = ownUid

	post, err := p.service.CreatePost(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
		return
	}
	ctx.JSON(http.StatusCreated, post)
}

func (p *PostController) GetOwnPosts(ctx *gin.Context) {
	ownUid, err := checkUid(ctx)
	if err != nil {
		return
	}

	posts, err := p.service.GetOwnPosts(ctx, ownUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func (p *PostController) GetTimelinePosts(ctx *gin.Context) {
	ownUid, err := checkUid(ctx)
	if err != nil {
		return
	}

	posts, err := p.service.GetTimelinePosts(ctx, ownUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, posts)
}
