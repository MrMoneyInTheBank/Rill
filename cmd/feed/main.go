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

	jwtToken, exists := os.LookupEnv("JWT_TOKEN")
	if !exists {
		log.Printf("JWT not provided. Can only connect to public channels")
	} else if jwtToken == "" {
		log.Printf("JWT unset. Can only connect to public channels")
	} else {
		log.Printf("JWT token set; Can attempt to connect to all channels")
	}

	var readLimit int64 = -1
	conn, err := stream.Connect(ctx, endpoint, &readLimit)
	if err != nil {
		log.Fatalf("Could not establish websocket connection: %v\n", err)
	}
	log.Println("Established websocket connection.")
	defer conn.Close(websocket.StatusNormalClosure, "Closing connection")

	heartbeatReq := messages.SubscribeReq{
		Type:    "subscribe",
		Channel: channels.HeartbeatsChannel,
		JWT:     jwtToken,
	}

	level2Req := messages.SubscribeReq{
		Type:       "subscribe",
		ProductIDs: []string{"BTC-USD"},
		Channel:    channels.Level2Channel,
	}

	marketTradesReq := messages.SubscribeReq{
		Type:       "subscribe",
		ProductIDs: []string{"BTC-USD"},
		Channel:    channels.MarketTradesChannel,
	}

	subscriptionReqs := []messages.SubscribeReq{
		heartbeatReq,
		level2Req,
		marketTradesReq,
	}

	for _, req := range subscriptionReqs {
		err = stream.SendMsg(ctx, conn, req)
		if err != nil {
			log.Fatalf("Error sending subscription message: %v\n", err)
		}
		log.Printf("Request subscription to %v channels for product %v\n", req.Channel, req.ProductIDs)
	}

	payloadSizes := make(map[channels.Channel][]int)
	for {
		msgType, data, err := conn.Read(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(ctx.Err(), context.Canceled) {
				log.Println("Shutdown requested")
				break
			}
			log.Fatalf("Error reading WebSocket message: %v\n", err)
		}

		if msgType != websocket.MessageText {
			log.Printf("Warning: received non-text message frame type (%v). Skipping payload.\n", msgType)
			continue
		}

		msg, err := messages.ParseMsg(data)
		if err != nil {
			log.Printf("Error parsing WebSocket message: %v\nError: %v", string(data), err)
			continue
		}

		payloadSizes[msg.GetChannel()] = append(payloadSizes[msg.GetChannel()], len(data))
		log.Printf("%#v", msg)

	}
	log.Printf("Payload sizes: %v", payloadSizes)
}
