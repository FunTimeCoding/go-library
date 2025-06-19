package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment),
		strings.ToInteger(
			environment.Get(constant.PortEnvironment),
			0,
		),
		true,
		environment.Get(constant.UserEnvironment),
		environment.Get(constant.PasswordEnvironment),
		"",
	)
}
