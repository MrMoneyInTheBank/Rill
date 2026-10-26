package messages

import "github.com/MrMoneyInTheBank/rill/internal/channels"

type WebsocketMsg struct {
	Type string `json:"type"`
}

type SubscribeMsg struct {
	Type       string             `json:"type"`
	ProductIDs []string           `json:"product_ids,omitempty"`
	Channels   []channels.Channel `json:"channels"`
}

type Msg interface {
	isMsg()
}

func (WebsocketMsg) isMsg() {}
func (SubscribeMsg) isMsg() {}
