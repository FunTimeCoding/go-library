package sentry

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/getsentry/sentry-go"
)

func Start(
	projectName string,
	locator string,
	environment string,
	version string,
) *sentry.Hub {
	errors.FatalOnEmpty(projectName, "projectName")
	errors.FatalOnEmpty(locator, "locator")

	if environment == "" {
		environment = UndefinedEnvironment
	}

	if version == "" {
		version = UndefinedVersion
	}

	h := sentry.CurrentHub()
	client, e := sentry.NewClient(
		sentry.ClientOptions{
			Dsn:         locator,
			Environment: environment,
			Release: fmt.Sprintf(
				"%s@%s",
				projectName,
				version,
			),
			TracesSampleRate: 1.0,
		},
	)
	errors.FatalOnError(e)
	h.BindClient(client)

	return h
}
