package sentry

import "github.com/funtimecoding/go-library/pkg/sentry/issue"

func (c *Client) TrackedIssues() []*issue.Issue {
	var result []*issue.Issue

	if c.organization == "" {
		return result
	}

	organization := c.Organization(c.organization)

	for _, element := range c.projects {
		result = append(
			result,
			c.Issues(
				organization,
				c.Project(organization, element),
				PeriodFortnight,
			)...,
		)
	}

	return result
}
