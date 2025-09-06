package kestra

import (
	"github.com/funtimecoding/go-library/pkg/kestra/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment),
		environment.Get(constant.TokenEnvironment),
	)
}
