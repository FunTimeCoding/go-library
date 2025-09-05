package netbox

func (c *Client) DeleteInternetAddress(identifier int32) {
	r, e := c.client.IpamAPI.IpamIpAddressesDestroy(
		c.context,
		identifier,
	).Execute()
	verifyDelete("internet", r, e)
}
