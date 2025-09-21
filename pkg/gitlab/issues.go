package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/issue"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Issues() []*issue.Issue {
	result, _, e := c.client.Issues.ListIssues(
		&gitlab.ListIssuesOptions{
			ListOptions: constant.DefaultListOptions,
		},
	)
	errors.PanicOnError(e)

	return issue.NewSlice(result)
}
