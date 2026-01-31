package client

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func New(
	host string,
	token string,
) *sentry.Client {
	result, e := sentry.NewClient(
		token,
		locator.New(host).Base(sentry.DefaultEndpoint).Trail().Pointer(),
		nil,
	)
	errors.PanicOnError(e)

	return result
}
