package gitlab

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustArtifactsFile(
	project int64,
	reference string,
	job string,
) string {
	result, e := c.ArtifactsFile(project, reference, job)
	errors.PanicOnError(e)

	return result
}
