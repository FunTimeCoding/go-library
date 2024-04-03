package sentry

import "testing"

func TestCaptureOnError(t *testing.T) {
	CaptureOnError(nil)
}
