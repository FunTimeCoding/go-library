package sentry

import "github.com/funtimecoding/go-library/pkg/sentry/issue"

func (c *Client) AllIssues() []*issue.Issue {
	var result []*issue.Issue

	for _, o := range c.Organizations() {
		for _, p := range c.OrganizationProjects(o) {
			result = append(result, c.Issues(o, p, PeriodFortnight)...)
		}
	}

	return result
}
