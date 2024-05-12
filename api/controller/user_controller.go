package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/service"
	"github.com/isaki-kaji/nijimas-api/util"
	"github.com/jackc/pgx/v5"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (u *UserController) Create(ctx *gin.Context) {
	var req service.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		fmt.Print(err)
		return
	}
	user, err := u.service.CreateUser(ctx, req)
	if err != nil {
		if err.Error() == util.UserAlreadyExists {
			ctx.JSON(http.StatusConflict, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		fmt.Print(err)
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (u *UserController) Get(ctx *gin.Context) {
	uid := ctx.Param("id")
	user, err := u.service.GetUser(ctx, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			fmt.Print(err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		fmt.Print(err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
