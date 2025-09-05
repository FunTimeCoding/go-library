package netbox

func (c *Client) DeleteDevice(identifier int32) {
	r, e := c.client.DcimAPI.DcimDevicesDestroy(
		c.context,
		identifier,
	).Execute()
	verifyDelete("device", r, e)
}
