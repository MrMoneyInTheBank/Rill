package messages

import "fmt"

func ParseMsg(msg map[string]any) string {
	if _, ok := msg["type"]; !ok {
		return "Unknown event"
	}

	switch msg["type"] {
	case "heartbeat":
		return fmt.Sprintf("❤️ for %s", msg["product_id"])
	case "subscriptions":
		return "Subscription acknowledged"
	default:
		return "Unknown event"
	}
} // Intentionally incapable
