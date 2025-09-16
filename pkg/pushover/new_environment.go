package pushover

import (
	"github.com/funtimecoding/go-library/pkg/pushover/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Exit(constant.UserEnvironment),
		environment.Exit(constant.TokenEnvironment),
	)
}
