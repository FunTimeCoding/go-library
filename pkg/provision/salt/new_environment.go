package salt

import (
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		strings.ToIntegerStrict(
			environment.Fallback(constant.PortEnvironment, "8000"),
		),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.PasswordEnvironment),
	)
}
