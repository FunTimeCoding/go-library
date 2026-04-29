package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/release"
)

func (c *Client) MustReleases(
	namespace string,
	repository string,
) []*release.Release {
	result, e := c.Releases(namespace, repository)
	errors.PanicOnError(e)

	return result
}
