package confluence

func (c *Client) DefaultPage() string {
	if c.defaultSpace == "" {
		panic("default page not set")
	}

	return c.defaultPage
}
