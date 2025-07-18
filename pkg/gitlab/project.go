package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) Project(identifier int) *project.Project {
	if p, okay := c.projectCache[identifier]; okay {
		return p
	}

	result, _, e := c.client.Projects.GetProject(identifier, nil)

	if e != nil {
		fmt.Printf("Project fail: %v\n", identifier)
	}

	errors.PanicOnError(e)
	c.projectCache[identifier] = project.New(result)

	return c.projectCache[identifier]
}
