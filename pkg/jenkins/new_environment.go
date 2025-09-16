package jenkins

import (
	"github.com/funtimecoding/go-library/pkg/jenkins/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Exit(constant.HostEnvironment),
		environment.Exit(constant.UserEnvironment),
		environment.Exit(constant.PasswordEnvironment),
	)
}
