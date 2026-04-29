package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) File(
	project int64,
	branch string,
	name string,
) (*gitlab.File, error) {
	result, r, e := c.client.RepositoryFiles.GetFile(
		project,
		name,
		&gitlab.GetFileOptions{Ref: &branch},
	)

	if r != nil && r.StatusCode == 404 {
		// Do not panic
		return nil, nil
	}

	return result, e
}
