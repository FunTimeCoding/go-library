package ollama

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web"
)

func NewEnvironment() *Client {
	return New(
		environment.GetDefault(
			constant.HostEnvironment,
			web.Localhost,
		),
	)
}
