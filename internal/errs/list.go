package errs

import "errors"

var (
	ErrNotfound                    = errors.New("not found")
	ErrUserNotFound                = errors.New("user not found")
	ErrInvalidUserID               = errors.New("invalid user id")
	ErrInvalidRequestBody          = errors.New("invalid request body")
	ErrInvalidFieldValuse          = errors.New("invalid field value")
	ErrInvalidUserName             = errors.New("invalid user name, min 4 symbols")
	ErrUsernameAlreadyExists       = errors.New("username already exists")
	ErrEmployeeNotFound            = errors.New("employee not found")
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password")
	ErrInvalidToken                = errors.New("invalid token")
	ErrSomethingWentWrong          = errors.New("something went wrong")
)
