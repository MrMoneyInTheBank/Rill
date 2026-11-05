package messages

import "errors"

var (
	ErrInvalidReq            = errors.New("Invalid subscription request")
	ErrTooManyProducts       = errors.New("Too many product ids")
	ErrInvalidL2EventType    = errors.New("Invalid level2 event type")
	ErrInvalidL2Side         = errors.New("Invalid level2 side")
	ErrInvalidTradeEventType = errors.New("Invalid trade event type")
	ErrInvalidTradeSide      = errors.New("Invalid trade side")
	ErrUnsupportedChannel    = errors.New("Unsupported channel for parsing")
)
