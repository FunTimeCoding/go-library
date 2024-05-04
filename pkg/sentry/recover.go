package sentry

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
)

func Recover(
	h *sentry.Hub,
	a any,
) {
	if a == nil {
		return
	}

	var e error

	switch t := a.(type) {
	case error:
		e = t
	case string:
		e = errors.New(t)
	default:
		e = fmt.Errorf("unexpected recover: %#v", a)
	}

	h.CaptureException(e)
}
