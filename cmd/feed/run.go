package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/coder/websocket"

	"github.com/MrMoneyInTheBank/rill/internal/config"
	"github.com/MrMoneyInTheBank/rill/internal/messages"
	"github.com/MrMoneyInTheBank/rill/internal/stream"
)

const (
	websocketEndpoint string = "wss://advanced-trade-ws.coinbase.com"
)

func run() error {
	parsedOpts, err := config.ParseOptions(os.Args[1:])
	if err != nil {
		if errors.Is(flag.ErrHelp, err) {
			return nil
		}
		return errors.Join(config.ErrBadFlags, err)
	}

	subscriptionReqs, err := messages.GetSubReqs(parsedOpts)
	if err != nil {
		return errors.Join(ErrBadSubReqs, err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer stop()

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
