package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) ProjectByName(
	namespace string,
	name string,
) (*project.Project, error) {
	result, _, e := c.client.Projects.ListProjects(
		&gitlab.ListProjectsOptions{Search: &name},
	)

	if e != nil {
		return nil, e
	}

	count := len(result)
	var p *gitlab.Project

	if count == 0 {
		return nil, nil
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
			return nil, fmt.Errorf(
				"ambiguous project: %d results for %s/%s",
				count,
				namespace,
				name,
			)
		}
	}

	if p != nil && p.Namespace.Path == namespace {
		return project.New(p), nil
	}

	return nil, nil
}
