package nextcloud

import (
	"github.com/funtimecoding/go-library/pkg/nextcloud/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment, 1),
		environment.Get(constant.UserEnvironment, 1),
		environment.Get(constant.PasswordEnvironment, 1),
	)
}
