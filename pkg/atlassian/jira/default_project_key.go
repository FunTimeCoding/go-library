package jira

func (c *Client) DefaultProjectKey() string {
	if c.defaultProjectKey == "" {
		panic("default project key not set")
	}

	return c.defaultProjectKey
}
