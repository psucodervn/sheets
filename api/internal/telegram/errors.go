package telegram

import (
	"errors"
)

var (
	ErrTokenInvalid = errors.New("token is invalid")
	ErrTokenExpired = errors.New("token was expired")
	ErrUserNotFound = errors.New("user not found")
)
