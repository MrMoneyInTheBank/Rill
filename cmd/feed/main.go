package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"

	"github.com/coder/websocket"

	"github.com/MrMoneyInTheBank/rill/internal/channels"
	"github.com/MrMoneyInTheBank/rill/internal/messages"
	"github.com/MrMoneyInTheBank/rill/internal/stream"
)

const endpoint string = "wss://advanced-trade-ws.coinbase.com"

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer stop()

	conn, err := stream.Connect(ctx, endpoint, nil)
	if err != nil {
		log.Fatalf("Could not establish websocket connection: %v\n", err)
	}
	log.Println("Established websocket connection.")
	defer conn.Close(websocket.StatusNormalClosure, "Closing connection")

	jwtToken := os.Getenv("JWT_TOKEN")

	subscribeReq := messages.SubscribeReq{
		Type:    "subscribe",
		Channel: channels.HeartbeatsChannel,
		JWT:     jwtToken,
	}

	err = stream.SendMsg(ctx, conn, subscribeReq)
	if err != nil {
		log.Fatalf("Error sending subscription message: %v\n", err)
	}
	log.Printf("Request subscription to %v channels for product %v\n", subscribeReq.Channel, subscribeReq.ProductIDs)

	for {
		msgType, data, err := conn.Read(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(ctx.Err(), context.Canceled) {
				log.Println("Shutdown requested")
				return
			}
			log.Fatalf("Error reading WebSocket message: %v\n", err)
		}

		if msgType != websocket.MessageText {
			log.Printf("Warning: received non-text message frame type (%v). Skipping payload.\n", msgType)
			continue
		}

		msg, err := messages.ParseMsg(data)
		if err != nil {
			log.Printf("Error parsing WebSocket message: %v\n", err)
		}

		switch m := msg.(type) {
		case messages.SubscriptionMsg:
			log.Printf("%#v", m)
		case messages.HeartbeatsMsg:
			log.Printf("%#v", m)
		case messages.TradesMsg:
			log.Printf("%#v", m)
		case messages.Level2Msg:
			log.Printf("%#v", m)
		}

	}
}
