package confluence

func (c *Client) PagesBasic() string {
	return c.basic.GetV2("/pages")
}
