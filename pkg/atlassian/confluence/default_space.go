package confluence

func (c *Client) DefaultSpace() string {
	if c.defaultSpace == "" {
		panic("default space not set")
	}

	return c.defaultSpace
}
