package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) AdminMode(on bool) *gitlab.Settings {
	// AdminMode is persisted
	result, r, e := c.client.Settings.UpdateSettings(
		&gitlab.UpdateSettingsOptions{AdminMode: new(on)},
	)
	panicOnError(r, e)

	return result
}
