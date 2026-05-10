package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v86/github"
)

func (c *Client) MustPullRequests(
	owner string,
	name string,
) []*github.PullRequest {
	result, e := c.PullRequests(owner, name)
	errors.PanicOnError(e)

	return result
}
