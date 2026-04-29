package proxmox

func Get[T any](
	c *Client,
	locator string,
	result *T,
) error {
	return c.client.Get(c.context, locator, result)
}
