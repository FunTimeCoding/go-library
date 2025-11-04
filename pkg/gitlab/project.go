package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) Project(identifier int) *project.Project {
	if p, okay := c.projectCache[identifier]; okay {
		return p
	}

	result, r, e := c.client.Projects.GetProject(identifier, nil)

	if e != nil {
		fmt.Printf("Project fail: %v\n", identifier)
	}

	panicOnError(r, e)
	c.projectCache[identifier] = project.New(result)

	return c.projectCache[identifier]
}
