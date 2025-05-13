package chromium

import (
	"github.com/funtimecoding/go-library/pkg/chromium/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment, 1),
		strings.ToIntegerStrict(
			environment.Get(constant.PortEnvironment, 1),
		),
	)
}
