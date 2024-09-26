package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/apperror"
	"github.com/isaki-kaji/nijimas-api/service"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var req service.CreateUserRequest
	ownUid, err := checkPostReq(ctx, &req)
	if err != nil {
		return
	}
	req.Uid = ownUid

	user, err := u.service.CreateUser(ctx, req)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			ctx.JSON(http.StatusConflict, apperror.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (u *UserController) GetUserByUid(ctx *gin.Context) {
	uid := ctx.Param("id")
	user, err := u.service.GetUserByUid(ctx, uid)
	if err != nil {
		if apperror.DataNotFound.Equal(err) {
			ctx.JSON(http.StatusNotFound, apperror.ErrorResponse(err))
			fmt.Print(err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(err))
		fmt.Print(err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var req service.UpdateUserParams
	ownUid, err := checkPostReq(ctx, &req)
	if err != nil {
		return
	}
	req.Uid = ownUid

	user, err := u.service.UpdateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apperror.ErrorResponse(err))
		fmt.Print(err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
