package controller

import "errors"

var ErrUidNotFound = errors.New("uid not found")
var ErrUidParamNotFound = errors.New("uid param not found")
var ErrUidTypeAssertionFailed = errors.New("uid type assertion failed")
var ErrInvalidYear = errors.New("year must be between 2000 and 2100")
var ErrInvalidMonth = errors.New("month must be between 1 and 12")
