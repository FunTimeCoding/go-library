package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustTree(project int64) []*gitlab.TreeNode {
	result, e := c.Tree(project)
	errors.PanicOnError(e)

	return result
}
