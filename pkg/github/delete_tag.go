package github

import "fmt"

func (c *Client) DeleteTag(
	owner string,
	repository string,
	name string,
) {
	r, e := c.client.Git.DeleteRef(
		c.context,
		owner,
		repository,
		fmt.Sprintf("refs/tags/%s", name),
	)
	panicOnError(r, e)
}
