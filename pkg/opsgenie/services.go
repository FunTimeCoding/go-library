package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/service"
)

func (c *Client) Services() []service.Service {
	result, e := c.userClient.Service.List(c.context, &service.ListRequest{})
	errors.PanicOnError(e)

	return result.Services
}
