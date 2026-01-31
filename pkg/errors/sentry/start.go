package sentry

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
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
		environment = constant.UndefinedEnvironment
	}

	if version == "" {
		version = constant.UndefinedVersion
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
			EnableTracing:    true,
			TracesSampleRate: 1.0,
			AttachStacktrace: true,
			SendDefaultPII:   true,
		},
	)
	errors.FatalOnError(e)
	h.BindClient(client)

	return h
}
