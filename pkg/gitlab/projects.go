package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Projects() []*project.Project {
	var result []*project.Project
	var number int

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
		result = append(result, project.NewSlice(page)...)

		if len(page) < constant.PerPage100 {
			break
		}

		number++
	}

	return result
}
