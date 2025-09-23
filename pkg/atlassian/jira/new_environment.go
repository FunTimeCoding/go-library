package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/constant"
	jira "github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Fallback(
		jira.DefaultProjectKeyEnvironment,
		"",
	); s != "" {
		o = append(o, WithDefaultProjectKey(s))
	}

	if s := environment.Fallback(
		jira.DefaultProjectNameEnvironment,
		"",
	); s != "" {
		o = append(o, WithDefaultProjectName(s))
	}

	if s := environment.Fallback(
		jira.DefaultIssueTypeEnvironment,
		"",
	); s != "" {
		o = append(o, WithDefaultIssueType(s))
	}

	if v := environment.Slice(jira.ClosedStatusEnvironment); len(v) > 0 {
		o = append(o, WithClosedStatus(v))
	}

	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.TokenEnvironment),
		o...,
	)
}
