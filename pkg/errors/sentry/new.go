package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/client"
)

func New(
	host string,
	token string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(token, "token")

	return &Client{
		basic:  basic.New(host, token),
		client: client.New(host, token),
	}
}
