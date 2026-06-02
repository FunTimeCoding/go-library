package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.FallbackInteger(constant.PortEnvironment, 0),
		!environment.Exists(constant.InsecureEnvironment),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.PasswordEnvironment),
		"",
		o...,
	)
}
