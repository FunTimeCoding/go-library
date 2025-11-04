package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) SearchBlob(query string) []*gitlab.Blob {
	var result []*gitlab.Blob

	for _, p := range c.Projects() {
		page, r, e := c.client.Search.BlobsByProject(
			p.Identifier,
			query,
			&gitlab.SearchOptions{},
		)
		panicOnError(r, e)
		result = append(result, page...)
	}

	return result
}
