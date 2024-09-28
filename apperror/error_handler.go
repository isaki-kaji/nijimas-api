package apperror

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func ErrorResponse(ctx *gin.Context, err error) gin.H {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = &AppError{
			ErrCode: UnKnown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceId, _ := ctx.Get("traceId")
	slog.Error(fmt.Sprintf("\"%s\"", appErr.Error()),
		"traceId", traceId,
		"error_code", appErr.ErrCode,
		"message", appErr.Message,
		"errors", appErr.Errors)

	return gin.H{
		"code":    appErr.ErrCode,
		"message": appErr.Message,
		"errors":  appErr.Errors,
	}
}
