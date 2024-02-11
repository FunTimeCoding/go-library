package gitlab

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) DeletePipeline(
	project int,
	pipeline int,
) {
	_, e := c.client.Pipelines.DeletePipeline(project, pipeline)
	errors.PanicOnError(e)
}
