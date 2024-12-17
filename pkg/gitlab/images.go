package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Images(
	project int,
	repository int,
) []*gitlab.RegistryRepositoryTag {
	result, _, e := c.client.ContainerRegistry.ListRegistryRepositoryTags(
		project,
		repository,
		&gitlab.ListRegistryRepositoryTagsOptions{PerPage: PerPage1000},
	)
	errors.PanicOnError(e)

	return result
}
