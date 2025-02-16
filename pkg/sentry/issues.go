package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sentry/helper"
	"github.com/funtimecoding/go-library/pkg/sentry/issue"
)

func (c *Client) Issues(
	organization sentry.Organization,
	project sentry.Project,
	period string,
) []*issue.Issue {
	helper.ValidateContains(Periods, period)
	result, _, e := c.client.GetIssues(
		organization,
		project,
		&period,
		nil,
		&UnresolvedFilter,
	)
	errors.PanicOnError(e)

	return issue.NewSlice(result)
}
