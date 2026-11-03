package main

import "errors"

var (
	ErrBadSubReqs         = errors.New("Error constructing subscription requests")
	ErrBadConnection      = errors.New("Error establishing websocket connection")
	ErrBadWrite           = errors.New("Error writing to websocket connection")
	ErrInvalidMessageType = errors.New("Invalid Message Type")
	ErrReadingMessage     = errors.New("Error reading WebSocket message")
	ErrParsingMessage     = errors.New("Error parsing WebSocket message")
)
