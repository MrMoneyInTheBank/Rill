package messages

import "errors"

var (
	ErrInvalidReq      = errors.New("Invalid subscription request")
	ErrTooManyProducts = errors.New("Too many product ids")
)
