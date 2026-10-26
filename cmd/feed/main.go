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

const endpoint string = "wss://ws-feed.exchange.coinbase.com"

func main() {
	logfile, err := os.OpenFile("feed.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatalf("Could not prepare log file: %v", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer stop()

	conn, err := stream.Connect(ctx, endpoint, nil)
	if err != nil {
		log.Printf("Could not establish websocket connection: %v\n", err)
		return
	}
	log.Println("Established websocket connection.")
	defer conn.Close(websocket.StatusNormalClosure, "Closing connection")

	subscribeMsg := messages.SubscribeMsg{
		Type:       "subscribe",
		ProductIDs: []string{"BTC-USD"},
		Channels: []channels.Channel{
			channels.HeartbeatChannel,
		},
	}

	err = stream.SendMsg(ctx, conn, subscribeMsg)
	if err != nil {
		log.Printf("Error sending subscription message: %v\n", err)
		return
	}
	log.Printf("Subscribed to %v channels for product %v\n", subscribeMsg.Channels, subscribeMsg.ProductIDs)

	for {
		var response map[string]any

		if err := stream.ReadMsg(ctx, conn, &response); err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(ctx.Err(), context.Canceled) {
				log.Println("Shutdown requested")
				return
			}

			log.Printf("Error reading WebSocket message: %v", err)
			return
		}

		log.Println(messages.ParseMsg(response))
	}
}
