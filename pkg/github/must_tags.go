package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v86/github"
)

func (c *Client) MustTags(
	owner string,
	repository string,
) []*github.RepositoryTag {
	result, e := c.Tags(owner, repository)
	errors.PanicOnError(e)

	return result
}
