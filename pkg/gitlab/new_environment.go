package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web"
)

func NewEnvironment() *Client {
	return New(
		web.TrimScheme(environment.Get(constant.HostEnvironment, 1)),
		environment.Get(constant.TokenEnvironment, 1),
		[]int{},
	)
}
