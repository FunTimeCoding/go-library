package iterm

import (
	"github.com/funtimecoding/go-library/pkg/iterm/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.HostEnvironment); s != "" {
		o = append(o, WithHost(s))
	}

	return New(o...)
}
