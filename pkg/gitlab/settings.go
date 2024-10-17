package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) Settings() *gitlab.Settings {
	result, _, e := c.client.Settings.GetSettings()
	errors.PanicOnError(e)

	return result
}