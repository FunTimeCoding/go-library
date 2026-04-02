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

	for _, p := range c.projects {
		proj := c.Project(c.organization, p)
		result = append(
			result,
			c.Issues(c.organization, proj.ID, constant.PeriodFortnight)...,
		)
	}

	return result
}
