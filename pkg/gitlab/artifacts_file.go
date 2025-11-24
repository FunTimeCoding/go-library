package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) ArtifactsFile(
	project int64,
	reference string,
	job string,
) string {
	// This only works for successful jobs
	reader, r, e := c.client.Jobs.DownloadArtifactsFile(
		project,
		reference,
		&gitlab.DownloadArtifactsFileOptions{Job: &job},
	)
	panicOnError(r, e)

	return string(system.ReadAll(reader))
}
