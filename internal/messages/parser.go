package messages

import (
	"encoding/json"
	"fmt"

	"github.com/MrMoneyInTheBank/rill/internal/channels"
)

func ParseMsg(data []byte) (Msg, error) {
	var baseMsg BaseMsg

	if err := json.Unmarshal(data, &baseMsg); err != nil {
		return nil, fmt.Errorf("Error parsing message: %v\n Error: %v", string(data), err)
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
		return l2Msg, nil
	case channels.MarketTradesChannel:
		var tradesMsg TradesMsg
		if err := json.Unmarshal(data, &tradesMsg); err != nil {
			return nil, err
		}
		return tradesMsg, nil
	default:
		return nil, fmt.Errorf("Unsupported channel for parsing: %v", baseMsg)
	}
}
