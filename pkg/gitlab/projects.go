package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Projects() []*project.Project {
	var result []*gitlab.Project
	var number int64

	if len(c.groups) > 0 {
		return c.GroupProjects(c.groups...)
	}

	for {
		page, r, e := c.client.Projects.ListProjects(
			&gitlab.ListProjectsOptions{
				ListOptions: gitlab.ListOptions{
					PerPage: constant.PerPage100,
					Page:    number,
				},
			},
		)
		panicOnError(r, e)
		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		number++
	}

	return project.NewSlice(result)
}
