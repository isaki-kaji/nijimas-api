package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/apperror"
	"github.com/isaki-kaji/nijimas-api/util"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId, err := util.GenerateSnowflakeID()
		if err != nil {
			err = apperror.OtherInternalErr.Wrap(err, "failed to generate trace id")
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, apperror.ErrorResponse(ctx, err))
			return
		}
		ctx.Set("traceId", traceId)
		slog.Info("Request", "trace_id", traceId, "path", ctx.Request.URL.Path, "method", ctx.Request.Method)

		ctx.Next()

		statusCode := ctx.Writer.Status()
		slog.Info("Response", "trace_id", traceId, "status_code", statusCode)

	}
}
