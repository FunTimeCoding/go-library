package salt

import (
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment, 1),
		strings.ToIntegerStrict(
			environment.GetDefault(constant.PortEnvironment, "8000"),
		),
		environment.Get(constant.UserEnvironment, 1),
		environment.Get(constant.PasswordEnvironment, 1),
	)
}
