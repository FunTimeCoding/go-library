package netbox

func (c *Client) DeviceTagNames(name string) ([]string, error) {
	d, e := c.DeviceByNameStrict(name)

	if e != nil {
		return nil, e
	}

	var result []string

	for _, t := range d.Raw.Tags {
		result = append(result, t.GetName())
	}

	return result, nil
}
