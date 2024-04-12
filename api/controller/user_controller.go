package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/domain"
)

type UserController struct {
	service domain.UserService
}

func NewUserController(service domain.UserService) *UserController {
	return &UserController{service: service}
}

func (u *UserController) Create(ctx *gin.Context) {
	var req domain.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		fmt.Print(err)
		return
	}
	user, err := u.service.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		fmt.Print(err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
