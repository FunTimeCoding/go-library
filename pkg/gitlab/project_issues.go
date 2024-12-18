package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/issue"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) ProjectIssues(project int) []*issue.Issue {
	result, _, e := c.client.Issues.ListProjectIssues(
		project,
		&gitlab.ListProjectIssuesOptions{ListOptions: constant.DefaultListOptions},
	)
	errors.PanicOnError(e)

	return issue.NewSlice(result)
}
