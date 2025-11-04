package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) Commit(
	project int,
	branch string,
	text string,
	path string,
	content string,
	update bool,
) *gitlab.Commit {
	var action gitlab.FileActionValue

	if update {
		action = gitlab.FileUpdate
	} else {
		action = gitlab.FileCreate
	}

	result, r, e := c.client.Commits.CreateCommit(
		project,
		&gitlab.CreateCommitOptions{
			Branch:        &branch,
			CommitMessage: &text,
			Actions: []*gitlab.CommitActionOptions{
				{
					Action:   &action,
					FilePath: &path,
					Content:  &content,
				},
			},
		},
	)
	panicOnError(r, e)

	return result
}
