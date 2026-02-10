package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) Variables(project int64) []*gitlab.ProjectVariable {
	result, r, e := c.client.ProjectVariables.ListVariables(
		project,
		&gitlab.ListProjectVariablesOptions{},
	)

	if r != nil && r.StatusCode == 403 {
		// Do not panic
		return result
	}

	panicOnError(r, e)

	return result
}
