package main

import (
	"context"
	"errors"
	"log"

	"github.com/coder/websocket"

	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

func readLoop(ctx context.Context, conn *websocket.Conn, msgCh chan messages.Msg, msgChCancel context.CancelFunc) error {
	for {
		msgType, data, err := conn.Read(ctx)

		if errors.Is(ctx.Err(), context.Canceled) {
			log.Println("Shutdown requested")
			msgChCancel()
			return nil
		}

		if err != nil {
			log.Printf("Error reading WebSocket message: %v\n", err)
			return errors.Join(ErrReadingMessage, err)
		}

		if msgType != websocket.MessageText {
			log.Printf("Warning: received non-text message frame type (%v). Skipping payload.\n", msgType)
		}

		msg, err := messages.ParseMsg(data)
		if err != nil {
			return errors.Join(ErrParsingMessage, err)
		}

		log.Printf("RECEIVED %#v", msg)
		msgCh <- msg

	}
}
