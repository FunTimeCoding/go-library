package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"github.com/xanzy/go-gitlab"
	"log"
)

func (c *Client) ProjectByName(
	namespace string,
	name string,
) *project.Project {
	result, _, e := c.client.Projects.ListProjects(
		&gitlab.ListProjectsOptions{Search: &name},
	)
	errors.PanicOnError(e)
	count := len(result)
	var p *gitlab.Project

	if count == 0 {
		return nil
	} else if count == 1 {
		p = result[0]
	} else if count > 1 {
		// Name could also be partial of another repository, with a longer name
		for _, element := range result {
			if p == nil && element.Name == name {
				p = element
			} else if p != nil && element.Name == name {
				// A second one with the same name is a problem, should not be possible
				p = element
			}
		}

		if p == nil {
			log.Panicf("unexpected: %d", count)
		}
	}

	if p.Namespace.Path == namespace {
		return project.New(p)
	}

	return nil
}
