package kestra

import "github.com/kestra-io/client-sdk/go-sdk"

func (c *Client) Tenants() []kestra_api_client.Tenant {
	r, s, e := c.client.TenantsAPI.Find(
		c.context,
	).Page(0).Size(10).Execute()
	panicOnError(e, s)

	return r.Results
}
