package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/service"
)

type FavoriteController struct {
	service service.FavoriteService
}

func NewFavoriteController(service service.FavoriteService) *FavoriteController {
	return &FavoriteController{service: service}
}

func (f *FavoriteController) ToggleFavorite(ctx *gin.Context) {
	var req db.GetFavoriteParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	favorite, flag, err := f.service.ToggleFavorite(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if flag == "created" {
		ctx.JSON(http.StatusCreated, favorite)
		return
	}

	ctx.JSON(http.StatusNoContent, favorite)
}
