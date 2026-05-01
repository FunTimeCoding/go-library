package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/power_feed"
)

func (c *Client) MustPowerFeeds() []*power_feed.Feed {
	result, e := c.PowerFeeds()
	errors.PanicOnError(e)

	return result
}
