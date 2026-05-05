package callback

import (
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web/constant"
)

func NewEnvironment(verbose bool) *Server {
	return New(
		environment.FallbackInteger(PortEnvironment, constant.ListenPort),
		verbose,
	)
}
