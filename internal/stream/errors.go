package stream

import "errors"

var ErrInvalidReadSize = errors.New("ReadSize must point to -1 or a positive integer")
