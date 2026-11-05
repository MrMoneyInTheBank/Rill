package messages

import (
	"encoding/json"
	"fmt"

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
				return nil, fmt.Errorf("invalid level2 event type %q", event.Type)
			}
			for _, update := range event.Updates {
				if update.Side != L2BidSide && update.Side != L2OfferSide {
					return nil, fmt.Errorf("invalid level2 side %q", update.Side)
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
				return nil, fmt.Errorf("invalid trade event type %q", event.Type)
			}
			for _, trade := range event.Trades {
				if trade.MakerSide != TradeBuySide && trade.MakerSide != TradeSellSide {
					return nil, fmt.Errorf("invalid trade side %q", trade.MakerSide)
				}
			}
		}
		return tradesMsg, nil
	default:
		return nil, fmt.Errorf("Unsupported channel for parsing: %v", baseMsg)
	}
}
