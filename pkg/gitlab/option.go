package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Option() *gitlab.Settings {
	result, _, e := c.client.Settings.GetSettings()
	errors.PanicOnError(e)

	return result
}
