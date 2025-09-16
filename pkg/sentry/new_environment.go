package sentry

import (
	"github.com/funtimecoding/go-library/pkg/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Exit(constant.HostEnvironment),
		environment.Exit(constant.TokenEnvironment),
	)
}
