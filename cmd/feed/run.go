package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/coder/websocket"

	"github.com/MrMoneyInTheBank/rill/internal/messages"
	"github.com/MrMoneyInTheBank/rill/internal/stream"
)

const (
	websocketEndpoint string = "wss://advanced-trade-ws.coinbase.com"
	BTCUSD            string = "BTC-USD"
)

func run() error {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer stop()

	subscriptionReqs, err := messages.GetSubReqs(BTCUSD)
	if err != nil {
		return errors.Join(ErrBadSubReqs, err)
	}
	conn, err := stream.Connect(ctx, websocketEndpoint, nil)
	if err != nil {
		return errors.Join(ErrBadConnection, err)
	}
	log.Println("Established websocket connection.")
	defer conn.Close(websocket.StatusNormalClosure, "Closing connection")

	subCtx, subCtxCancel := context.WithTimeout(ctx, time.Second*5)
	defer subCtxCancel()

	err = stream.SendMsgs(subCtx, conn, subscriptionReqs)
	if err != nil {
		return errors.Join(ErrBadWrite, err)
	}
	subCtxCancel()

	err = readLoop(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}
