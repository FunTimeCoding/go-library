package netbox

func (c *Client) DeleteInternet(identifier int32) {
	r, e := c.client.IpamAPI.IpamIpAddressesDestroy(
		c.context,
		identifier,
	).Execute()
	verifyDelete("internet", r, e)
}
