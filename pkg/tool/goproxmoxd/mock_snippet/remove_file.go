package mock_snippet

func (c *Client) RemoveFile(path string) {
	delete(c.files, path)
}
