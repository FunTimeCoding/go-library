package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

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

	if e != nil && r != nil && r.StatusCode == 404 {
		return nil
	}

	errors.PanicOnError(e)

	return result
}
