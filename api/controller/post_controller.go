package controller

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/service"
)

type PostController struct {
	service service.PostService
}

func NewPostController(service service.PostService) *PostController {
	return &PostController{service: service}
}

func (p *PostController) CreatePost(ctx *gin.Context) {
	var req service.CreatePostRequest
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

func (p *PostController) GetPostsByQuery(ctx *gin.Context) {
	uid := ctx.Query("uid")
	mainCategory := ctx.Query("main-category")
	myUid, exists := ctx.Get("myUid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("own uid is required")))
	}

	if uid != "" {
		param := db.GetPostsByUidParams{
			Uid:   myUid.(string),
			Uid_2: uid,
		}
		posts, err := p.service.GetPostsByUid(ctx, param)
		if err != nil {
			slog.Warn("failed to get posts because of internal server error")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, posts)
		return
	}

	if mainCategory != "" {
		param := db.GetPostsByMainCategoryParams{
			MainCategory: mainCategory,
			Uid:          myUid.(string),
		}

		posts, err := p.service.GetPostsByMainCategory(ctx, param)
		if err != nil {
			slog.Warn("failed to get posts because of internal server error")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, posts)
		return
	}
}
