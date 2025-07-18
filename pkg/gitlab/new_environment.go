package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web"
)

func NewEnvironment(o ...OptionFunc) *Client {
	if s := environment.GetSliceInteger(
		constant.GroupEnvironment,
	); len(s) > 0 {
		o = append(o, WithGroups(s))
	}

	return New(
		web.TrimScheme(environment.Get(constant.HostEnvironment)),
		environment.Get(constant.TokenEnvironment),
		o...,
	)
}
