package confluence

func (c *Client) UserBasic() string {
	return c.basic.Get("/user/current")
}
