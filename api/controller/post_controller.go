package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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
		// バリデーションエラーの詳細をログに記録
		//型アサーションを使って、エラーがバリデーションエラーかどうかを判定
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, vErr := range validationErrs {
				// 各フィールドのエラーをログに出力
				slog.Warn("Validation error on field '%s': %s", vErr.Field(), vErr.ActualTag())
			}
		}
		// バリデーションエラーの詳細をクライアントに送信
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := p.service.CreatePost(ctx, req)
	if err != nil {
		slog.Warn("failed to create post because of internal server error")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, post)
}
