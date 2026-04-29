package gitlab

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeletePipeline(
	project int64,
	pipeline int64,
) {
	errors.PanicOnError(c.DeletePipeline(project, pipeline))
}
