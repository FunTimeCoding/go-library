package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
)

func (c *Client) Trace(
	project int,
	job int,
) string {
	result, _, e := c.client.Jobs.GetTraceFile(project, job)
	errors.PanicOnError(e)

	return string(system.ReadAll(result))
}
