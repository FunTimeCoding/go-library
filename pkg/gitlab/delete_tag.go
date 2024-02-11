package gitlab

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) DeleteTag(
	project int,
	name string,
) {
	_, e := c.client.Tags.DeleteTag(project, name)
	errors.PanicOnError(e)
}
