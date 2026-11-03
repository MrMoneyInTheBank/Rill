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

func NewSubscribeReq(channel channels.Channel, product_ids ...string) (Req, error) {
	if len(product_ids) > 1 {
		return SubscribeReq{}, ErrTooManyProducts
	}
	if channel == channels.HeartbeatsChannel && len(product_ids) != 0 {
		return SubscribeReq{}, ErrInvalidReq
	}

	if channel != channels.HeartbeatsChannel && len(product_ids) == 0 {
		return SubscribeReq{}, ErrInvalidReq
	}

	return SubscribeReq{
		Type:       "subscribe",
		Channel:    channel,
		ProductIDs: product_ids,
	}, nil
}

func GetSubReqs(product string, opts ...string) ([]Req, error) {
	heartbeatReq, err := NewSubscribeReq(channels.HeartbeatsChannel)
	if err != nil {
		return nil, err
	}
	level2Req, err := NewSubscribeReq(channels.Level2Channel, product)
	if err != nil {
		return nil, err
	}
	marketTradesReq, err := NewSubscribeReq(channels.MarketTradesChannel, product)
	if err != nil {
		return nil, err
	}

	return []Req{heartbeatReq, level2Req, marketTradesReq}, nil
}
