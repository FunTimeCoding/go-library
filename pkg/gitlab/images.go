package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/image"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Images(
	project int64,
	repository int64,
) []*gitlab.RegistryRepositoryTag {
	result, r, e := c.client.ContainerRegistry.ListRegistryRepositoryTags(
		project,
		repository,
		&gitlab.ListRegistryRepositoryTagsOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: constant.PerPage1000,
			},
		},
	)
	panicOnError(r, e)

	return image.Sort(result)
}
