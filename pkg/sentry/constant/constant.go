package constant

const (
	UndefinedEnvironment = "undefined"
	UndefinedVersion     = "undefined"

	HostEnvironment         = "SENTRY_HOST"
	TokenEnvironment        = "SENTRY_TOKEN"
	OrganizationEnvironment = "SENTRY_ORGANIZATION"
	ProjectEnvironment      = "SENTRY_PROJECT"
	LocatorEnvironment      = "SENTRY_LOCATOR"

	ErrorType = "error"
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
