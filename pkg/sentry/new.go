package sentry

import (
	"fmt"
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sentry/basic_client"
)

func New(
	host string,
	token string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(token, "token")
	host2 := fmt.Sprintf(
		"https://%s%s",
		host,
		sentry.DefaultEndpoint,
	)
	client, e := sentry.NewClient(token, &host2, nil)
	errors.PanicOnError(e)

	return &Client{
		basicClient: basic_client.New(host, token),
		client:      client,
	}
}
