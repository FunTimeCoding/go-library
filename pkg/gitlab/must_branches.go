package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/branch"
)

func (c *Client) MustBranches(project int64) []*branch.Branch {
	result, e := c.Branches(project)
	errors.PanicOnError(e)

	return result
}
