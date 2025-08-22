package ollama

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...OptionFunc) *Client {
	if s := environment.GetDefault(
		constant.HostEnvironment,
		"",
	); s != "" {
		o = append(o, WithHost(s))
	}

	if s := environment.GetDefault(
		constant.PortEnvironment,
		"",
	); s != "" {
		o = append(o, WithPort(strings.ToIntegerStrict(s)))
	}

	return New(o...)
}
