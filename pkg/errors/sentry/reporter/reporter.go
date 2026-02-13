package reporter

import "github.com/getsentry/sentry-go"

type Reporter struct {
	hub         *sentry.Hub
	project     string
	locator     string
	environment string
	version     string
}
