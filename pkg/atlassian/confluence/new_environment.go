package confluence

import (
	confluence "github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Default(
		confluence.DefaultSpaceEnvironment,
		"",
	); s != "" {
		o = append(o, WithDefaultSpace(s))
	}

	if s := environment.Default(
		confluence.DefaultPageEnvironment,
		"",
	); s != "" {
		o = append(o, WithDefaultPage(s))
	}

	if v := environment.Slice(confluence.LabelEnvironment); len(v) > 0 {
		o = append(o, WithLabel(v))
	}

	return New(
		environment.Exit(constant.HostEnvironment),
		environment.Exit(constant.UserEnvironment),
		environment.Exit(constant.TokenEnvironment),
		o...,
	)
}
