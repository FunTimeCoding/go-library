package technitium

func (c *Client) get(path string, result any) error {
	r, e := c.do(path)

	if e != nil {
		return e
	}

	return c.decode(r, result)
}
