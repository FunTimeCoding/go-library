package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

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

	result, _, e := c.client.Commits.CreateCommit(
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
	errors.PanicOnError(e)

	return result
}
