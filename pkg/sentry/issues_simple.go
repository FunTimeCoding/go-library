package sentry

import (
	"github.com/funtimecoding/go-library/pkg/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/sentry/issue"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func (c *Client) IssuesSimple() []*issue.Issue {
	o := c.Organization(
		environment.Get(constant.OrganizationEnvironment, 1),
	)

	return c.Issues(
		o,
		c.Project(o, environment.Get(constant.ProjectEnvironment, 1)),
		constant.PeriodFortnight,
	)
}
