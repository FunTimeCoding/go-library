package mock_client

func (c *Client) Commits() []*Commit {
	return c.commits
}
