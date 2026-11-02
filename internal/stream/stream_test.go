package stream_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/coder/websocket"

	"github.com/MrMoneyInTheBank/rill/internal/channels"
	"github.com/MrMoneyInTheBank/rill/internal/messages"
	"github.com/MrMoneyInTheBank/rill/internal/stream"
)

func ptr[T any](t T) *T {
	return &t
}

func setupMockWSServer(t *testing.T, send bool, read bool) (string, func(), chan error, chan error) {
	t.Helper()

	writeCh := make(chan error, 1)
	readCh := make(chan error, 1)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			return
		}
		defer conn.CloseNow()

		if send {
			data := make([]byte, 4<<10)
			for i := range data {
				data[i] = 'A'
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()
			err := conn.Write(ctx, websocket.MessageText, data)
			writeCh <- err

		}

		if read {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			_, _, err := conn.Read(ctx)
			readCh <- err
		}
	}))

	wsURL := strings.Replace(server.URL, "http", "ws", 1)

	teardown := func() {
		server.Close()
	}

	return wsURL, teardown, writeCh, readCh
}

func TestInvalidReadLimit(t *testing.T) {
	cases := []struct {
		name         string
		readSize     int64
		expectedConn *websocket.Conn
		expectedErr  error
	}{
		{"zero", 0, nil, stream.ErrInvalidReadSize},
		{"negativeBarMinusOne", -2, nil, stream.ErrInvalidReadSize},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			conn, err := stream.Connect(context.Background(), "", &tc.readSize)
			if conn != nil {
				defer conn.CloseNow()
			}

			if conn != tc.expectedConn || !errors.Is(err, tc.expectedErr) {
				t.Errorf("Read size %d does not produce error", tc.readSize)
			}
		})
	}
}

func TestSuccessfulHandshake(t *testing.T) {
	endpoint, teardown, _, _ := setupMockWSServer(t, false, false)
	defer teardown()

	cases := []struct {
		name        string
		readSize    *int64
		expectedErr error
	}{
		{"ExplicitMaxReadSize", ptr(int64(4 << 10)), nil},
		{"DisableReadSize", ptr(int64(-1)), nil},
		{"CustomDefaultReadSize", nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			conn, err := stream.Connect(context.Background(), endpoint, tc.readSize)
			if conn != nil {
				defer conn.CloseNow()
			}

			if conn == nil || err != nil {
				if tc.readSize == nil {
					t.Errorf("Default max read size fails to give successful handshake")
				} else {
					t.Errorf("Max read size %d fails to give successful handshake", *tc.readSize)
				}
			}
		})
	}
}

func TestReadSizeConfig(t *testing.T) {
	endpoint, teardown, ch, _ := setupMockWSServer(t, true, false)
	defer teardown()

	cases := []struct {
		name        string
		readSize    int64
		expectedErr error
	}{
		{"SufficientSize", 8 << 10, nil},
		{"IndufficientSize", 1, websocket.ErrMessageTooBig},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			conn, err := stream.Connect(context.Background(), endpoint, &tc.readSize)
			if err != nil {
				t.Fatal("Could not establish connection")
			}
			defer conn.CloseNow()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()

			_, _, err = conn.Read(ctx)

			select {
			case writeErr := <-ch:
				if writeErr != nil {
					t.Logf("Server write error: %v", writeErr)
				}
			case <-ctx.Done():
				t.Fatalf("Timed out waiting for server write result: %v", ctx.Err())
			}

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Failed read for max read size %d", tc.readSize)
			}
		})
	}
}

func TestSendMsg(t *testing.T) {
	endpoint, teardown, _, ch := setupMockWSServer(t, false, true)
	defer teardown()

	req := messages.SubscribeReq{
		Type:    "subscribe",
		Channel: channels.HeartbeatsChannel,
	}

	cases := []struct {
		name         string
		shouldCancel bool
		expectedErr  error
	}{
		{"RegularSend", false, nil},
		{"CancelledContextSend", true, context.Canceled},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			conn, err := stream.Connect(context.Background(), endpoint, nil)
			if err != nil {
				t.Fatalf("Could not establish connection: %v", err)
			}
			defer conn.CloseNow()

			sendCtx, cancelSend := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelSend()

			if tc.shouldCancel {
				cancelSend()
			}

			err = stream.SendMsg(sendCtx, conn, req)
			if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("SendMsg error = %v; want %v", err, tc.expectedErr)
			}

			if tc.shouldCancel {
				conn.CloseNow()
			}

			waitCtx, cancelWait := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancelWait()

			select {
			case readErr := <-ch:
				if !tc.shouldCancel && readErr != nil {
					t.Fatalf("Server could not receive message: %v", readErr)
				}
			case <-waitCtx.Done():
				t.Fatalf("Timed out waiting for server read result: %v", waitCtx.Err())
			}
		})
	}
}
