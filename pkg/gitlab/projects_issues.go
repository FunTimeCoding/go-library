package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/issue"

func (c *Client) ProjectsIssues() []*issue.Issue {
	var result []*issue.Issue

	for _, identifier := range c.projects {
		for _, element := range c.ProjectIssues(identifier) {
			if element.Done() {
				continue
			}

			result = append(result, element)
		}
	}

	return result
}
