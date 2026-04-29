package github

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeleteTag(
	owner string,
	repository string,
	name string,
) {
	errors.PanicOnError(c.DeleteTag(owner, repository, name))
}
