package kestra

import (
	"github.com/funtimecoding/go-library/pkg/kestra/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Fallback(
		constant.TokenEnvironment,
		"",
	); s != "" {
		o = append(o, WithToken(s))
	}

	if s := environment.Fallback(
		constant.UserEnvironment,
		"",
	); s != "" {
		o = append(o, WithUser(s))
	}

	if s := environment.Fallback(
		constant.PasswordEnvironment,
		"",
	); s != "" {
		o = append(o, WithPassword(s))
	}

	return New(environment.Required(constant.HostEnvironment), o...)
}
