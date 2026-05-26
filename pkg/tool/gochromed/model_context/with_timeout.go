package model_context

import (
	"context"
	"time"
)

func withTimeout[T any](
	timeout time.Duration,
	f func() (T, error),
) (T, error) {
	type result struct {
		value T
		e     error
	}
	done := make(chan result, 1)
	go func() {
		v, e := f()
		done <- result{value: v, e: e}
	}()

	select {
	case r := <-done:
		return r.value, r.e
	case <-time.After(timeout):
		var zero T

		return zero, context.DeadlineExceeded
	}
}
