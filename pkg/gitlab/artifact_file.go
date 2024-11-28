package gitlab

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) ArtifactFile(
	project int,
	job int,
	path string,
) *bytes.Reader {
	result, _, e := c.client.Jobs.DownloadSingleArtifactsFile(
		project,
		job,
		path,
	)
	errors.PanicOnError(e)

	return result
}
