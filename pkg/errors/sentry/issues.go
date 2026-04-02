package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/helper"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
)

func (c *Client) Issues(
	organization string,
	projectIdentifier string,
	period string,
) []*issue.Issue {
	helper.ValidateContains(constant.Periods, period)
	query := map[string]string{
		"project": projectIdentifier,
		"query":   constant.UnresolvedFilter,
	}

	if period != "" {
		query["statsPeriod"] = period
	}

	var result []response.Issue
	errors.PanicOnError(
		json.Unmarshal(
			c.basic.GetBytes(
				fmt.Sprintf("organizations/%s/issues", organization),
				query,
			),
			&result,
		),
	)

	return issue.NewSlice(result)
}
