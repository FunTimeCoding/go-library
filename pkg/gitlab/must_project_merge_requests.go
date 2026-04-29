package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/merge_request"
)

func (c *Client) MustProjectMergeRequests(
	project int64,
	all bool,
) []*merge_request.Request {
	result, e := c.ProjectMergeRequests(project, all)
	errors.PanicOnError(e)

	return result
}
