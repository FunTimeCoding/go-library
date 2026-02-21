package mock_client

func (c *Client) Remove(fingerprint string) {
	var result = c.alerts[:0]

	for _, a := range c.alerts {
		if a.Fingerprint != fingerprint {
			result = append(result, a)
		}
	}

	c.alerts = result
}
