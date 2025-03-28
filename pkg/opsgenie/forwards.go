package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/forwarding_rule"
)

func (c *Client) Forwards() []forwarding_rule.ForwardingRule {
	result, e := c.userClient.Forward.List(
		c.context,
		&forwarding_rule.ListRequest{},
	)
	errors.PanicOnError(e)

	return result.ForwardingRule
}
