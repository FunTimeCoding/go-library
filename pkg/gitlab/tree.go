package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) Tree(project int) []*gitlab.TreeNode {
	result, r, e := c.client.Repositories.ListTree(
		project,
		&gitlab.ListTreeOptions{},
	)

	if r != nil && r.StatusCode == 404 {
		// Do not panic
		return []*gitlab.TreeNode{}
	}

	panicOnError(r, e)

	return result
}
