package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/apperror"
	"github.com/isaki-kaji/nijimas-api/application/service"
)

type FollowController struct {
	service service.FollowService
}

func NewFollowController(service service.FollowService) *FollowController {
	return &FollowController{service: service}
}

func (f *FollowController) ToggleFollow(ctx *gin.Context) {
	var req service.ToggleFollowParams
	ownUid, err := checkPostReq(ctx, &req)
	if err != nil {
		return
	}
	req.Uid = ownUid

	follow, flag, err := f.service.ToggleFollow(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
		return
	}

	if flag == service.FlagCreated {
		ctx.JSON(http.StatusCreated, follow)
		return
	}

	ctx.JSON(http.StatusNoContent, follow)
}
