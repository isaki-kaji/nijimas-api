package middleware

import "errors"

var ErrAuthorizationHeaderRequired = errors.New("authorization header is required")
var ErrBearerTokenRequired = errors.New("authorization header must be a bearer token")
