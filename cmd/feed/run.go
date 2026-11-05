package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

const (
	websocketEndpoint string = "wss://advanced-trade-ws.coinbase.com"
)

func NewLogger(logfileName string) (*log.Logger, error, func()) {
	logfile, err := os.OpenFile(logfileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	teardown := func() {
		logfile.Close()
	}
	if err != nil {
		return nil, err, nil
	}

	logger := log.New(logfile, log.Default().Prefix(), log.Default().Flags())
	return logger, nil, teardown
}

func run() error {
	msgCh := make(chan messages.Msg)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	normLogger, err, teardown := NewLogger("norm.log")
	if err != nil {
		panic(err)
	}
	defer teardown()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		norm_handler(ctx, msgCh, normLogger)
	}()

	if err := message_handler(websocketEndpoint, msgCh, cancel); err != nil {
		return err
	}
	cancel()

	wg.Wait()
	return nil
}
