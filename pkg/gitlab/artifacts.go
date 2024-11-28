package gitlab

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Artifacts(
	project int,
	job int,
) *bytes.Reader {
	result, _, e := c.client.Jobs.GetJobArtifacts(project, job)
	errors.PanicOnError(e)

	return result
}
