package query

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func Open(
	c *jira.Client,
	project string,
) []*issue.Issue {
	return c.Search(
		"project = '%s' AND status NOT IN (%s) ORDER BY key DESC",
		project,
		join.Comma(Quote(c.IssueOption().ClosedStatus)),
	)
}
