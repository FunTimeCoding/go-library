package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) Projects() []*gitlab.Project {
	var result []*gitlab.Project
	var number int

	for {
		page, _, e := c.client.Projects.ListProjects(
			&gitlab.ListProjectsOptions{
				ListOptions: gitlab.ListOptions{
					PerPage: PerPage,
					Page:    number,
				},
			},
		)
		errors.PanicOnError(e)
		result = append(result, page...)

		if len(page) < PerPage {
			break
		}

		number++
	}

	return result
}
