package controller

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
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

func (u *UserController) CreateUser(ctx *gin.Context) {
	var req service.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		fmt.Print(err)
		return
	}

	myUid, exists := ctx.Get("myUid")
	if !exists {
		slog.Warn("own uid is required")
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("own uid is required")))
		return
	}

	if req.Uid != myUid.(string) {
		slog.Warn("uid in request body must be the same as the uid in the token")
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("uid in request body must be the same as the uid in the token")))
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

func (u *UserController) GetUserById(ctx *gin.Context) {
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

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var req db.UpdateUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		fmt.Print(err)
		return
	}

	myUid, exists := ctx.Get("myUid")
	if !exists {
		slog.Warn("own uid is required")
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("own uid is required")))
		return
	}

	if req.Uid != myUid.(string) {
		slog.Warn("uid in request body must be the same as the uid in the token")
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("uid in request body must be the same as the uid in the token")))
		return
	}

	user, err := u.service.UpdateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		fmt.Print(err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
