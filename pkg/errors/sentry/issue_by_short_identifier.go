package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
	"strings"
)

func (c *Client) IssueByShortIdentifier(
	o sentry.Organization,
	identifier string,
) *issue.Issue {
	b := c.basic.GetBytes(
		fmt.Sprintf(
			"projects/%s/%s/issues",
			*o.Slug,
			strings.ToLower(strings.SplitN(identifier, "-", 2)[0]),
		),
		map[string]string{"shortIdLookup": "1", "query": identifier},
	)
	var result []sentry.Issue
	errors.PanicOnError(json.Unmarshal(b, &result))

	if len(result) == 0 {
		return nil
	}

	return issue.New(&result[0])
}
