package sentry

import (
	"github.com/funtimecoding/go-library/pkg/log"
	"github.com/getsentry/sentry-go"
	"time"
)

func Flush() {
	if !sentry.Flush(2 * time.Second) {
		log.Warning("sentry flush fail")
	}
}
