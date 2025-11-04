package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/notification_group"
)

func (c *Client) NotificationGroups() []*notification_group.Group {
	result, r, e := c.client.ExtrasAPI.ExtrasNotificationGroupsList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return notification_group.NewSlice(result.Results)
}
