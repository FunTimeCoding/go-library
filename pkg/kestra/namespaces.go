package kestra

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/kestra/constant"
)

func (c *Client) Namespaces() []string {
	result, r, e := c.client.FlowsAPI.ListDistinctNamespaces(
		c.context,
		constant.MainTenant,
	).Execute()
	errors.PanicOnWebError(r, e)

	return result
}
