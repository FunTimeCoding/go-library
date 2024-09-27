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
					PerPage: PerPage1000,
					Page:    number,
				},
			},
		)
		errors.PanicOnError(e)
		result = append(result, page...)

		if len(page) < PerPage1000 {
			break
		}

		number++
	}

	return result
}
