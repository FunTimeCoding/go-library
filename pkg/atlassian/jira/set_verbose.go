package jira

func (c *Client) SetVerbose(b bool) *Client {
	c.verbose = b

	if c.issueOption != nil {
		c.issueOption.Verbose = b
	}

	return c
}
