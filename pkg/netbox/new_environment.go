package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(p ...Option) *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.TokenEnvironment),
		p...,
	)
}
