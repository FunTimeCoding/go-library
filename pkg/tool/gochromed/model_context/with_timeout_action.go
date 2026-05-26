package model_context

import (
	"context"
	"time"
)

func withTimeoutAction(
	timeout time.Duration,
	f func() error,
) error {
	done := make(chan error, 1)
	go func() {
		done <- f()
	}()

	select {
	case e := <-done:
		return e
	case <-time.After(timeout):
		return context.DeadlineExceeded
	}
}
