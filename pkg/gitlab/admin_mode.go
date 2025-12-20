package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) AdminMode(on bool) *gitlab.Settings {
	// AdminMode is persisted
	result, r, e := c.client.Settings.UpdateSettings(
		&gitlab.UpdateSettingsOptions{AdminMode: ptr.To(on)},
	)
	panicOnError(r, e)

	return result
}
