package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Exit(constant.HostEnvironment),
		strings.ToInteger(
			environment.Exit(constant.PortEnvironment),
			0,
		),
		true,
		environment.Exit(constant.UserEnvironment),
		environment.Exit(constant.PasswordEnvironment),
		"",
	)
}
