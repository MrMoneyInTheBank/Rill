package rest

import "errors"

var (
	ErrReqTimedOut  = errors.New("Request timed out")
	ErrReqFailed    = errors.New("Request failed")
	ErrBadReq       = errors.New("Could not construct request")
	ErrCantValidate = errors.New("Unable to validate product")
)
