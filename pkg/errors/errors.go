package customErrors

import "errors"

var (
	ErrParsingObjectID = errors.New("object id parsing error")
)