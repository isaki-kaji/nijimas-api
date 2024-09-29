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

// func (p *PostController) GetPostsByQuery(ctx *gin.Context) {
// 	uid := ctx.Query("uid")
// 	mainCategory := ctx.Query("main-category")
// 	myUid, exists := ctx.Get("myUid")
// 	if !exists {
// 		ctx.JSON(http.StatusUnauthorized, apperror.ErrorResponse(errors.New("own uid is required")))
// 	}

// 	if uid != "" {
// 		param := db.GetPostsByUidParams{
// 			Uid:   myUid.(string),
// 			Uid_2: uid,
// 		}
// 		posts, err := p.service.GetPostsByUid(ctx, param)
// 		if err != nil {
// 			slog.Warn("failed to get posts because of internal server error")
// 			ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, posts)
// 		return
// 	}

// 	if mainCategory != "" {
// 		param := db.GetPostsByMainCategoryParams{
// 			MainCategory: mainCategory,
// 			Uid:          myUid.(string),
// 		}

// 		posts, err := p.service.GetPostsByMainCategory(ctx, param)
// 		if err != nil {
// 			slog.Warn("failed to get posts because of internal server error")
// 			ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, posts)
// 		return
// 	}
// }
