package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"github.com/gpustack/gguf-parser-go/util/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) GroupProjects(identifier ...int) []*project.Project {
	var result []*gitlab.Project

	for _, i := range identifier {
		var number int

		for {
			page, r, e := c.client.Groups.ListGroupProjects(
				i,
				&gitlab.ListGroupProjectsOptions{
					ListOptions: gitlab.ListOptions{
						PerPage: constant.PerPage100,
						Page:    number,
					},
					Archived: ptr.Bool(false),
				},
			)
			panicOnError(r, e)
			result = append(result, page...)

			if len(page) < constant.PerPage100 {
				break
			}

			number++
		}
	}

	return project.NewSlice(result)
}
