package server

import (
	"context"
	"fmt"
	"github.com/coder/websocket"
	"golang.org/x/time/rate"
	"io"
	"time"
)

// echo reads from the WebSocket connection and then writes
// the received message back to it.
// The entire function has 10s to complete.
func echo(
	c *websocket.Conn,
	l *rate.Limiter,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if e := l.Wait(ctx); e != nil {
		return e
	}

	messageType, r, readFail := c.Reader(ctx)

	if readFail != nil {
		return readFail
	}

	w, writeFail := c.Writer(ctx, messageType)

	if writeFail != nil {
		return writeFail
	}

	if _, e := io.Copy(w, r); e != nil {
		return fmt.Errorf("failed to io.Copy: %w", e)
	}

	return w.Close()
}
