package github

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) DeleteTag(
	owner string,
	repository string,
	name string,
) {
	_, e := c.client.Git.DeleteRef(
		c.context,
		owner,
		repository,
		fmt.Sprintf("refs/tags/%s", name),
	)
	errors.PanicOnError(e)
}
