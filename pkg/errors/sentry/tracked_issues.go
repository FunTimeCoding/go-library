package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
)

func (c *Client) TrackedIssues() []*issue.Issue {
	var result []*issue.Issue

	if c.organization == "" {
		return result
	}

	organization := c.Organization(c.organization)

	for _, p := range c.projects {
		result = append(
			result,
			c.Issues(
				organization,
				c.Project(organization, p),
				constant.PeriodFortnight,
			)...,
		)
	}

	return result
}
