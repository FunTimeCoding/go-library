package sentry

import "github.com/funtimecoding/go-library/pkg/system/environment"

func NewEnvironment() *Client {
	return New(
		environment.Get(HostEnvironment, 1),
		environment.Get(TokenEnvironment, 1),
	)
}
