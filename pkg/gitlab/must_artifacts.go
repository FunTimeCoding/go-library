package gitlab

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustArtifacts(
	project int64,
	job int64,
) *bytes.Reader {
	result, e := c.Artifacts(project, job)
	errors.PanicOnError(e)

	return result
}
