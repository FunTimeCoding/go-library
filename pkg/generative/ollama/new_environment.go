package ollama

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.HostEnvironment); s != "" {
		o = append(o, WithHost(s))
	}

	if s := environment.Optional(constant.PortEnvironment); s != "" {
		o = append(o, WithPort(strings.ToIntegerStrict(s)))
	}

	return New(o...)
}
