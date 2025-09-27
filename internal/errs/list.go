package errs

import "errors"

var (
	ErrNotfound           = errors.New("not found")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidFieldValuse = errors.New("invalid field value")
)
