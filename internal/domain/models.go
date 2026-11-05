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
	Type          OrderType
	ProductID     string
	OrderTime     time.Time
	OrderRecvTime time.Time
	Side          OrderSide
	PriceLevel    decimal.Decimal
	NewQty        decimal.Decimal
}

type Heartbeat struct {
	Counter int
}
