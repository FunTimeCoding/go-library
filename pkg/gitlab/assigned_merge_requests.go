package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/merge_request"
	"github.com/xanzy/go-gitlab"
	"k8s.io/utils/ptr"
)

func (c *Client) AssignedMergeRequests(all bool) []*merge_request.Request {
	requests, _, e := c.client.MergeRequests.ListMergeRequests(
		&gitlab.ListMergeRequestsOptions{
			AssigneeID:  gitlab.AssigneeID(c.user.ID),
			State:       ptr.To[string](OpenedState),
			ListOptions: DefaultListOptions,
		},
		nil,
	)
	errors.PanicOnError(e)
	result := merge_request.NewSlice(requests)

	if all {
		return result
	}

	return merge_request.FilterDone(result)
}
