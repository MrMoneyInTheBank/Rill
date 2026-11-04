package config

import "errors"

var (
	ErrUnknownFlag = errors.New("Error parsing unknown flag")
	ErrBadFlags    = errors.New("Could not parse input flags")
)
