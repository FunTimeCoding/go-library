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
	var start int
	var result []*issue.Issue

	if len(a) > 0 {
		query = fmt.Sprintf(query, a)
	}

	for {
		page, _, e := c.client.Issue.Search(
			query,
			&jira.SearchOptions{
				StartAt:    start,
				MaxResults: constant.SearchLimit,
			},
		)
		errors.PanicOnError(e)
		result = append(result, issue.NewSlice(page, c.IssueOption())...)

		if len(page) < constant.SearchLimit {
			break
		}

		start += constant.SearchLimit
	}

	return result
}
