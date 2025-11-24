package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) MergeRequestCommits(
	project int64,
	request int64,
) []*commit.Commit {
	result, r, e := c.client.MergeRequests.GetMergeRequestCommits(
		project,
		request,
		&gitlab.GetMergeRequestCommitsOptions{},
	)
	panicOnError(r, e)

	return commit.NewSlice(result)
}
