package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) AdminMode(on bool) *gitlab.Settings {
	// AdminMode is persisted
	result, _, e := c.client.Settings.UpdateSettings(
		&gitlab.UpdateSettingsOptions{AdminMode: ptr.To[bool](on)},
	)
	errors.PanicOnError(e)

	return result
}
