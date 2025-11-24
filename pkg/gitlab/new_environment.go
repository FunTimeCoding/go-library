package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.SliceInteger64(
		constant.GroupEnvironment,
	); len(s) > 0 {
		o = append(o, WithGroups(s))
	}

	return New(
		web.TrimScheme(environment.Required(constant.HostEnvironment)),
		environment.Required(constant.TokenEnvironment),
		o...,
	)
}
