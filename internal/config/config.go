package config

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type Opts struct {
	All          bool
	Heartbeats   bool
	Level2       bool
	MarketTrades bool
	Product      string
}

func ParseOptions(args []string) (Opts, error) {
	var parsed Opts

	flags := flag.NewFlagSet("rill", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	flags.BoolVar(&parsed.All, "all", false, "Subscribe to all available channels")
	flags.BoolVar(&parsed.Heartbeats, "h", true, "Subscribe to heartbeats channel")
	flags.BoolVar(&parsed.Level2, "l", false, "Subscribe to level 2 channel")
	flags.BoolVar(&parsed.MarketTrades, "m", false, "Subscribe to market trades channel")
	flags.StringVar(&parsed.Product, "product", "BTC-USD", "Specify a product")

	if err := flags.Parse(args); err != nil {
		if errors.Is(flag.ErrHelp, err) {
			flags.SetOutput(os.Stderr)
			flags.Usage()
		}
		return parsed, err
	}

	if parsed.All {
		parsed.Heartbeats = true
		parsed.Level2 = true
		parsed.MarketTrades = true
	}

	if flags.NArg() != 0 {
		return parsed, errors.Join(ErrUnknownFlag, fmt.Errorf("unexpected argument %q", flags.Arg(0)))
	}

	return parsed, nil
}

func ValidateOptions(opts Opts) error {
	return nil // will handle validation logic later
}
