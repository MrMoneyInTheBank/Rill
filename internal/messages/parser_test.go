package messages_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

func TestParseMsgValidatesEventTypesAndSides(t *testing.T) {
	for _, feed := range []struct {
		channel  string
		items    string
		fields   string
		sides    []string
		eventErr error
		sideErr  error
	}{
		{"l2_data", "updates", `"event_time":"2026-11-05T00:00:00Z","price_level":"100","new_quantity":"0"`, []string{"bid", "offer"}, messages.ErrInvalidL2EventType, messages.ErrInvalidL2Side},
		{"market_trades", "trades", `"time":"2026-11-05T00:00:00Z","price":"100","size":"1","trade_id":"1"`, []string{"BUY", "SELL"}, messages.ErrInvalidTradeEventType, messages.ErrInvalidTradeSide},
	} {
		for _, eventType := range []string{"snapshot", "update", "", "unexpected"} {
			for _, side := range append(feed.sides, "", "unexpected") {
				t.Run(feed.channel+"/"+eventType+"/"+side, func(t *testing.T) {
					payload := fmt.Sprintf(`{"channel":%q,"events":[{"type":%q,%q:[{"side":%q,%s}]}]}`, feed.channel, eventType, feed.items, side, feed.fields)
					_, err := messages.ParseMsg([]byte(payload))
					var wantErr error
					if eventType != "snapshot" && eventType != "update" {
						wantErr = feed.eventErr
					} else if side != feed.sides[0] && side != feed.sides[1] {
						wantErr = feed.sideErr
					}
					if !errors.Is(err, wantErr) {
						t.Fatalf("got error %v, want %v", err, wantErr)
					}
				})
			}
		}
		for _, eventType := range []string{"", "unexpected"} {
			payload := fmt.Sprintf(`{"channel":%q,"events":[{"type":%q,%q:[]}]}`, feed.channel, eventType, feed.items)
			if _, err := messages.ParseMsg([]byte(payload)); !errors.Is(err, feed.eventErr) {
				t.Fatalf("event type %q without items: got %v, want %v", eventType, err, feed.eventErr)
			}
		}
	}
}

func TestParseMsgUnsupportedChannel(t *testing.T) {
	_, err := messages.ParseMsg([]byte(`{"channel":"unexpected"}`))
	if !errors.Is(err, messages.ErrUnsupportedChannel) {
		t.Fatalf("got %v, want %v", err, messages.ErrUnsupportedChannel)
	}
}
