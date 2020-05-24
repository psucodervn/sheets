package user

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrAuthNotFound = errors.New("auth user not found")
)
