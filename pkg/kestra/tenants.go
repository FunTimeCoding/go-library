package kestra

import "github.com/kestra-io/client-sdk/go-sdk"

func (c *Client) Tenants() []kestra_api_client.Tenant {
	result, r, e := c.client.TenantsAPI.Find(
		c.context,
	).Page(0).Size(10).Execute()
	panicOnError(e, r)

	return result.Results
}
