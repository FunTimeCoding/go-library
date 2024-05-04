package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/getsentry/sentry-go"
	"time"
)

func Flush(h *sentry.Hub) {
	if !h.Flush(2 * time.Second) {
		errors.Warning("sentry flush fail")
	}
}
