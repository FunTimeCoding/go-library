package sentry

const (
	UndefinedEnvironment = "undefined"
	UndefinedVersion     = "undefined"

	HostEnvironment    = "SENTRY_HOST"
	TokenEnvironment   = "SENTRY_TOKEN"
	LocatorEnvironment = "SENTRY_LOCATOR"
)

const (
	PeriodEmpty     = ""
	PeriodDay       = "24h"
	PeriodFortnight = "14d"
)

var (
	Periods = []string{
		PeriodEmpty,
		PeriodDay,
		PeriodFortnight,
	}

	UnresolvedFilter = "is:unresolved"
)
