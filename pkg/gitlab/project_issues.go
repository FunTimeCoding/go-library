package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/issue"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) ProjectIssues(project int64) []*issue.Issue {
	result, r, e := c.client.Issues.ListProjectIssues(
		project,
		&gitlab.ListProjectIssuesOptions{
			ListOptions: constant.DefaultListOptions,
		},
	)
	panicOnError(r, e)

	return issue.NewSlice(result)
}
