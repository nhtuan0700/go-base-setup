package logic

import "errors"

var (
	ErrNotFound     = errors.New("Resource not found")
	ErrUnauthorized = errors.New("Unauthorized access")
	ErrInternal     = errors.New("Internal server error")
)
