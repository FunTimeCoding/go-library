package kestra

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kestra/constant"
	"github.com/kestra-io/client-sdk/go-sdk"
)

func (c *Client) Flows(namespace string) []kestra_api_client.Flow {
	result, r, e := c.client.FlowsAPI.ListFlowsByNamespace(
		c.context,
		namespace,
		constant.MainTenant,
	).Execute()
	errors.PanicOnWebError(r, e)

	return result
}
