package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Projects() []*gitlab.Project {
	var result []*gitlab.Project
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
		result = append(result, page...)

		if len(page) < constant.PerPage100 {
			break
		}

		number++
	}

	return result
}
