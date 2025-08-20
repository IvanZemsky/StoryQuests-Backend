package customErrors

import "errors"

var (
	ErrLoginUserNotFound = errors.New("user with this login not found")
	ErrUserAlreadyExists = errors.New("user with this login already exists")
	ErrMismatchedPassword = errors.New("invalid password")
)