package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.PasswordEnvironment),
		prometheus.NewEnvironment(),
	)
}
