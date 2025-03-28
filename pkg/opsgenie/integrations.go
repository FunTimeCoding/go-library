package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/integration"
)

func (c *Client) Integrations() []integration.GenericFields {
	result, e := c.userClient.Integration.List(c.context)
	errors.PanicOnError(e)

	return result.Integrations
}
