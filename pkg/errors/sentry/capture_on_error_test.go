package sentry

import (
	"github.com/getsentry/sentry-go"
	"testing"
)

func TestCaptureOnError(t *testing.T) {
	CaptureOnError(sentry.CurrentHub(), nil)
}
