package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) ReadCommit(
	project int,
	sha string,
) *gitlab.Commit {
	result, _, e := c.client.Commits.GetCommit(
		project,
		sha,
		&gitlab.GetCommitOptions{},
	)
	errors.PanicOnError(e)

	return result
}
