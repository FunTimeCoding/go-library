package open_webui

import (
	"github.com/funtimecoding/go-library/pkg/generative/open_webui/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment, 1),
		environment.Get(constant.TokenEnvironment, 1),
	)
}
