package stream

import (
	"context"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"

	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

func Connect(ctx context.Context, endpoint string, readSize *int64) (*websocket.Conn, error) {
	dialCtx, dialCancel := context.WithTimeout(ctx, time.Second*10)
	conn, _, err := websocket.Dial(dialCtx, endpoint, nil)
	dialCancel()

	if err != nil {
		return nil, err
	}

	conn.SetReadLimit(4 << 20) // temp buffer size 4 MiB
	if readSize != nil {
		conn.SetReadLimit(*readSize)
	}

	return conn, nil
}

func SendMsg(ctx context.Context, conn *websocket.Conn, msg messages.Msg) error {
	err := wsjson.Write(ctx, conn, msg)
	if err != nil {
		return err
	}
	return nil
}
