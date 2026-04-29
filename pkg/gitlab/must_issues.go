package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/issue"
)

func (c *Client) MustIssues() []*issue.Issue {
	result, e := c.Issues()
	errors.PanicOnError(e)

	return result
}
