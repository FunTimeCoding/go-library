package kestra

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/kestra-io/client-sdk/go-sdk"
)

func (c *Client) Tenants() []kestra_api_client.Tenant {
	result, r, e := c.client.TenantsAPI.Find(
		c.context,
	).Page(0).Size(10).Execute()
	errors.PanicOnWebError(r, e)

	return result.Results
}
