package errors

import "errors"

var (
	ErrIndexOutOfBounds = errors.New("index out of bounds")
	ErrEmptyList        = errors.New("list is empty")
	ErrElementNotFound  = errors.New("element not found")
)
