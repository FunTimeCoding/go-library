package kestra

import (
	"github.com/funtimecoding/go-library/pkg/kestra/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Default(
		constant.TokenEnvironment,
		"",
	); s != "" {
		o = append(o, WithToken(s))
	}

	if s := environment.Default(
		constant.UserEnvironment,
		"",
	); s != "" {
		o = append(o, WithUser(s))
	}

	if s := environment.Default(
		constant.PasswordEnvironment,
		"",
	); s != "" {
		o = append(o, WithPassword(s))
	}

	return New(environment.Exit(constant.HostEnvironment), o...)
}
