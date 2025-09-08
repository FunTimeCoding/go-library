package kestra

import (
	"github.com/funtimecoding/go-library/pkg/kestra/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...OptionFunc) *Client {
	if s := environment.GetDefault(
		constant.TokenEnvironment,
		"",
	); s != "" {
		o = append(o, WithToken(s))
	}

	if s := environment.GetDefault(
		constant.UserEnvironment,
		"",
	); s != "" {
		o = append(o, WithUser(s))
	}

	if s := environment.GetDefault(
		constant.PasswordEnvironment,
		"",
	); s != "" {
		o = append(o, WithPassword(s))
	}

	return New(environment.Get(constant.HostEnvironment), o...)
}
