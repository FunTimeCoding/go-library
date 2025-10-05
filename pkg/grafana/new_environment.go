package grafana

import (
	"github.com/funtimecoding/go-library/pkg/grafana/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.PortEnvironment),
		environment.Required(constant.TokenEnvironment),
	)
}
