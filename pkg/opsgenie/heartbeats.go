package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/heartbeat"
)

func (c *Client) Heartbeats() []heartbeat.Heartbeat {
	result, e := c.userClient.Heartbeat.List(c.context)
	errors.PanicOnError(e)

	return result.Heartbeats
}
