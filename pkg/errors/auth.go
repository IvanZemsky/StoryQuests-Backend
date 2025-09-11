package commonErrors

var (
	ErrLoginUserNotFound  = NewError("USER_NOT_FOUND", "user not found")
	ErrMismatchedPassword = NewError("WRONG_PASSWORD", "wrong password")
	ErrUserAlreadyExists  = NewError("USER_ALREADY_EXISTS", "user already exists")
)
