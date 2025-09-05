package netbox

func (c *Client) DeletePhysical(identifier int32) {
	r, e := c.client.DcimAPI.DcimMacAddressesDestroy(
		c.context,
		identifier,
	).Execute()
	verifyDelete("physical", r, e)
}
