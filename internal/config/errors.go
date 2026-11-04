package config

import "errors"

var (
	ErrUnknownFlag    = errors.New("Error parsing unknown flag")
	ErrBadFlags       = errors.New("Could not parse input flags")
	ErrMissingProduct = errors.New("Need a product to subscribe to level2 or market_trades channel")
)
