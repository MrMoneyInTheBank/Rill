package main

import (
	"context"
	"log"
	"os"

	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

const (
	websocketEndpoint string = "wss://advanced-trade-ws.coinbase.com"
)

func NewLogger(logfileName string) (*log.Logger, error) {
	logfile, err := os.OpenFile(logfileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, err
	}

	logger := log.New(logfile, log.Default().Prefix(), log.Default().Flags())
	return logger, nil
}

func run() error {
	msgCh := make(chan messages.Msg)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	normLogger, err := NewLogger("norm.log")
	if err != nil {
		panic(err)
	}

	go norm_handler(ctx, msgCh, normLogger)

	if err := message_handler(websocketEndpoint, msgCh, cancel); err != nil {
		return err
	}
	return nil
}
