package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/sentry/issue"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"slices"
)

func (c *Client) IssuesSimple() []*issue.Issue {
	if o := environment.GetDefault(
		constant.OrganizationEnvironment,
		"",
	); o != "" {
		r := c.Organization(o)
		var projects []sentry.Project

		if p := environment.GetDefault(
			constant.ProjectEnvironment,
			"",
		); p != "" {
			names := split.Comma(p)

			for _, j := range c.OrganizationProjects(r) {
				if slices.Contains(names, j.Name) {
					projects = append(projects, j)
				}
			}
		} else {
			projects = append(projects, c.OrganizationProjects(r)...)
		}

		var result []*issue.Issue

		for _, p := range projects {
			result = append(
				result,
				c.Issues(r, p, constant.PeriodFortnight)...,
			)
		}

		return result
	}

	return c.AllIssues()
}
