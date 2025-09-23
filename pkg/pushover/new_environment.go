package pushover

import (
	"github.com/funtimecoding/go-library/pkg/pushover/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.TokenEnvironment),
	)
}
