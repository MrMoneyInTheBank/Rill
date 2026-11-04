package domain

import (
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

func NormalizeTrade(t messages.Trade, recvTime time.Time) (Trade, error) {
	trade := Trade{
		TradeID:       t.TradeID,
		ProductID:     t.ProductID,
		TradeTime:     t.Time,
		TradeRecvTime: recvTime,
		Price:         t.Price,
		Size:          t.Size,
	}

	if trade.TradeID == "" {
		return Trade{}, ErrInvalidTradeID
	}

	if trade.TradeTime.IsZero() {
		return Trade{}, ErrInvalidTradeTime
	}

	if trade.TradeRecvTime.IsZero() {
		return Trade{}, ErrInvalidTradeRecvTime
	}

	if trade.Price.LessThanOrEqual(decimal.Zero) {
		return Trade{}, ErrInvalidTradePrice
	}

	if trade.Size.LessThanOrEqual(decimal.Zero) {
		return Trade{}, ErrInvalidTradeSize
	}

	return trade, nil
}

func NormalizeTrades(msg messages.TradesMsg) ([]Trade, error) {
	if msg.Events[0].Type == string(messages.TradeSnapshotEvent) {
		return nil, ErrTradeSnapshot
	}

	cap := func() int {
		maxTrades := 0
		for _, event := range msg.Events {
			maxTrades = max(maxTrades, len(event.Trades))
		}
		return len(msg.Events) * maxTrades
	}()

	res := make([]Trade, 0, cap)
	msgTime := msg.Timestamp

	for _, event := range msg.Events {
		for _, t := range event.Trades {
			trade, err := NormalizeTrade(t, msgTime)
			if err != nil {
				return res, err
			}
			res = append(res, trade)

		}
	}

	return res, nil
}
