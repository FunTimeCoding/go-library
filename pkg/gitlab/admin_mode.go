package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
	"k8s.io/utils/ptr"
)

func (c *Client) AdminMode(on bool) *gitlab.Settings {
	// AdminMode is persisted
	result, _, e := c.client.Settings.UpdateSettings(
		&gitlab.UpdateSettingsOptions{AdminMode: ptr.To[bool](on)},
	)
	errors.PanicOnError(e)

	return result
}
