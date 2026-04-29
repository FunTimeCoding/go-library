package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/merge_request"
)

func (c *Client) MustBranchRequest(
	project int64,
	branch string,
) *merge_request.Request {
	result, e := c.BranchRequest(project, branch)
	errors.PanicOnError(e)

	return result
}
