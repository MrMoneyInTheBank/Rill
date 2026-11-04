package messages

import (
	"github.com/MrMoneyInTheBank/rill/internal/channels"
	"github.com/MrMoneyInTheBank/rill/internal/config"
)

type SubscribeReq struct {
	Type       string           `json:"type"`
	ProductIDs []string         `json:"product_ids,omitempty"`
	Channel    channels.Channel `json:"channel"`
}

type UnsubscribeReq struct {
	Type       string           `json:"type"`
	ProductIDs []string         `json:"product_ids,omitempty"`
	Channel    channels.Channel `json:"channel"`
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

func GetSubReqs(opts config.Opts) ([]Req, error) {
	requests := make([]Req, 0, 3)

	if opts.Heartbeats {
		heartbeatsReq, err := NewSubscribeReq(channels.HeartbeatsChannel)
		if err != nil {
			return nil, err
		}
		requests = append(requests, heartbeatsReq)

	}
	if opts.Level2 {

		level2Req, err := NewSubscribeReq(channels.Level2Channel, opts.Product)
		if err != nil {
			return nil, err
		}

		requests = append(requests, level2Req)
	}

	if opts.MarketTrades {

		marketTradesReq, err := NewSubscribeReq(channels.MarketTradesChannel, opts.Product)
		if err != nil {
			return nil, err
		}
		requests = append(requests, marketTradesReq)
	}

	return requests, nil
}
