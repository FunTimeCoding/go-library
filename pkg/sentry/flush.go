package sentry

import (
	"github.com/getsentry/sentry-go"
	"time"
)

func Flush() {
	sentry.Flush(2 * time.Second)
}
