package telegram

import (
	"errors"
)

var (
	ErrTokenInvalid     = errors.New("token is invalid")
	ErrTokenExpired     = errors.New("token was expired")
	ErrUserNotFound     = errors.New("user not found")
	ErrAlreadyCheckedIn = errors.New("user already checked in")
	ErrNotCheckedIn     = errors.New("user have not checked in yet")
	ErrDatabase         = errors.New("database error")
)
