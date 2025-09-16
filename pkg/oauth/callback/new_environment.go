package callback

import (
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(verbose bool) *Server {
	return New(
		strings.ToIntegerStrict(
			environment.Default(PortEnvironment, "8080"),
		),
		verbose,
	)
}
