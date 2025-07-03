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
		r := c.Organization(environment.Get(constant.OrganizationEnvironment))
		var projects []sentry.Project

		if s := environment.GetDefault(
			constant.ProjectEnvironment,
			"",
		); s != "" {
			names := split.Comma(s)

			for _, p := range c.OrganizationProjects(r) {
				if slices.Contains(names, p.Name) {
					projects = append(projects, p)
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
