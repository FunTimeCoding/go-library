package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/commit"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) MergeRequestCommits(
	project int,
	request int,
) []*commit.Commit {
	result, _, e := c.client.MergeRequests.GetMergeRequestCommits(
		project,
		request,
		&gitlab.GetMergeRequestCommitsOptions{},
	)
	errors.PanicOnError(e)

	return commit.NewSlice(result)
}
