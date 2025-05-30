package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/escalation"
)

func (c *Client) Escalations() []escalation.Escalation {
	result, e := c.userClient.Escalation.List(c.context)
	errors.PanicOnError(e)

	return result.Escalations
}
