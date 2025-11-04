package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go"
	"log"
)

func (c *Client) ProjectByName(
	namespace string,
	name string,
) *project.Project {
	result, r, e := c.client.Projects.ListProjects(
		&gitlab.ListProjectsOptions{Search: &name},
	)
	panicOnError(r, e)
	count := len(result)
	var p *gitlab.Project

	if count == 0 {
		return nil
	} else if count == 1 {
		p = result[0]
	} else if count > 1 {
		// Name could also be partial of another repository, with a longer name
		for _, l := range result {
			if l.Namespace.Path == namespace && l.Name == name {
				p = l

				break
			}
		}

		if p == nil {
			log.Panicf("unexpected: %d %+v", count, result)
		}
	}

	if p != nil && p.Namespace.Path == namespace {
		return project.New(p)
	}

	return nil
}
