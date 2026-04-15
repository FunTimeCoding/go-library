package netb

import (
	"github.com/funtimecoding/go-library/pkg/netb/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.RequiredInteger(constant.PortEnvironment),
		environment.Exists(constant.InsecureEnvironment),
	)
}
