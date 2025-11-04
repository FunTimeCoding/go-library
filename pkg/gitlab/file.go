package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) File(
	project int,
	branch string,
	name string,
) *gitlab.File {
	result, r, e := c.client.RepositoryFiles.GetFile(
		project,
		name,
		&gitlab.GetFileOptions{Ref: &branch},
	)

	if r != nil && r.StatusCode == 404 {
		// Do not panic
		return nil
	}

	panicOnError(r, e)

	return result
}
