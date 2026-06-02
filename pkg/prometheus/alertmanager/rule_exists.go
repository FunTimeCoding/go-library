package alertmanager

func (c *Client) RuleExists(name string) (bool, error) {
	v, e := c.Rule(name)

	if e != nil {
		return false, e
	}

	return v != nil, nil
}
