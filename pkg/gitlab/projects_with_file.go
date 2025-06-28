package gitlab

import (
	"fmt"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) ProjectsWithFile(path string) []*gitlab.Project {
	var result []*gitlab.Project

	for _, p := range c.Projects() {
		if false {
			fmt.Printf("Project: %s\n", p.NameWithNamespace)
		}

		for _, n := range c.Tree(p.ID) {
			if path == n.Path {
				result = append(result, p)

				break
			}
		}
	}

	return result
}
