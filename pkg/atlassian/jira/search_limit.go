package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) SearchLimit(
	limit int,
	query string,
	a ...any,
) []*issue.Issue {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	page, _, e := c.client.Issue.Search(
		query,
		&jira.SearchOptions{MaxResults: limit},
	)
	errors.PanicOnError(e)

	result := issue.NewSlice(page, c.IssueOption())

	return c.enrichMany(result)
}
