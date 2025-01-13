package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/merge_request"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) AssignedReviews(all bool) []*merge_request.Request {
	requests, _, e := c.client.MergeRequests.ListMergeRequests(
		&gitlab.ListMergeRequestsOptions{
			ReviewerID:  gitlab.ReviewerID(c.user.ID),
			State:       ptr.To[string](constant.OpenedState),
			ListOptions: constant.DefaultListOptions,
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
