package n8n

import (
	"github.com/funtimecoding/go-library/pkg/generative/n8n/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment),
		environment.Get(constant.TokenEnvironment),
	)
}
