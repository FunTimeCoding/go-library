package query

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
)

func Open(
	c *jira.Client,
	project string,
) []*issue.Issue {
	return c.Search(
		"project = '%s' AND status != %s ORDER BY key DESC",
		project,
		constant.Closed,
	)
}
