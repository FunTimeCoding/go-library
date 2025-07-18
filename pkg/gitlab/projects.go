package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Projects() []*project.Project {
	var result []*gitlab.Project
	var number int

	if len(c.groups) > 0 {
		return c.GroupProjects(c.groups...)
	}

	for {
		page, _, e := c.client.Projects.ListProjects(
			&gitlab.ListProjectsOptions{
				ListOptions: gitlab.ListOptions{
					PerPage: constant.PerPage100,
					Page:    number,
				},
			},
		)
		errors.PanicOnError(e)
		result = append(result, page...)

		if len(page) < constant.PerPage100 {
			break
		}

		number++
	}

	return project.NewSlice(result)
}
