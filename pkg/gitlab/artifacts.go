package gitlab

import "bytes"

func (c *Client) Artifacts(
	project int,
	job int,
) *bytes.Reader {
	result, r, e := c.client.Jobs.GetJobArtifacts(project, job)
	panicOnError(r, e)

	return result
}
