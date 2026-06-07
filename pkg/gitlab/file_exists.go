package gitlab

import (
	"errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) FileExists(
	p *project.Project,
	branch string,
	file string,
) (bool, error) {
	_, e := c.File(p.Identifier, branch, file)

	if e != nil {
		if errors.Is(e, constant.ErrorNotFound) {
			return false, nil
		}

		return false, e
	}

	return true, nil
}
