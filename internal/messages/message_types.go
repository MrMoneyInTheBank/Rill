package messages

import (
	"github.com/MrMoneyInTheBank/rill/internal/channels"
)

type BaseMsg struct {
	Channel     channels.Channel `json:"channel"`
	Timestamp   string           `json:"timestamp"`
	SequenceNum int              `json:"sequence_num"`
}

type SubscriptionMsg struct {
	BaseMsg
	Events []SubscriptionEvent `json:"events"`
}

type HeartbeatsMsg struct {
	BaseMsg
	Events []HeartbeatEvent `json:"events"`
}

type TradesMsg struct {
	BaseMsg
	Events []TradeEvent `json:"events"`
}

type Level2Msg struct {
	BaseMsg
	Events []L2Event `json:"events"`
}

type Msg interface {
	isMsg()
}

func (BaseMsg) isMsg() {}
