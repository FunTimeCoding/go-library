package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/notification"
)

func (c *Client) NotificationRules(user string) []notification.SimpleNotificationRuleResult {
	result, e := c.userClient.Notification.ListRule(
		c.context,
		&notification.ListRuleRequest{UserIdentifier: user},
	)
	errors.PanicOnError(e)

	return result.SimpleNotificationRules
}
