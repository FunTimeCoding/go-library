package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/merge_request"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) ProjectMergeRequests(
	project int64,
	all bool,
) []*merge_request.Request {
	var result []*gitlab.BasicMergeRequest
	var number int64

	for {
		o := &gitlab.ListProjectMergeRequestsOptions{
			State: ptr.To[string](constant.OpenedState),
			ListOptions: gitlab.ListOptions{
				PerPage: constant.PerPage100,
				Page:    number,
			},
		}

		if all {
			o.State = ptr.To[string]("all")
		}

		page, r, e := c.client.MergeRequests.ListProjectMergeRequests(
			project,
			o,
		)
		panicOnError(r, e)
		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		number++
	}

	return merge_request.NewSlice(result)
}
