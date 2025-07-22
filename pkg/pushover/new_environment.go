package pushover

import (
	"github.com/funtimecoding/go-library/pkg/pushover/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.UserEnvironment),
		environment.Get(constant.TokenEnvironment),
	)
}
