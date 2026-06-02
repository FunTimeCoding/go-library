package alertmanager

func (c *Client) SilenceExists(name string) (bool, error) {
	v, e := c.SilenceByRule(name)

	if e != nil {
		return false, e
	}

	return v != nil, nil
}
