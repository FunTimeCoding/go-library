package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/project"

func (c *Client) FileExists(
	p *project.Project,
	branch string,
	file string,
) bool {
	return c.File(p.Identifier, branch, file) != nil
}
