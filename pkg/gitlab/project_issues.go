package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/issue"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) ProjectIssues(project int) []*issue.Issue {
	result, _, e := c.client.Issues.ListProjectIssues(
		project,
		&gitlab.ListProjectIssuesOptions{ListOptions: DefaultListOptions},
	)
	errors.PanicOnError(e)

	return issue.NewSlice(result)
}
