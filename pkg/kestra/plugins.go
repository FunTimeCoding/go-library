package kestra

import "github.com/kestra-io/client-sdk/go-sdk"

func (c *Client) Plugins() []kestra_api_client.Plugin {
	result, r, e := c.client.PluginsAPI.ListPlugins(c.context).Execute()
	panicOnError(e, r)

	return result
}
