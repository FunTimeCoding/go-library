package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web"
)

func NewEnvironment() *Client {
	return New(
		web.TrimScheme(environment.Get(HostEnvironment, 1)),
		environment.Get(TokenEnvironment, 1),
		[]int{},
	)
}
