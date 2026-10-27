package messages

import (
	"encoding/json"
	"fmt"

	"github.com/MrMoneyInTheBank/rill/internal/channels"
)

func ParseMsg(data []byte) (string, string, error) {
	var frame MsgFrame

	if err := json.Unmarshal(data, &frame); err != nil {
		return "Failed parsing message", string(data), err
	}

	switch frame.Channel {
	case channels.SubscriptionaChannel:
		return "Subscription acknowledged", string(data), nil
	case channels.HeartbeatsChannel:
		var hbEvents []HeartbeatEvent
		if err := json.Unmarshal(frame.Events, &hbEvents); err != nil {
			return "Failed parsing heartbeats", "", err
		}
		if len(hbEvents) > 0 {
			msg := fmt.Sprintf("[Seq: %d] Heartbeat counter: %v", frame.SequenceNum, hbEvents[0].HeartbeatCounter)
			return msg, string(data), nil
		} else {
			return "No heartbeats", string(data), nil
		}
	default:
		return "Unknown message", string(data), nil
	}
} // Intentionally incapable
