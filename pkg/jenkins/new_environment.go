package jenkins

import (
	"github.com/funtimecoding/go-library/pkg/jenkins/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment),
		environment.Get(constant.UserEnvironment),
		environment.Get(constant.PasswordEnvironment),
	)
}
