package chroma

import (
	"github.com/funtimecoding/go-library/pkg/generative/chroma/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		strings.ToIntegerStrict(
			environment.Required(constant.PortEnvironment),
		),
	)
}
