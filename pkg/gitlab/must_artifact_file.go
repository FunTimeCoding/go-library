package gitlab

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustArtifactFile(
	project int64,
	job int64,
	path string,
) *bytes.Reader {
	result, e := c.ArtifactFile(project, job, path)
	errors.PanicOnError(e)

	return result
}
