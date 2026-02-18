package github

import (
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/github/pull_request"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/google/go-github/v83/github"
)

func (c *Client) BranchRequest(
	owner string,
	repository string,
	branch string,
) *pull_request.Request {
	result, r, e := c.client.PullRequests.List(
		c.context,
		owner,
		repository,
		&github.PullRequestListOptions{
			State:       constant.All,
			Head:        join.Colon(owner, branch),
			ListOptions: github.ListOptions{PerPage: 1},
		},
	)

	if r != nil && r.StatusCode == 404 {
		// Repository not found, do not panic

		return nil
	}

	panicOnError(r, e)

	if len(result) == 0 {
		// Branch not found, do not panic

		return nil
	}

	return pull_request.New(result[0])
}
