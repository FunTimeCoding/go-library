package netbox

func (c *Client) DeleteInterface(identifier int32) {
	r, e := c.client.DcimAPI.DcimInterfacesDestroy(
		c.context,
		identifier,
	).Execute()
	verifyDelete("interface", r, e)
}
