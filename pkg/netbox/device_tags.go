package netbox

func (c *Client) DeviceTagNames(name string) []string {
	var result []string

	for _, t := range c.DeviceByName(name, true).Raw.Tags {
		result = append(result, t.GetName())
	}

	return result
}
