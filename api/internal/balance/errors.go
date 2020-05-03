package balance

import "errors"

var (
	ErrInvalidDataFormat   = errors.New("invalid data format")
	ErrEmptyTransaction    = errors.New("empty transaction")
	ErrNotFound            = errors.New("not found")
	ErrTransactionNotFound = errors.New("transaction not found")
)
