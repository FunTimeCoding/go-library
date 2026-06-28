package mock_snippet

func (c *Client) WriteFile(
	path string,
	content []byte,
) {
	c.files[path] = content
}
