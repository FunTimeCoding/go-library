package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) ArtifactsFile(
	project int,
	reference string,
	job string,
) string {
	// This only works for successful jobs
	reader, _, e := c.client.Jobs.DownloadArtifactsFile(
		project,
		reference,
		&gitlab.DownloadArtifactsFileOptions{Job: &job},
	)
	errors.PanicOnError(e)

	return string(system.ReadAll(reader))
}
