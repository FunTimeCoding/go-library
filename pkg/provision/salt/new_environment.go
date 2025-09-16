package salt

import (
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Exit(constant.HostEnvironment),
		strings.ToIntegerStrict(
			environment.Default(constant.PortEnvironment, "8000"),
		),
		environment.Exit(constant.UserEnvironment),
		environment.Exit(constant.PasswordEnvironment),
	)
}
