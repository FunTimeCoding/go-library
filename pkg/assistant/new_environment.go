package assistant

import (
	"github.com/funtimecoding/go-library/pkg/assistant/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.TokenEnvironment),
		o...,
	)
}
