package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/isaki-kaji/nijimas-api/db/sqlc"
)

type createUserRequest struct {
	Uid      string `json:"uid" binding:"required"`
	Username string `json:"username" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

func (server *Server) Signup(ctx *gin.Context) {
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

	user, err := server.service.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
