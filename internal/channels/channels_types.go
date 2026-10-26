package channels

type ChannelName string

const (
	HeartbeatChannel   ChannelName = "heartbeat"
	Level2BatchChannel ChannelName = "level2_batch"
)

type ChannelSpec struct {
	Name       ChannelName `json:"name"`
	ProductIDs []string    `json:"product_ids,omitempty"`
}

type Channel interface {
	isChannel()
}

func (ChannelName) isChannel() {}
func (ChannelSpec) isChannel() {}
