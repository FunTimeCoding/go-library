package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) ReadCommit(
	project int,
	sha string,
) *gitlab.Commit {
	result, r, e := c.client.Commits.GetCommit(
		project,
		sha,
		&gitlab.GetCommitOptions{},
	)
	panicOnError(r, e)

	return result
}
