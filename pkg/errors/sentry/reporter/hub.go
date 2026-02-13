package reporter

import "github.com/getsentry/sentry-go"

func (r *Reporter) Hub() *sentry.Hub {
	return r.hub
}
