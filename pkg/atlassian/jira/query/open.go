package query

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
)

func Open(c *jira.Client) []*issue.Issue {
	return c.Search("status != Closed")
}
