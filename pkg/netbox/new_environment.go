package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(p ...OptionFunc) *Client {
	return New(
		environment.Get(constant.HostEnvironment),
		environment.Get(constant.TokenEnvironment),
		p...,
	)
}
