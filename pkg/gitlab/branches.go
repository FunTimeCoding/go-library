package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/branch"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Branches(project int) []*branch.Branch {
	result, _, e := c.client.Branches.ListBranches(
		project,
		&gitlab.ListBranchesOptions{},
	)
	errors.PanicOnError(e)

	return branch.NewSlice(result)
}
