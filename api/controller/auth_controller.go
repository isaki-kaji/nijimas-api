package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
	"github.com/isaki-kaji/nijimas-api/domain"
)

type AuthController struct {
	service domain.AuthService
}

func NewAuthController(service domain.AuthService) *AuthController {
	return &AuthController{service: service}
}

type createUserRequest struct {
	Uid      string `json:"uid" binding:"required"`
	Username string `json:"username" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

func (u *AuthController) Signup(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Uid:      req.Uid,
		Username: req.Username,
		Currency: req.Currency,
	}

	user, err := u.service.SignupUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
