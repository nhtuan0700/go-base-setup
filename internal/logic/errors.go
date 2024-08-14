package logic

import "errors"

var (
	ErrNotFound       = errors.New("resource not found")
	ErrUnauthorized   = errors.New("unauthorized access")
	ErrInternal       = errors.New("internal server error")
	ErrDuplicateEmail = errors.New("email was taken")
)
