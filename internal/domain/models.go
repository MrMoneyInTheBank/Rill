package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	OrderSide string
	OrderType string
)

const (
	BID   OrderSide = "BID"
	OFFER OrderSide = "OFFER"
)

const (
	SNAPSHOT OrderType = "SNAPSHOT"
	UPDATE   OrderType = "UPDATE"
)

type Trade struct {
	TradeID       string
	ProductID     string
	TradeTime     time.Time
	TradeRecvTime time.Time
	Price         decimal.Decimal
	Size          decimal.Decimal
}

type Order struct {
	OrderTime  time.Time
	Side       OrderSide
	PriceLevel decimal.Decimal
	NewQty     decimal.Decimal
}

// BookEvent preserves one complete snapshot or update, including an empty snapshot.
type BookEvent struct {
	Type          OrderType
	ProductID     string
	OrderRecvTime time.Time
	Levels        []Order
}

type Heartbeat struct {
	Counter int
}
