package messages

import "github.com/MrMoneyInTheBank/rill/internal/channels"

type SubscribeReq struct {
	Type       string           `json:"type"`
	ProductIDs []string         `json:"product_ids,omitempty"`
	Channel    channels.Channel `json:"channel"`
	JWT        string           `json:"jwt,omitempty"`
}

type UnsubscribeReq struct {
	Type       string           `json:"type"`
	ProductIDs []string         `json:"product_ids,omitempty"`
	Channel    channels.Channel `json:"channel"`
	JWT        string           `json:"jwt,omitempty"`
}

type Req interface {
	isReq()
}

func (SubscribeReq) isReq()   {}
func (UnsubscribeReq) isReq() {}
