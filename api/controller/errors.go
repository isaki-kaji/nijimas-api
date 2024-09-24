package controller

import "errors"

var ErrUidNotFound = errors.New("uid not found")
var ErrUidTypeAssertionFailed = errors.New("uid type assertion failed")
