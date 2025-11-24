package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/branch"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Branches(project int64) []*branch.Branch {
	result, r, e := c.client.Branches.ListBranches(
		project,
		&gitlab.ListBranchesOptions{},
	)
	panicOnError(r, e)

	return branch.NewSlice(result)
}
