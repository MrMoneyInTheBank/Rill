package channels

type Channel string

const (
	HeartbeatsChannel    Channel = "heartbeats"
	Level2Channel        Channel = "level2"
	MarketTradesChannel  Channel = "market_trades"
	SubscriptionaChannel Channel = "subscriptions"
)
