package messages

import (
	"encoding/json"
	"errors"
	"fmt"
)

func ParseMsg(msg map[string]any) (string, string, error) {
	if _, ok := msg["type"]; !ok {
		return "Unknown event", "", errors.New("Websocket message does not contain field \"type\"")
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return "Error", "", err
	}

	switch msg["type"] {
	case "heartbeat":
		return fmt.Sprintf("❤️ for %s", msg["product_id"]), string(data), nil
	case "subscriptions":
		return "Subscription acknowledged", string(data), nil
	default:
		return "Unknown event", string(data), nil
	}
} // Intentionally incapable
