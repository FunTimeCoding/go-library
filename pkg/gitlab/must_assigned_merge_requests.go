package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/merge_request"
)

func (c *Client) MustAssignedMergeRequests(all bool) []*merge_request.Request {
	result, e := c.AssignedMergeRequests(all)
	errors.PanicOnError(e)

	return result
}
