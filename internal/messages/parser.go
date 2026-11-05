package messages

import (
	"encoding/json"

	"github.com/MrMoneyInTheBank/rill/internal/channels"
)

func ParseMsg(data []byte) (Msg, error) {
	var baseMsg BaseMsg

	if err := json.Unmarshal(data, &baseMsg); err != nil {
		return nil, err
	}

	switch baseMsg.Channel {
	case channels.SubscriptionsChannel:
		var subMsg SubscriptionMsg
		if err := json.Unmarshal(data, &subMsg); err != nil {
			return nil, err
		}
		return subMsg, nil
	case channels.HeartbeatsChannel:
		var hbMsg HeartbeatsMsg
		if err := json.Unmarshal(data, &hbMsg); err != nil {
			return nil, err
		}
		return hbMsg, nil
	case channels.Level2DataChannel:
		var l2Msg Level2Msg
		if err := json.Unmarshal(data, &l2Msg); err != nil {
			return nil, err
		}
		for _, event := range l2Msg.Events {
			if event.Type != L2SnapshotEvent && event.Type != L2UpdateEvent {
				return nil, ErrInvalidL2EventType
			}
			for _, update := range event.Updates {
				if update.Side != L2BidSide && update.Side != L2OfferSide {
					return nil, ErrInvalidL2Side
				}
			}
		}
		return l2Msg, nil
	case channels.MarketTradesChannel:
		var tradesMsg TradesMsg
		if err := json.Unmarshal(data, &tradesMsg); err != nil {
			return nil, err
		}
		for _, event := range tradesMsg.Events {
			if event.Type != TradeSnapshotEvent && event.Type != TradeUpdateEvent {
				return nil, ErrInvalidTradeEventType
			}
			for _, trade := range event.Trades {
				if trade.MakerSide != TradeBuySide && trade.MakerSide != TradeSellSide {
					return nil, ErrInvalidTradeSide
				}
			}
		}
		return tradesMsg, nil
	default:
		return nil, ErrUnsupportedChannel
	}
}
