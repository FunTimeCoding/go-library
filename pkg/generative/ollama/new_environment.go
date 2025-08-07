package ollama

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
)

func NewEnvironment() *Client {
	return New(
		environment.GetDefault(constant.HostEnvironment, web.Localhost),
	)
}
