package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/apperror"
	"github.com/isaki-kaji/nijimas-api/service"
)

type FavoriteController struct {
	service service.FavoriteService
}

func NewFavoriteController(service service.FavoriteService) *FavoriteController {
	return &FavoriteController{service: service}
}

func (f *FavoriteController) ToggleFavorite(ctx *gin.Context) {
	var req service.ToggleFavoriteParams
	ownUid, err := checkPostReq(ctx, &req)
	if err != nil {
		return
	}
	req.Uid = ownUid

	favorite, flag, err := f.service.ToggleFavorite(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
		return
	}

	if flag == service.FlagCreated {
		ctx.JSON(http.StatusCreated, favorite)
		return
	}

	ctx.JSON(http.StatusNoContent, favorite)
}
