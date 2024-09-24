package apperror

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error) gin.H {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = &AppError{
			ErrCode: UnKnown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	return gin.H{
		"code":    appErr.ErrCode,
		"message": appErr.Message,
		"errors":  appErr.Errors,
	}
}
