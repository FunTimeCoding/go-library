package confluence

func (c *Client) SpacesBasic() string {
	return c.basic.GetV2("/spaces")
}
