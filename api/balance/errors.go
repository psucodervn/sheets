package balance

import "errors"

var (
	ErrInvalidDataFormat = errors.New("invalid data format")
	ErrEmptyTransaction  = errors.New("empty transaction")
)
