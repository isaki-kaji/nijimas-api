package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/apperror"
	"github.com/isaki-kaji/nijimas-api/application/service"
)

type PostSearchController struct {
	service service.PostSearchService
}

func NewPostSearchController(service service.PostSearchService) *PostSearchController {
	return &PostSearchController{service: service}
}

func (p *PostSearchController) GetPostsByQuery(ctx *gin.Context) {
	ownUid, err := checkUid(ctx)
	if err != nil {
		return
	}

	uid := ctx.Query("uid")
	mainCategory := ctx.Query("main-category")
	subCategory := ctx.Query("sub-category")

	if uid != "" && mainCategory != "" && subCategory != "" {
		return
	}

	if uid != "" {
		posts, err := p.service.GetPostsByUid(ctx, ownUid, uid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, posts)
		return
	}

	if mainCategory != "" {
		posts, err := p.service.GetPostsByMainCategory(ctx, ownUid, mainCategory)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, posts)
		return
	}

	if subCategory != "" {
		posts, err := p.service.GetPostsBySubCategory(ctx, ownUid, subCategory)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, posts)
		return
	}
}
