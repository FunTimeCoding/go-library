package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustImages(
	project int64,
	repository int64,
) []*gitlab.RegistryRepositoryTag {
	result, e := c.Images(project, repository)
	errors.PanicOnError(e)

	return result
}
