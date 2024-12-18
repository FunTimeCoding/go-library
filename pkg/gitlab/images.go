package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/image"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Images(
	project int,
	repository int,
) []*gitlab.RegistryRepositoryTag {
	result, _, e := c.client.ContainerRegistry.ListRegistryRepositoryTags(
		project,
		repository,
		&gitlab.ListRegistryRepositoryTagsOptions{PerPage: constant.PerPage1000},
	)
	errors.PanicOnError(e)

	return image.Sort(result)
}
