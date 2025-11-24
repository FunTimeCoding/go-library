package gitlab

import "github.com/funtimecoding/go-library/pkg/system"

func (c *Client) Trace(
	project int64,
	job int64,
) string {
	result, r, e := c.client.Jobs.GetTraceFile(project, job)
	panicOnError(r, e)

	return string(system.ReadAll(result))
}
