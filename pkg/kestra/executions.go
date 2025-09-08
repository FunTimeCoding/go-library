package kestra

import (
	"github.com/funtimecoding/go-library/pkg/kestra/constant"
	"github.com/kestra-io/client-sdk/go-sdk"
)

func (c *Client) Executions(namespace string) []kestra_api_client.FlowForExecution {
	result, r, e := c.client.ExecutionsAPI.ListFlowExecutionsByNamespace(
		c.context,
		namespace,
		constant.MainTenant,
	).Execute()
	panicOnError(e, r)

	return result
}
