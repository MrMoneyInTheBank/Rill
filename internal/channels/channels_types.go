package channels

type Channel string

const (
	HeartbeatsChannel    Channel = "heartbeats"
	Level2Channel        Channel = "level2"
	Level2DataChannel    Channel = "l2_data"
	MarketTradesChannel  Channel = "market_trades"
	SubscriptionaChannel Channel = "subscriptions"
)
