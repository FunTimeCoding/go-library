package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
)

func (c *Client) AllIssues() []*issue.Issue {
	var result []*issue.Issue

	for _, o := range c.Organizations() {
		for _, p := range c.OrganizationProjects(o.Slug) {
			result = append(
				result,
				c.Issues(o.Slug, p.ID, constant.PeriodFortnight)...,
			)
		}
	}

	return result
}
