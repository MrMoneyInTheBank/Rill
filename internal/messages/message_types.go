package messages

import (
	"encoding/json"

	"github.com/MrMoneyInTheBank/rill/internal/channels"
)

type BaseMsg struct {
	Channel     channels.Channel `json:"channel"`
	Timestamp   string           `json:"timestamp"`
	SequenceNum int              `json:"sequence_num"`
}

type SubscribeMsg struct {
	Type       string           `json:"type"`
	ProductIDs []string         `json:"product_ids,omitempty"`
	Channel    channels.Channel `json:"channel"`
	JWT        string           `json:"jwt,omitempty"`
}

type UnsubscribeMsg struct {
	Type       string           `json:"type"`
	ProductIDs []string         `json:"product_ids,omitempty"`
	Channel    channels.Channel `json:"channel"`
	JWT        string           `json:"jwt,omitempty"`
}

type MsgFrame struct {
	BaseMsg
	Events json.RawMessage `json:"events"`
}

type HeartbeatsMsg struct {
	BaseMsg
	Events []HeartbeatEvent `json:"events"`
}

type TradesMsg struct {
	BaseMsg
	Events TradeEvent `json:"events"`
}

type Level2Msg struct {
	BaseMsg
	Events []L2Event `json:"events"`
}

type Msg interface {
	isMsg()
}

func (BaseMsg) isMsg()      {}
func (SubscribeMsg) isMsg() {}
