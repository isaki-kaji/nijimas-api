package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/isaki-kaji/nijimas-api/apperror"
)

func AuthMiddleware(authClient *auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")
		if authorizationHeader == "" {
			err := apperror.Unauthorized.Wrap(ErrAuthorizationHeaderRequired, "Authorization header is required")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apperror.ErrorResponse(ctx, err))
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authorizationHeader, bearerPrefix) {
			err := apperror.Unauthorized.Wrap(ErrBearerTokenRequired, "Authorization header must be a bearer token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apperror.ErrorResponse(ctx, err))
			return
		}

		idToken := strings.TrimPrefix(authorizationHeader, bearerPrefix)
		decodedToken, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			err = apperror.Unauthorized.Wrap(err, "failed to verify ID token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apperror.ErrorResponse(ctx, err))
			return
		}

		ctx.Set("ownUid", decodedToken.UID)
	}
}
