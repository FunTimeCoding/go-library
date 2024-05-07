package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) ReadCommit(
	project int,
	sha string,
) *gitlab.Commit {
	result, _, e := c.client.Commits.GetCommit(project, sha)
	errors.PanicOnError(e)

	return result
}
