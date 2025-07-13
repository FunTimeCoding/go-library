package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...OptionFunc) *Client {
	return New(
		environment.Get(constant.HostEnvironment),
		environment.Get(constant.UserEnvironment),
		environment.Get(constant.TokenEnvironment),
		o...,
	)
}
