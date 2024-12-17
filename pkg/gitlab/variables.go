package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Variables(project int) []*gitlab.ProjectVariable {
	result, r, e := c.client.ProjectVariables.ListVariables(
		project,
		&gitlab.ListProjectVariablesOptions{},
	)

	if r != nil && r.StatusCode == 403 {
		return result
	}

	errors.PanicOnError(e)

	return result
}
