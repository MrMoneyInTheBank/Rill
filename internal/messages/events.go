package messages

// Heartbeat

type HeartbeatEvent struct {
	CurrentTime      string `json:"current_time"`
	HeartbeatCounter int    `json:"heartbeat_counter"`
}

// Market Trades

type Trade struct {
	TradeID   string `json:"trade_id"`
	ProductID string `json:"product_id"`
	Price     string `json:"price"`
	Size      string `json:"size"`
	MakerSide string `json:"side"`
	Time      string `json:"time"`
}

type TradeEvent struct {
	Type   string `json:"type"`
	Trades []Trade
}

// L2

type L2Update struct {
	Side       string `json:"side"`
	EventTime  string `json:"event_time"`
	PriceLevel string `json:"price_level"`
	NewQty     string `json:"new_quantity"`
}

type L2Event struct {
	Type      string     `json:"type"`
	ProductID string     `json:"product_id"`
	Updates   []L2Update `json:"updates"`
}
