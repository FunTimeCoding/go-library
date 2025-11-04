package github

import "github.com/funtimecoding/go-library/pkg/github/release"

func (c *Client) Releases(
	namespace string,
	repository string,
) []*release.Release {
	result, r, e := c.client.Repositories.ListReleases(
		c.context,
		namespace,
		repository,
		nil,
	)
	panicOnError(r, e)

	return release.NewSlice(result)
}
