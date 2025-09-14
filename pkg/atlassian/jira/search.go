package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Search(
	query string,
	a ...any,
) []*issue.Issue {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	var result []*issue.Issue
	var token string

	for {
		page, r, e := c.client.Issue.SearchV2JQL(
			query,
			&jira.SearchOptionsV2{
				NextPageToken: token,
				Fields:        []string{"*all"},
				MaxResults:    constant.SearchLimit,
			},
		)
		errors.PanicOnError(e)
		token = r.NextPageToken
		result = append(result, issue.NewSlice(page, c.IssueOption())...)

		if r.IsLast {
			break
		}
	}

	return c.enrichMany(result)
}
