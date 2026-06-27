package mock_snippet

func (c *Client) ReadFile(path string) []byte {
	return c.files[path]
}
