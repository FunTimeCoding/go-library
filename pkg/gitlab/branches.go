package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) Branches(project int) []*gitlab.Branch {
	result, _, e := c.client.Branches.ListBranches(
		project,
		&gitlab.ListBranchesOptions{},
	)
	errors.PanicOnError(e)

	return result
}
