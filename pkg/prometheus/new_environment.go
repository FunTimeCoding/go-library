package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment, 1),
		strings.ToInteger(
			environment.Get(constant.PortEnvironment, 1),
			0,
		),
		true,
		environment.Get(constant.UserEnvironment, 1),
		environment.Get(constant.PasswordEnvironment, 1),
	)
}
