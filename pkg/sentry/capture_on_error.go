package sentry

import "github.com/getsentry/sentry-go"

func CaptureOnError(e error) {
	if e != nil {
		sentry.CaptureException(e)
	}
}
