package salt

import (
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment),
		strings.ToIntegerStrict(
			environment.GetDefault(constant.PortEnvironment, "8000"),
		),
		environment.Get(constant.UserEnvironment),
		environment.Get(constant.PasswordEnvironment),
	)
}
