package gitlab

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) DeleteImage(
	project int,
	repository int,
	tag string,
) {
	_, e := c.client.ContainerRegistry.DeleteRegistryRepositoryTag(
		project,
		repository,
		tag,
	)
	errors.PanicOnError(e)
}
