package stream

import (
	"context"
	"log"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"

	"github.com/MrMoneyInTheBank/rill/internal/messages"
)

const defaultMaxReadSize int64 = 8 << 20

func Connect(ctx context.Context, endpoint string, maxReadSize *int64) (*websocket.Conn, error) {
	// nil maxReadSize falls back to default ReadLimit of 8MiB
	// Size chosen based on empirical results during testing
	var readLimit int64 = defaultMaxReadSize
	if maxReadSize != nil {
		if *maxReadSize <= 0 && *maxReadSize != -1 {
			return nil, ErrInvalidReadSize
		}
		readLimit = *maxReadSize
	}

	dialCtx, dialCancel := context.WithTimeout(ctx, time.Second*10)
	conn, _, err := websocket.Dial(dialCtx, endpoint, nil)
	dialCancel()

	if err != nil {
		return nil, err
	}

	conn.SetReadLimit(readLimit)

	return conn, nil
}

func SendMsg(ctx context.Context, conn *websocket.Conn, msg messages.Req) error {
	return wsjson.Write(ctx, conn, msg)
}

func SendMsgs(ctx context.Context, conn *websocket.Conn, msgs []messages.Req) error {
	for _, msg := range msgs {
		if err := SendMsg(ctx, conn, msg); err != nil {
			return err
		}
		log.Printf("SENT request msg: %#v", msg)
	}

	return nil
}
