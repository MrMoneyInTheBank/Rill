package main

import (
	"context"
	"log"

	"github.com/MrMoneyInTheBank/rill/internal/domain"
	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

func norm_handler(ctx context.Context, msgCh chan messages.Msg, logger *log.Logger) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-msgCh:
			if !ok {
				return
			}
			normalizedMsg, err := domain.Normalizer(msg)
			if err != nil {
				logger.Println("Error normalizing message")
			}
			logger.Printf("NORM %#v", normalizedMsg)
		}
	}
}
