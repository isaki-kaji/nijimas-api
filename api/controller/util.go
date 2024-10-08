package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/isaki-kaji/nijimas-api/apperror"
)

func checkPostReq(ctx *gin.Context, req any) (string, error) {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			err = apperror.WrapValidationErr(validationErrs)
		} else {
			err = apperror.ReqBodyDecodeFailed.Wrap(err, "failed to decode request body")
		}
		ctx.JSON(http.StatusBadRequest, apperror.ErrorResponse(ctx, err))
		return "", err
	}

	uidStr, err := checkUid(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, apperror.ErrorResponse(ctx, err))
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

func getTimezone(ctx *gin.Context) string {
	timezone := ctx.GetHeader("Time-zone")
	if timezone == "" {
		timezone = "Asia/Tokyo"
	}
	return timezone
}
