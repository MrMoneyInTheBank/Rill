package messages_test

import (
	"fmt"
	"testing"

	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

func TestParseMsgValidatesEventTypesAndSides(t *testing.T) {
	for _, feed := range []struct {
		channel string
		items   string
		fields  string
		sides   []string
	}{
		{"l2_data", "updates", `"event_time":"2026-11-05T00:00:00Z","price_level":"100","new_quantity":"0"`, []string{"bid", "offer"}},
		{"market_trades", "trades", `"time":"2026-11-05T00:00:00Z","price":"100","size":"1","trade_id":"1"`, []string{"BUY", "SELL"}},
	} {
		for _, eventType := range []string{"snapshot", "update", "", "unexpected"} {
			for _, side := range append(feed.sides, "", "unexpected") {
				t.Run(feed.channel+"/"+eventType+"/"+side, func(t *testing.T) {
					payload := fmt.Sprintf(`{"channel":%q,"events":[{"type":%q,%q:[{"side":%q,%s}]}]}`, feed.channel, eventType, feed.items, side, feed.fields)
					_, err := messages.ParseMsg([]byte(payload))
					valid := (eventType == "snapshot" || eventType == "update") && (side == feed.sides[0] || side == feed.sides[1])
					if (err == nil) != valid {
						t.Fatalf("valid=%v, got error %v", valid, err)
					}
				})
			}
		}
		for _, eventType := range []string{"", "unexpected"} {
			payload := fmt.Sprintf(`{"channel":%q,"events":[{"type":%q,%q:[]}]}`, feed.channel, eventType, feed.items)
			if _, err := messages.ParseMsg([]byte(payload)); err == nil {
				t.Fatalf("accepted invalid event type %q without items", eventType)
			}
		}
	}
}
