package github

import "fmt"

func (c *Client) DeleteTag(
	owner string,
	repository string,
	name string,
) error {
	_, e := c.client.Git.DeleteRef(
		c.context,
		owner,
		repository,
		fmt.Sprintf("refs/tags/%s", name),
	)

	return e
}
