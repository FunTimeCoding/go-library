package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/policy"
)

func (c *Client) NotificationPolicies(team string) []policy.PolicyProps {
	r, e := c.userClient.Policy.ListNotificationPolicies(
		c.context,
		&policy.ListNotificationPoliciesRequest{TeamId: team},
	)
	errors.PanicOnError(e)

	return r.Policies
}
