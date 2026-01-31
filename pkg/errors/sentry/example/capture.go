package example

import (
	"fmt"
	"github.com/getsentry/sentry-go"
)

func Capture() {
	sentry.CaptureException(fmt.Errorf("example error"))
}
