package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		strings.ToInteger(
			environment.Required(constant.PortEnvironment),
			0,
		),
		true,
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.PasswordEnvironment),
		"",
		o...,
	)
}
