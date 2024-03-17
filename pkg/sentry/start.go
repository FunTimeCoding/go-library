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
) {
	errors.FatalOnError(
		sentry.Init(
			sentry.ClientOptions{
				Dsn:         locator,
				Environment: environment,
				Release:     fmt.Sprintf("%s@%s", projectName, version),
			},
		),
	)
}
