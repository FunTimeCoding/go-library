package sentry

import "github.com/getsentry/sentry-go"

func CaptureOnError(
	h *sentry.Hub,
	e error,
) {
	if e != nil {
		h.CaptureException(e)
	}
}
