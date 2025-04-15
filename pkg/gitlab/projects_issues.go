package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/issue"

func (c *Client) ProjectsIssues() []*issue.Issue {
	var result []*issue.Issue

	for _, identifier := range c.projects {
		for _, e := range c.ProjectIssues(identifier) {
			if e.Done() {
				continue
			}

			result = append(result, e)
		}
	}

	return result
}
