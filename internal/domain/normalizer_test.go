package domain_test

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"github.com/MrMoneyInTheBank/rill/internal/domain"
	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

func TestNormalizeL2PreservesEvents(t *testing.T) {
	ts := time.Date(2026, 11, 5, 0, 0, 0, 0, time.UTC)
	msg := messages.Level2Msg{
		BaseMsg: messages.BaseMsg{Timestamp: ts},
		Events: []messages.L2Event{
			{Type: messages.L2SnapshotEvent, ProductID: "BTC-USD", Updates: []messages.L2Update{
				{Side: messages.L2BidSide, EventTime: ts, PriceLevel: decimal.NewFromInt(100), NewQty: decimal.NewFromInt(2)},
				{Side: messages.L2OfferSide, EventTime: ts, PriceLevel: decimal.NewFromInt(101), NewQty: decimal.NewFromInt(3)},
			}},
			{Type: messages.L2UpdateEvent, ProductID: "BTC-USD", Updates: []messages.L2Update{
				{Side: messages.L2BidSide, EventTime: ts, PriceLevel: decimal.NewFromInt(100), NewQty: decimal.Zero},
			}},
			{Type: messages.L2SnapshotEvent, ProductID: "ETH-USD"},
			{Type: messages.L2SnapshotEvent, ProductID: "BTC-USD"},
		},
	}
	got, err := domain.NormalizeL2(msg)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 4 {
		t.Fatalf("got %d events, want 4", len(got))
	}
	for i, want := range []struct {
		kind    domain.OrderType
		product string
		levels  int
	}{
		{domain.SNAPSHOT, "BTC-USD", 2},
		{domain.UPDATE, "BTC-USD", 1},
		{domain.SNAPSHOT, "ETH-USD", 0},
		{domain.SNAPSHOT, "BTC-USD", 0},
	} {
		if got[i].Type != want.kind || got[i].ProductID != want.product || len(got[i].Levels) != want.levels || !got[i].OrderRecvTime.Equal(ts) {
			t.Errorf("event %d: unexpected grouping or metadata: %#v", i, got[i])
		}
	}
	if got[0].Levels[0].Side != domain.BID || got[0].Levels[1].Side != domain.OFFER {
		t.Fatal("snapshot sides were not preserved")
	}
	if !got[0].Levels[1].PriceLevel.Equal(decimal.NewFromInt(101)) || !got[0].Levels[1].NewQty.Equal(decimal.NewFromInt(3)) || !got[0].Levels[1].OrderTime.Equal(ts) {
		t.Fatal("snapshot level values were not preserved")
	}
	if !got[1].Levels[0].NewQty.IsZero() {
		t.Fatal("zero-quantity removal was not preserved")
	}
}

func TestNormalizeL2RejectsInvalidValues(t *testing.T) {
	for _, event := range []messages.L2Event{
		{Type: "unexpected"},
		{Type: messages.L2UpdateEvent, Updates: []messages.L2Update{{Side: "unexpected"}}},
	} {
		got, err := domain.NormalizeL2(messages.Level2Msg{Events: []messages.L2Event{event}})
		if err == nil || got != nil {
			t.Fatalf("invalid event accepted: %#v", event)
		}
	}
}
