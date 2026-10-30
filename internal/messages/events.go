package messages

// Subscription (placeholder)

type SubscriptionEvent struct {
	Subscriptions map[string][]string `json:"subscriptions"`
}

// Heartbeat

type HeartbeatEvent struct {
	CurrentTime      string `json:"current_time"`
	HeartbeatCounter int    `json:"heartbeat_counter"`
}

// Market Trades

type (
	TradeEventType string
	TradeSide      string
)

const (
	TradeSnapshotEvent TradeEventType = "snapshot"
	TradeUpdateEvent   TradeEventType = "update"
)

const (
	TradeBuySide  TradeSide = "BUY"
	TradeSellSide TradeSide = "SELL"
)

type Trade struct {
	TradeID   string    `json:"trade_id"`
	ProductID string    `json:"product_id"`
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	MakerSide TradeSide `json:"side"`
	Time      string    `json:"time"`
}

type TradeEvent struct {
	Type   string  `json:"type"`
	Trades []Trade `json:"trades"`
}

// L2

type (
	L2EventType  string
	L2UpdateSide string
)

const (
	L2SnapshotEvent L2EventType = "snapshot"
	L2UpdateEvent   L2EventType = "update"
)

const (
	L2BidSide   L2UpdateSide = "bid"
	L2OfferSide L2UpdateSide = "offer"
)

type L2Update struct {
	Side       L2UpdateSide `json:"side"`
	EventTime  string       `json:"event_time"`
	PriceLevel string       `json:"price_level"`
	NewQty     string       `json:"new_quantity"`
}

type L2Event struct {
	Type      L2EventType `json:"type"`
	ProductID string      `json:"product_id"`
	Updates   []L2Update  `json:"updates"`
}
