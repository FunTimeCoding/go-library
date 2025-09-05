package netbox

func (c *Client) DeletePhysicalAddress(identifier int32) {
	r, e := c.client.DcimAPI.DcimMacAddressesDestroy(
		c.context,
		identifier,
	).Execute()
	verifyDelete("physical", r, e)
}
