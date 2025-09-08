package kestra

import "github.com/funtimecoding/go-library/pkg/kestra/constant"

func (c *Client) Namespaces() []string {
	result, r, e := c.client.FlowsAPI.ListDistinctNamespaces(
		c.context,
		constant.MainTenant,
	).Execute()
	panicOnError(e, r)

	return result
}
