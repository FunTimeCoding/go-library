package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) RegistryRepositories(project int) []*gitlab.RegistryRepository {
	result, _, e := c.client.ContainerRegistry.ListProjectRegistryRepositories(
		project,
		nil,
	)
	errors.PanicOnError(e)

	return result
}
