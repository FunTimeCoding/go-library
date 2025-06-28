package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Tree(project int) []*gitlab.TreeNode {
	result, r, e := c.client.Repositories.ListTree(
		project,
		&gitlab.ListTreeOptions{},
	)

	if r != nil && r.StatusCode == 404 {
		return []*gitlab.TreeNode{}
	}

	errors.PanicOnError(e)

	return result
}
