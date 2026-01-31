package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/helper"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
)

func (c *Client) Issues(
	o sentry.Organization,
	p sentry.Project,
	period string,
) []*issue.Issue {
	helper.ValidateContains(constant.Periods, period)
	result, _, e := c.client.GetIssues(
		o,
		p,
		&period,
		nil,
		&constant.UnresolvedFilter,
	)
	errors.PanicOnError(e)

	return issue.NewSlice(result)
}
