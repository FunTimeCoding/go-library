package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/merge_request"
	"gitlab.com/gitlab-org/api/client-go"
	"k8s.io/utils/ptr"
)

func (c *Client) ProjectMergeRequests(
	project int,
	all bool,
) []*merge_request.Request {
	var result []*gitlab.MergeRequest
	var number int

	for {
		o := &gitlab.ListProjectMergeRequestsOptions{
			State: ptr.To[string](OpenedState),
			ListOptions: gitlab.ListOptions{
				PerPage: PerPage100,
				Page:    number,
			},
		}

		if all {
			o.State = ptr.To[string]("all")
		}

		page, _, e := c.client.MergeRequests.ListProjectMergeRequests(
			project,
			o,
		)
		errors.PanicOnError(e)
		result = append(result, page...)

		if len(page) < PerPage100 {
			break
		}

		number++
	}

	return merge_request.NewSlice(result)
}
