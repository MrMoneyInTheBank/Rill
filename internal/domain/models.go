package domain

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"

	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

type Trade struct {
	TradeID       string
	ProductID     string
	TradeTime     time.Time
	TradeRecvTime time.Time
	Price         decimal.Decimal
	Size          decimal.Decimal
}

func NormalizeTrades(msg messages.TradesMsg) ([]Trade, error) {
	if msg.Events[0].Type == string(messages.TradeSnapshotEvent) {
		return nil, errors.New("Reject normalizing market_trades snapshot event")
	}

	res := make([]Trade, len(msg.Events[0].Trades), len(msg.Events[0].Trades))
	msgTime := msg.Timestamp

	for idx, t := range msg.Events[0].Trades {

		trade := Trade{
			TradeID:       t.TradeID,
			ProductID:     t.ProductID,
			TradeTime:     t.Time,
			TradeRecvTime: msgTime,
			Price:         t.Price,
			Size:          t.Size,
		}

		res[idx] = trade
	}

	return res, nil
}
