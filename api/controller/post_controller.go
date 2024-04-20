package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/domain"
)

type PostController struct {
	service domain.PostService
}

func NewPostController(service domain.PostService) *PostController {
	return &PostController{service: service}
}

func (p *PostController) Create(ctx *gin.Context) {
	var req domain.CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		slog.Warn("failed to bind json to CreatePostRequest")
		return
	}
	post, err := p.service.CreatePost(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		slog.Warn("failed to create post because of internal server error")
		return
	}
	ctx.JSON(http.StatusCreated, post)
}
