package gitlab

import "bytes"

func (c *Client) ArtifactFile(
	project int64,
	job int64,
	path string,
) *bytes.Reader {
	result, r, e := c.client.Jobs.DownloadSingleArtifactsFile(
		project,
		job,
		path,
	)
	panicOnError(r, e)

	return result
}
