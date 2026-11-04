package messages

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

const goTimeLayout = "2006-01-02 15:04:05.999999999 -0700 MST"

func (e *HeartbeatEvent) UnmarshalJSON(b []byte) error {
	var raw struct {
		CurrentTime      string `json:"current_time"`
		HeartbeatCounter int    `json:"heartbeat_counter"`
	}

	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	s := raw.CurrentTime
	if before, _, found := strings.Cut(s, " m="); found {
		s = before
	}

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		if t, err = time.Parse(goTimeLayout, s); err != nil {
			return fmt.Errorf("parsing current_time %q: %w", raw.CurrentTime, err)
		}
	}

	e.CurrentTime = t
	e.HeartbeatCounter = raw.HeartbeatCounter

	return nil
}

func (e *L2Update) UnmarshalJSON(b []byte) error {
	var raw struct {
		Side       L2UpdateSide    `json:"side"`
		EventTime  string          `json:"event_time"`
		PriceLevel decimal.Decimal `json:"price_level"`
		NewQty     decimal.Decimal `json:"new_quantity"`
	}

	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	s := raw.EventTime
	if before, _, found := strings.Cut(s, " m="); found {
		s = before
	}

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		if t, err = time.Parse(goTimeLayout, s); err != nil {
			return fmt.Errorf("parsing current_time %q: %w", raw.EventTime, err)
		}
	}

	e.Side = raw.Side
	e.EventTime = t
	e.PriceLevel = raw.PriceLevel
	e.NewQty = raw.NewQty

	return nil
}
