package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/project"

func (c *Client) FileExists(
	p *project.Project,
	branch string,
	file string,
) (bool, error) {
	result, e := c.File(p.Identifier, branch, file)

	if e != nil {
		return false, e
	}

	return result != nil, nil
}
