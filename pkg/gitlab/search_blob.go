package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) SearchBlob(query string) []*gitlab.Blob {
	var result []*gitlab.Blob

	for _, p := range c.Projects() {
		page, _, e := c.client.Search.BlobsByProject(
			p.ID,
			query,
			&gitlab.SearchOptions{},
		)
		errors.PanicOnError(e)
		result = append(result, page...)
	}

	return result
}
