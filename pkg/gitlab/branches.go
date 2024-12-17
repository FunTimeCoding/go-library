package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Branches(project int) []*gitlab.Branch {
	result, _, e := c.client.Branches.ListBranches(
		project,
		&gitlab.ListBranchesOptions{},
	)
	errors.PanicOnError(e)

	return result
}
