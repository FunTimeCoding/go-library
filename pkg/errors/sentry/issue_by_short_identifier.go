package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
	"strings"
)

func (c *Client) IssueByShortIdentifier(
	organization string,
	identifier string,
) *issue.Issue {
	var result []response.Issue
	errors.PanicOnError(
		json.Unmarshal(
			c.basic.GetBytes(
				fmt.Sprintf(
					"projects/%s/%s/issues",
					organization,
					strings.ToLower(
						strings.SplitN(identifier, "-", 2)[0],
					),
				),
				map[string]string{"shortIdLookup": "1", "query": identifier},
			),
			&result,
		),
	)

	if len(result) == 0 {
		return nil
	}

	return issue.New(&result[0])
}
