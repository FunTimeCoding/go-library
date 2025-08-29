package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/power_feed"
)

func (c *Client) PowerFeeds() []*power_feed.Feed {
	result, _, e := c.client.DcimAPI.DcimPowerFeedsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return power_feed.NewSlice(result.Results)
}
