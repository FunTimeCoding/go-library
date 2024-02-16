package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/release"
)

func (c *Client) Releases(
	namespace string,
	repository string,
) []*release.Release {
	result, _, e := c.client.Repositories.ListReleases(
		c.context,
		namespace,
		repository,
		nil,
	)
	errors.PanicOnError(e)

	return release.NewSlice(result)
}
