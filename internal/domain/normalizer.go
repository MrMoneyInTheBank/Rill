package domain

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

func NormalizeHeartbeats(msg messages.HeartbeatsMsg) (Heartbeat, error) {
	if len(msg.Events) == 0 {
		return Heartbeat{}, nil
	}
	hb := Heartbeat{
		Counter: msg.Events[0].HeartbeatCounter,
	}

	return hb, nil
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
	if len(msg.Events) == 0 {
		return []Trade{}, nil
	}
	if msg.Events[0].Type == messages.TradeSnapshotEvent {
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

func NormalizeOrderUpdate(update messages.L2Update) (Order, error) {
	var side OrderSide

	switch update.Side {
	case messages.L2BidSide:
		side = BID
	case messages.L2OfferSide:
		side = OFFER
	default:
		return Order{}, fmt.Errorf("invalid level2 side %q", update.Side)
	}

	return Order{
		OrderTime:  update.EventTime,
		Side:       side,
		PriceLevel: update.PriceLevel,
		NewQty:     update.NewQty,
	}, nil
}

func NormalizeL2(msg messages.Level2Msg) ([]BookEvent, error) {
	res := make([]BookEvent, 0, len(msg.Events))
	for _, event := range msg.Events {
		var eventType OrderType
		switch event.Type {
		case messages.L2SnapshotEvent:
			eventType = SNAPSHOT
		case messages.L2UpdateEvent:
			eventType = UPDATE
		default:
			return nil, fmt.Errorf("invalid level2 event type %q", event.Type)
		}

		bookEvent := BookEvent{
			Type:          eventType,
			ProductID:     event.ProductID,
			OrderRecvTime: msg.Timestamp,
			Levels:        make([]Order, 0, len(event.Updates)),
		}
		for _, update := range event.Updates {
			order, err := NormalizeOrderUpdate(update)
			if err != nil {
				return nil, err
			}
			bookEvent.Levels = append(bookEvent.Levels, order)
		}
		res = append(res, bookEvent)
	}

	return res, nil
}

func Normalizer(msg messages.Msg) (any, error) {
	switch m := msg.(type) {
	case messages.HeartbeatsMsg:
		return NormalizeHeartbeats(m)
	case messages.TradesMsg:
		return NormalizeTrades(m)
	case messages.Level2Msg:
		return NormalizeL2(m)
	}

	return nil, nil
}
