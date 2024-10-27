package service

import "errors"

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrFollowAlreadyExists = errors.New("follow already exists")
var ErrFollowRequestAlreadyExists = errors.New("follow request already exists")
