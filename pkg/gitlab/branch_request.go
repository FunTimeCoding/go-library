package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/merge_request"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) BranchRequest(
	project int64,
	branch string,
) *merge_request.Request {
	result, r, e := c.client.MergeRequests.ListProjectMergeRequests(
		project,
		&gitlab.ListProjectMergeRequestsOptions{
			SourceBranch: ptr.To(branch),
			ListOptions:  gitlab.ListOptions{PerPage: 1},
		},
	)

	if r != nil && r.StatusCode == 404 {
		// Project not found, do not panic

		return nil
	}

	panicOnError(r, e)

	if len(result) == 0 {
		// Branch not found, do not panic

		return nil
	}

	return merge_request.New(result[0])
}
