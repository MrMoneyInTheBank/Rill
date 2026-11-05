package domain

import "errors"

var (
	ErrInvalidTradeID       = errors.New("Invalid Trade ID")
	ErrInvalidTradePrice    = errors.New("Invalid Trade Price")
	ErrInvalidTradeTime     = errors.New("Invalid Trade Time")
	ErrInvalidTradeRecvTime = errors.New("Invalid TradeRecv Time")
	ErrInvalidTradeSize     = errors.New("Invalid Trade Size")
	ErrTradeSnapshot        = errors.New("Reject normalizing market_trades snapshot event")
	ErrInvalidL2Side        = errors.New("Invalid level2 side")
	ErrInvalidL2EventType   = errors.New("Invalid level2 event type")
)
