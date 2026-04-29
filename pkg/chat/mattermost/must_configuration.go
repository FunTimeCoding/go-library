package mattermost

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustConfiguration() map[string]string {
	result, e := c.Configuration()
	errors.PanicOnError(e)

	return result
}
