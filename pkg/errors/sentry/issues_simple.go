package sentry

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"slices"
)

func (c *Client) IssuesSimple(verbose bool) []*issue.Issue {
	if o := environment.Optional(constant.OrganizationEnvironment); o != "" {
		var projects []response.Project

		if p := environment.Optional(constant.ProjectEnvironment); p != "" {
			names := split.Comma(p)

			for _, j := range c.OrganizationProjects(o) {
				if slices.Contains(names, j.Name) {
					projects = append(projects, j)
				}
			}
		} else {
			projects = append(projects, c.OrganizationProjects(o)...)
		}

		var result []*issue.Issue

		for _, p := range projects {
			if verbose {
				fmt.Printf("Project: %s\n", p.Name)
			}

			result = append(
				result,
				c.Issues(o, p.ID, constant.PeriodFortnight)...,
			)
		}

		return result
	}

	return c.AllIssues()
}
