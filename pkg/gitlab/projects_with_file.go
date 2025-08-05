package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"strings"
)

func (c *Client) ProjectsWithFile(
	path string,
	caseInsensitive bool,
) []*project.Project {
	var result []*project.Project

	if caseInsensitive {
		path = strings.ToLower(path)
	}

	for _, p := range c.Projects() {
		if c.verbose {
			fmt.Printf("Project: %s\n", p.Raw.NameWithNamespace)
		}

		for _, n := range c.Tree(p.Identifier) {
			if path == n.Path ||
				(caseInsensitive && path == strings.ToLower(n.Path)) {
				result = append(result, p)

				break
			}
		}
	}

	return result
}
