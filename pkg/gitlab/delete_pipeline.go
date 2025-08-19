package gitlab

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) DeletePipeline(
	project int,
	pipeline int,
) {
	r, e := c.client.Pipelines.DeletePipeline(project, pipeline)

	if r != nil && r.StatusCode == 404 {
		// Do not fail if pipeline does not exist
		return
	}

	errors.PanicOnError(e)
}
