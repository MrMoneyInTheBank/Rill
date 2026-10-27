package main

import (
	"context"
	"errors"
	"fmt"
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

	jwtToken := os.Getenv("JWT_TOKEN")

	subscribeMsg := messages.SubscribeMsg{
		Type:    "subscribe",
		Channel: channels.HeartbeatsChannel,
		JWT:     jwtToken,
	}

	err = stream.SendMsg(ctx, conn, subscribeMsg)
	if err != nil {
		log.Printf("Error sending subscription message: %v\n", err)
		return
	}
	log.Printf("Request subscription to %v channels for product %v\n", subscribeMsg.Channel, subscribeMsg.ProductIDs)

	for {
		msgType, data, err := conn.Read(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(ctx.Err(), context.Canceled) {
				log.Println("Shutdown requested")
				return
			}
			log.Printf("Error reading WebSocket message: %v\n", err)
			return
		}

		if msgType != websocket.MessageText {
			log.Printf("Warning: received non-text message frame type (%v). Skipping payload.\n", msgType)
			continue
		}

		consoleMsg, jsonData, err := messages.ParseMsg(data)
		if err != nil {
			log.Printf("Error parsing WebSocket message: %v\n", err)
		}

		fmt.Println(consoleMsg)
		log.Println(jsonData)
	}
}
