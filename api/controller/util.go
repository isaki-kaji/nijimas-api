package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/isaki-kaji/nijimas-api/apperror"
)

func checkPostReq(ctx *gin.Context, req any) (string, error) {
	if err := ctx.ShouldBindJSON(req); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			err = apperror.WrapValidationErr(validationErrs)
		} else {
			err = apperror.ReqBodyDecodeFailed.Wrap(err, "failed to decode request body")
		}
		ctx.JSON(http.StatusBadRequest, apperror.ErrorResponse(err))
		return "", err
	}

	uidStr, err := checkUid(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, apperror.ErrorResponse(err))
		return "", err
	}

	return uidStr, nil
}

func checkUid(ctx *gin.Context) (string, error) {
	ownUid, exists := ctx.Get("ownUid")
	if !exists {
		err := apperror.Unauthorized.Wrap(ErrUidNotFound, "own uid is required")
		return "", err
	}

	uidStr, ok := ownUid.(string)
	if !ok {
		err := apperror.Unauthorized.Wrap(ErrUidTypeAssertionFailed, "failed to assert uid type")
		return "", err
	}
	return uidStr, nil
}
