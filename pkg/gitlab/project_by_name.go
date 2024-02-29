package gitlab

import (
	"fmt"
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
			if element.Namespace.Name == namespace && element.Name == name {
				fmt.Printf(
					"match: %s/%s\n",
					element.Namespace.Name,
					element.Name,
				)
				p = element

				break
			} else {
				fmt.Printf(
					"warning: not matching %s/%s\n",
					element.Namespace.Name,
					element.Name,
				)
			}
		}

		if p == nil {
			log.Panicf("unexpected: %d %+v", count, result)
		}
	}

	if p.Namespace.Path == namespace {
		return project.New(p)
	}

	return nil
}
