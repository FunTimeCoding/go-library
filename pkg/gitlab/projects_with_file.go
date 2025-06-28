package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) ProjectsWithFile(path string) []*project.Project {
	var result []*project.Project

	for _, p := range c.Projects() {
		if false {
			fmt.Printf("Project: %s\n", p.Raw.NameWithNamespace)
		}

		for _, n := range c.Tree(p.Identifier) {
			if path == n.Path {
				result = append(result, p)

				break
			}
		}
	}

	return result
}
