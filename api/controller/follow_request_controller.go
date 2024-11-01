package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/apperror"
	"github.com/isaki-kaji/nijimas-api/application/service"
)

type FollowRequestController struct {
	service service.FollowRequestService
}

func NewFollowRequestController(service service.FollowRequestService) *FollowRequestController {
	return &FollowRequestController{service: service}
}

func (c *FollowRequestController) DoFollowRequest(ctx *gin.Context) {
	var req service.DoFollowRequestParams
	ownUid, err := checkPostReq(ctx, &req)
	if err != nil {
		return
	}
	req.Uid = ownUid

	fr, err := c.service.DoFollowRequest(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
		return
	}

	ctx.JSON(http.StatusOK, fr)
}

func (c *FollowRequestController) CancelFollowRequest(ctx *gin.Context) {
	var req service.CancelFollowRequestParams
	ownUid, err := checkPostReq(ctx, &req)
	if err != nil {
		return
	}
	req.Uid = ownUid

	fr, err := c.service.CancelFollowRequest(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
		return
	}

	ctx.JSON(http.StatusNoContent, fr)
}
