package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/constant"
	jira "github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...OptionFunc) *Client {
	if s := environment.GetDefault(
		jira.DefaultProjectKeyEnvironment,
		"",
	); s != "" {
		o = append(o, WithDefaultProjectKey(s))
	}

	if s := environment.GetDefault(
		jira.DefaultProjectNameEnvironment,
		"",
	); s != "" {
		o = append(o, WithDefaultProjectName(s))
	}

	if s := environment.GetDefault(
		jira.DefaultIssueTypeEnvironment,
		"",
	); s != "" {
		o = append(o, WithDefaultIssueType(s))
	}

	return New(
		environment.Get(constant.HostEnvironment),
		environment.Get(constant.UserEnvironment),
		environment.Get(constant.TokenEnvironment),
		o...,
	)
}
