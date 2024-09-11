package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/merge_request"
	"github.com/xanzy/go-gitlab"
	"k8s.io/utils/ptr"
)

func (c *Client) ProjectMergeRequests(
	project int,
) []*merge_request.Request {
	result, _, e := c.client.MergeRequests.ListProjectMergeRequests(
		project,
		&gitlab.ListProjectMergeRequestsOptions{
			State:       ptr.To[string](OpenedState),
			ListOptions: DefaultListOptions,
		},
	)
	errors.PanicOnError(e)

	return merge_request.NewSlice(result)
}
