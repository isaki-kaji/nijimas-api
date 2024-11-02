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
	var req service.FollowRequestParams
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
	var req service.FollowRequestParams
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

func (c *FollowRequestController) HandleFollowRequest(ctx *gin.Context) {
	status := ctx.Query("status")
	var req service.FollowRequestParams
	ownUid, err := checkPostReq(ctx, &req)
	if err != nil {
		return
	}
	req.Uid = ownUid

	if status == "accept" {
		f, err := c.service.AcceptFollowRequest(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, f)
		return
	}

	if status == "reject" {
		fr, err := c.service.RejectFollowRequest(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, fr)
		return
	}

	err = apperror.BadQueryParam.Wrap(ErrInvalidStatus, "invalid status")
	ctx.JSON(http.StatusBadRequest, apperror.ErrorResponse(ctx, err))
}
